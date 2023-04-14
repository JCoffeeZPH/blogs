package main

import (
	commonConfig "blogs/common/config"
	"flag"
	"fmt"
	"log"

	"blogs/app/admin/api/internal/handler"
	"blogs/app/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var (
	nacosConfigFile = "app/admin/api/etc/nacos.yaml"
	configFile      = "app/admin/api/etc/config.yaml"
)

func main() {
	flag.Parse()

	var c commonConfig.Config
	nacosConfig := commonConfig.MustLoad(nacosConfigFile, configFile, &c)
	err := nacosConfig.Listen(&c)
	if err != nil {
		log.Fatalf("nacos config listen failed, err: %+v", err)
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c, nacosConfig)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

type T struct {
	Data struct {
		ViewsCount   int `json:"viewsCount"`
		MessageCount int `json:"messageCount"`
		UserCount    int `json:"userCount"`
		ArticleCount int `json:"articleCount"`
		CategoryDTOs []struct {
			Id           int    `json:"id"`
			CategoryName string `json:"categoryName"`
			ArticleCount int    `json:"articleCount"`
		} `json:"categoryDTOs"`
		TagDTOs []struct {
			Id      int         `json:"id"`
			TagName string      `json:"tagName"`
			Count   interface{} `json:"count"`
		} `json:"tagDTOs"`
		ArticleStatisticsDTOs []struct {
			Date  string `json:"date"`
			Count int    `json:"count"`
		} `json:"articleStatisticsDTOs"`
		UniqueViewDTOs []struct {
			Day        string `json:"day"`
			ViewsCount int    `json:"viewsCount"`
		} `json:"uniqueViewDTOs"`
		ArticleRankDTOs []struct {
			ArticleTitle string `json:"articleTitle"`
			ViewsCount   int    `json:"viewsCount"`
		} `json:"articleRankDTOs"`
	} `json:"data"`
}
