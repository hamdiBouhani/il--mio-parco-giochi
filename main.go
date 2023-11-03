package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hamdiBouhani/il-mio-parco-giochi/stack"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func EvaluatePostfixExpressionHandler(c *fiber.Ctx, expression string) error {
	result, err := stack.EvaluatePostfixExpression(expression)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString("Result: " + fmt.Sprintf("%v", result))
}

func main() {

	// Connect to database
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_DATABASE"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	app.Get("/evaluate-postfix-expression", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Request().URI().QueryString()))
		expression := c.Query("expression")
		return EvaluatePostfixExpressionHandler(c, expression)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
