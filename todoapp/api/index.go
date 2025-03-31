package handler

import (
	"net/http"
	"path/filepath"

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
var nextID = 4

func Handler(w http.ResponseWriter, r *http.Request) {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// API routes
	api := router.Group("/api")
	{
		api.GET("/todos", func(c *gin.Context) {
			c.JSON(http.StatusOK, todos)
		})

		api.POST("/todos", func(c *gin.Context) {
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
		})

		api.DELETE("/todos/:id", func(c *gin.Context) {
			id := c.Param("id")
			for i, todo := range todos {
				if string(todo.ID) == id {
					todos = append(todos[:i], todos[i+1:]...)
					c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
					return
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		})
	}

	// Handle root path specially
	router.GET("/", func(c *gin.Context) {
		c.File(filepath.Join("src", "html", "index.html"))
	})

	// Serve static files
	router.Static("/images", filepath.Join("src", "images"))
	router.Static("/html", filepath.Join("src", "html"))

	router.ServeHTTP(w, r)
}
