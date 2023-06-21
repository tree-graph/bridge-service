// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RouterPeerInfo is an auto generated low-level Go binding around an user-defined struct.
type RouterPeerInfo struct {
	Op         uint8
	Eip        uint8
	Timestamp  *big.Int
	Registerer common.Address
	Enable     bool
	UriMode    uint8
}

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localContract\",\"type\":\"address\"}],\"name\":\"ArrivalConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userNonce\",\"type\":\"uint256\"}],\"name\":\"CrossRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"local\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"remoteContract\",\"type\":\"address\"}],\"name\":\"DepartureConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"localContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userNonce_\",\"type\":\"uint256\"}],\"name\":\"claimByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"index\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"peer\",\"type\":\"address\"}],\"name\":\"getArrivalInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumRouter.OP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"enumRouter.EIP\",\"name\":\"eip\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"registerer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"},{\"internalType\":\"enumRouter.URI_MODE\",\"name\":\"uriMode\",\"type\":\"uint8\"}],\"internalType\":\"structRouter.PeerInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"index\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"peer\",\"type\":\"address\"}],\"name\":\"getDepartureInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumRouter.OP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"enumRouter.EIP\",\"name\":\"eip\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"registerer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"},{\"internalType\":\"enumRouter.URI_MODE\",\"name\":\"uriMode\",\"type\":\"uint8\"}],\"internalType\":\"structRouter.PeerInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"}],\"name\":\"getUserNextClaimNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"index\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"listArrivalChainIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"listArrivalIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"index\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"listArrivalPeerIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"index\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"listDepartureChainIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"listDepartureIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"index\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"listDeparturePeerIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"remoteContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumRouter.OP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"enumRouter.URI_MODE\",\"name\":\"uriMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"localContract\",\"type\":\"address\"}],\"name\":\"registerArrival\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"local\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumRouter.OP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"enumRouter.URI_MODE\",\"name\":\"uriMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"remoteContract\",\"type\":\"address\"}],\"name\":\"registerDeparture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vault *VaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vault *VaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vault.Contract.DEFAULTADMINROLE(&_Vault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vault *VaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vault.Contract.DEFAULTADMINROLE(&_Vault.CallOpts)
}

// Claim is a free data retrieval call binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes ) pure returns()
func (_Vault *VaultCaller) Claim(opts *bind.CallOpts, arg0 []byte) error {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "claim", arg0)

	if err != nil {
		return err
	}

	return err

}

// Claim is a free data retrieval call binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes ) pure returns()
func (_Vault *VaultSession) Claim(arg0 []byte) error {
	return _Vault.Contract.Claim(&_Vault.CallOpts, arg0)
}

// Claim is a free data retrieval call binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes ) pure returns()
func (_Vault *VaultCallerSession) Claim(arg0 []byte) error {
	return _Vault.Contract.Claim(&_Vault.CallOpts, arg0)
}

// GetArrivalInfo is a free data retrieval call binding the contract method 0x0dd643f3.
//
// Solidity: function getArrivalInfo(address index, uint256 chain, address peer) view returns((uint8,uint8,uint256,address,bool,uint8))
func (_Vault *VaultCaller) GetArrivalInfo(opts *bind.CallOpts, index common.Address, chain *big.Int, peer common.Address) (RouterPeerInfo, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getArrivalInfo", index, chain, peer)

	if err != nil {
		return *new(RouterPeerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterPeerInfo)).(*RouterPeerInfo)

	return out0, err

}

// GetArrivalInfo is a free data retrieval call binding the contract method 0x0dd643f3.
//
// Solidity: function getArrivalInfo(address index, uint256 chain, address peer) view returns((uint8,uint8,uint256,address,bool,uint8))
func (_Vault *VaultSession) GetArrivalInfo(index common.Address, chain *big.Int, peer common.Address) (RouterPeerInfo, error) {
	return _Vault.Contract.GetArrivalInfo(&_Vault.CallOpts, index, chain, peer)
}

// GetArrivalInfo is a free data retrieval call binding the contract method 0x0dd643f3.
//
// Solidity: function getArrivalInfo(address index, uint256 chain, address peer) view returns((uint8,uint8,uint256,address,bool,uint8))
func (_Vault *VaultCallerSession) GetArrivalInfo(index common.Address, chain *big.Int, peer common.Address) (RouterPeerInfo, error) {
	return _Vault.Contract.GetArrivalInfo(&_Vault.CallOpts, index, chain, peer)
}

// GetDepartureInfo is a free data retrieval call binding the contract method 0x03610a5c.
//
// Solidity: function getDepartureInfo(address index, uint256 chain, address peer) view returns((uint8,uint8,uint256,address,bool,uint8))
func (_Vault *VaultCaller) GetDepartureInfo(opts *bind.CallOpts, index common.Address, chain *big.Int, peer common.Address) (RouterPeerInfo, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDepartureInfo", index, chain, peer)

	if err != nil {
		return *new(RouterPeerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(RouterPeerInfo)).(*RouterPeerInfo)

	return out0, err

}

// GetDepartureInfo is a free data retrieval call binding the contract method 0x03610a5c.
//
// Solidity: function getDepartureInfo(address index, uint256 chain, address peer) view returns((uint8,uint8,uint256,address,bool,uint8))
func (_Vault *VaultSession) GetDepartureInfo(index common.Address, chain *big.Int, peer common.Address) (RouterPeerInfo, error) {
	return _Vault.Contract.GetDepartureInfo(&_Vault.CallOpts, index, chain, peer)
}

// GetDepartureInfo is a free data retrieval call binding the contract method 0x03610a5c.
//
// Solidity: function getDepartureInfo(address index, uint256 chain, address peer) view returns((uint8,uint8,uint256,address,bool,uint8))
func (_Vault *VaultCallerSession) GetDepartureInfo(index common.Address, chain *big.Int, peer common.Address) (RouterPeerInfo, error) {
	return _Vault.Contract.GetDepartureInfo(&_Vault.CallOpts, index, chain, peer)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vault *VaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vault *VaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vault.Contract.GetRoleAdmin(&_Vault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vault *VaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vault.Contract.GetRoleAdmin(&_Vault.CallOpts, role)
}

// GetUserNextClaimNonce is a free data retrieval call binding the contract method 0x9aa4b2fd.
//
// Solidity: function getUserNextClaimNonce(address issuer, uint256 srcChainId) view returns(uint256)
func (_Vault *VaultCaller) GetUserNextClaimNonce(opts *bind.CallOpts, issuer common.Address, srcChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getUserNextClaimNonce", issuer, srcChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNextClaimNonce is a free data retrieval call binding the contract method 0x9aa4b2fd.
//
// Solidity: function getUserNextClaimNonce(address issuer, uint256 srcChainId) view returns(uint256)
func (_Vault *VaultSession) GetUserNextClaimNonce(issuer common.Address, srcChainId *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetUserNextClaimNonce(&_Vault.CallOpts, issuer, srcChainId)
}

// GetUserNextClaimNonce is a free data retrieval call binding the contract method 0x9aa4b2fd.
//
// Solidity: function getUserNextClaimNonce(address issuer, uint256 srcChainId) view returns(uint256)
func (_Vault *VaultCallerSession) GetUserNextClaimNonce(issuer common.Address, srcChainId *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetUserNextClaimNonce(&_Vault.CallOpts, issuer, srcChainId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vault *VaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vault *VaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vault.Contract.HasRole(&_Vault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vault *VaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vault.Contract.HasRole(&_Vault.CallOpts, role, account)
}

// ListArrivalChainIndex is a free data retrieval call binding the contract method 0x36cda99b.
//
// Solidity: function listArrivalChainIndex(address index, uint256 offset, uint256 size) view returns(uint256[], uint256 total)
func (_Vault *VaultCaller) ListArrivalChainIndex(opts *bind.CallOpts, index common.Address, offset *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "listArrivalChainIndex", index, offset, size)

	if err != nil {
		return *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListArrivalChainIndex is a free data retrieval call binding the contract method 0x36cda99b.
//
// Solidity: function listArrivalChainIndex(address index, uint256 offset, uint256 size) view returns(uint256[], uint256 total)
func (_Vault *VaultSession) ListArrivalChainIndex(index common.Address, offset *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	return _Vault.Contract.ListArrivalChainIndex(&_Vault.CallOpts, index, offset, size)
}

// ListArrivalChainIndex is a free data retrieval call binding the contract method 0x36cda99b.
//
// Solidity: function listArrivalChainIndex(address index, uint256 offset, uint256 size) view returns(uint256[], uint256 total)
func (_Vault *VaultCallerSession) ListArrivalChainIndex(index common.Address, offset *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	return _Vault.Contract.ListArrivalChainIndex(&_Vault.CallOpts, index, offset, size)
}

// ListArrivalIndex is a free data retrieval call binding the contract method 0x7ef848c2.
//
// Solidity: function listArrivalIndex(uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultCaller) ListArrivalIndex(opts *bind.CallOpts, offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "listArrivalIndex", offset, size)

	if err != nil {
		return *new([]common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListArrivalIndex is a free data retrieval call binding the contract method 0x7ef848c2.
//
// Solidity: function listArrivalIndex(uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultSession) ListArrivalIndex(offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Vault.Contract.ListArrivalIndex(&_Vault.CallOpts, offset, size)
}

// ListArrivalIndex is a free data retrieval call binding the contract method 0x7ef848c2.
//
// Solidity: function listArrivalIndex(uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultCallerSession) ListArrivalIndex(offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Vault.Contract.ListArrivalIndex(&_Vault.CallOpts, offset, size)
}

// ListArrivalPeerIndex is a free data retrieval call binding the contract method 0xa7b69cb5.
//
// Solidity: function listArrivalPeerIndex(address index, uint256 chain, uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultCaller) ListArrivalPeerIndex(opts *bind.CallOpts, index common.Address, chain *big.Int, offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "listArrivalPeerIndex", index, chain, offset, size)

	if err != nil {
		return *new([]common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListArrivalPeerIndex is a free data retrieval call binding the contract method 0xa7b69cb5.
//
// Solidity: function listArrivalPeerIndex(address index, uint256 chain, uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultSession) ListArrivalPeerIndex(index common.Address, chain *big.Int, offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Vault.Contract.ListArrivalPeerIndex(&_Vault.CallOpts, index, chain, offset, size)
}

// ListArrivalPeerIndex is a free data retrieval call binding the contract method 0xa7b69cb5.
//
// Solidity: function listArrivalPeerIndex(address index, uint256 chain, uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultCallerSession) ListArrivalPeerIndex(index common.Address, chain *big.Int, offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Vault.Contract.ListArrivalPeerIndex(&_Vault.CallOpts, index, chain, offset, size)
}

// ListDepartureChainIndex is a free data retrieval call binding the contract method 0x1ef9da61.
//
// Solidity: function listDepartureChainIndex(address index, uint256 offset, uint256 size) view returns(uint256[], uint256 total)
func (_Vault *VaultCaller) ListDepartureChainIndex(opts *bind.CallOpts, index common.Address, offset *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "listDepartureChainIndex", index, offset, size)

	if err != nil {
		return *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListDepartureChainIndex is a free data retrieval call binding the contract method 0x1ef9da61.
//
// Solidity: function listDepartureChainIndex(address index, uint256 offset, uint256 size) view returns(uint256[], uint256 total)
func (_Vault *VaultSession) ListDepartureChainIndex(index common.Address, offset *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	return _Vault.Contract.ListDepartureChainIndex(&_Vault.CallOpts, index, offset, size)
}

// ListDepartureChainIndex is a free data retrieval call binding the contract method 0x1ef9da61.
//
// Solidity: function listDepartureChainIndex(address index, uint256 offset, uint256 size) view returns(uint256[], uint256 total)
func (_Vault *VaultCallerSession) ListDepartureChainIndex(index common.Address, offset *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	return _Vault.Contract.ListDepartureChainIndex(&_Vault.CallOpts, index, offset, size)
}

// ListDepartureIndex is a free data retrieval call binding the contract method 0x1a10c5fe.
//
// Solidity: function listDepartureIndex(uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultCaller) ListDepartureIndex(opts *bind.CallOpts, offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "listDepartureIndex", offset, size)

	if err != nil {
		return *new([]common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListDepartureIndex is a free data retrieval call binding the contract method 0x1a10c5fe.
//
// Solidity: function listDepartureIndex(uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultSession) ListDepartureIndex(offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Vault.Contract.ListDepartureIndex(&_Vault.CallOpts, offset, size)
}

// ListDepartureIndex is a free data retrieval call binding the contract method 0x1a10c5fe.
//
// Solidity: function listDepartureIndex(uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultCallerSession) ListDepartureIndex(offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Vault.Contract.ListDepartureIndex(&_Vault.CallOpts, offset, size)
}

// ListDeparturePeerIndex is a free data retrieval call binding the contract method 0x5e2fdcbc.
//
// Solidity: function listDeparturePeerIndex(address index, uint256 chain, uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultCaller) ListDeparturePeerIndex(opts *bind.CallOpts, index common.Address, chain *big.Int, offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "listDeparturePeerIndex", index, chain, offset, size)

	if err != nil {
		return *new([]common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListDeparturePeerIndex is a free data retrieval call binding the contract method 0x5e2fdcbc.
//
// Solidity: function listDeparturePeerIndex(address index, uint256 chain, uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultSession) ListDeparturePeerIndex(index common.Address, chain *big.Int, offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Vault.Contract.ListDeparturePeerIndex(&_Vault.CallOpts, index, chain, offset, size)
}

// ListDeparturePeerIndex is a free data retrieval call binding the contract method 0x5e2fdcbc.
//
// Solidity: function listDeparturePeerIndex(address index, uint256 chain, uint256 offset, uint256 size) view returns(address[], uint256 total)
func (_Vault *VaultCallerSession) ListDeparturePeerIndex(index common.Address, chain *big.Int, offset *big.Int, size *big.Int) ([]common.Address, *big.Int, error) {
	return _Vault.Contract.ListDeparturePeerIndex(&_Vault.CallOpts, index, chain, offset, size)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vault *VaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vault *VaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vault.Contract.SupportsInterface(&_Vault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vault *VaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vault.Contract.SupportsInterface(&_Vault.CallOpts, interfaceId)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xc11bec28.
//
// Solidity: function batchTransferFrom(uint256 toChainId, address targetContract, address assetContract, uint256[] tokenIds) returns()
func (_Vault *VaultTransactor) BatchTransferFrom(opts *bind.TransactOpts, toChainId *big.Int, targetContract common.Address, assetContract common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "batchTransferFrom", toChainId, targetContract, assetContract, tokenIds)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xc11bec28.
//
// Solidity: function batchTransferFrom(uint256 toChainId, address targetContract, address assetContract, uint256[] tokenIds) returns()
func (_Vault *VaultSession) BatchTransferFrom(toChainId *big.Int, targetContract common.Address, assetContract common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Vault.Contract.BatchTransferFrom(&_Vault.TransactOpts, toChainId, targetContract, assetContract, tokenIds)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xc11bec28.
//
// Solidity: function batchTransferFrom(uint256 toChainId, address targetContract, address assetContract, uint256[] tokenIds) returns()
func (_Vault *VaultTransactorSession) BatchTransferFrom(toChainId *big.Int, targetContract common.Address, assetContract common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Vault.Contract.BatchTransferFrom(&_Vault.TransactOpts, toChainId, targetContract, assetContract, tokenIds)
}

// ClaimByAdmin is a paid mutator transaction binding the contract method 0xc4eb2460.
//
// Solidity: function claimByAdmin(uint256 srcChainId, address srcContract, address localContract, uint256[] tokenIds, uint256[] amounts, string[] uris, address issuer, uint256 userNonce_) returns()
func (_Vault *VaultTransactor) ClaimByAdmin(opts *bind.TransactOpts, srcChainId *big.Int, srcContract common.Address, localContract common.Address, tokenIds []*big.Int, amounts []*big.Int, uris []string, issuer common.Address, userNonce_ *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "claimByAdmin", srcChainId, srcContract, localContract, tokenIds, amounts, uris, issuer, userNonce_)
}

// ClaimByAdmin is a paid mutator transaction binding the contract method 0xc4eb2460.
//
// Solidity: function claimByAdmin(uint256 srcChainId, address srcContract, address localContract, uint256[] tokenIds, uint256[] amounts, string[] uris, address issuer, uint256 userNonce_) returns()
func (_Vault *VaultSession) ClaimByAdmin(srcChainId *big.Int, srcContract common.Address, localContract common.Address, tokenIds []*big.Int, amounts []*big.Int, uris []string, issuer common.Address, userNonce_ *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.ClaimByAdmin(&_Vault.TransactOpts, srcChainId, srcContract, localContract, tokenIds, amounts, uris, issuer, userNonce_)
}

// ClaimByAdmin is a paid mutator transaction binding the contract method 0xc4eb2460.
//
// Solidity: function claimByAdmin(uint256 srcChainId, address srcContract, address localContract, uint256[] tokenIds, uint256[] amounts, string[] uris, address issuer, uint256 userNonce_) returns()
func (_Vault *VaultTransactorSession) ClaimByAdmin(srcChainId *big.Int, srcContract common.Address, localContract common.Address, tokenIds []*big.Int, amounts []*big.Int, uris []string, issuer common.Address, userNonce_ *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.ClaimByAdmin(&_Vault.TransactOpts, srcChainId, srcContract, localContract, tokenIds, amounts, uris, issuer, userNonce_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vault *VaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vault *VaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.Contract.GrantRole(&_Vault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vault *VaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.Contract.GrantRole(&_Vault.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vault *VaultTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vault *VaultSession) Initialize() (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vault *VaultTransactorSession) Initialize() (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_Vault *VaultTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "onERC1155BatchReceived", arg0, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_Vault *VaultSession) OnERC1155BatchReceived(arg0 common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.Contract.OnERC1155BatchReceived(&_Vault.TransactOpts, arg0, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_Vault *VaultTransactorSession) OnERC1155BatchReceived(arg0 common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.Contract.OnERC1155BatchReceived(&_Vault.TransactOpts, arg0, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_Vault *VaultTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "onERC1155Received", arg0, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_Vault *VaultSession) OnERC1155Received(arg0 common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.Contract.OnERC1155Received(&_Vault.TransactOpts, arg0, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_Vault *VaultTransactorSession) OnERC1155Received(arg0 common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.Contract.OnERC1155Received(&_Vault.TransactOpts, arg0, from, id, value, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Vault *VaultTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "onERC721Received", arg0, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Vault *VaultSession) OnERC721Received(arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.Contract.OnERC721Received(&_Vault.TransactOpts, arg0, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Vault *VaultTransactorSession) OnERC721Received(arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Vault.Contract.OnERC721Received(&_Vault.TransactOpts, arg0, from, tokenId, data)
}

// RegisterArrival is a paid mutator transaction binding the contract method 0x364f2de0.
//
// Solidity: function registerArrival(address remoteContract, uint256 remoteChainId, uint8 op, uint8 uriMode, address localContract) returns()
func (_Vault *VaultTransactor) RegisterArrival(opts *bind.TransactOpts, remoteContract common.Address, remoteChainId *big.Int, op uint8, uriMode uint8, localContract common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "registerArrival", remoteContract, remoteChainId, op, uriMode, localContract)
}

// RegisterArrival is a paid mutator transaction binding the contract method 0x364f2de0.
//
// Solidity: function registerArrival(address remoteContract, uint256 remoteChainId, uint8 op, uint8 uriMode, address localContract) returns()
func (_Vault *VaultSession) RegisterArrival(remoteContract common.Address, remoteChainId *big.Int, op uint8, uriMode uint8, localContract common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RegisterArrival(&_Vault.TransactOpts, remoteContract, remoteChainId, op, uriMode, localContract)
}

// RegisterArrival is a paid mutator transaction binding the contract method 0x364f2de0.
//
// Solidity: function registerArrival(address remoteContract, uint256 remoteChainId, uint8 op, uint8 uriMode, address localContract) returns()
func (_Vault *VaultTransactorSession) RegisterArrival(remoteContract common.Address, remoteChainId *big.Int, op uint8, uriMode uint8, localContract common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RegisterArrival(&_Vault.TransactOpts, remoteContract, remoteChainId, op, uriMode, localContract)
}

// RegisterDeparture is a paid mutator transaction binding the contract method 0x03cda6e0.
//
// Solidity: function registerDeparture(address local, uint256 targetChainId, uint8 op, uint8 uriMode, address remoteContract) returns()
func (_Vault *VaultTransactor) RegisterDeparture(opts *bind.TransactOpts, local common.Address, targetChainId *big.Int, op uint8, uriMode uint8, remoteContract common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "registerDeparture", local, targetChainId, op, uriMode, remoteContract)
}

// RegisterDeparture is a paid mutator transaction binding the contract method 0x03cda6e0.
//
// Solidity: function registerDeparture(address local, uint256 targetChainId, uint8 op, uint8 uriMode, address remoteContract) returns()
func (_Vault *VaultSession) RegisterDeparture(local common.Address, targetChainId *big.Int, op uint8, uriMode uint8, remoteContract common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RegisterDeparture(&_Vault.TransactOpts, local, targetChainId, op, uriMode, remoteContract)
}

// RegisterDeparture is a paid mutator transaction binding the contract method 0x03cda6e0.
//
// Solidity: function registerDeparture(address local, uint256 targetChainId, uint8 op, uint8 uriMode, address remoteContract) returns()
func (_Vault *VaultTransactorSession) RegisterDeparture(local common.Address, targetChainId *big.Int, op uint8, uriMode uint8, remoteContract common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RegisterDeparture(&_Vault.TransactOpts, local, targetChainId, op, uriMode, remoteContract)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vault *VaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vault *VaultSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RenounceRole(&_Vault.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vault *VaultTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RenounceRole(&_Vault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vault *VaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vault *VaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RevokeRole(&_Vault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vault *VaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RevokeRole(&_Vault.TransactOpts, role, account)
}

// VaultArrivalConfiguredIterator is returned from FilterArrivalConfigured and is used to iterate over the raw logs and unpacked data for ArrivalConfigured events raised by the Vault contract.
type VaultArrivalConfiguredIterator struct {
	Event *VaultArrivalConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultArrivalConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultArrivalConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultArrivalConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultArrivalConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultArrivalConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultArrivalConfigured represents a ArrivalConfigured event raised by the Vault contract.
type VaultArrivalConfigured struct {
	RemoteContract common.Address
	RemoteChainId  *big.Int
	LocalContract  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterArrivalConfigured is a free log retrieval operation binding the contract event 0xe9fb7d7e00118c8e027ac046db156d6e48ef637bd79b3d7fa30708d7b3b5e58b.
//
// Solidity: event ArrivalConfigured(address indexed remoteContract, uint256 indexed remoteChainId, address localContract)
func (_Vault *VaultFilterer) FilterArrivalConfigured(opts *bind.FilterOpts, remoteContract []common.Address, remoteChainId []*big.Int) (*VaultArrivalConfiguredIterator, error) {

	var remoteContractRule []interface{}
	for _, remoteContractItem := range remoteContract {
		remoteContractRule = append(remoteContractRule, remoteContractItem)
	}
	var remoteChainIdRule []interface{}
	for _, remoteChainIdItem := range remoteChainId {
		remoteChainIdRule = append(remoteChainIdRule, remoteChainIdItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "ArrivalConfigured", remoteContractRule, remoteChainIdRule)
	if err != nil {
		return nil, err
	}
	return &VaultArrivalConfiguredIterator{contract: _Vault.contract, event: "ArrivalConfigured", logs: logs, sub: sub}, nil
}

// WatchArrivalConfigured is a free log subscription operation binding the contract event 0xe9fb7d7e00118c8e027ac046db156d6e48ef637bd79b3d7fa30708d7b3b5e58b.
//
// Solidity: event ArrivalConfigured(address indexed remoteContract, uint256 indexed remoteChainId, address localContract)
func (_Vault *VaultFilterer) WatchArrivalConfigured(opts *bind.WatchOpts, sink chan<- *VaultArrivalConfigured, remoteContract []common.Address, remoteChainId []*big.Int) (event.Subscription, error) {

	var remoteContractRule []interface{}
	for _, remoteContractItem := range remoteContract {
		remoteContractRule = append(remoteContractRule, remoteContractItem)
	}
	var remoteChainIdRule []interface{}
	for _, remoteChainIdItem := range remoteChainId {
		remoteChainIdRule = append(remoteChainIdRule, remoteChainIdItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "ArrivalConfigured", remoteContractRule, remoteChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultArrivalConfigured)
				if err := _Vault.contract.UnpackLog(event, "ArrivalConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseArrivalConfigured is a log parse operation binding the contract event 0xe9fb7d7e00118c8e027ac046db156d6e48ef637bd79b3d7fa30708d7b3b5e58b.
//
// Solidity: event ArrivalConfigured(address indexed remoteContract, uint256 indexed remoteChainId, address localContract)
func (_Vault *VaultFilterer) ParseArrivalConfigured(log types.Log) (*VaultArrivalConfigured, error) {
	event := new(VaultArrivalConfigured)
	if err := _Vault.contract.UnpackLog(event, "ArrivalConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultCrossRequestIterator is returned from FilterCrossRequest and is used to iterate over the raw logs and unpacked data for CrossRequest events raised by the Vault contract.
type VaultCrossRequestIterator struct {
	Event *VaultCrossRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultCrossRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultCrossRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultCrossRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultCrossRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultCrossRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultCrossRequest represents a CrossRequest event raised by the Vault contract.
type VaultCrossRequest struct {
	Asset          common.Address
	From           common.Address
	TokenIds       []*big.Int
	Amounts        []*big.Int
	Uris           []string
	ToChainId      *big.Int
	TargetContract common.Address
	UserNonce      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCrossRequest is a free log retrieval operation binding the contract event 0x3863b4c2b672e3221642422fc9637df01babed1d750409520a373b4a787afc59.
//
// Solidity: event CrossRequest(address indexed asset, address indexed from, uint256[] tokenIds, uint256[] amounts, string[] uris, uint256 toChainId, address targetContract, uint256 userNonce)
func (_Vault *VaultFilterer) FilterCrossRequest(opts *bind.FilterOpts, asset []common.Address, from []common.Address) (*VaultCrossRequestIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "CrossRequest", assetRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &VaultCrossRequestIterator{contract: _Vault.contract, event: "CrossRequest", logs: logs, sub: sub}, nil
}

// WatchCrossRequest is a free log subscription operation binding the contract event 0x3863b4c2b672e3221642422fc9637df01babed1d750409520a373b4a787afc59.
//
// Solidity: event CrossRequest(address indexed asset, address indexed from, uint256[] tokenIds, uint256[] amounts, string[] uris, uint256 toChainId, address targetContract, uint256 userNonce)
func (_Vault *VaultFilterer) WatchCrossRequest(opts *bind.WatchOpts, sink chan<- *VaultCrossRequest, asset []common.Address, from []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "CrossRequest", assetRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultCrossRequest)
				if err := _Vault.contract.UnpackLog(event, "CrossRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossRequest is a log parse operation binding the contract event 0x3863b4c2b672e3221642422fc9637df01babed1d750409520a373b4a787afc59.
//
// Solidity: event CrossRequest(address indexed asset, address indexed from, uint256[] tokenIds, uint256[] amounts, string[] uris, uint256 toChainId, address targetContract, uint256 userNonce)
func (_Vault *VaultFilterer) ParseCrossRequest(log types.Log) (*VaultCrossRequest, error) {
	event := new(VaultCrossRequest)
	if err := _Vault.contract.UnpackLog(event, "CrossRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDepartureConfiguredIterator is returned from FilterDepartureConfigured and is used to iterate over the raw logs and unpacked data for DepartureConfigured events raised by the Vault contract.
type VaultDepartureConfiguredIterator struct {
	Event *VaultDepartureConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDepartureConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDepartureConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDepartureConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDepartureConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepartureConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDepartureConfigured represents a DepartureConfigured event raised by the Vault contract.
type VaultDepartureConfigured struct {
	Local          common.Address
	TargetChainId  *big.Int
	RemoteContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDepartureConfigured is a free log retrieval operation binding the contract event 0x0e5a0f5238edad743f0bbb7140294ec2214d83e62fc9a2816d3f79163981cb53.
//
// Solidity: event DepartureConfigured(address indexed local, uint256 indexed targetChainId, address remoteContract)
func (_Vault *VaultFilterer) FilterDepartureConfigured(opts *bind.FilterOpts, local []common.Address, targetChainId []*big.Int) (*VaultDepartureConfiguredIterator, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DepartureConfigured", localRule, targetChainIdRule)
	if err != nil {
		return nil, err
	}
	return &VaultDepartureConfiguredIterator{contract: _Vault.contract, event: "DepartureConfigured", logs: logs, sub: sub}, nil
}

// WatchDepartureConfigured is a free log subscription operation binding the contract event 0x0e5a0f5238edad743f0bbb7140294ec2214d83e62fc9a2816d3f79163981cb53.
//
// Solidity: event DepartureConfigured(address indexed local, uint256 indexed targetChainId, address remoteContract)
func (_Vault *VaultFilterer) WatchDepartureConfigured(opts *bind.WatchOpts, sink chan<- *VaultDepartureConfigured, local []common.Address, targetChainId []*big.Int) (event.Subscription, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DepartureConfigured", localRule, targetChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDepartureConfigured)
				if err := _Vault.contract.UnpackLog(event, "DepartureConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepartureConfigured is a log parse operation binding the contract event 0x0e5a0f5238edad743f0bbb7140294ec2214d83e62fc9a2816d3f79163981cb53.
//
// Solidity: event DepartureConfigured(address indexed local, uint256 indexed targetChainId, address remoteContract)
func (_Vault *VaultFilterer) ParseDepartureConfigured(log types.Log) (*VaultDepartureConfigured, error) {
	event := new(VaultDepartureConfigured)
	if err := _Vault.contract.UnpackLog(event, "DepartureConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Vault contract.
type VaultRoleAdminChangedIterator struct {
	Event *VaultRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRoleAdminChanged represents a RoleAdminChanged event raised by the Vault contract.
type VaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vault *VaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VaultRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VaultRoleAdminChangedIterator{contract: _Vault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vault *VaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRoleAdminChanged)
				if err := _Vault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vault *VaultFilterer) ParseRoleAdminChanged(log types.Log) (*VaultRoleAdminChanged, error) {
	event := new(VaultRoleAdminChanged)
	if err := _Vault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Vault contract.
type VaultRoleGrantedIterator struct {
	Event *VaultRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRoleGranted represents a RoleGranted event raised by the Vault contract.
type VaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vault *VaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultRoleGrantedIterator{contract: _Vault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vault *VaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRoleGranted)
				if err := _Vault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vault *VaultFilterer) ParseRoleGranted(log types.Log) (*VaultRoleGranted, error) {
	event := new(VaultRoleGranted)
	if err := _Vault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Vault contract.
type VaultRoleRevokedIterator struct {
	Event *VaultRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRoleRevoked represents a RoleRevoked event raised by the Vault contract.
type VaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vault *VaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultRoleRevokedIterator{contract: _Vault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vault *VaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRoleRevoked)
				if err := _Vault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vault *VaultFilterer) ParseRoleRevoked(log types.Log) (*VaultRoleRevoked, error) {
	event := new(VaultRoleRevoked)
	if err := _Vault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
