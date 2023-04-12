package config

import (
	"blogs/common/utils"
	"bytes"
	"github.com/mitchellh/mapstructure"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"sync"
)

var (
	nacosOnce    sync.Once
	configClient config_client.IConfigClient
	env          string
)

type Config struct {
	rest.RestConf
	Mysql MysqlConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}

type MysqlConfig struct {
	UserName     string
	Password     string
	Host         string
	Port         uint64
	DatabaseName string
}

type RedisConfig struct {
	Host         string
	Port         int
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
	MaxRetries   int
}

type NacosServerConfig struct {
	Servers     []Server
	Group       string
	DataID      string
	ExtDataIDs  []string
	NameSpaceID string
	RPCGroup    string
}

type Server struct {
	Addr string
	Port uint64
}

func init() {
	env = os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
}

func MustLoad(nacosConfigFilePath, localConfigPath string, v interface{}) *NacosServerConfig {
	serverConfig := &NacosServerConfig{}
	loadLocalConfig(nacosConfigFilePath, serverConfig)

	err := serverConfig.initConfigClient()
	if err != nil {
		log.Fatalf("init config client error: %+v", err)
	}

	loadLocalConfig(localConfigPath, v)

	config, err := serverConfig.GetConfig()
	if err != nil {
		log.Fatalf("get config error: %+v", err)
	}

	err = conf.LoadFromYamlBytes([]byte(config), v)
	if err != nil {
		log.Fatalf("load from yml bytes error: %+v", err)
	}

	return serverConfig
}
func loadLocalConfig(nacosConfigPath string, v interface{}) {
	f, err := os.Open(nacosConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	input, _ := io.ReadAll(f)
	resultMap := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(input, &resultMap); err != nil {
		log.Fatalln(err)
	}

	err = mapstructure.Decode(resultMap[env], v)
	if err != nil {
		log.Fatalln(err)
	}
}

func (conf *NacosServerConfig) initConfigClient() (err error) {
	nacosOnce.Do(func() {
		serverConfigs := make([]constant.ServerConfig, 0)
		for _, config := range conf.Servers {
			serverConfigs = append(serverConfigs, constant.ServerConfig{
				IpAddr: config.Addr, Port: config.Port,
			})
		}

		configClient, err = clients.NewConfigClient(vo.NacosClientParam{
			ClientConfig:  &constant.ClientConfig{TimeoutMs: 5000, NamespaceId: conf.NameSpaceID},
			ServerConfigs: serverConfigs,
		})
	})
	return
}

func (conf *NacosServerConfig) GetConfig() (string, error) {
	configMap := make(map[interface{}]interface{})

	mainConfig, err := configClient.GetConfig(vo.ConfigParam{DataId: conf.DataID, Group: conf.Group})
	if err != nil {
		return "", err
	}
	mainMap, err := utils.UnmarshalYamlToMap(mainConfig)
	if err != nil {
		return "", err
	}

	extMap := make(map[interface{}]interface{})
	for _, dataID := range conf.ExtDataIDs {
		extConfig, err := configClient.GetConfig(vo.ConfigParam{DataId: dataID, Group: conf.Group})
		if err != nil {
			return "", err
		}

		tmpExtMap, err := utils.UnmarshalYamlToMap(extConfig)
		if err != nil {
			return "", err
		}
		extMap = utils.MergeMap(extMap, tmpExtMap)
	}

	configMap = utils.MergeMap(configMap, extMap)
	configMap = utils.MergeMap(configMap, mainMap)

	return utils.MarshalObjectToYamlString(configMap)
}

func (conf *NacosServerConfig) Listen(v interface{}) error {
	return configClient.ListenConfig(vo.ConfigParam{
		DataId: conf.DataID,
		Group:  conf.Group,
		OnChange: func(namespace, group, dataId, data string) {
			err := viper.ReadConfig(bytes.NewBufferString(data))
			if err != nil {
				log.Printf("viper read config failed: %v", err)
			}
			log.Printf("config changed: %+v", v)
			err = viper.Unmarshal(v)
			if err != nil {
				log.Printf("viper unmarshal config failed: %v", err)
			}
		},
	})
}
