package config

import (
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

var V *viper.Viper

const (
	EnvKey      = "ENV"
	configFile  = "system"
	configType  = "json"
	defaultPath = "$GOPATH/src/github.com/atulsinha007/uber/config/"
)

var environments = map[string]bool{
	"local": true,
}

func init() {
	V = viper.New()
	err := V.BindEnv(EnvKey)
	if err != nil {
		log.L.Fatal("ENV load failed")
	}

	envName := strings.ToLower(V.GetString(EnvKey))
	log.L.With(zap.String("env", envName)).Info("env set")
	if _, ok := environments[envName]; !ok {
		log.L.With(zap.String("env", envName)).Fatal("invalid ENV variable")
	}

	V.AutomaticEnv()

	err = loadConfigFile(envName)

	if err != nil {
		log.L.Fatal("Fatal error config file")
	}
}

func loadConfigFile(envName string) error {
	basePath := defaultPath
	basePath = filepath.Join(basePath, envName)

	V.AddConfigPath(basePath)
	V.SetConfigName(configFile)
	V.SetConfigType(configType)
	return V.ReadInConfig()
}
