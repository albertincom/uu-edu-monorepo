package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()

	// Serve static files from the React build directory
	r.Static("/assets", "./static/assets")

	// Serve the React index.html file
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Catch-all route to handle React routing
	r.NoRoute(func(c *gin.Context) {
		if filepath.Ext(c.Request.URL.Path) == "" {
			c.File("./static/index.html")
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		}
	})

	// Start the server on port 8080
	r.Run(":8080")
}
