package main

import (
	"github.com/ethan-stone/gin-todo/db"
	"github.com/ethan-stone/gin-todo/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)


func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	db.Connect()	

	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/todo", router.PostTodo)
	r.GET("/todo/:id", router.GetTodo)
	r.GET("/todo", router.GetTodos)
	r.PATCH("/todo/:id", router.PatchTodo)
	r.Run() // listen and serve on 0.0.0.0:8080
}