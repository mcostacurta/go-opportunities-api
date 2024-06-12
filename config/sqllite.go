package config

import (
	"os"

	"github.com/mcostacurta/gooportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLLite() (*gorm.DB, error) {
	logger := GetLogger("sqllite")

	dbPath := "./db/main.db"

	//Check  if db exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("File not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("folder creating error: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("file creating error: %v", err)
			return nil, err
		}
		file.Close()
	}

	//create db and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqllite opening error: %v", err)
		return nil, err
	}

	// migrate schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqllite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
