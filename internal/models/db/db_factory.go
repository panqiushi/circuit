package db

import (
	"os"
	"strings"

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

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	if strings.HasSuffix(wd, "server") {
		viper.AddConfigPath("../..")
	} else {
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	driver := viper.GetString("database.default")
	connection := viper.GetStringMapString("database.connection." + driver)
	config := "host=" + connection["host"] + " port=" + connection["port"] + " user=" + connection["username"] + " dbname=" + connection["database"] + " password=" + connection["password"] + " sslmode=" + connection["sslmode"]
	db, err := gorm.Open(driver, config)
	if err != nil {
		panic(err)
	}
	instance = db
	return instance
}
