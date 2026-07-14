package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	router := gin.Default()
	router.GET("/get", getDetails)
	router.Run("localhost:8080")
}
