package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/microcosm-cc/bluemonday"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/znss1989/mm-news-apis/docs"
)

var db *sql.DB
var p *bluemonday.Policy

// @title MM News APIs
// @version 1.0
// @description This is web server for channel news

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
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
	router.GET("/", healthCheck)
	router.GET("/channel", getChannels)
	router.POST("/channel", addChannel)
	router.GET("/channel/:id", getArticles)
	router.POST("/channel/:id", addArticle)

	// policy creation for HTML sanitizer
	p = bluemonday.StrictPolicy()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080")
}
