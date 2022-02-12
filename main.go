package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/techagentng/GOAUTH-JJWT/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	//Calling the external route function and parsing gin router
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success":"Access granted for ap-1"})
	})

	router.GET("api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success":"Access granted for ap-2"})
	})

	router.Run(":" + port)
}

