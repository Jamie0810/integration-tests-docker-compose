package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Grpc      Grpc   `mapstructure:"grpc"`
	Pubsub    Pubsub `mapstructure:"pubsub"`
	LogLevel  string `mapstructure:"log_level"`
	LogFormat string `mapstructure:"log_format"`
}

type Grpc struct {
	Addr string `mapstructure:"addr"`
}

type Pubsub struct {
	ProjectID           string `mapstructure:"project_id"`
	CredentialsFilePath string `mapstructure:"credentials_file_path"`
	Deposit             string `mapstructure:"topic_deposit"`
}

func InitConfig(configPath string) (config Config, err error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	/* default */
	v.SetDefault("log_level", "INFO")
	v.SetDefault("log_format", "console")

	defaultPath := `./configs`

	if configPath == "" {
		configPath = defaultPath
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AddConfigPath(configPath)

	files, _ := ioutil.ReadDir(configPath)
	index := 0

	for _, file := range files {
		if filepath.Ext("./"+file.Name()) != ".yaml" && filepath.Ext("./"+file.Name()) != ".yml" {
			continue
		}

		v.SetConfigName(file.Name())
		var err error
		if index == 0 {
			err = v.ReadInConfig()
		} else {
			err = v.MergeInConfig()
		}
		if err == nil {
			index++
		}
	}

	if err = v.Unmarshal(&config); err != nil {
		return
	}

	return
}
