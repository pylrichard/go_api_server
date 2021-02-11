package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
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
	c.initLog()
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
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) initLog() {
	cfg := log.PassLagerCfg{
		Writers:		viper.GetString("log.writers"),
		LoggerLevel:	viper.GetString("log.logger_level"),
		LoggerFile:		viper.GetString("log.logger_file"),
		LogFormatText:	viper.GetBool("log.log_format_text"),
		RollingPolicy:	viper.GetString("log.rollingPolicy"),
		LogRotateDate:	viper.GetInt("log.log_rotate_date"),
		LogRotateSize:	viper.GetInt("log.log_rotate_size"),
		LogBackupCount:	viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&cfg)
}

// 监听配置文件并热加载
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
