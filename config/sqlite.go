package config

import (
	"os"

	"github.com/galvao0/fitplan/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/fitplan.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("databse file not found, creating...")

		if err := os.MkdirAll("./db", os.ModePerm); err != nil {
			logger.Errorf("sqlite creating directory error: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("sqlite creating file error %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Workout{}, &schemas.Exercise{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
