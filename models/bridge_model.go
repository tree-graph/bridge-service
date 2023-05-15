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
	Id        int        `json:"id" gorm:"primary_key"`
	ChainId   int64      `json:"chain_id" gorm:"index"`
	Hash      string     `json:"hash" binding:"required" gorm:"unique"`
	Result    string     `json:"result" binding:"required" gorm:"varchar(255) default ''"`
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

func (e *Config) TableName() string {
	return "config"
}

// chain config
type Chain struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	Name      string     `json:"hash" binding:"required" gorm:"unique"`
	Rpc       string     `json:"rpc"  binding:"required" gorm:"unique"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

func GetChain(id int64) (Chain, error) {
	var bean Chain
	err := DB.Where("id=?", id).Take(&bean).Error
	return bean, err
}

const InvalidHash = 500
const EmptyEvent = 501

// cross detail for each transaction
type CrossInfo struct {
	Id             int    `json:"id" gorm:"primary_key"`
	SourceChain    int64  `json:"chain_id" gorm:"index"`
	Hash           string `json:"hash" binding:"required" gorm:"unique"`
	Asset          string `json:"asset" gorm:"type:char(42)"`
	From           string `json:"from" gorm:"type:char(42)"`
	TargetChain    int64  `json:"to_chain_id" gorm:""`
	TargetContract string `json:"target_contract" gorm:"type:char(66)"`
	UserNonce      int64  `json:"user_nonce" gorm:""`
	BlockNumber    uint64 `json:"block_number"`
	BlockTime      *time.Time `json:"block_time,string,omitempty"`
	//UpdatedAt      *time.Time `json:"updated_at,string,omitempty"`
}
type CrossItem struct {
	Id          int    `json:"id" gorm:"primary_key"`
	CrossInfoId int    `json:"cross_info_id" gorm:""`
	TokenId     string `json:"token_id" binding:"required" gorm:"type:varchar(66)"`
	Amount      string `json:"amount"  binding:"required" gorm:"type:varchar(66)"`
	Uri         string `json:"uri"  binding:"required" gorm:"type:longtext"`
	//CreatedAt   *time.Time `json:"created_at,string,omitempty"`
	//UpdatedAt   *time.Time `json:"updated_at,string,omitempty"`
}
