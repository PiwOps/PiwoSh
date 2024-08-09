package datastorer

import (
	"fmt"
	"os"

	"smuggr.xyz/piwosh/common/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Logger = logger.NewCustomLogger("data")

var DB *gorm.DB

func getDSN() string {
	return "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=" + os.Getenv("DB_SSL_MODE") +
		" TimeZone=" + os.Getenv("DB_TIMEZONE")
}

func Initialize() {
	Logger.Info(logger.MsgInitializing)

	db, err := gorm.Open(postgres.Open(getDSN()))
	if err != nil {
		Logger.Log(logger.ErrInitializing.Format(err.Error()))
	}

	DB = db
	
	if err := DB.AutoMigrate(&User{}, &CommandsConfig{}); err != nil {
		Logger.Log(logger.ErrInitializing.Format(err.Error()))
	}

	Logger.Success(fmt.Sprintf(
		"connected to database %s on: %s:%s as user: %s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
	))
}

func Cleanup() {
	Logger.Log(logger.MsgCleaningUp)

	sqlDB, err := DB.DB()
	if err != nil {
		Logger.Log(logger.ErrCleaningUp.Format(err.Error()))
		return
	}

	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			Logger.Log(logger.ErrCleaningUp.Format(err.Error()))
			return
		}
	}
}
