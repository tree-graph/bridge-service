package blockchain

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
)

/**
Differences between Consortium chain and Conflux Chain are here.
Different branch could changes code here, to avoid conflicts.
*/
func PatchChain(client *sdk.Client) uint32 {
	forceNetworkId := uint32(1029)
	// consortium chain
	client.SetNetworkId(forceNetworkId)
	// Conflux chain
	//forceNetworkId, err := client.GetNetworkID()
	//helpers.CheckFatalError("PatchChain GetNetworkID", err)
	return forceNetworkId
}
