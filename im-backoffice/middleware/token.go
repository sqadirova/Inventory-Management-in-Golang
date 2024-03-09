package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"im-backoffice/CONSTANTS"
	"im-backoffice/keycloak"
	"log"
	"strings"
)

var newKeyCloak = keycloak.NewKeycloak()

func ExtractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func VerifyAccessToken(accessToken string) (string, jwt.MapClaims, error) {
	token := ExtractBearerToken(accessToken)

	if token == "" {
		return "", nil, errors.New("bearer_token_missing")
	}

	result, err := newKeyCloak.Gocloak.RetrospectToken(context.Background(), token, newKeyCloak.ClientId, newKeyCloak.ClientSecret, newKeyCloak.Realm)

	if err != nil {
		log.Println(err)
		return "", nil, errors.New("invalid_token")
	}

	if !*result.Active {
		return "", nil, errors.New("access_token_expired")
	}

	_, claims, err := ParseToken(token)
	if err != nil {
		return "", nil, err
	}

	return token, claims, nil
}

// ParseToken Parse token and return role and claims from token
func ParseToken(accessToken string) (string, jwt.MapClaims, error) {
	if token, _ := jwt.Parse(accessToken, nil); token != nil {
		claims := token.Claims.(jwt.MapClaims)
		resourceAccess := claims["resource_access"].(map[string]interface{})
		realmAccess := resourceAccess[newKeyCloak.ClientId].(map[string]interface{})
		roles := realmAccess["roles"].([]interface{})

		for _, role := range roles {
			if role == CONSTANTS.WORKER || role == CONSTANTS.ADMIN || role == CONSTANTS.INSPECTOR {
				return role.(string), claims, nil
			}
		}
	}

	return "", nil, errors.New("invalid_token")
}
