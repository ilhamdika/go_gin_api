package main

import (
	"github.com/gin-gonic/gin"
	"nayaka/routes"
	"net/http"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	})

	r.Run(":8080")
}
