// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openzeppelin

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
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

// BeaconProxyABI is the input ABI used to generate the binding from.
const BeaconProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// BeaconProxyBin is the compiled bytecode used for deploying new contracts.
var BeaconProxyBin = "0x608060405260405162000f3f38038062000f3f8339818101604052810190620000299190620006a6565b6200003d828260006200004560201b60201c565b5050620009c7565b62000056836200013c60201b60201c565b8273ffffffffffffffffffffffffffffffffffffffff167f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e60405160405180910390a2600082511180620000a75750805b156200013757620001358373ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200012391906200070c565b83620002db60201b620000371760201c565b505b505050565b62000152816200031160201b620000641760201c565b62000194576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200018b90620007c5565b60405180910390fd5b6200021b8173ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200020a91906200070c565b6200031160201b620000641760201c565b6200025d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000254906200085d565b60405180910390fd5b80620002977fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6200033460201b620000871760201c565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b606062000309838360405180606001604052806027815260200162000f18602791396200033e60201b60201c565b905092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000819050919050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516200036a9190620008cc565b600060405180830381855af49150503d8060008114620003a7576040519150601f19603f3d011682016040523d82523d6000602084013e620003ac565b606091505b5091509150620003c586838387620003d060201b60201c565b925050509392505050565b60608315620004405760008351036200043757620003f4856200031160201b60201c565b62000436576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200042d9062000935565b60405180910390fd5b5b82905062000453565b6200045283836200045b60201b60201c565b5b949350505050565b6000825111156200046f5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620004a59190620009a3565b60405180910390fd5b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620004ef82620004c2565b9050919050565b6200050181620004e2565b81146200050d57600080fd5b50565b6000815190506200052181620004f6565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200057c8262000531565b810181811067ffffffffffffffff821117156200059e576200059d62000542565b5b80604052505050565b6000620005b3620004ae565b9050620005c1828262000571565b919050565b600067ffffffffffffffff821115620005e457620005e362000542565b5b620005ef8262000531565b9050602081019050919050565b60005b838110156200061c578082015181840152602081019050620005ff565b60008484015250505050565b60006200063f6200063984620005c6565b620005a7565b9050828152602081018484840111156200065e576200065d6200052c565b5b6200066b848285620005fc565b509392505050565b600082601f8301126200068b576200068a62000527565b5b81516200069d84826020860162000628565b91505092915050565b60008060408385031215620006c057620006bf620004b8565b5b6000620006d08582860162000510565b925050602083015167ffffffffffffffff811115620006f457620006f3620004bd565b5b620007028582860162000673565b9150509250929050565b600060208284031215620007255762000724620004b8565b5b6000620007358482850162000510565b91505092915050565b600082825260208201905092915050565b7f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e60008201527f7472616374000000000000000000000000000000000000000000000000000000602082015250565b6000620007ad6025836200073e565b9150620007ba826200074f565b604082019050919050565b60006020820190508181036000830152620007e0816200079e565b9050919050565b7f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960008201527f73206e6f74206120636f6e747261637400000000000000000000000000000000602082015250565b6000620008456030836200073e565b91506200085282620007e7565b604082019050919050565b60006020820190508181036000830152620008788162000836565b9050919050565b600081519050919050565b600081905092915050565b6000620008a2826200087f565b620008ae81856200088a565b9350620008c0818560208601620005fc565b80840191505092915050565b6000620008da828462000895565b915081905092915050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b60006200091d601d836200073e565b91506200092a82620008e5565b602082019050919050565b6000602082019050818103600083015262000950816200090e565b9050919050565b600081519050919050565b60006200096f8262000957565b6200097b81856200073e565b93506200098d818560208601620005fc565b620009988162000531565b840191505092915050565b60006020820190508181036000830152620009bf818462000962565b905092915050565b61054180620009d76000396000f3fe6080604052366100135761001161001d565b005b61001b61001d565b005b610025610091565b610035610030610093565b610110565b565b606061005c83836040518060600160405280602781526020016104e560279139610136565b905092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000819050919050565b565b600061009d6101bc565b73ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061010b919061033b565b905090565b3660008037600080366000845af43d6000803e8060008114610131573d6000f35b3d6000fd5b60606000808573ffffffffffffffffffffffffffffffffffffffff168560405161016091906103d9565b600060405180830381855af49150503d806000811461019b576040519150601f19603f3d011682016040523d82523d6000602084013e6101a0565b606091505b50915091506101b186838387610213565b925050509392505050565b60006101ea7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b610087565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060831561027557600083510361026d5761022d85610064565b61026c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102639061044d565b60405180910390fd5b5b829050610280565b61027f8383610288565b5b949350505050565b60008251111561029b5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102cf91906104c2565b60405180910390fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610308826102dd565b9050919050565b610318816102fd565b811461032357600080fd5b50565b6000815190506103358161030f565b92915050565b600060208284031215610351576103506102d8565b5b600061035f84828501610326565b91505092915050565b600081519050919050565b600081905092915050565b60005b8381101561039c578082015181840152602081019050610381565b60008484015250505050565b60006103b382610368565b6103bd8185610373565b93506103cd81856020860161037e565b80840191505092915050565b60006103e582846103a8565b915081905092915050565b600082825260208201905092915050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b6000610437601d836103f0565b915061044282610401565b602082019050919050565b600060208201905081810360008301526104668161042a565b9050919050565b600081519050919050565b6000601f19601f8301169050919050565b60006104948261046d565b61049e81856103f0565b93506104ae81856020860161037e565b6104b781610478565b840191505092915050565b600060208201905081810360008301526104dc8184610489565b90509291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e2d3045857e1cab9624680014fb7a7a0406fa90030c0970e6d8486d66eb4fcd764736f6c63430008120033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564"

// DeployBeaconProxy deploys a new Conflux contract, binding an instance of BeaconProxy to it.
func DeployBeaconProxy(auth *bind.TransactOpts, backend bind.ContractBackend, beacon common.Address, data []byte) (*types.UnsignedTransaction, *types.Hash, *BeaconProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(BeaconProxyABI))
	if err != nil {
		return nil, nil, nil, err
	}

	tx, hash, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BeaconProxyBin), backend, beacon, data)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, hash, &BeaconProxy{BeaconProxyCaller: BeaconProxyCaller{contract: contract}, BeaconProxyTransactor: BeaconProxyTransactor{contract: contract}, BeaconProxyFilterer: BeaconProxyFilterer{contract: contract}}, nil
}

// BeaconProxy is an auto generated Go binding around an Conflux contract.
type BeaconProxy struct {
	BeaconProxyCaller     // Read-only binding to the contract
	BeaconProxyTransactor // Write-only binding to the contract
	BeaconProxyFilterer   // Log filterer for contract events
}

// BeaconProxyCaller is an auto generated read-only Go binding around an Conflux contract.
type BeaconProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconProxyBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type BeaconProxyBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconProxyTransactor is an auto generated write-only Go binding around an Conflux contract.
type BeaconProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconProxyBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type BeaconProxyBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconProxyFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type BeaconProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconProxySession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type BeaconProxySession struct {
	Contract     *BeaconProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeaconProxyCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type BeaconProxyCallerSession struct {
	Contract *BeaconProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BeaconProxyTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type BeaconProxyTransactorSession struct {
	Contract     *BeaconProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BeaconProxyRaw is an auto generated low-level Go binding around an Conflux contract.
type BeaconProxyRaw struct {
	Contract *BeaconProxy // Generic contract binding to access the raw methods on
}

// BeaconProxyCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type BeaconProxyCallerRaw struct {
	Contract *BeaconProxyCaller // Generic read-only contract binding to access the raw methods on
}

// BeaconProxyTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type BeaconProxyTransactorRaw struct {
	Contract *BeaconProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeaconProxy creates a new instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxy(address types.Address, backend bind.ContractBackend) (*BeaconProxy, error) {
	contract, err := bindBeaconProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeaconProxy{BeaconProxyCaller: BeaconProxyCaller{contract: contract}, BeaconProxyTransactor: BeaconProxyTransactor{contract: contract}, BeaconProxyFilterer: BeaconProxyFilterer{contract: contract}}, nil
}

// NewBeaconProxyCaller creates a new read-only instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxyCaller(address types.Address, caller bind.ContractCaller) (*BeaconProxyCaller, error) {
	contract, err := bindBeaconProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyCaller{contract: contract}, nil
}

// NewBeaconProxyTransactor creates a new write-only instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxyTransactor(address types.Address, transactor bind.ContractTransactor) (*BeaconProxyTransactor, error) {
	contract, err := bindBeaconProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyTransactor{contract: contract}, nil
}

// NewBeaconProxyFilterer creates a new log filterer instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxyFilterer(address types.Address, filterer bind.ContractFilterer) (*BeaconProxyFilterer, error) {
	contract, err := bindBeaconProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyFilterer{contract: contract}, nil
}

// NewBeaconProxyCaller creates a new read-only instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxyBulkCaller(address types.Address, caller bind.ContractCaller) (*BeaconProxyBulkCaller, error) {
	contract, err := bindBeaconProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyBulkCaller{contract: contract}, nil
}

// NewBeaconProxyBulkTransactor creates a new write-only instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxyBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*BeaconProxyBulkTransactor, error) {
	contract, err := bindBeaconProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyBulkTransactor{contract: contract}, nil
}

// bindBeaconProxy binds a generic wrapper to an already deployed contract.
func bindBeaconProxy(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BeaconProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconProxy *BeaconProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconProxy.Contract.BeaconProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconProxy *BeaconProxyRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.Contract.BeaconProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconProxy *BeaconProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.Contract.BeaconProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconProxy *BeaconProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconProxy *BeaconProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconProxy *BeaconProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BeaconProxy *BeaconProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BeaconProxy *BeaconProxySession) Fallback(calldata []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.Contract.Fallback(&_BeaconProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BeaconProxy *BeaconProxyTransactorSession) Fallback(calldata []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.Contract.Fallback(&_BeaconProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BeaconProxy *BeaconProxyTransactor) Receive(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BeaconProxy *BeaconProxySession) Receive() (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.Contract.Receive(&_BeaconProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BeaconProxy *BeaconProxyTransactorSession) Receive() (*types.UnsignedTransaction, *types.Hash, error) {
	return _BeaconProxy.Contract.Receive(&_BeaconProxy.TransactOpts)
}

// BeaconProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BeaconProxy contract.
type BeaconProxyAdminChangedIterator struct {
	Event *BeaconProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *BeaconProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconProxyAdminChanged)
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
		it.Event = new(BeaconProxyAdminChanged)
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
func (it *BeaconProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconProxyAdminChanged represents a AdminChanged event raised by the BeaconProxy contract.
type BeaconProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// BeaconProxyAdminChangedOrChainReorg represents a AdminChanged subscription event raised by the BeaconProxy contract.
type BeaconProxyAdminChangedOrChainReorg struct {
	Event      *BeaconProxyAdminChanged
	ChainReorg *types.ChainReorg
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BeaconProxy *BeaconProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BeaconProxyAdminChangedIterator, error) {

	logs, err := _BeaconProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BeaconProxyAdminChangedIterator{contract: _BeaconProxy.contract, event: "AdminChanged", logs: logs}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BeaconProxy *BeaconProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BeaconProxyAdminChangedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _BeaconProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconProxyAdminChangedOrChainReorg)
				event.Event = new(BeaconProxyAdminChanged)

				if log.ChainReorg == nil {
					if err := _BeaconProxy.contract.UnpackLog(event.Event, "AdminChanged", *log.Log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BeaconProxy *BeaconProxyFilterer) ParseAdminChanged(log types.Log) (*BeaconProxyAdminChanged, error) {
	event := new(BeaconProxyAdminChanged)
	if err := _BeaconProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeaconProxyBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BeaconProxy contract.
type BeaconProxyBeaconUpgradedIterator struct {
	Event *BeaconProxyBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BeaconProxyBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconProxyBeaconUpgraded)
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
		it.Event = new(BeaconProxyBeaconUpgraded)
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
func (it *BeaconProxyBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconProxyBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconProxyBeaconUpgraded represents a BeaconUpgraded event raised by the BeaconProxy contract.
type BeaconProxyBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// BeaconProxyBeaconUpgradedOrChainReorg represents a BeaconUpgraded subscription event raised by the BeaconProxy contract.
type BeaconProxyBeaconUpgradedOrChainReorg struct {
	Event      *BeaconProxyBeaconUpgraded
	ChainReorg *types.ChainReorg
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BeaconProxy *BeaconProxyFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BeaconProxyBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, err := _BeaconProxy.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyBeaconUpgradedIterator{contract: _BeaconProxy.contract, event: "BeaconUpgraded", logs: logs}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BeaconProxy *BeaconProxyFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BeaconProxyBeaconUpgradedOrChainReorg, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BeaconProxy.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconProxyBeaconUpgradedOrChainReorg)
				event.Event = new(BeaconProxyBeaconUpgraded)

				if log.ChainReorg == nil {
					if err := _BeaconProxy.contract.UnpackLog(event.Event, "BeaconUpgraded", *log.Log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BeaconProxy *BeaconProxyFilterer) ParseBeaconUpgraded(log types.Log) (*BeaconProxyBeaconUpgraded, error) {
	event := new(BeaconProxyBeaconUpgraded)
	if err := _BeaconProxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeaconProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BeaconProxy contract.
type BeaconProxyUpgradedIterator struct {
	Event *BeaconProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *BeaconProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconProxyUpgraded)
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
		it.Event = new(BeaconProxyUpgraded)
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
func (it *BeaconProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconProxyUpgraded represents a Upgraded event raised by the BeaconProxy contract.
type BeaconProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// BeaconProxyUpgradedOrChainReorg represents a Upgraded subscription event raised by the BeaconProxy contract.
type BeaconProxyUpgradedOrChainReorg struct {
	Event      *BeaconProxyUpgraded
	ChainReorg *types.ChainReorg
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeaconProxy *BeaconProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BeaconProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, err := _BeaconProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyUpgradedIterator{contract: _BeaconProxy.contract, event: "Upgraded", logs: logs}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeaconProxy *BeaconProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BeaconProxyUpgradedOrChainReorg, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BeaconProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconProxyUpgradedOrChainReorg)
				event.Event = new(BeaconProxyUpgraded)

				if log.ChainReorg == nil {
					if err := _BeaconProxy.contract.UnpackLog(event.Event, "Upgraded", *log.Log); err != nil {
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
func (_BeaconProxy *BeaconProxyFilterer) ParseUpgraded(log types.Log) (*BeaconProxyUpgraded, error) {
	event := new(BeaconProxyUpgraded)
	if err := _BeaconProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
