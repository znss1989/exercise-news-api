package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type channel struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type article struct {
	ID        string `json:"id"`
	ChannelID string `json:"channelID"`
	Url       string `json:"url"`
	WordCount int    `json:"wordCount"`
}

var channels = []channel{
	{ID: "1", Title: "fashion"},
	{ID: "2", Title: "science"},
	{ID: "3", Title: "auto"},
}

func main() {
	router := gin.Default()
	router.GET("/channels", getChannels)

	router.Run("localhost:8080")
}

func getChannels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, channels)
}
