package routes

import (
	"todo-app/configs"
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)


func TodoRoute(r *gin.RouterGroup) {
    todoCollection := configs.GetCollection(configs.DB, "todos")
    todoController := controllers.TodoController{ Repo: todoCollection }

    r.GET("", todoController.GetTodos)
    r.POST("", todoController.PostTodo)
    r.GET("/:id", todoController.GetTodoById)
}
