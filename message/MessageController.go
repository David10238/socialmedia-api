package message

import (
	"encoding/json"
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

	request := &messageRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.SendStatus(http.StatusNotAcceptable)
	}

	// todo load from database
	res := newMessageResponse(
		[]message{
			{
				Id:       5,
				AuthorId: userId,
				Message:  "This is a short message",
				IsPublic: false,
			},
			{
				Id:       4,
				AuthorId: 4,
				Message:  "This is a slightly longer message",
				IsPublic: false,
			},
			{
				Id:       2,
				AuthorId: 9,
				Message:  "This is the longest message because I want to see if the messages properly wrap around and expand the message box, so that it can fit as many lines as the api returns or if I will have to alter the css,so that it actually performs well.",
				IsPublic: false,
			},
		})

	bytes, _ := json.Marshal(res)
	return c.Send(bytes)
}
