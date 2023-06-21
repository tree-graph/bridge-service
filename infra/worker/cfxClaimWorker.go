package worker

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/tree-graph/bridge-service/models"
)

type CfxClaimWorker struct {
	Chain         models.Chain
	ChainId       int64
	CfxClient     *sdk.Client
	address       types.Address
	DelayForError int
}

func (worker CfxClaimWorker) GetChainId() int64 {
	return worker.ChainId
}
func (CfxClaimWorker) DoWork() (int, error) {
	return 0, nil
}
