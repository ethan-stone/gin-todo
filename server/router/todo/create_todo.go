package todo

import (
	"net/http"

	"github.com/ethan-stone/gin-todo/db"
	"github.com/ethan-stone/gin-todo/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CreateTodoInput struct {
	Description string `json:"description" binding:"required"`
}

func Create(c *gin.Context) {
	claims := c.MustGet("Claims").(*middleware.Claims)

	log.WithFields(log.Fields{
		"claims": claims,
	}).Info("User")

	var body CreateTodoInput 

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := db.Todo{Description: body.Description}
	result := db.DB.Create(&todo)

	if result.Error != nil {
		log.Error(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	log.Infof("Todo with ID: %v created", todo.ID)
	c.JSON(http.StatusOK, gin.H{"data": todo})
	return
}

