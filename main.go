package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	timestamp := time.Now().Unix()
	isDev := os.Getenv("APP_ENV") == "dev"
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.tmpl.html",
			gin.H{
				"isDev":            isDev,
				"initialTimestamp": timestamp,
			})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "pong",
			"timestamp": timestamp,
		})
	})

	router.Run(":" + port)
}
