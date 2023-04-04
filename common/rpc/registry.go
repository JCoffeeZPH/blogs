package rpc

import (
	"blogs/common/config"
	"blogs/common/constants"
	"fmt"
	"github.com/JCoffeeZPH/zero-contrib/zrpc/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"os"
	"strings"
)

var (
	env = "dev"
)

func init() {
	environment := os.Getenv("ENV")
	if len(environment) != 0 {
		env = strings.ToLower(environment)
	}
}

func MustRegister(nacosConfig *config.NacosServerConfig, rpcConfig *zrpc.RpcServerConf) {
	serverConfigs, clientConfig := getServerClientConfigs(nacosConfig)

	opts := nacos.NewNacosConfig(rpcConfig.Name, rpcConfig.ListenOn, serverConfigs, clientConfig)
	opts.Group = nacosConfig.Group

	if err := nacos.RegisterService(opts); err != nil {
		log.Fatalf("register service failed: %+v", err)
	}
}

func NewZrpcClient(serverName, clientName string, nacosConfig *config.NacosServerConfig) zrpc.Client {
	return zrpc.MustNewClient(zrpc.RpcClientConf{
		//Target: fmt.Sprintf("nacos://%s:%d/%s?timeout=%s&namespace_id=%s&group_name=%s&app_name=%s", "10.58.233.164", 8841, serverName, "3s", constants.GetNamespace(env).ToString(), "rpc-dev", clientName),
		//Target: fmt.Sprintf("nacos://%s:%d/%s?timeout=%s&namespace_id=%s&group_name=%s&app_name=%s",
		//	instance.Ip, instance.Port, serverName, "3s", constants.GetNamespace(env).ToString(), nacosConfig.Group, clientName),

		Target: fmt.Sprintf("nacos://%s:%d/%s?timeout=%s&namespace_id=%s&group_name=%s&app_name=%s", nacosConfig.Servers[0].Addr, nacosConfig.Servers[0].Port, serverName, "3s", nacosConfig.NameSpaceID, nacosConfig.BlogRPCGroup, clientName),
	})
}

func getServerClientConfigs(nacosConfig *config.NacosServerConfig) ([]constant.ServerConfig, *constant.ClientConfig) {
	serverConfigs := make([]constant.ServerConfig, 0)
	for _, conf := range nacosConfig.Servers {
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			IpAddr: conf.Addr, Port: conf.Port,
		})
	}

	clientConfig := &constant.ClientConfig{
		NamespaceId:         constants.GetNamespace(env).ToString(),
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
		LogDir:              "/log",
	}
	return serverConfigs, clientConfig
}
