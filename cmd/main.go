package main

import (
	"fmt"
	"math/rand"
	"time"

	Expense "github.com/SudipDevnath/SplitExp-backend/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user/:value", func(c *fiber.Ctx) error {

		Expenses := [20]Expense.Expense{}
		for i := range Expenses {
			Expenses[i] = Expense.NewExpense(randomDate(), randomCategory(), randomDescription(), randomName(), randomType(), randomAmount())
		}

		return c.JSON(Expenses)
		// Expenses := []Expense.Expense{}
		// return c.JSON(Expenses)
	})

	app.Listen(":8000")
}

// Generate a random date string in the format "dd-mm-yyyy"
func randomDate() string {
	min := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC).Unix()
	seconds := rand.Int63n(max-min) + min
	return time.Unix(seconds, 0).Format("02-01-2006")
}

// Generate a random category
func randomCategory() string {
	categories := []string{"Groceries", "Transport", "Utilities", "Entertainment", "Other"}
	return categories[rand.Intn(len(categories))]
}

// Generate a random description
func randomDescription() string {
	descriptions := []string{"Electricity bill", "Grocery shopping", "Bus fare", "Movie tickets", "Dinner"}
	return descriptions[rand.Intn(len(descriptions))]
}

// Generate a random name
func randomName() string {
	names := []string{"Sudip", "Anil", "Rita", "Mina", "Sita"}
	return names[rand.Intn(len(names))]
}

// Generate a random type
func randomType() string {
	types := []string{"lent", "borrowed"}
	return types[rand.Intn(len(types))]
}

// Generate a random amount
func randomAmount() string {
	amount := rand.Intn(1000) + 1
	return fmt.Sprintf("Rs.%d", amount)
}
