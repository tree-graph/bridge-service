package blockchain

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	evmTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/contracts/tokens"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
	"time"
)

type CfxFetcher struct {
	Client     sdk.Client
	filter     *tokens.TokenVaultFilterer
	filterOpts *bind.FilterOpts
	vaultAddr  string
}

func GetCfxClient() *sdk.Client {
	return nil
}

func NewCfxFetcher(chain int64, vaultAddr string, client *sdk.Client) (*CfxFetcher, error) {
	filter, err := tokens.NewTokenVaultFilterer(cfxaddress.MustNewFromHex(vaultAddr, uint32(chain)), client)
	if err != nil {
		logrus.WithError(err).Error("NewTokenVaultFilterer by cfx fail")
		return nil, err
	}
	fetcher := CfxFetcher{
		Client: *client, filter: filter, vaultAddr: vaultAddr,
	}
	fetcher.filterOpts = &bind.FilterOpts{
		Start: nil,
		End:   nil,
	}
	return &fetcher, nil
}

func (fetcher CfxFetcher) BlockNumber() (uint64, error) {
	epoch, err := fetcher.Client.GetEpochNumber()
	if err != nil {
		return 0, err
	}
	return epoch.ToInt().Uint64(), nil
}

func (fetcher CfxFetcher) Fetch(epoch uint64) (*time.Time, []*vault.VaultCrossRequest, error) {
	fetcher.filterOpts.Start = types.NewEpochNumberUint64(epoch)
	fetcher.filterOpts.End = fetcher.filterOpts.Start

	iter, err := fetcher.filter.FilterCrossRequest(fetcher.filterOpts, nil, nil)
	if err != nil {
		logrus.WithError(err).Error("query cfx logs fail, epoch ", epoch)
	}
	var matchedLogs []*vault.VaultCrossRequest
	for iter.Next() {
		log := iter.Event
		if log.Raw.Address.GetHexAddress() != fetcher.vaultAddr {
			continue
		}
		matchedLogs = append(matchedLogs, convertLog(log))
	}
	blockObj, err := fetcher.Client.GetBlockByEpoch(fetcher.filterOpts.Start)
	if err != nil {
		logrus.Debug("GetBlockByEpoch fail , epoch ", epoch)
		return nil, nil, err
	}
	blockTime := time.Unix(blockObj.Timestamp.ToInt().Int64(), 0)
	return &blockTime, matchedLogs, nil
}

func convertLog(log *tokens.TokenVaultCrossRequest) *vault.VaultCrossRequest {
	return &vault.VaultCrossRequest{
		Raw: evmTypes.Log{
			TxHash:      *log.Raw.TransactionHash.ToCommonHash(),
			BlockNumber: log.Raw.EpochNumber.ToInt().Uint64(),
		},
		Asset:          common.HexToAddress(log.Asset.Hex()),
		From:           common.HexToAddress(log.From.Hex()),
		ToChainId:      log.ToChainId,
		TargetContract: common.HexToAddress(log.TargetContract.Hex()),
		UserNonce:      log.UserNonce,
		TokenIds:       log.TokenIds,
		Amounts:        log.Amounts,
		Uris:           log.Uris,
	}
}
