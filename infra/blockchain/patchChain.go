package blockchain

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/tree-graph/bridge-service/helpers"
)

/**
Differences between Consortium chain and Conflux Chain are here.
Different branch could changes code here, to avoid conflicts.
*/
func PatchChain(client sdk.Client) uint32 {
	forceNetworkId := uint32(0)
	// consortium chain
	//client.SetNetworkId(uint32(forceNetworkId))
	// Conflux chain
	forceNetworkId, err := client.GetNetworkID()
	helpers.CheckFatalError("PatchChain GetNetworkID", err)
	return forceNetworkId
}
