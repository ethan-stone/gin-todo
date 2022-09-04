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
		log.WithFields(log.Fields{
			"resource": "todos",
			"todo_id": nil,
		}).Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := db.Todo{Description: body.Description}
	result := db.DB.Create(&todo)

	if result.Error != nil {
		log.WithFields(log.Fields{
			"resource": "todos",
			"todo_id": nil, 
		}).Error(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	log.WithFields(log.Fields{
		"resource": "todos",
		"todo_id": todo.ID,
	}).Infof("Todo with ID: %v created", todo.ID)
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
		log.WithFields(log.Fields{
			"resource": "todos",
			"todo_id": id,
		}).Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := db.Todo{ID: uuid.MustParse(id)}
	findResult := db.DB.First(&todo)

	if findResult.Error != nil {
		log.WithFields(log.Fields{
			"resource": "todos",
			"todo_id": id,
		}).Error(findResult.Error)
		if errors.Is(findResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	updateResult := db.DB.Model(&todo).Updates(&db.Todo{Description: body.Description})

	if updateResult.Error != nil {
		log.WithFields(log.Fields{
			"resource": "todos",
			"todo_id": id,
		}).Error(updateResult.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	log.WithFields(log.Fields{
		"resource": "todos",
		"todo_id": todo.ID,
		"updates": body,
	}).Infof("Todo with ID: %v updated", id)

	c.JSON(http.StatusOK, gin.H{"data": todo})
	return
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := db.RetrieveTodo(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

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
			c.JSON(http.StatusBadRequest, gin.H{"error": limitParseErr.Error()})
			return
		}
		limit =limitParse 
	}

	todos, err := db.ListTodos(&db.Todo{}, skip, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": todos})
	return
}