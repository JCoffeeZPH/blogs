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
