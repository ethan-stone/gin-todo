package main

import (
	"github.com/ethan-stone/gin-todo/db"
	"github.com/ethan-stone/gin-todo/router/todo"
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
	r.POST("/todo", todo.Create)
	r.GET("/todo/:id", todo.Get)
	r.GET("/todo", todo.List)
	r.PATCH("/todo/:id", todo.Update)
	r.Run() // listen and serve on 0.0.0.0:8080
}