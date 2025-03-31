package api

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
var nextID = 4

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Serve static files
	router.Static("/images", "./src/images")
	router.Static("/html", "./src/html")
	router.StaticFile("/", "./src/html/index.html")

	// CRUD API
	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	router.POST("/todos", func(c *gin.Context) {
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

	router.DELETE("/todos/:id", func(c *gin.Context) {
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

	router.ServeHTTP(w, r)
}
