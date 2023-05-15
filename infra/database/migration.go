package database

import (
	"github.com/tree-graph/bridge-service/models"
)

var migrationModels = []interface{}{
	&models.CrossRequest{},
	&models.Config{},
	&models.Chain{},
	&models.CrossInfo{},
	&models.CrossItem{},
}
