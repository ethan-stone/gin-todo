package todo

import (
	"errors"

	"github.com/ethan-stone/go-todo/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Get(c *fiber.Ctx) error {
	id := c.Params("id")

	todo := db.Todo{ID: uuid.MustParse(id)}
	result := db.DB.First(&todo)

	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	log.Info().Msgf("Todo with ID: %v retrieved", id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todo,
	})
}
