package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// healthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func healthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"msg": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

func getChannels(c *gin.Context) {
	channels, err := queryChannels()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(500)
	}
	c.JSON(http.StatusOK, channels)
}

func addChannel(c *gin.Context) {
	var chn Channel
	if err := c.BindJSON(&chn); err != nil {
		c.AbortWithStatus(400)
		return
	}
	id, err := insertChannel(chn)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(500)
		return
	}
	c.JSON(http.StatusOK, id)
}

func getArticles(c *gin.Context) {
	channelID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(400)
		return
	}
	var ftr Filter
	if err := c.BindJSON(&ftr); err != nil {
		fmt.Println((err.Error()))
		ftr = Filter{
			Lo: 0,
			Hi: -1,
		}
	}
	fmt.Printf("Filter: %v\n", ftr)
	var articles []Article
	if ftr.Lo > ftr.Hi {
		articles, err = queryArticles(channelID)
		fmt.Printf("get articles in channel %v\n", channelID)
	} else {
		articles, err = queryArticlesFiltering(channelID, ftr)
		fmt.Printf("get articles in channel with filtering %v, %v\n", channelID, ftr)

	}

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(500)
		return
	}
	c.JSON(http.StatusOK, articles)
}

func addArticle(c *gin.Context) {
	channelID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	var atc Article
	if err := c.BindJSON(&atc); err != nil {
		c.AbortWithStatus(400)
		return
	}
	id, err := insertArticle(channelID, atc)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(500)
	}

	// async update article
	go updateArticleWordCount(id)

	c.JSON(http.StatusAccepted, id)
}
