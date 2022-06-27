package user

import (
	"encoding/json"
	"github.com/David10238/socialmedia-api/db"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

func Route(server *fiber.App) {
	server.Put("/user/login", handleLogin)
}

func handleLogin(c *fiber.Ctx) error {
	credentials := &userCredentials{}
	if err := c.BodyParser(credentials); err != nil || len(credentials.Email) == 0 || len(credentials.Password) == 0 {
		return c.SendStatus(http.StatusNotAcceptable)
	}

	if len(credentials.Name) == 0 {
		user := db.User{}
		foundUser := db.Conn.First(&user, "`email` = ?", credentials.Email).RowsAffected != 0

		if !foundUser || user.Password != credentials.Password {
			return c.SendStatus(http.StatusForbidden)
		}

		dat, _ := json.Marshal(userKeys{
			UserId:  user.ID,
			UserKey: user.Key,
		})
		return c.Send(dat)
	}

	// user is trying to create a new account
	user := db.User{}
	foundUser := db.Conn.First(&user, "`email` = ?", credentials.Email).RowsAffected != 0
	if foundUser { // make sure email isn't being used
		return c.SendStatus(http.StatusForbidden)
	}

	user = db.User{
		LastKeyUpdateAt: time.Now(),
		Key:             newKey(),
		Email:           credentials.Email,
		Password:        credentials.Password,
		Name:            credentials.Name,
	}
	db.Conn.Create(&user)

	dat, _ := json.Marshal(userKeys{
		UserId:  user.ID,
		UserKey: user.Key,
	})
	return c.Send(dat)
}
