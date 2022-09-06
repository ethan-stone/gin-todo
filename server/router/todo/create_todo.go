package todo

import (
	"github.com/ethan-stone/go-todo/db"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type CreateTodoInput struct {
	Description string `json:"description" binding:"required"`
}

func Create(c *fiber.Ctx) error {
	body := new(CreateTodoInput)

	if err := c.BodyParser(body); err != nil {
		log.Error().Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	todo := db.Todo{Description: body.Description}
	result := db.DB.Create(&todo)

	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	log.Info().Msgf("Todo with ID: %v created", todo.ID)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todo,
	})
}
