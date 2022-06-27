package message

import (
	"encoding/json"
	"github.com/David10238/socialmedia-api/db"
	"github.com/David10238/socialmedia-api/user"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Route(server *fiber.App) {
	server.Put("/messages/request", handleRequest)
}

func handleRequest(c *fiber.Ctx) error {
	userId, authError := user.CheckAuthentication(c)
	if authError != nil {
		return authError
	}

	request := messageRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.SendStatus(http.StatusNotAcceptable)
	}

	messages := []db.Message{}

	// user is loading from own posts
	if userId == request.TargetId && !request.SearchFriends {
		db.Conn.Find(&messages, "author_id = ?", request.TargetId)
	}

	// user is loading from among friends
	if userId == request.TargetId && !request.SearchFriends {
		// todo implement
	}

	// user is loading from another user
	if userId == request.TargetId && !request.SearchFriends {
		// todo implement
	}

	res := newMessageResponse(messages)
	bytes, _ := json.Marshal(res)
	return c.Send(bytes)
}
