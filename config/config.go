package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 配置文件
type Config struct {
	Name string
}

// Init 初始化配置
func Init(name string) error {
	c := Config{
		Name: name,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

// 初始化配置文件
func (c *Config) initConfig() error {
	if c.Name != "" {
		// 指定了配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 没有指定则解析默认配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式
	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 设置环境变量前缀
	viper.SetEnvPrefix("TINY_HTTP_SERVER")
	replacer := stirngs.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监听配置文件并热加载
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
