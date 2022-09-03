package main

import (
	"github.com/ethan-stone/gin-todo/db"
	"github.com/ethan-stone/gin-todo/router"
	"github.com/gin-gonic/gin"
)


func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	db.Connect()	

	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/todo", router.PostTodo)
	r.GET("/todo/:id", router.GetTodo)
	r.PATCH("/todo/:id", router.PatchTodo)
	r.Run() // listen and serve on 0.0.0.0:8080
}