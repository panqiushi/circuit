package db

import (
	"log"
	"os"
	"strings"

	"circuit.io/circuit/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	if driver == "postgres" {
		connection := viper.GetStringMapString("database.connection." + driver)
		config := "host=" + connection["host"] + " port=" + connection["port"] + " user=" + connection["username"] + " dbname=" + connection["database"] + " password=" + connection["password"] + " sslmode=" + connection["sslmode"]
		db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = db
	}

	if driver == "sqlite" {
		database := viper.GetString("database.connection." + driver + ".database")
		db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = db
	}

	if err := instance.AutoMigrate(&models.User{}); err != nil {
		log.Println(err)
	}

	if err := instance.AutoMigrate(&models.Project{}); err != nil {
		log.Println(err)
	}

	if err := instance.AutoMigrate(&models.VersionRelease{}); err != nil {
		log.Println(err)
	}

	return instance
}
