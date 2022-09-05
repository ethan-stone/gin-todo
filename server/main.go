package main

import (
	"github.com/ethan-stone/go-todo/db"
	"github.com/ethan-stone/go-todo/middleware/logger"
	"github.com/ethan-stone/go-todo/middleware/supabaseauth"
	"github.com/ethan-stone/go-todo/router/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "pong",
	})
}

func main() {
	db.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
	}))
	app.Use(logger.New())
	app.Use(supabaseauth.New())

	app.Get("/ping", ping)
	app.Post("/todo", todo.Create)
	app.Get("/todo/:id", todo.Get)
	app.Get("/todo", todo.List)
	app.Patch("/todo/:id", todo.Update)
	app.Listen(":8080")
}
