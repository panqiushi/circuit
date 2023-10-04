package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var instance *gorm.DB

func DB() *gorm.DB {
	if instance != nil {
		return instance
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	defaultConfig := viper.GetStringMapString("database.connections." + viper.GetString("database.default"))
	driver := defaultConfig["driver"]
	config := "host=" + defaultConfig["host"] + " port=" + defaultConfig["port"] + " user=" + defaultConfig["user"] + " dbname=" + defaultConfig["dbname"] + " password=" + defaultConfig["password"] + " sslmode=" + defaultConfig["sslmode"]
	db, err := gorm.Open(driver, config)
	if err != nil {
		panic(err)
	}
	instance = db
	return instance
}
