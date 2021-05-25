package db

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"hawk.wie.gg/models"
)

var DB *gorm.DB

func CreateDB(path string) {
	var home string
	var err error

	if path == "" {
		home, err = homedir.Dir()
		cobra.CheckErr(err)
	} else {
		home = path
	}

	DB, err = gorm.Open(sqlite.Open(home+"/db.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		cobra.CheckErr(err)
	}

	if DB == nil {
		cobra.CompErrorln("Couldnt open DB")
	}

	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Entry{})
	DB.AutoMigrate(&models.Setting{})
}
