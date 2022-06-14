package main

import (
	"fmt"
	"github.com/David10238/socialmedia-api/db"
	"github.com/David10238/socialmedia-api/message"
	"github.com/David10238/socialmedia-api/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	db.Initialize()

	fmt.Println(db.Conn.Name())

	server := fiber.New()

	server.Use(cors.New()) // allows client to connect

	user.Route(server)
	message.Route(server)

	log.Fatal(server.Listen(":5000"))
}
