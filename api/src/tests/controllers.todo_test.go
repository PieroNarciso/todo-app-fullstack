package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PieroNarciso/todo-app-fullstack/src/src/configs"
	"github.com/PieroNarciso/todo-app-fullstack/src/src/controllers"
	"github.com/PieroNarciso/todo-app-fullstack/src/src/models"
	"github.com/PieroNarciso/todo-app-fullstack/src/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)


type TodoTestSuite struct {
    suite.Suite
    coll *mongo.Collection
    router *gin.Engine
}

func (suite *TodoTestSuite) SetupTest() {
    suite.coll = configs.GetCollection(configs.DB, "todos")
    suite.router = GetRouter()
    routes.TodoRoute(suite.router.Group("/todos"))
}

func (suite *TodoTestSuite) AfterTest(_, _ string) {
    ctx, _ := controllers.InitContext()
    suite.coll.Drop(ctx)
}

func (suite *TodoTestSuite) TestGetOne() {
    newTodo := models.Todo{
        Title: "Title",
        Description: "description",
    }
    ctx, _ := controllers.InitContext()
    suite.coll.InsertOne(ctx, newTodo)
    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/todos", nil)
    suite.router.ServeHTTP(w, req)
    var todos []models.Todo
    json.NewDecoder(w.Body).Decode(&todos)
    assert.Equal(suite.T(), http.StatusOK, w.Code)
    assert.Equal(suite.T(), 1, len(todos))
}

func (suite *TodoTestSuite) TestGetLenZero() {
    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/todos", nil)
    suite.router.ServeHTTP(w, req)
    var todos []models.Todo
    json.NewDecoder(w.Body).Decode(&todos)
    assert.Equal(suite.T(), 0, len(todos))
}

func (suite *TodoTestSuite) TestCreateOne() {
    w := httptest.NewRecorder()
    todoPost := models.Todo{ Title: "Hola" }
    var buf bytes.Buffer
    _ = json.NewEncoder(&buf).Encode(todoPost)
    req := httptest.NewRequest(http.MethodPost, "/todos", &buf)
    suite.router.ServeHTTP(w, req)
    var todo models.Todo
    json.NewDecoder(w.Body).Decode(&todo)
    assert.NotEmpty(suite.T(), todo.Id)
    assert.Equal(suite.T(), http.StatusCreated, w.Code)
}

func TestTodoTestSuite(t *testing.T) {
    suite.Run(t, new(TodoTestSuite))
}
