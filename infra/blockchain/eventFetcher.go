package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
)

type EventFetcher struct {
	Handler    EvmHandler
	Filter     *vault.VaultFilterer
	FilterOpts *bind.FilterOpts
}

func NewEventFetcher(chain int64, vaultAddr string) (*EventFetcher, error) {
	evmHandler := GetEvmHandler(chain)
	contract, err := vault.NewVaultFilterer(common.HexToAddress(vaultAddr), evmHandler.Client)
	if err != nil {
		return nil, err
	}
	fetcher := &EventFetcher{
		Filter:  contract,
		Handler: evmHandler,
	}
	fetcher.FilterOpts = &bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: context.Background(),
	}
	return fetcher, nil
}

func (fetcher EventFetcher) Fetch(block uint64) ([]*vault.VaultCrossRequest, error) {
	fetcher.FilterOpts.Start = block
	fetcher.FilterOpts.End = &block

	logIterator, err := fetcher.Filter.FilterCrossRequest(fetcher.FilterOpts, nil, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = logIterator.Close()
	}()

	var logs []*vault.VaultCrossRequest
	for logIterator.Next() {
		logs = append(logs, logIterator.Event)
	}

	if err := logIterator.Error(); err != nil {
		return nil, err
	}

	return logs, nil
}
