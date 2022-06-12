package user

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Route(server *fiber.App) {
	server.Put("/user/login", handleLogin)
}

func handleLogin(c *fiber.Ctx) error {
	credentials := &userCredentials{}
	if err := c.BodyParser(credentials); err != nil || len(credentials.Email) == 0 || len(credentials.Password) == 0 {
		return c.SendStatus(http.StatusNotAcceptable)
	}

	keys := userKeys{}
	// todo check from a database
	if len(credentials.Name) == 0 {
		// user is trying to login
		keys.UserId = 0
		keys.AuthKey = "LoginKey"

		if credentials.Password == "forbid" {
			return c.SendStatus(http.StatusForbidden)
		}

		dat, _ := json.Marshal(keys)
		return c.Send(dat)
	}

	// user is trying to create a new account
	keys.UserId = 0
	keys.AuthKey = "SignupKey"
	if credentials.Password == "forbid" {
		return c.SendStatus(http.StatusForbidden)
	}
	dat, _ := json.Marshal(keys)
	return c.Send(dat)
}
