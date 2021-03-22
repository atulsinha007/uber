package config

import (
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

const (
	EnvKey      = "ENV"
	configFile  = "system"
	configType  = "json"
	defaultPath = "$GOPATH/src/github.com/atulsinha007/uber/config/"
)

var environments = map[string]bool{
	"local": true,
}

func SetUpEnv() {
	err := viper.BindEnv(EnvKey)
	if err != nil {
		log.L.Fatal("ENV load failed")
	}

	envName := strings.ToLower(viper.GetString(EnvKey))
	log.L.With(zap.String("env", envName)).Info("env set")
	if _, ok := environments[envName]; !ok {
		log.L.With(zap.String("env", envName)).Fatal("invalid ENV variable")
	}

	err = loadConfigFile(envName)

	if err != nil {
		log.L.Fatal("Fatal error config file")
	}
}

func loadConfigFile(envName string) error {
	basePath := defaultPath
	basePath = filepath.Join(basePath, envName)

	viper.AddConfigPath(basePath)
	viper.SetConfigName(configFile)
	viper.SetConfigType(configType)
	return viper.ReadInConfig()
}
