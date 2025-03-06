package server

import (
	"log"

	internal "github.com/njavilas2015/godis/internal"

	"github.com/gin-gonic/gin"
)

func HTTPServer() {

	r := gin.Default()

	hash := r.Group("/hash")
	{

		hash.POST("/", func(c *gin.Context) {

			var json struct {
				Args []string `json:"args"`
			}

			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(400, gin.H{"error": "Invalid JSON"})
				return
			}

			args := json.Args

			responseChan := internal.HandlerHashStore("HSET", args)

			response := <-responseChan

			c.String(200, response)
		})

		hash.GET("/", func(c *gin.Context) {

			var json struct {
				Args []string `json:"args"`
			}

			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(400, gin.H{"error": "Invalid JSON"})
				return
			}

			args := json.Args

			responseChan := internal.HandlerHashStore("HGET", args)

			response := <-responseChan

			c.String(200, response)
		})

		hash.GET("/all", func(c *gin.Context) {

			responseChan := internal.HandlerHashStore("ALL", []string{})

			response := <-responseChan

			c.String(200, response)
		})

		hash.DELETE("/", func(c *gin.Context) {

			responseChan := internal.HandlerHashStore("DROP", []string{})

			response := <-responseChan

			c.String(200, response)
		})
	}

	err := r.Run(":8080")

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
