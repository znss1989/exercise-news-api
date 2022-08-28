package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getChannels(c *gin.Context) {
	channels, err := queryChannels()
	if err != nil {
		fmt.Println(err.Error())
		return
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

func getArticles(c *gin.Context) {
	channelID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	articles, err := queryArticles(channelID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, articles)
}

func addArticle(c *gin.Context) {
	channelID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	var atc Article
	if err := c.BindJSON(&atc); err != nil {
		return
	}
	id, err := insertArticle(channelID, atc)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusAccepted, id)
}
