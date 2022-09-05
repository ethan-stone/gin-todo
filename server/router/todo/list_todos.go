package todo

import (
	"strconv"

	"github.com/ethan-stone/go-todo/db"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func List(c *fiber.Ctx) error {
	skipQuery := c.Query("skip")
	limitQuery := c.Query("limit")

	var skip int

	if skipQuery == "" {
		skip = 0
	} else {
		skipParse, skipParseErr := strconv.Atoi(skipQuery)
		if skipParseErr != nil {
			log.Error().Msg(skipParseErr.Error())
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": skipParseErr.Error(),
			})
		}
		skip = skipParse
	}

	var limit int

	if limitQuery == "" {
		limit = 50
	} else {
		limitParse, limitParseErr := strconv.Atoi(limitQuery)
		if limitParseErr != nil {
			log.Error().Msg(limitParseErr.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": limitParseErr.Error(),
			})
		}
		limit = limitParse
	}

	var todos []db.Todo

	result := db.DB.Limit(limit).Offset(skip).Find(&todos)

	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	log.Info().Msg("Todos retrieved")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": todos})
}
