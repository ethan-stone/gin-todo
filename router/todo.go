package router

import (
	"errors"
	"net/http"

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