package tokens

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

var TokenFactoryInterface, _ = abi.JSON(strings.NewReader(TokenFactoryABI))
var TokenVaultInterface, _ = abi.JSON(strings.NewReader(TokenVaultABI))

func GetEventID(name string) common.Hash {
	return TokenFactoryInterface.Events[name].ID
}

const (
	EipNotSet uint8 = iota
	EIP20
	EIP721
	EIP1155
)

const (
	OpNotSet uint8 = iota
	MINT
	BURN20
	BURN721
	BURN1155
	TRANSFER
)

const (
	UriModeNotSet uint8 = iota
	BaseUri
	STORAGE
)
