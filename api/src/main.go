package main

import (
	"github.com/PieroNarciso/todo-app-fullstack/src/src/routes"
	"github.com/gin-gonic/gin"
)




func main() {

    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    todoGroup := r.Group("/todos")
    routes.TodoRoute(todoGroup)

    r.Run()
}
