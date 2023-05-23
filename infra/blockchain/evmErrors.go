package blockchain

import (
	"fmt"
	"math/big"
	"strings"
)

type ErrNotEnoughCashes map[string]ErrNotEnoughCash

func (e ErrNotEnoughCashes) Error() string {
	msgs := []string{}
	for user, err := range e {
		msgs = append(msgs, fmt.Sprintf("%v %v", user, err.Error()))
	}
	return strings.Join(msgs, "\n")
}

type ErrNotEnoughCash struct {
	Need *big.Int
	Got  *big.Int
}

func (e ErrNotEnoughCash) Error() string {
	return fmt.Sprintf("out of balance, need %v, got %v", e.Need, e.Got)
}

var (
	TX_PENDING_REASON_NOT_ENOUGH_CASH_BUT_NOT = "notEnoughCash but balance enough"
	TX_PENDING_REASON_NOT_ENOUGH_CASH         = "notEnoughCash"
	TX_PENDING_REASON_FUTURE_NONCE            = "futureNonce"
	TX_PENDING_REASON_OLD_EPOCH_HEIGHT        = "oldEpochHeight"
	TX_PENDING_REASON_OUT_DATED_STATUS        = "outdatedStatus"
)

type TxRpcError int

const (
	TX_ERR_RPC_OUT_OF_BALANCE TxRpcError = iota
	TX_ERR_RPC_TXPOOL_FULL
	TX_ERR_RPC_INSERT_VALIDATION_FAIL
	TX_ERR_RPC_INSERT_FAIL_QUOTA_UNENOUGH
	TX_ERR_RPC_SAME_NONCE_ALREADY_INSERTED
	TX_ERR_RPC_STALE_NONCE
	TX_ERR_RPC_EXCEED_MAX_GAS
	TX_ERR_RPC_POOL_INCONSISTENT
	TX_ERR_RPC_FAIL_READ_ACCOUNT_CACHE
	TX_ERR_RPC_NONCE_DISTANT
	TX_ERR_RPC_FAILED_IMPORT_TO_DEFFER_POOL
	TX_ERR_RPC_TX_ALREADY_EXIST

	TX_ERR_RPC_TOO_MANY_REQUEST
	TX_ERR_RPC_CONFURA_QPS_EXCEED
	TX_ERR_RPC_CONFURA_DAILYQ_EXCEED
	TX_ERR_RPC_MISSING_BEST_HASH
	TX_ERR_RPC_EPOCH_HEIGHT_OUT_BOUND
	// 包括 vm revert, not enough cash
	TX_ERR_ESTIMATE_OTHERS
	TX_ERR_RPC_ESTIMATE_NOT_ENOUGH_CASH
	TX_ERR_RPC_ESTIMATE_VM_ERR
	TX_ERR_DROPPED
	TX_ERR_RPC_OTHER
)

func (e *TxRpcError) NeedResendWhenSentErr() (needSend, needUpperGas bool) {
	switch *e {
	case TX_ERR_RPC_OUT_OF_BALANCE:
		return false, false
	case TX_ERR_RPC_TXPOOL_FULL:
		return true, false
	case TX_ERR_RPC_INSERT_VALIDATION_FAIL:
		return false, false
	case TX_ERR_RPC_INSERT_FAIL_QUOTA_UNENOUGH:
		return false, false
	case TX_ERR_RPC_SAME_NONCE_ALREADY_INSERTED:
		return true, true
	case TX_ERR_RPC_STALE_NONCE:
		return false, false
	case TX_ERR_RPC_EXCEED_MAX_GAS:
		return false, false
	case TX_ERR_RPC_POOL_INCONSISTENT:
		return true, false
	case TX_ERR_RPC_FAIL_READ_ACCOUNT_CACHE:
		return true, false
	case TX_ERR_RPC_NONCE_DISTANT:
		return false, false
	case TX_ERR_RPC_FAILED_IMPORT_TO_DEFFER_POOL:
		return true, false
	case TX_ERR_RPC_TX_ALREADY_EXIST:
		return false, false
	case TX_ERR_RPC_TOO_MANY_REQUEST:
		fallthrough
	case TX_ERR_RPC_CONFURA_QPS_EXCEED:
		fallthrough
	case TX_ERR_RPC_CONFURA_DAILYQ_EXCEED:
		return true, false
	case TX_ERR_RPC_EPOCH_HEIGHT_OUT_BOUND: //EpochHeightOutOfBound
		return false, false
	default:
		return false, false
	}
}

// "txpool is full"
// Transaction Pool is full
// failed to insert tx into pool (validation failed), hash = {:?}, error = {:?}
// failed to insert tx into pool (quota not enough), hash = {:?}
// transaction gas {} exceeds the maximum value {:?}, the half of pivot block gas limit
// "transaction gas price {} less than the minimum value {}"
// "Tx with same nonce already inserted. To replace it, you need to specify a gas price > {}"
// ready_account_pool and deferred_pool are inconsistent! ready_tx={:?} first_pending={:?}
// "Failed to read account_cache from storage: {}"
// "Transaction {:?} is discarded due to in too distant future",
// "Transaction {:?} is discarded due to a too stale nonce",
// "Transaction {:?} is discarded due to out of balance, needs {:?} but account balance is {:?}",
// "Failed imported to deferred pool
// "tx already exist"

func GetRpcErrorType(errorStr string) TxRpcError {
	if strings.Contains(errorStr, "txpool is full") {
		return TX_ERR_RPC_TXPOOL_FULL
	}
	if strings.Contains(errorStr, "Transaction Pool is full") {
		return TX_ERR_RPC_TXPOOL_FULL
	}
	if strings.Contains(errorStr, "failed to insert tx into pool (validation failed)") {
		return TX_ERR_RPC_INSERT_VALIDATION_FAIL
	}
	if strings.Contains(errorStr, "failed to insert tx into pool (quota not enough)") {
		return TX_ERR_RPC_INSERT_FAIL_QUOTA_UNENOUGH
	}
	if strings.Contains(errorStr, "exceeds the maximum value") {
		return TX_ERR_RPC_EXCEED_MAX_GAS
	}
	if strings.Contains(errorStr, "Tx with same nonce already inserted") {
		return TX_ERR_RPC_SAME_NONCE_ALREADY_INSERTED
	}
	if strings.Contains(errorStr, "ready_account_pool and deferred_pool are inconsistent!") {
		return TX_ERR_RPC_POOL_INCONSISTENT
	}
	if strings.Contains(errorStr, "Failed to read account_cache from storage") {
		return TX_ERR_RPC_FAIL_READ_ACCOUNT_CACHE
	}
	if strings.Contains(errorStr, "discarded due to in too distant future") {
		return TX_ERR_RPC_NONCE_DISTANT
	}
	if strings.Contains(errorStr, "discarded due to a too stale nonce") {
		return TX_ERR_RPC_STALE_NONCE
	}
	if strings.Contains(errorStr, "discarded due to out of balance") {
		return TX_ERR_RPC_OUT_OF_BALANCE
	}
	if strings.Contains(errorStr, "NotEnoughCash") {
		return TX_ERR_RPC_ESTIMATE_NOT_ENOUGH_CASH
	}
	if strings.Contains(errorStr, "block_number is missing for best_hash") {
		return TX_ERR_RPC_MISSING_BEST_HASH
	}
	if strings.Contains(errorStr, "Failed imported to deferred pool") {
		return TX_ERR_RPC_FAILED_IMPORT_TO_DEFFER_POOL
	}
	if strings.Contains(errorStr, "tx already exist") {
		return TX_ERR_RPC_TX_ALREADY_EXIST
	}
	if strings.Contains(errorStr, "Vm reverted") || strings.Contains(errorStr, "VmError") || strings.Contains(errorStr, "Transaction reverted") {
		return TX_ERR_RPC_ESTIMATE_VM_ERR
	}
	if strings.Contains(errorStr, "too many requests") {
		return TX_ERR_RPC_TOO_MANY_REQUEST
	}
	if strings.Contains(errorStr, "allowed qps exceeded") {
		return TX_ERR_RPC_CONFURA_QPS_EXCEED
	}
	if strings.Contains(errorStr, "daily request count exceeded") {
		return TX_ERR_RPC_CONFURA_DAILYQ_EXCEED
	}
	if strings.Contains(errorStr, "EpochHeightOutOfBound") {
		return TX_ERR_RPC_EPOCH_HEIGHT_OUT_BOUND
	}

	return TX_ERR_RPC_OTHER
}

type TxNormalError int

const (
	TX_ERR_NORMAL_OTHER TxNormalError = iota
	TX_ERR_NORMAL_TIMEOUT
	TX_ERR_NORMAL_PENDING_DROPPED_TX_EMPTY
	TX_ERR_NORMAL_PENDING_DROPPED_OUT_EPOCH_HEIGHT
	TX_ERR_NORMAL_DISCONNECT
	TX_ERR_NORMAL_NO_SUCH_HOST
)
