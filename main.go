package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID     string `json:"id"`
	Nama   string `json:"nama"`
	Status bool   `json:"status"`
}

var todos = []todo{
	{ID: "1", Nama: "Abik", Status: true},
	{ID: "2", Nama: "Ketir", Status: false},
	{ID: "3", Nama: "Hilwa", Status: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var NewTodo todo

	if err := context.BindJSON(&NewTodo); err != nil {
		return
	}

	todos = append(todos, NewTodo)
	context.IndentedJSON(http.StatusCreated, NewTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toogleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}

	todo.Status = !todo.Status
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo Not Found")
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toogleTodoStatus)
	router.POST("/todos", addTodo)
	router.DELETE("/todos", deleteTodo)
	router.Run("localhost:9090")
}
