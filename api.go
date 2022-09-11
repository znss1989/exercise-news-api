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

// getChannels godoc
// @Summary Get all channels
// @Description Get a list of all the channels for news articles.
// @Tags channels
// @Produce json
// @Success 200 {array} Channel
// @Router /channel [get]
func getChannels(c *gin.Context) {
	channels, err := queryChannels()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(500)
	}
	c.JSON(http.StatusOK, channels)
}

// addChannel godoc
// @Summary Add a new channel
// @Description Add a new channel in records for news articles
// @Tags channels
// @Param title body ChannelRequest true  "Title in JSON"
// @Produce json
// @Success 200 {integer} integer
// @Router /channel [post]
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

// getArticles godoc
// @Summary Get articles of a channel
// @Description Get a list of all articles under a channel. Without query parameters, this will return all the articles of the channel.
// @Description If the lower (lo) and upper bound (hi) bound of word count provided, the articles are filterd by word count accordingly before returned.
// @Tags articles
// @Produce json
// @Param id path int true "Channel ID"
// @Param lo query int false "[Optional] lower bound of word count"
// @Param hi query int false "[Optional] upper bound of word count"
// @Success 200 {array} Article
// @Router /channel/{id} [get]
func getArticles(c *gin.Context) {
	channelID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(400)
		return
	}
	lo, err1 := strconv.ParseInt(c.Query("lo"), 10, 64)
	hi, err2 := strconv.ParseInt(c.Query("hi"), 10, 64)

	var ftr Filter
	if err1 != nil || err2 != nil {
		ftr = Filter{
			Lo: 0,
			Hi: -1,
		}
	} else {
		ftr = Filter{
			Lo: lo,
			Hi: hi,
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

// addArticle godoc
// @Summary Add a new article
// @Description Add a new article of a channel
// @Tags articles
// @Accept json
// @Produce json
// @Param id path int true "Channel ID"
// @Param url body ArticleRequest true  "Url in JSON"
// @Success 200 {integer} integer
// @Router /channel/{id} [post]
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
