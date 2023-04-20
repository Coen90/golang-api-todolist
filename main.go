package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos []todo = []todo{
	{ID: "1", Item: "clean bed", Completed: false},
	{ID: "2", Item: "make crud", Completed: false},
	{ID: "3", Item: "connect with db", Completed: false},
}

func getTodos(context *gin.Context) {
	fmt.Println(len(todos))
	if cnt := len(todos); cnt <= 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "There is no todos"})
		return
	}
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo
	err := context.BindJSON(&newTodo)

	if err != nil {
		fmt.Println("error!!", err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Cannot add todo"})
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id string) (*todo, int, error) {
	for i, v := range todos {
		if v.ID == id {
			return &todos[i], i, nil
		}
	}
	return nil, 0, errors.New("There is no todo")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, _,err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func patchTodo(context *gin.Context) {
	id := context.Param("id")
	todo, _, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	_, index, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	todos = append(todos[:index], todos[index+1:]...)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted completely"})
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", patchTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.POST("/todos", addTodos)

	router.Run("localhost:8080")
}

// package main

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type todo struct {
// 	ID        string `json:"id"`
// 	Item      string `json:"item"`
// 	Completed bool   `json:"completed"`
// }

// var todos = []todo{
// 	{ID: "1", Item: "Clean Room", Completed: false},
// 	{ID: "2", Item: "Read Book", Completed: false},
// 	{ID: "3", Item: "Record Video", Completed: false},
// }

// func getTodos(context *gin.Context) {
// 	context.IndentedJSON(http.StatusOK, todos)
// }

// func addTodo(context *gin.Context) {
// 	var newTodo todo

// 	if err := context.BindJSON(&newTodo); err != nil {
// 		return
// 	}
// 	todos = append(todos, newTodo)

// 	context.IndentedJSON(http.StatusOK, newTodo)
// }

// func getTodoById(id string) (*todo, error) {
// 	for i, t := range todos {
// 		if t.ID == id {
// 			return &todos[i], nil
// 		}
// 	}

// 	return nil, errors.New("todo not found")
// }

// func getTodo(context *gin.Context) {
// 	id := context.Param("id")
// 	todo, err := getTodoById(id)

// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
// 		return
// 	}

// 	context.IndentedJSON(http.StatusOK, todo)
// }

// func toggleTodoStatus(context *gin.Context) {
// 	id := context.Param("id")
// 	todo, err := getTodoById(id)

// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, "Todo Not Found")
//      return
// 	}

// 	todo.Completed = !todo.Completed

// 	context.IndentedJSON(http.StatusOK, todo)
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/todos", getTodos)
// 	router.GET("/todos/:id", getTodo)
// 	router.PATCH("/todos/:id", toggleTodoStatus)
// 	router.POST("/todos", addTodo)
// 	router.Run("localhost:8080")
// }
