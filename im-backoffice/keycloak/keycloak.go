package keycloak

import (
	"github.com/Nerzal/gocloak/v12"
	"im-backoffice/config"
)

type Keycloak struct {
	Gocloak      *gocloak.GoCloak
	ClientId     string
	ClientSecret string
	Realm        string
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		Gocloak:      gocloak.NewClient(config.Configuration.KeycloakConfig.Path),
		ClientId:     config.Configuration.KeycloakConfig.ClientName,
		ClientSecret: config.Configuration.KeycloakConfig.ClientSecretKey,
		Realm:        config.Configuration.KeycloakConfig.RealmName,
	}
}
