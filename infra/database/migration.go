package database

import (
	"github.com/tree-graph/bridge-service/models"
)

//Add list of model add for migrations
//var migrationModels = []interface{}{&ex_models.Example{}, &model.Example{}, &model.Address{})}
var migrationModels = []interface{}{&models.Example{}, &models.User{}, &models.CreditCard{},
	&models.CrossRequest{},
	&models.Config{},
	&models.Chain{},
	&models.CrossInfo{},
	&models.CrossItem{},
}
