package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/helpers"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
	"time"
)

type IEventFetcher interface {
	Fetch(blockOrEpoch uint64) (*time.Time, []*vault.VaultCrossRequest, error)
	// cfx chain uses epoch
	BlockNumber() (uint64, error)
}

type EventFetcher struct {
	Handler    EvmHandler
	Filter     *vault.VaultFilterer
	FilterOpts *bind.FilterOpts
	vaultAddr  string
}

func NewEventFetcher(chain int64, vaultAddr string) (*EventFetcher, error) {
	evmHandler := GetEvmHandler(chain)
	contract, err := vault.NewVaultFilterer(common.HexToAddress(vaultAddr), evmHandler.Client)
	if err != nil {
		return nil, err
	}
	fetcher := &EventFetcher{
		Filter:    contract,
		Handler:   evmHandler,
		vaultAddr: vaultAddr,
	}
	fetcher.FilterOpts = &bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: context.Background(),
	}
	return fetcher, nil
}

func (fetcher EventFetcher) BlockNumber() (uint64, error) {
	return fetcher.Handler.Client.BlockNumber(context.Background())
}

func (fetcher EventFetcher) Fetch(block uint64) (*time.Time, []*vault.VaultCrossRequest, error) {
	fetcher.FilterOpts.Start = block
	fetcher.FilterOpts.End = &block

	logIterator, err := fetcher.Filter.FilterCrossRequest(fetcher.FilterOpts, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = logIterator.Close()
	}()

	var logs []*vault.VaultCrossRequest
	for logIterator.Next() {
		logs = append(logs, logIterator.Event)
	}

	if err := logIterator.Error(); err != nil {
		return nil, nil, err
	}

	if len(logs) < 1 {
		return nil, logs, nil
	}

	blockObj, blockErr := fetcher.Handler.BlockByNumber(helpers.Uint64ToBigInt(logs[0].Raw.BlockNumber))
	if blockErr != nil {
		logrus.WithError(blockErr).Error("fetch block fail, block is ", block)
		return nil, nil, nil
	}
	time := time.Unix(int64(blockObj.Time()), 0)

	return &time, logs, nil
}
