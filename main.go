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

// @contact.name Lei Wu
// @contact.email znss1989@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
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

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", healthCheck)
		v1.GET("/channel", getChannels)
		v1.POST("/channel", addChannel)
		v1.GET("/channel/:id", getArticles)
		v1.POST("/channel/:id", addArticle)
	}

	// policy creation for HTML sanitizer
	p = bluemonday.StrictPolicy()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("0.0.0.0:8080")
}
