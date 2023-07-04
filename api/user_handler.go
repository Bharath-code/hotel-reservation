package api

import (
	"github.com/Bharath-code/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

// name of the function starts with caps because to make it public so that we can access in other parts of the program
func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "bharath",
		LastName:  "kumar",
	}
	c.JSON(u)
	return nil
}

func HandleGetUser(c *fiber.Ctx) error {
	c.JSON("Zelda")
	return nil
}
