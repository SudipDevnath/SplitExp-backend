package main

import (
	Expense "github.com/SudipDevnath/SplitExp-backend/internal"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user/:value", func(c *fiber.Ctx) error {

		Expense := Expense.NewExpense("02-08-2024", "Other", "Electricity bill", "Sudip", "lent", "Rs.100")

		return c.JSON(Expense)
		// => Get request with value: hello world
	})

	app.Listen(":8000")
}
