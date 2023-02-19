// installing go - new pc / dubai 

package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get(path:"/", func (c *fiber.Ctx) error {
		return c.SendString(body:"hi there")
	})

	app.Listen(addr: ":3000")
}