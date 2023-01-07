package main

import (
	"fmt"

	"github.com/Killayt/news-site/internal/config"
	"github.com/Killayt/news-site/internal/server"
)

func main() {
	conf, _ := config.LoadConf("internal/config/config.yaml")
	fmt.Println(conf.Port, "\t", conf.Database.DBLogin, "\t", conf.Database.DBPassword)
	server.Run()
}
