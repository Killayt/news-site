package server

import (
	"log"

	"github.com/Killayt/news-site/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// var conf *config.Conf
	// conf, _ = config.LoadConf("internal/config/config.yaml")

	r.GET("/", handlers.IndexHandler)
	r.GET("/news{id:1-99}", handlers.NewsHandler)

	gin.SetMode(gin.ReleaseMode)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
