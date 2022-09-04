package todo

import (
	"errors"
	"net/http"

	"github.com/ethan-stone/gin-todo/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UpdateTodoInput struct {
	Description string `json:"description"`
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var body UpdateTodoInput 

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := db.Todo{ID: uuid.MustParse(id)}
	findResult := db.DB.First(&todo)

	if findResult.Error != nil {
		log.Error(findResult.Error)
		if errors.Is(findResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	updateResult := db.DB.Model(&todo).Updates(&db.Todo{Description: body.Description})

	if updateResult.Error != nil {
		log.Error(updateResult.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	log.Infof("Todo with ID: %v updated", id)
	c.JSON(http.StatusOK, gin.H{"data": todo})
	return
}


