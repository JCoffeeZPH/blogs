package main

import (
	"blogs/app/core/api/internal/config"
	"blogs/app/core/api/internal/handler"
	"blogs/app/core/api/internal/svc"
	commonConfig "blogs/common/config"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

var (
	nacosConfigFile = "app/core/api/etc/nacos.yaml"
	configFile      = "app/core/api/etc/config.yaml"
)

func main() {
	flag.Parse()

	var c config.Config
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
