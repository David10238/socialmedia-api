package main

import (
	"github.com/David10238/socialmedia-api/user"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	server := fiber.New()

	user.Route(server)

	log.Fatal(server.Listen(":5000"))
}
