package database

import (
	"errors"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/models"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	DBErr error
)

func Init() {
	config := mysql.MustNewConfigFromViper()
	db := config.MustOpenOrCreate()
	err := db.AutoMigrate(migrationModels...)
	if err != nil {
		logrus.Fatalln("migrate database fail:", err.Error())
	}

	DB = db
	models.SetDB(DB)

	// auto add hardhat local node
	if _, err := models.GetChain(models.HardhatLocalChain); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&models.Chain{
				Id:   models.HardhatLocalChain,
				Name: "Hardhat",
				Rpc:  "http://127.0.0.1:8545",
			})
		}
	}

}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}
