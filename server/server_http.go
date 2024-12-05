package server

import (
	"log"

	internal "github.com/njavilas2015/godis/internal"

	"github.com/gin-gonic/gin"
)

func HandlerHSet(args []string) <-chan string {
	response := make(chan string)

	go func() {
		defer close(response)

		if len(args) < 3 || len(args)%2 == 0 {
			response <- "ERROR: HSET incorrectly configured"
			return
		}

		key := args[0]

		for i := 1; i < len(args); i += 2 {
			field := args[i]
			value := args[i+1]

			response <- internal.Hs.AddJobHSet(key, field, value)
		}
	}()

	return response
}

func HandlerHGet(args []string) <-chan string {
	response := make(chan string)

	go func() {
		defer close(response)

		if len(args) != 2 {
			response <- "ERROR: Invalid number of arguments for HGET"
			return
		}

		key := args[0]
		field := args[1]

		response <- internal.Hs.AddJobHGet(key, field)
	}()

	return response
}

func HTTPServer() {

	r := gin.Default()

	r.POST("/hset", func(c *gin.Context) {

		var json struct {
			Args []string `json:"args"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		args := json.Args

		responseChan := HandlerHSet(args)
		response := <-responseChan

		c.String(200, response)
	})

	r.POST("/hget", func(c *gin.Context) {

		var json struct {
			Args []string `json:"args"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		args := json.Args

		responseChan := HandlerHGet(args)
		response := <-responseChan

		c.String(200, response)
	})

	err := r.Run(":8080")

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
