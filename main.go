package main

import (
	"os"
	"github.com/eniworoeva/JWT-auth/routes"
	"github.com/gin-gonic/gin"
)


func main()  {
	port := os.Getenv("PORT")
	if port == ""{
		port = "800 0"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		 c.JSON(200, gin.H{"success":"Access granted for API-1"})  
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success":"Access granted for API-2"})
	})

	router.Run(":" + port)
}