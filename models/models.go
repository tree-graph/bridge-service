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

const CfxChain = "cfx"
const EvmChain = "evm"

// chain config
type Chain struct {
	Id int64 `json:"id" gorm:"primary_key"`
	// chain id for consortium chain may be duplicate. It's 1029 for now(2023.6.20).
	ChainId   int64  `json:"chain_id" binding:"required" gorm:"type:int not null"`
	Name      string `json:"hash" binding:"required" gorm:"unique,not null"`
	VaultAddr string `json:"vault_addr" binding:"required" gorm:"type:varchar(66) not null"`
	Rpc       string `json:"rpc"  binding:"required" gorm:"type:varchar(256) not null"`
	// possible values: cfx, evm
	ChainType  string     `json:"chain_type"  binding:"required" gorm:"type:varchar(16) not null default 'cfx'"`
	DelayBlock int        `json:"delay_block"  binding:"required" gorm:"not null"`
	CreatedAt  *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,string,omitempty"`
}

func GetChain(id int64) (Chain, error) {
	var bean Chain
	err := DB.Where("id=?", id).Take(&bean).Error
	return bean, err
}

// Store confidential information in separate data sheets
type Secret struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	ChainId   int64      `json:"chain_id" gorm:"unique"`
	Private   string     `json:"private" binding:"required" gorm:"char(66),not null"`
	Address   string     `json:"address" binding:"" gorm:"type:varchar(66)"`
	Comment   string     `json:"comment" binding:"" gorm:"type:varchar(1024)"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

func GetSecret(id int64) (Secret, error) {
	var bean Secret
	err := DB.Where("chain_id=?", id).Take(&bean).Error
	return bean, err
}

// track fetching event cursor of each chain.
type ChainCursor struct {
	Id          int64      `json:"id" gorm:"primary_key"`
	Block       int64      `json:"block" gorm:"not null"`
	LatestBlock int64      `json:"latest_block" gorm:""`
	UpdatedAt   *time.Time `json:"updated_at,string,omitempty"`
}

const TxNotFound = 404
const InvalidHash = 500
const EmptyEvent = 501

// Cross detail for each transaction. Corresponds to CrossEvent in the contract.
type CrossInfo struct {
	Id             int        `json:"id" gorm:"primary_key,autoIncrement,index:idx_target_chain,priority:2"`
	SourceChain    int64      `json:"chain_id" gorm:"index, not null"`
	TxnHash        string     `json:"txn_hash" binding:"required" gorm:"unique,not null"`
	Asset          string     `json:"asset" gorm:"type:char(42) not null"`
	From           string     `json:"from" gorm:"type:char(42) not null"`
	TargetChain    int64      `json:"to_chain_id" gorm:"not null,index:idx_target_chain,priority:1"`
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

const ClaimStepSendingTx = "sending_tx"
const ClaimStepWaitingTx = "waiting_tx"
const ClaimStepError = "error"

// Keep undergoing claiming records
type ClaimPool struct {
	Id             int        `json:"id" gorm:"primary_key,autoIncrement"`
	CrossInfoId    int        `json:"cross_info_id" gorm:"not null"`
	TargetChain    int64      `json:"target_chain_id" gorm:"not null"`
	TargetContract string     `json:"target_contract" gorm:"type:char(66) not null"`
	TxnHash        string     `json:"txn_hash" binding:"required" gorm:"unique,not null"`
	From           string     `json:"from" gorm:"type:char(42) not null"`
	Nonce          uint64     `json:"user_nonce" gorm:"not null"`
	Step           string     `json:"step" gorm:"type:varchar(16) not null"`
	Status         *int64     `json:"status" gorm:""`
	CreatedAt      *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,string,omitempty"`
}

// Keep claiming history, include all transactions had been sent.
type ClaimHistory struct {
	Id             int        `json:"id" gorm:"primary_key,autoIncrement"`
	CrossInfoId    int        `json:"cross_info_id" gorm:"index,not null"`
	TargetChain    int64      `json:"target_chain_id" gorm:"not null"`
	TargetContract string     `json:"target_contract" gorm:"type:char(66) not null"`
	TxnHash        string     `json:"txn_hash" binding:"required" gorm:"unique,not null"`
	From           string     `json:"from" gorm:"type:char(42) not null"`
	Nonce          uint64     `json:"user_nonce" gorm:"not null"`
	Comment        string     `json:"comment" binding:"" gorm:""`
	Status         uint64     `json:"status" gorm:""`
	CreatedAt      *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,string,omitempty"`
}

// Each chain has a claim cursor, indicates which cross_info had been processed.
type ClaimCursor struct {
	TargetChain int64      `json:"target_chain_id" gorm:"unique"`
	CrossInfoId int64      `json:"cross_info_id" gorm:"not null"`
	UpdatedAt   *time.Time `json:"updated_at,string,omitempty"`
}
