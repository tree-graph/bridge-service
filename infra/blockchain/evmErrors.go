package blockchain

import (
	"fmt"
	"math/big"
	"strings"
)

type ErrNotEnoughCashes map[string]ErrNotEnoughCash

func (e ErrNotEnoughCashes) Error() string {
	var msgArray []string
	for user, err := range e {
		msgArray = append(msgArray, fmt.Sprintf("%v %v", user, err.Error()))
	}
	return strings.Join(msgArray, "\n")
}

type ErrNotEnoughCash struct {
	Need *big.Int
	Got  *big.Int
}

func (e ErrNotEnoughCash) Error() string {
	return fmt.Sprintf("out of balance, need %v, got %v", e.Need, e.Got)
}

type TxRpcError int

const (
	TxErrRpcOutOfBalance TxRpcError = iota
	TxErrRpcTxPoolFull
	TxErrRpcInsertValidationFail
	TxErrRpcInsertFailQuotaNotEnough
	TxErrRpcSameNonceAlreadyInserted
	TxErrRpcStaleNonce
	TxErrRpcExceedMaxGas
	TxErrRpcPoolInconsistent
	TxErrRpcFailReadAccountCache
	TxErrRpcNonceDistant
	TxErrRpcFailedImportToDefferPool
	TxAlreadyExist

	TxErrRpcTooManyRequest
	TxErrRpcConfuraQpsExceed
	TxErrRpcConfuraDailyQpsExceed
	TxErrRpcMissingBestHash
	TxErrRpcEpochHeightOutBound
	// includes: vm revert, not enough cash
	TxErrEstimateOthers
	TxErrRpcEstimateNotEnoughCash
	TxErrRpcEstimateVmErr
	TxErrDropped
	TxErrRpcOther
)

func (e TxRpcError) CheckTxErrorStatus() (needSend, needUpperGas bool) {
	switch e {
	case TxErrRpcOutOfBalance:
		return false, false
	case TxErrRpcTxPoolFull:
		return true, false
	case TxErrRpcInsertValidationFail:
		return false, false
	case TxErrRpcInsertFailQuotaNotEnough:
		return false, false
	case TxErrRpcSameNonceAlreadyInserted:
		return true, true
	case TxErrRpcStaleNonce:
		return false, false
	case TxErrRpcExceedMaxGas:
		return false, false
	case TxErrRpcPoolInconsistent:
		return true, false
	case TxErrRpcFailReadAccountCache:
		return true, false
	case TxErrRpcNonceDistant:
		return false, false
	case TxErrRpcFailedImportToDefferPool:
		return true, false
	case TxAlreadyExist:
		return false, false
	case TxErrRpcTooManyRequest:
		fallthrough
	case TxErrRpcConfuraQpsExceed:
		fallthrough
	case TxErrRpcConfuraDailyQpsExceed:
		return true, false
	case TxErrRpcEpochHeightOutBound: //EpochHeightOutOfBound
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

func ParseRpcError(errorStr string) TxRpcError {
	if strings.Contains(errorStr, "txpool is full") {
		return TxErrRpcTxPoolFull
	}
	if strings.Contains(errorStr, "Transaction Pool is full") {
		return TxErrRpcTxPoolFull
	}
	if strings.Contains(errorStr, "failed to insert tx into pool (validation failed)") {
		return TxErrRpcInsertValidationFail
	}
	if strings.Contains(errorStr, "failed to insert tx into pool (quota not enough)") {
		return TxErrRpcInsertFailQuotaNotEnough
	}
	if strings.Contains(errorStr, "exceeds the maximum value") {
		return TxErrRpcExceedMaxGas
	}
	if strings.Contains(errorStr, "Tx with same nonce already inserted") {
		return TxErrRpcSameNonceAlreadyInserted
	}
	if strings.Contains(errorStr, "ready_account_pool and deferred_pool are inconsistent!") {
		return TxErrRpcPoolInconsistent
	}
	if strings.Contains(errorStr, "Failed to read account_cache from storage") {
		return TxErrRpcFailReadAccountCache
	}
	if strings.Contains(errorStr, "discarded due to in too distant future") {
		return TxErrRpcNonceDistant
	}
	if strings.Contains(errorStr, "discarded due to a too stale nonce") {
		return TxErrRpcStaleNonce
	}
	if strings.Contains(errorStr, "discarded due to out of balance") {
		return TxErrRpcOutOfBalance
	}
	if strings.Contains(errorStr, "NotEnoughCash") {
		return TxErrRpcEstimateNotEnoughCash
	}
	if strings.Contains(errorStr, "block_number is missing for best_hash") {
		return TxErrRpcMissingBestHash
	}
	if strings.Contains(errorStr, "Failed imported to deferred pool") {
		return TxErrRpcFailedImportToDefferPool
	}
	if strings.Contains(errorStr, "tx already exist") {
		return TxAlreadyExist
	}
	if strings.Contains(errorStr, "Vm reverted") || strings.Contains(errorStr, "VmError") || strings.Contains(errorStr, "Transaction reverted") {
		return TxErrRpcEstimateVmErr
	}
	if strings.Contains(errorStr, "too many requests") {
		return TxErrRpcTooManyRequest
	}
	if strings.Contains(errorStr, "allowed qps exceeded") {
		return TxErrRpcConfuraQpsExceed
	}
	if strings.Contains(errorStr, "daily request count exceeded") {
		return TxErrRpcConfuraDailyQpsExceed
	}
	if strings.Contains(errorStr, "EpochHeightOutOfBound") {
		return TxErrRpcEpochHeightOutBound
	}

	return TxErrRpcOther
}
