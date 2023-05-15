package database

import (
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/tree-graph/bridge-service/models"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	DBErr error
)

func Init() {
	config := mysql.MustNewConfigFromViper()
	db := config.MustOpenOrCreate(migrationModels...)

	DB = db
	models.SetDB(DB)
}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}
