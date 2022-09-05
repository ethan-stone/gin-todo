package todo

import (
	"errors"

	"github.com/ethan-stone/go-todo/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UpdateTodoInput struct {
	Description string `json:"description"`
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(UpdateTodoInput)

	if err := c.BodyParser(&body); err != nil {
		log.Error().Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	todo := db.Todo{ID: uuid.MustParse(id)}
	findResult := db.DB.First(&todo)

	if findResult.Error != nil {
		log.Error().Msg(findResult.Error.Error())
		if errors.Is(findResult.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	updateResult := db.DB.Model(&todo).Updates(&db.Todo{Description: body.Description})

	if updateResult.Error != nil {
		log.Error().Msg(updateResult.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	log.Info().Msgf("Todo with ID: %v updated", id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todo,
	})
}
