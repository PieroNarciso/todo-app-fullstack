package routes

import (
	"github.com/PieroNarciso/todo-app-fullstack/src/src/configs"
	"github.com/PieroNarciso/todo-app-fullstack/src/src/controllers"
	"github.com/gin-gonic/gin"
)


func TodoRoute(r *gin.RouterGroup) {
    todoCollection := configs.GetCollection(configs.DB, "todos")
    todoController := controllers.TodoController{ Repo: todoCollection }

    r.GET("", todoController.GetTodos)
    r.POST("", todoController.PostTodo)
    r.GET("/:id", todoController.GetTodoById)
}
