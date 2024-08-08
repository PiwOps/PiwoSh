// configurator.go
package configurator

import (
	"os"

	"smuggr.xyz/piwosh/common/logger"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Logger = logger.NewCustomLogger("conf")
var Config GlobalConfig

func loadEnv() error {
	Logger.Info(logger.MsgLoadingResource.Format("*.env", logger.ResourceEnv))

	err := godotenv.Load()
	if err != nil {
		return logger.ErrReadingResource.Format("", logger.ResourceEnv)
	}

	return nil
}

func loadConfig(config *GlobalConfig) error {
	Logger.Info(logger.MsgLoadingResource.Format(os.Getenv("CONFIG_PATH") + "*." + os.Getenv("CONFIG_TYPE"), logger.ResourceConfig))

	if err := viper.ReadInConfig(); err != nil {
		return logger.ErrReadingResource.Format("", logger.ResourceConfig)
	}

	err := viper.Unmarshal(config)
	if err != nil {
		return logger.ErrFormattingResource.Format(logger.ResourceConfig)
	}

	return nil
}

func Initialize() {
	Logger.Info(logger.MsgInitializing)

	if err := loadEnv(); err != nil {
		Logger.Error(err)
	}

	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))
	viper.SetConfigType(os.Getenv("CONFIG_TYPE"))

	if err := loadConfig(&Config); err != nil {
		Logger.Error(err)
	}
}
