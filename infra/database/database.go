package database

import (
	"github.com/akmamun/gin-boilerplate-examples/config"
	infraLogger "github.com/akmamun/gin-boilerplate-examples/infra/logger"
	"github.com/akmamun/gin-boilerplate-examples/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

func Init() {
	masterDSN, replicaDSN := config.DbConfiguration()

	if err := DBConnection(masterDSN, replicaDSN); err != nil {
		infraLogger.Infof("masterDSN %s %s\n", masterDSN, replicaDSN)
		infraLogger.Fatalf("database DbConnection error: %s", err)
	}
}

// DBConnection create database connection
func DBConnection(masterDSN, replicaDSN string) error {
	var db = DB

	logMode := viper.GetBool("DB_LOG_MODE")
	replicaDB := viper.GetBool("REPLICA_DB_ENABLE")

	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	db, err = gorm.Open(mysql.Open(masterDSN), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})
	if replicaDB {
		db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				mysql.Open(replicaDSN),
			},
			Policy: dbresolver.RandomPolicy{},
		}))
	}

	if err != nil {
		DBErr = err
		log.Println("Db connection error")
		return err
	}

	err = db.AutoMigrate(migrationModels...)

	if err != nil {
		return err
	}
	DB = db
	models.SetDB(DB)

	return nil
}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}

// GetDBError connection error
func GetDBError() error {
	return DBErr
}
