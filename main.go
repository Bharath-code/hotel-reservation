package main

import (
	"flag"

	"github.com/Bharath-code/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//
	//  make build -> ./bin/api --listenAddr :4000 -> this command will run the api in desired port
	listenAddr := flag.String("listenAddr", ":3000", "The listen address to api server")
	flag.Parse()
	app := fiber.New()
	//Group is like versioning our api like v1, v2 ...
	apiv1 := app.Group("/api/v1")
	app.Get("/foo", handleFoo)
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	c.JSON(map[string]string{"msg": "working just fine!"})
	return nil
}
