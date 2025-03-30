package main

import (
	"net/http"
	"strconv"

	uuid "github.com/arcleife/ids-generator/internal/uuid"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Get uuid
	r.GET("/uuid", func(c *gin.Context) {
		delimiter := c.DefaultQuery("delimiter", "-")

		response := uuid.New().All(delimiter)
		c.String(http.StatusOK, response)
	})

	// Get uuids
	r.GET("/uuid/bulk", func(c *gin.Context) {
		n, err := strconv.Atoi(c.Query("n"))
		delimiter := c.DefaultQuery("delimiter", "-")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"value":  nil,
				"status": "n must be integer"})
		} else {
			response := uuid.Bulk(n, delimiter)
			c.JSON(http.StatusOK, gin.H{
				"value":  response,
				"status": "OK"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")
}
