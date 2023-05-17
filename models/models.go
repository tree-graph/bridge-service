package models

import (
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
)

func SetDB(db *gorm.DB) {
	DB = db
}

// save user's request
type CrossRequest struct {
	Id        int        `json:"id" gorm:"primary_key,autoIncrement"`
	ChainId   int64      `json:"chain_id" binding:"required" gorm:"index,not null"`
	Hash      string     `json:"hash" binding:"required" gorm:"unique,not null"`
	Result    string     `json:"result" gorm:"varchar(255) default ''"`
	Status    int        `json:"status" gorm:""`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

// TableName Database Table Name of this model
func (e *CrossRequest) TableName() string {
	return "cross_requests"
}

const CrossReqId = "CROSS_REQ_ID"

// system configuration
type Config struct {
	Name    string `json:"name" binding:"required" gorm:"primary_key"`
	Content string `json:"content" binding:"required"`
}

// chain config
type Chain struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	Name      string     `json:"hash" binding:"required" gorm:"unique,not null"`
	Rpc       string     `json:"rpc"  binding:"required" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

const HardhatLocalChain = 31337

func GetChain(id int64) (Chain, error) {
	var bean Chain
	err := DB.Where("id=?", id).Take(&bean).Error
	return bean, err
}

const DuplicateEntryError = 1062 // mysql use this code
const TxNotFound = 404
const InvalidHash = 500
const EmptyEvent = 501

// Cross detail for each transaction. Corresponds to CrossEvent in the contract.
type CrossInfo struct {
	Id             int        `json:"id" gorm:"primary_key,autoIncrement"`
	SourceChain    int64      `json:"chain_id" gorm:"index, not null"`
	TxnHash        string     `json:"txn_hash" binding:"required" gorm:"unique,not null"`
	Asset          string     `json:"asset" gorm:"type:char(42) not null"`
	From           string     `json:"from" gorm:"type:char(42) not null"`
	TargetChain    int64      `json:"to_chain_id" gorm:"not null"`
	TargetContract string     `json:"target_contract" gorm:"type:char(66) not null"`
	UserNonce      int64      `json:"user_nonce" gorm:"not null"`
	BlockNumber    uint64     `json:"block_number" gorm:"not null"`
	BlockTime      *time.Time `json:"block_time,string,omitempty" gorm:"not null"`
}

// There is batch style crossing request, this table keeps track of each token for each request.
type CrossItem struct {
	Id          int    `json:"id" gorm:"primary_key,autoIncrement"`
	CrossInfoId int    `json:"cross_info_id" gorm:"not null"`
	TokenId     string `json:"token_id" binding:"required" gorm:"type:varchar(66)"`
	Amount      string `json:"amount"  binding:"required" gorm:"type:varchar(66)"`
	Uri         string `json:"uri"  binding:"required" gorm:"type:longtext"`
}
