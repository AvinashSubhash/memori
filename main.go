package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTopicRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func createTopic(c *gin.Context) {
	var request CreateTopicRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var topicInterval int = 1
	var topicStatus string = "incomplete"
	println(request.Description, topicInterval, request.Name, topicStatus)
	// Feeder function to pass to database - Not Ready
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

type album struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var albums = []album{
	{ID: "1", Name: "avinash"},
	{ID: "2", Name: "akash"},
}

func getDetails(call *gin.Context) {
	call.IndentedJSON(http.StatusOK, albums[1])
}

func main() {
	println(albums)
	ConnectDatabase()
	router := gin.Default()
	router.GET("/get", getDetails)
	router.POST("/create", createTopic)
	router.Run("localhost:8080")
}
