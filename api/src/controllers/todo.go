package controllers

import (
	"fmt"
	"net/http"

	"github.com/PieroNarciso/todo-app-fullstack/src/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoController struct {
    Repo *mongo.Collection
}

func (cc *TodoController) GetTodos(c *gin.Context) {
    ctx, cancel := InitContext()
    defer cancel()
    todos := []models.Todo{}
    results, err := cc.Repo.Find(ctx, bson.M{})
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, bson.M{})
        return
    }
    defer results.Close(ctx)
    for results.Next(ctx) {
        var todo models.Todo
        if err = results.Decode(&todo); err != nil {
            c.IndentedJSON(http.StatusInternalServerError, bson.M{})
            return
        }
        todos = append(todos, todo)
    }
    c.IndentedJSON(http.StatusOK, todos)
}

func (cc *TodoController) PostTodo(c *gin.Context) {
    ctx, cancel := InitContext()
    defer cancel()
    var todo models.Todo
    if err := c.BindJSON(&todo); err != nil {
        c.IndentedJSON(http.StatusBadRequest, bson.M{})
        return
    }
    newTodo := models.Todo{
        Id: primitive.NewObjectID(),
        Title: todo.Title,
        Description: todo.Description,
        Completed: false,
    }
    _, err := cc.Repo.InsertOne(ctx, newTodo)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, bson.M{})
        return
    }
    fmt.Println(newTodo)
    c.IndentedJSON(http.StatusCreated, newTodo)
}

func (cc *TodoController) GetTodoById(c *gin.Context) {
    var todo models.Todo
    ctx, cancel := InitContext()
    defer cancel()
    id := c.Param("id")
    objId, _ := primitive.ObjectIDFromHex(id)
    err := cc.Repo.FindOne(ctx, bson.M{"_id": objId}).Decode(&todo)
    if err != nil {
        fmt.Println(err)
        c.IndentedJSON(
            http.StatusNotFound,
            bson.M{},
        )
        return
    }
    c.IndentedJSON(http.StatusOK, todo)
}
