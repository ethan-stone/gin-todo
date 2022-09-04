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

func Get(c *gin.Context) {
	id := c.Param("id")

	todo := db.Todo{ID: uuid.MustParse(id)}
	result := db.DB.First(&todo)

	if result.Error != nil {
		log.Error(result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	log.Infof("Todo with ID: %v retrieved", id)
	c.JSON(http.StatusOK, gin.H{"data": todo})
	return
}

