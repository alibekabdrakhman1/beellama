package main

import (
	"fmt"
	"github.com/alibekabdrakhman1/beellama/app/internal/applicator"
	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"github.com/alibekabdrakhman1/beellama/app/internal/logger"
	"github.com/spf13/viper"
)

// @title          tinyllama
// @version        1.0

// @host           localhost:8080
// @BasePath       /api
func main() {
	cfg, err := loadConfig("/config")
	if err != nil {
		panic(fmt.Sprintf("ошибка при загрузке конфига: %v", err))
	}
	l, err := logger.InitLogger(cfg.Logging.LogDir)
	if err != nil {
		panic(err)
	}

	app := applicator.New(l, &cfg)
	err = app.Run()
	fmt.Println(err)
}

func loadConfig(path string) (config config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
