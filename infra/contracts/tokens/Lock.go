// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// LockABI is the input ABI used to generate the binding from.
const LockABI = "[{\"inputs\":[],\"name\":\"test\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LockBin is the compiled bytecode used for deploying new contracts.
var LockBin = "0x608060405234801561001057600080fd5b5060b58061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063f8a8fd6d14602d575b600080fd5b60336047565b604051603e91906066565b60405180910390f35b600042905090565b6000819050919050565b606081604f565b82525050565b6000602082019050607960008301846059565b9291505056fea26469706673582212207c78c6bb148e0beb1b8f30a042b4d2723dbd26ec1566e2804a12015a5e034f1c64736f6c63430008120033"

// DeployLock deploys a new Conflux contract, binding an instance of Lock to it.
func DeployLock(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.UnsignedTransaction, *types.Hash, *Lock, error) {
	parsed, err := abi.JSON(strings.NewReader(LockABI))
	if err != nil {
		return nil, nil, nil, err
	}

	tx, hash, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockBin), backend)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, hash, &Lock{LockCaller: LockCaller{contract: contract}, LockTransactor: LockTransactor{contract: contract}, LockFilterer: LockFilterer{contract: contract}}, nil
}

// Lock is an auto generated Go binding around an Conflux contract.
type Lock struct {
	LockCaller     // Read-only binding to the contract
	LockTransactor // Write-only binding to the contract
	LockFilterer   // Log filterer for contract events
}

// LockCaller is an auto generated read-only Go binding around an Conflux contract.
type LockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type LockBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockTransactor is an auto generated write-only Go binding around an Conflux contract.
type LockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type LockBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type LockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type LockSession struct {
	Contract     *Lock             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type LockCallerSession struct {
	Contract *LockCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LockTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type LockTransactorSession struct {
	Contract     *LockTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockRaw is an auto generated low-level Go binding around an Conflux contract.
type LockRaw struct {
	Contract *Lock // Generic contract binding to access the raw methods on
}

// LockCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type LockCallerRaw struct {
	Contract *LockCaller // Generic read-only contract binding to access the raw methods on
}

// LockTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type LockTransactorRaw struct {
	Contract *LockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLock creates a new instance of Lock, bound to a specific deployed contract.
func NewLock(address types.Address, backend bind.ContractBackend) (*Lock, error) {
	contract, err := bindLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lock{LockCaller: LockCaller{contract: contract}, LockTransactor: LockTransactor{contract: contract}, LockFilterer: LockFilterer{contract: contract}}, nil
}

// NewLockCaller creates a new read-only instance of Lock, bound to a specific deployed contract.
func NewLockCaller(address types.Address, caller bind.ContractCaller) (*LockCaller, error) {
	contract, err := bindLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockCaller{contract: contract}, nil
}

// NewLockTransactor creates a new write-only instance of Lock, bound to a specific deployed contract.
func NewLockTransactor(address types.Address, transactor bind.ContractTransactor) (*LockTransactor, error) {
	contract, err := bindLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockTransactor{contract: contract}, nil
}

// NewLockFilterer creates a new log filterer instance of Lock, bound to a specific deployed contract.
func NewLockFilterer(address types.Address, filterer bind.ContractFilterer) (*LockFilterer, error) {
	contract, err := bindLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockFilterer{contract: contract}, nil
}

// NewLockCaller creates a new read-only instance of Lock, bound to a specific deployed contract.
func NewLockBulkCaller(address types.Address, caller bind.ContractCaller) (*LockBulkCaller, error) {
	contract, err := bindLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockBulkCaller{contract: contract}, nil
}

// NewLockBulkTransactor creates a new write-only instance of Lock, bound to a specific deployed contract.
func NewLockBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*LockBulkTransactor, error) {
	contract, err := bindLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockBulkTransactor{contract: contract}, nil
}

// bindLock binds a generic wrapper to an already deployed contract.
func bindLock(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lock *LockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lock.Contract.LockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lock *LockRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Lock.Contract.LockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lock *LockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Lock.Contract.LockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lock *LockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lock *LockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Lock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lock *LockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Lock.Contract.contract.Transact(opts, method, params...)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Lock *LockCaller) Test(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _Lock.contract.Call(opts, &out, "test")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Lock *LockBulkCaller) Test(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Lock.contract.GenRequest(opts, "test")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Lock.contract.DecodeOutput(&out, rawOut, "test")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Lock *LockSession) Test() (*big.Int, error) {
	return _Lock.Contract.Test(&_Lock.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_Lock *LockCallerSession) Test() (*big.Int, error) {
	return _Lock.Contract.Test(&_Lock.CallOpts)
}
