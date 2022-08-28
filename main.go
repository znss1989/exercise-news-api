package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Channel struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// type article struct {
// 	ID        int64 `json:"id"`
// 	ChannelID string `json:"channelID"`
// 	Url       string `json:"url"`
// 	WordCount int    `json:"wordCount"`
// }

// var channels = []Channel{
// 	{ID: "1", Title: "fashion"},
// 	{ID: "2", Title: "science"},
// 	{ID: "3", Title: "auto"},
// }

func main() {
	// database set up
	var err error
	db, err = sql.Open("sqlite3", "./news.db")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("##### SQLite Connected! #####")

	defer db.Close()

	// gin server
	router := gin.Default()
	router.GET("/channel", getChannels)
	router.POST("/channel", addChannel)

	router.Run("localhost:8080")
}

// APIs
func getChannels(c *gin.Context) {
	channels, err := queryChannels()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, channels)
}

func addChannel(c *gin.Context) {
	var chn Channel
	if err := c.BindJSON(&chn); err != nil {
		return
	}
	id, err := insertChannel(chn)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, id)
}

// Queries
func queryChannels() ([]Channel, error) {
	var channels []Channel

	rows, err := db.Query("SELECT title FROM channels")
	if err != nil {
		return nil, fmt.Errorf("queryChannels: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chn Channel
		if err := rows.Scan(&chn.Title); err != nil {
			return nil, fmt.Errorf("queryChannels: %v", err)
		}
		channels = append(channels, chn)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryChannels: %v", err)
	}
	return channels, nil
}

func insertChannel(chn Channel) (int64, error) {
	result, err := db.Exec("INSERT INTO channels (title) VALUES (?)", chn.Title)
	if err != nil {
		return 0, fmt.Errorf("insertChannel: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insertChannel: %v", err)
	}
	return id, nil
}
