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
	cfg          *Config
)

type Config struct {
	rest.RestConf
	Mysql MysqlConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
	OSS   OSS         `yaml:"oss"`
}

type MysqlConfig struct {
	UserName     string `yaml:"user_name"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         uint64 `yaml:"port"`
	DatabaseName string `yaml:"database_name"`
}

type RedisConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
	MaxRetries   int    `yaml:"max_retries"`
}

type NacosServerConfig struct {
	Servers     []Server `yaml:"servers"`
	Group       string   `yaml:"group"`
	DataID      string   `yaml:"data_id"`
	ExtDataIDs  []string `yaml:"ext_data_ids"`
	NameSpaceID string   `yaml:"namespace_id"`
	RPCGroup    string   `yaml:"rpc_group"`
}

type OSS struct {
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	Endpoint        string `yaml:"endpoint"`
	BucketUrl       string `yaml:"bucket_url"`
	BucketName      string `yaml:"bucket_name"`
}

type Server struct {
	Addr string `yaml:"addr"`
	Port uint64 `yaml:"port"`
}

func init() {
	env = os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
}

func GetCfg() *Config {
	return cfg
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

	err = yaml.Unmarshal(config, v)
	if err != nil {
		log.Fatalf("load from yml bytes error: %+v", err)
	}

	cfg = v.(*Config)

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
	out, _ := yaml.Marshal(resultMap[env])
	err = yaml.Unmarshal(out, v)
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

func (conf *NacosServerConfig) GetConfig() ([]byte, error) {
	configMap := make(map[interface{}]interface{})

	mainConfig, err := configClient.GetConfig(vo.ConfigParam{DataId: conf.DataID, Group: conf.Group})
	if err != nil {
		return []byte{}, err
	}
	mainMap, err := utils.UnmarshalYamlToMap(mainConfig)
	if err != nil {
		return []byte{}, err
	}

	extMap := make(map[interface{}]interface{})
	for _, dataID := range conf.ExtDataIDs {
		extConfig, err := configClient.GetConfig(vo.ConfigParam{DataId: dataID, Group: conf.Group})
		if err != nil {
			return []byte{}, err
		}

		tmpExtMap, err := utils.UnmarshalYamlToMap(extConfig)
		if err != nil {
			return []byte{}, err
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
