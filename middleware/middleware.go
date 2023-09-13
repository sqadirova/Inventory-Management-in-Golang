package middleware

import (
	"IM-Golang/CONSTANTS"
	my_errors "IM-Golang/errors"
	"github.com/gofiber/fiber/v2"
	//userPack "im-backoffice/user"
)

var Common = func(c *fiber.Ctx) error {
	c.Request().Header.Add("Content-Type", "application/json")
	return c.Next()
}

func Protect(args ...string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := ExtractBearerToken(c.GetReqHeaders()["Authorization"])

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).
				JSON(my_errors.NewResponseByKey("bearer_token_missing", CONSTANTS.LANGUAGE))
		}

		accessToken, _, err := VerifyAccessToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(my_errors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
		}

		roleFromToken, claims, err := ParseToken(accessToken)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(my_errors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
		}

		c.Locals("jwt", accessToken)
		c.Locals("props", claims)

		if len(args) == 0 {
			return c.Next()
		}

		for _, uRole := range args {
			if uRole == roleFromToken {
				return c.Next()
			}
		}

		for _, role := range args {
			if role != roleFromToken {
				return c.Status(fiber.StatusUnauthorized).
					JSON(my_errors.NewResponseByKey("access_not_authorized", CONSTANTS.LANGUAGE))
			}
		}

		return c.Next()
	}
}
