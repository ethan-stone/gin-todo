package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ethan-stone/gin-todo/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostTodoInput struct {
	Description string `json:"description" binding:"required"`
}

func PostTodo(c *gin.Context) {
	var body PostTodoInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := db.InsertTodo(&db.Todo{Description: body.Description})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
	return
}

type PatchTodoInput struct {
	Description string `json:"description"`
}

func PatchTodo(c *gin.Context) {
	var body PatchTodoInput 
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
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

	updatedTodo, updateError := db.UpdateTodo(todo.ID.String(), &db.Todo{Description: body.Description})


	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedTodo})
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