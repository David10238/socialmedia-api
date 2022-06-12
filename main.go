package main

import (
	"github.com/David10238/socialmedia-api/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	server := fiber.New()

	server.Use(cors.New()) // allows client to connect

	user.Route(server)

	log.Fatal(server.Listen(":5000"))
}
