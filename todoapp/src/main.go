package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{
	{ID: 1, Title: "Call Mark the AI Real Estate Investor for the Full Stack Job - 405-963-2596", Status: "pending"},
	{ID: 2, Title: "Hire Mark", Status: "pending"},
	{ID: 3, Title: "Check out AIREINVESTOR.COM", Status: "pending"},
}
var nextID = 4 // Set the next ID to 4 since there are 3 placeholder tasks

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/images", "./images") // Serve images folder
	r.Static("/html", "./html")      // Serve HTML and CSS
	r.StaticFile("/", "./html/index.html") // Serve HTML at root

	// CRUD API
	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Run(":8080")
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo.ID = nextID
	nextID++
	newTodo.Status = "pending"
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range todos {
		if string(todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
