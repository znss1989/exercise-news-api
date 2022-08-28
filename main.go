package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type channel struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// type article struct {
// 	ID        string `json:"id"`
// 	ChannelID string `json:"channelID"`
// 	Url       string `json:"url"`
// 	WordCount int    `json:"wordCount"`
// }

var channels = []channel{
	{ID: "1", Title: "fashion"},
	{ID: "2", Title: "science"},
	{ID: "3", Title: "auto"},
}

func main() {
	// database set up
	db, err := sql.Open("sqlite3", "./news.db")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("SQLite Connected!")

	defer db.Close()

	// gin server
	router := gin.Default()
	router.GET("/channels", getChannels)

	router.Run("localhost:8080")
}

func getChannels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, channels)
}
