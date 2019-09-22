package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":3000"
	}
	return ":" + port
}

func main () {
	fmt.Println(getPort())
	router := gin.Default()
	router.Static("/css", "public/css")
	router.Static("/images", "public/images")
	router.Static("/js", "public/js")
	router.StaticFile("/", "public/index.html")
	router.Run(getPort())
}