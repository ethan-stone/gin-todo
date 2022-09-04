package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ethan-stone/gin-todo/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostTodoInput struct {
	Description string `json:"description" binding:"required"`
}

func PostTodo(c *gin.Context) {
	var body PostTodoInput

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

type PatchTodoInput struct {
	Description string `json:"description"`
}

func PatchTodo(c *gin.Context) {
	id := c.Param("id")

	var body PatchTodoInput 

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

func GetTodo(c *gin.Context) {
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

func GetTodos(c *gin.Context) {
	skipQuery := c.Query("skip")
	limitQuery := c.Query("limit")

	var skip int

	if skipQuery == "" {
		skip = 0
	} else {
		skipParse, skipParseErr := strconv.Atoi(skipQuery)
		if skipParseErr != nil {
			log.Error(skipParseErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": skipParseErr.Error()})
			return
		}
		skip = skipParse
	}

	var limit int

	if limitQuery == "" {
		limit = 50 
	} else {
		limitParse, limitParseErr := strconv.Atoi(limitQuery)
		if limitParseErr != nil {
			log.Error(limitParseErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": limitParseErr.Error()})
			return
		}
		limit =limitParse 
	}

	var todos []db.Todo

	result := db.DB.Limit(limit).Offset(skip).Find(&todos)

	if result.Error != nil {
		log.Error(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	
	log.Infof("Todos retrieved")
	c.JSON(http.StatusOK, gin.H{"data": todos})
	return
}