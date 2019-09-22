package main

import (
	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	router.Static("/css", "public/css")
	router.Static("/images", "public/images")
	router.Static("/js", "public/js")
	router.StaticFile("/", "public/index.html")
	router.Run()
}