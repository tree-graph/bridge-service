// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openzeppelin

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"

	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// UpgradeableBeaconABI is the input ABI used to generate the binding from.
const UpgradeableBeaconABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UpgradeableBeaconBin is the compiled bytecode used for deploying new contracts.
var UpgradeableBeaconBin = "0x608060405234801561001057600080fd5b506040516109cf3803806109cf8339818101604052810190610032919061024b565b61004e61004361006360201b60201c565b61006b60201b60201c565b61005d8161012f60201b60201c565b5061031b565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610142816101c560201b61021b1760201c565b610181576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610178906102fb565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610218826101ed565b9050919050565b6102288161020d565b811461023357600080fd5b50565b6000815190506102458161021f565b92915050565b600060208284031215610261576102606101e8565b5b600061026f84828501610236565b91505092915050565b600082825260208201905092915050565b7f5570677261646561626c65426561636f6e3a20696d706c656d656e746174696f60008201527f6e206973206e6f74206120636f6e747261637400000000000000000000000000602082015250565b60006102e5603383610278565b91506102f082610289565b604082019050919050565b60006020820190508181036000830152610314816102d8565b9050919050565b6106a58061032a6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633659cfe61461005c5780635c60da1b14610078578063715018a6146100965780638da5cb5b146100a0578063f2fde38b146100be575b600080fd5b61007660048036038101906100719190610477565b6100da565b005b610080610131565b60405161008d91906104b3565b60405180910390f35b61009e61015b565b005b6100a861016f565b6040516100b591906104b3565b60405180910390f35b6100d860048036038101906100d39190610477565b610198565b005b6100e261023e565b6100eb816102bc565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61016361023e565b61016d6000610348565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6101a061023e565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361020f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020690610551565b60405180910390fd5b61021881610348565b50565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b61024661040c565b73ffffffffffffffffffffffffffffffffffffffff1661026461016f565b73ffffffffffffffffffffffffffffffffffffffff16146102ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b1906105bd565b60405180910390fd5b565b6102c58161021b565b610304576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fb9061064f565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061044482610419565b9050919050565b61045481610439565b811461045f57600080fd5b50565b6000813590506104718161044b565b92915050565b60006020828403121561048d5761048c610414565b5b600061049b84828501610462565b91505092915050565b6104ad81610439565b82525050565b60006020820190506104c860008301846104a4565b92915050565b600082825260208201905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061053b6026836104ce565b9150610546826104df565b604082019050919050565b6000602082019050818103600083015261056a8161052e565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006105a76020836104ce565b91506105b282610571565b602082019050919050565b600060208201905081810360008301526105d68161059a565b9050919050565b7f5570677261646561626c65426561636f6e3a20696d706c656d656e746174696f60008201527f6e206973206e6f74206120636f6e747261637400000000000000000000000000602082015250565b60006106396033836104ce565b9150610644826105dd565b604082019050919050565b600060208201905081810360008301526106688161062c565b905091905056fea26469706673582212209c6355948815669edbba12ebf56954fd72cd132116e4dcf95fc773dce57055bd64736f6c63430008120033"

// DeployUpgradeableBeacon deploys a new Conflux contract, binding an instance of UpgradeableBeacon to it.
func DeployUpgradeableBeacon(auth *bind.TransactOpts, backend bind.ContractBackend, implementation_ common.Address) (*types.UnsignedTransaction, *types.Hash, *UpgradeableBeacon, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeableBeaconABI))
	if err != nil {
		return nil, nil, nil, err
	}

	tx, hash, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradeableBeaconBin), backend, implementation_)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, hash, &UpgradeableBeacon{UpgradeableBeaconCaller: UpgradeableBeaconCaller{contract: contract}, UpgradeableBeaconTransactor: UpgradeableBeaconTransactor{contract: contract}, UpgradeableBeaconFilterer: UpgradeableBeaconFilterer{contract: contract}}, nil
}

// UpgradeableBeacon is an auto generated Go binding around an Conflux contract.
type UpgradeableBeacon struct {
	UpgradeableBeaconCaller     // Read-only binding to the contract
	UpgradeableBeaconTransactor // Write-only binding to the contract
	UpgradeableBeaconFilterer   // Log filterer for contract events
}

// UpgradeableBeaconCaller is an auto generated read-only Go binding around an Conflux contract.
type UpgradeableBeaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableBeaconBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type UpgradeableBeaconBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableBeaconTransactor is an auto generated write-only Go binding around an Conflux contract.
type UpgradeableBeaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableBeaconBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type UpgradeableBeaconBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableBeaconFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type UpgradeableBeaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableBeaconSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type UpgradeableBeaconSession struct {
	Contract     *UpgradeableBeacon // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UpgradeableBeaconCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type UpgradeableBeaconCallerSession struct {
	Contract *UpgradeableBeaconCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UpgradeableBeaconTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type UpgradeableBeaconTransactorSession struct {
	Contract     *UpgradeableBeaconTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UpgradeableBeaconRaw is an auto generated low-level Go binding around an Conflux contract.
type UpgradeableBeaconRaw struct {
	Contract *UpgradeableBeacon // Generic contract binding to access the raw methods on
}

// UpgradeableBeaconCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type UpgradeableBeaconCallerRaw struct {
	Contract *UpgradeableBeaconCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeableBeaconTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type UpgradeableBeaconTransactorRaw struct {
	Contract *UpgradeableBeaconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeableBeacon creates a new instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeacon(address types.Address, backend bind.ContractBackend) (*UpgradeableBeacon, error) {
	contract, err := bindUpgradeableBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeacon{UpgradeableBeaconCaller: UpgradeableBeaconCaller{contract: contract}, UpgradeableBeaconTransactor: UpgradeableBeaconTransactor{contract: contract}, UpgradeableBeaconFilterer: UpgradeableBeaconFilterer{contract: contract}}, nil
}

// NewUpgradeableBeaconCaller creates a new read-only instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeaconCaller(address types.Address, caller bind.ContractCaller) (*UpgradeableBeaconCaller, error) {
	contract, err := bindUpgradeableBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconCaller{contract: contract}, nil
}

// NewUpgradeableBeaconTransactor creates a new write-only instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeaconTransactor(address types.Address, transactor bind.ContractTransactor) (*UpgradeableBeaconTransactor, error) {
	contract, err := bindUpgradeableBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconTransactor{contract: contract}, nil
}

// NewUpgradeableBeaconFilterer creates a new log filterer instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeaconFilterer(address types.Address, filterer bind.ContractFilterer) (*UpgradeableBeaconFilterer, error) {
	contract, err := bindUpgradeableBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconFilterer{contract: contract}, nil
}

// NewUpgradeableBeaconCaller creates a new read-only instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeaconBulkCaller(address types.Address, caller bind.ContractCaller) (*UpgradeableBeaconBulkCaller, error) {
	contract, err := bindUpgradeableBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconBulkCaller{contract: contract}, nil
}

// NewUpgradeableBeaconBulkTransactor creates a new write-only instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeaconBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*UpgradeableBeaconBulkTransactor, error) {
	contract, err := bindUpgradeableBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconBulkTransactor{contract: contract}, nil
}

// bindUpgradeableBeacon binds a generic wrapper to an already deployed contract.
func bindUpgradeableBeacon(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeableBeaconABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableBeacon *UpgradeableBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableBeacon.Contract.UpgradeableBeaconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableBeacon *UpgradeableBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.UpgradeableBeaconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableBeacon *UpgradeableBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.UpgradeableBeaconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableBeacon *UpgradeableBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableBeacon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableBeacon *UpgradeableBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableBeacon *UpgradeableBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _UpgradeableBeacon.contract.Call(opts, &out, "implementation")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconBulkCaller) Implementation(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _UpgradeableBeacon.contract.GenRequest(opts, "implementation")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _UpgradeableBeacon.contract.DecodeOutput(&out, rawOut, "implementation")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconSession) Implementation() (common.Address, error) {
	return _UpgradeableBeacon.Contract.Implementation(&_UpgradeableBeacon.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconCallerSession) Implementation() (common.Address, error) {
	return _UpgradeableBeacon.Contract.Implementation(&_UpgradeableBeacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _UpgradeableBeacon.contract.Call(opts, &out, "owner")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconBulkCaller) Owner(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _UpgradeableBeacon.contract.GenRequest(opts, "owner")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _UpgradeableBeacon.contract.DecodeOutput(&out, rawOut, "owner")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconSession) Owner() (common.Address, error) {
	return _UpgradeableBeacon.Contract.Owner(&_UpgradeableBeacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconCallerSession) Owner() (common.Address, error) {
	return _UpgradeableBeacon.Contract.Owner(&_UpgradeableBeacon.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableBeacon *UpgradeableBeaconBulkTransactor) RenounceOwnership(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _UpgradeableBeacon.contract.GenUnsignedTransaction(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableBeacon *UpgradeableBeaconSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.RenounceOwnership(&_UpgradeableBeacon.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactorSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.RenounceOwnership(&_UpgradeableBeacon.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableBeacon *UpgradeableBeaconBulkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) types.UnsignedTransaction {
	return _UpgradeableBeacon.contract.GenUnsignedTransaction(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableBeacon *UpgradeableBeaconSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.TransferOwnership(&_UpgradeableBeacon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactorSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.TransferOwnership(&_UpgradeableBeacon.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UpgradeableBeacon *UpgradeableBeaconBulkTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) types.UnsignedTransaction {
	return _UpgradeableBeacon.contract.GenUnsignedTransaction(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UpgradeableBeacon *UpgradeableBeaconSession) UpgradeTo(newImplementation common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.UpgradeTo(&_UpgradeableBeacon.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactorSession) UpgradeTo(newImplementation common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _UpgradeableBeacon.Contract.UpgradeTo(&_UpgradeableBeacon.TransactOpts, newImplementation)
}

// UpgradeableBeaconOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UpgradeableBeacon contract.
type UpgradeableBeaconOwnershipTransferredIterator struct {
	Event *UpgradeableBeaconOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UpgradeableBeaconOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableBeaconOwnershipTransferred)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradeableBeaconOwnershipTransferred)
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
func (it *UpgradeableBeaconOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableBeaconOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableBeaconOwnershipTransferred represents a OwnershipTransferred event raised by the UpgradeableBeacon contract.
type UpgradeableBeaconOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// UpgradeableBeaconOwnershipTransferredOrChainReorg represents a OwnershipTransferred subscription event raised by the UpgradeableBeacon contract.
type UpgradeableBeaconOwnershipTransferredOrChainReorg struct {
	Event      *UpgradeableBeaconOwnershipTransferred
	ChainReorg *types.ChainReorg
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UpgradeableBeaconOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, err := _UpgradeableBeacon.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconOwnershipTransferredIterator{contract: _UpgradeableBeacon.contract, event: "OwnershipTransferred", logs: logs}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UpgradeableBeaconOwnershipTransferredOrChainReorg, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradeableBeacon.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableBeaconOwnershipTransferredOrChainReorg)
				event.Event = new(UpgradeableBeaconOwnershipTransferred)

				if log.ChainReorg == nil {
					if err := _UpgradeableBeacon.contract.UnpackLog(event.Event, "OwnershipTransferred", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) ParseOwnershipTransferred(log types.Log) (*UpgradeableBeaconOwnershipTransferred, error) {
	event := new(UpgradeableBeaconOwnershipTransferred)
	if err := _UpgradeableBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableBeaconUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UpgradeableBeacon contract.
type UpgradeableBeaconUpgradedIterator struct {
	Event *UpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableBeaconUpgraded)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradeableBeaconUpgraded)
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
func (it *UpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableBeaconUpgraded represents a Upgraded event raised by the UpgradeableBeacon contract.
type UpgradeableBeaconUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// UpgradeableBeaconUpgradedOrChainReorg represents a Upgraded subscription event raised by the UpgradeableBeacon contract.
type UpgradeableBeaconUpgradedOrChainReorg struct {
	Event      *UpgradeableBeaconUpgraded
	ChainReorg *types.ChainReorg
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UpgradeableBeaconUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, err := _UpgradeableBeacon.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconUpgradedIterator{contract: _UpgradeableBeacon.contract, event: "Upgraded", logs: logs}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UpgradeableBeaconUpgradedOrChainReorg, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableBeacon.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableBeaconUpgradedOrChainReorg)
				event.Event = new(UpgradeableBeaconUpgraded)

				if log.ChainReorg == nil {
					if err := _UpgradeableBeacon.contract.UnpackLog(event.Event, "Upgraded", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) ParseUpgraded(log types.Log) (*UpgradeableBeaconUpgraded, error) {
	event := new(UpgradeableBeaconUpgraded)
	if err := _UpgradeableBeacon.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
