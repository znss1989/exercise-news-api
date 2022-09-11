package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHealthCheck(t *testing.T) {
	r := SetUpRouter()
	r.GET("/", healthCheck)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	mockResponse := `{"msg":"Server is up and running"}`
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddArticleFail(t *testing.T) {
	r := SetUpRouter()
	r.POST("/channel/:id", addArticle)
	id := "7"
	req, _ := http.NewRequest("POST", "/channel/"+id, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
