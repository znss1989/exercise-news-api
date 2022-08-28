package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
