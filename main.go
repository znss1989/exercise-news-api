package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

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
	fmt.Println("========= SQLite Connected! =========")

	defer db.Close()

	// gin server
	router := gin.Default()
	router.GET("/channel", getChannels)
	router.POST("/channel", addChannel)
	router.GET("/channel/:id", getArticles)
	router.POST("/channel/:id", addArticle)

	router.Run("localhost:8080")
}
