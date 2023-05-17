// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

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

// RouterABI is the input ABI used to generate the binding from.
const RouterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localContract\",\"type\":\"address\"}],\"name\":\"ArrivalConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"local\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"remoteContract\",\"type\":\"address\"}],\"name\":\"DepartureConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"local\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"}],\"name\":\"DepartureRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"remoteContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumRouter.OP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"enumRouter.URI_MODE\",\"name\":\"uriMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"localContract\",\"type\":\"address\"}],\"name\":\"registerArrival\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"local\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumRouter.OP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"enumRouter.URI_MODE\",\"name\":\"uriMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"remoteContract\",\"type\":\"address\"}],\"name\":\"registerDeparture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Router *RouterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Router *RouterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Router.Contract.DEFAULTADMINROLE(&_Router.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Router *RouterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Router.Contract.DEFAULTADMINROLE(&_Router.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Router *RouterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Router *RouterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Router.Contract.GetRoleAdmin(&_Router.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Router *RouterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Router.Contract.GetRoleAdmin(&_Router.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Router *RouterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Router *RouterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Router.Contract.HasRole(&_Router.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Router *RouterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Router.Contract.HasRole(&_Router.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Router *RouterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Router *RouterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Router.Contract.SupportsInterface(&_Router.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Router *RouterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Router.Contract.SupportsInterface(&_Router.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Router *RouterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Router *RouterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.Contract.GrantRole(&_Router.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Router *RouterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.Contract.GrantRole(&_Router.TransactOpts, role, account)
}

// RegisterArrival is a paid mutator transaction binding the contract method 0x364f2de0.
//
// Solidity: function registerArrival(address remoteContract, uint256 remoteChainId, uint8 op, uint8 uriMode, address localContract) returns()
func (_Router *RouterTransactor) RegisterArrival(opts *bind.TransactOpts, remoteContract common.Address, remoteChainId *big.Int, op uint8, uriMode uint8, localContract common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "registerArrival", remoteContract, remoteChainId, op, uriMode, localContract)
}

// RegisterArrival is a paid mutator transaction binding the contract method 0x364f2de0.
//
// Solidity: function registerArrival(address remoteContract, uint256 remoteChainId, uint8 op, uint8 uriMode, address localContract) returns()
func (_Router *RouterSession) RegisterArrival(remoteContract common.Address, remoteChainId *big.Int, op uint8, uriMode uint8, localContract common.Address) (*types.Transaction, error) {
	return _Router.Contract.RegisterArrival(&_Router.TransactOpts, remoteContract, remoteChainId, op, uriMode, localContract)
}

// RegisterArrival is a paid mutator transaction binding the contract method 0x364f2de0.
//
// Solidity: function registerArrival(address remoteContract, uint256 remoteChainId, uint8 op, uint8 uriMode, address localContract) returns()
func (_Router *RouterTransactorSession) RegisterArrival(remoteContract common.Address, remoteChainId *big.Int, op uint8, uriMode uint8, localContract common.Address) (*types.Transaction, error) {
	return _Router.Contract.RegisterArrival(&_Router.TransactOpts, remoteContract, remoteChainId, op, uriMode, localContract)
}

// RegisterDeparture is a paid mutator transaction binding the contract method 0x03cda6e0.
//
// Solidity: function registerDeparture(address local, uint256 targetChainId, uint8 op, uint8 uriMode, address remoteContract) returns()
func (_Router *RouterTransactor) RegisterDeparture(opts *bind.TransactOpts, local common.Address, targetChainId *big.Int, op uint8, uriMode uint8, remoteContract common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "registerDeparture", local, targetChainId, op, uriMode, remoteContract)
}

// RegisterDeparture is a paid mutator transaction binding the contract method 0x03cda6e0.
//
// Solidity: function registerDeparture(address local, uint256 targetChainId, uint8 op, uint8 uriMode, address remoteContract) returns()
func (_Router *RouterSession) RegisterDeparture(local common.Address, targetChainId *big.Int, op uint8, uriMode uint8, remoteContract common.Address) (*types.Transaction, error) {
	return _Router.Contract.RegisterDeparture(&_Router.TransactOpts, local, targetChainId, op, uriMode, remoteContract)
}

// RegisterDeparture is a paid mutator transaction binding the contract method 0x03cda6e0.
//
// Solidity: function registerDeparture(address local, uint256 targetChainId, uint8 op, uint8 uriMode, address remoteContract) returns()
func (_Router *RouterTransactorSession) RegisterDeparture(local common.Address, targetChainId *big.Int, op uint8, uriMode uint8, remoteContract common.Address) (*types.Transaction, error) {
	return _Router.Contract.RegisterDeparture(&_Router.TransactOpts, local, targetChainId, op, uriMode, remoteContract)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Router *RouterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Router *RouterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.Contract.RenounceRole(&_Router.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Router *RouterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.Contract.RenounceRole(&_Router.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Router *RouterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Router *RouterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.Contract.RevokeRole(&_Router.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Router *RouterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Router.Contract.RevokeRole(&_Router.TransactOpts, role, account)
}

// RouterArrivalConfiguredIterator is returned from FilterArrivalConfigured and is used to iterate over the raw logs and unpacked data for ArrivalConfigured events raised by the Router contract.
type RouterArrivalConfiguredIterator struct {
	Event *RouterArrivalConfigured // Event containing the contract specifics and raw log

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
func (it *RouterArrivalConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterArrivalConfigured)
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
		it.Event = new(RouterArrivalConfigured)
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
func (it *RouterArrivalConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterArrivalConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterArrivalConfigured represents a ArrivalConfigured event raised by the Router contract.
type RouterArrivalConfigured struct {
	RemoteContract common.Address
	RemoteChainId  *big.Int
	LocalContract  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterArrivalConfigured is a free log retrieval operation binding the contract event 0xe9fb7d7e00118c8e027ac046db156d6e48ef637bd79b3d7fa30708d7b3b5e58b.
//
// Solidity: event ArrivalConfigured(address indexed remoteContract, uint256 indexed remoteChainId, address localContract)
func (_Router *RouterFilterer) FilterArrivalConfigured(opts *bind.FilterOpts, remoteContract []common.Address, remoteChainId []*big.Int) (*RouterArrivalConfiguredIterator, error) {

	var remoteContractRule []interface{}
	for _, remoteContractItem := range remoteContract {
		remoteContractRule = append(remoteContractRule, remoteContractItem)
	}
	var remoteChainIdRule []interface{}
	for _, remoteChainIdItem := range remoteChainId {
		remoteChainIdRule = append(remoteChainIdRule, remoteChainIdItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "ArrivalConfigured", remoteContractRule, remoteChainIdRule)
	if err != nil {
		return nil, err
	}
	return &RouterArrivalConfiguredIterator{contract: _Router.contract, event: "ArrivalConfigured", logs: logs, sub: sub}, nil
}

// WatchArrivalConfigured is a free log subscription operation binding the contract event 0xe9fb7d7e00118c8e027ac046db156d6e48ef637bd79b3d7fa30708d7b3b5e58b.
//
// Solidity: event ArrivalConfigured(address indexed remoteContract, uint256 indexed remoteChainId, address localContract)
func (_Router *RouterFilterer) WatchArrivalConfigured(opts *bind.WatchOpts, sink chan<- *RouterArrivalConfigured, remoteContract []common.Address, remoteChainId []*big.Int) (event.Subscription, error) {

	var remoteContractRule []interface{}
	for _, remoteContractItem := range remoteContract {
		remoteContractRule = append(remoteContractRule, remoteContractItem)
	}
	var remoteChainIdRule []interface{}
	for _, remoteChainIdItem := range remoteChainId {
		remoteChainIdRule = append(remoteChainIdRule, remoteChainIdItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "ArrivalConfigured", remoteContractRule, remoteChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterArrivalConfigured)
				if err := _Router.contract.UnpackLog(event, "ArrivalConfigured", log); err != nil {
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
func (_Router *RouterFilterer) ParseArrivalConfigured(log types.Log) (*RouterArrivalConfigured, error) {
	event := new(RouterArrivalConfigured)
	if err := _Router.contract.UnpackLog(event, "ArrivalConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterDepartureConfiguredIterator is returned from FilterDepartureConfigured and is used to iterate over the raw logs and unpacked data for DepartureConfigured events raised by the Router contract.
type RouterDepartureConfiguredIterator struct {
	Event *RouterDepartureConfigured // Event containing the contract specifics and raw log

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
func (it *RouterDepartureConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterDepartureConfigured)
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
		it.Event = new(RouterDepartureConfigured)
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
func (it *RouterDepartureConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterDepartureConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterDepartureConfigured represents a DepartureConfigured event raised by the Router contract.
type RouterDepartureConfigured struct {
	Local          common.Address
	TargetChainId  *big.Int
	RemoteContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDepartureConfigured is a free log retrieval operation binding the contract event 0x0e5a0f5238edad743f0bbb7140294ec2214d83e62fc9a2816d3f79163981cb53.
//
// Solidity: event DepartureConfigured(address indexed local, uint256 indexed targetChainId, address remoteContract)
func (_Router *RouterFilterer) FilterDepartureConfigured(opts *bind.FilterOpts, local []common.Address, targetChainId []*big.Int) (*RouterDepartureConfiguredIterator, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "DepartureConfigured", localRule, targetChainIdRule)
	if err != nil {
		return nil, err
	}
	return &RouterDepartureConfiguredIterator{contract: _Router.contract, event: "DepartureConfigured", logs: logs, sub: sub}, nil
}

// WatchDepartureConfigured is a free log subscription operation binding the contract event 0x0e5a0f5238edad743f0bbb7140294ec2214d83e62fc9a2816d3f79163981cb53.
//
// Solidity: event DepartureConfigured(address indexed local, uint256 indexed targetChainId, address remoteContract)
func (_Router *RouterFilterer) WatchDepartureConfigured(opts *bind.WatchOpts, sink chan<- *RouterDepartureConfigured, local []common.Address, targetChainId []*big.Int) (event.Subscription, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "DepartureConfigured", localRule, targetChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterDepartureConfigured)
				if err := _Router.contract.UnpackLog(event, "DepartureConfigured", log); err != nil {
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
func (_Router *RouterFilterer) ParseDepartureConfigured(log types.Log) (*RouterDepartureConfigured, error) {
	event := new(RouterDepartureConfigured)
	if err := _Router.contract.UnpackLog(event, "DepartureConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterDepartureRemovedIterator is returned from FilterDepartureRemoved and is used to iterate over the raw logs and unpacked data for DepartureRemoved events raised by the Router contract.
type RouterDepartureRemovedIterator struct {
	Event *RouterDepartureRemoved // Event containing the contract specifics and raw log

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
func (it *RouterDepartureRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterDepartureRemoved)
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
		it.Event = new(RouterDepartureRemoved)
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
func (it *RouterDepartureRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterDepartureRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterDepartureRemoved represents a DepartureRemoved event raised by the Router contract.
type RouterDepartureRemoved struct {
	Local         common.Address
	TargetChainId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepartureRemoved is a free log retrieval operation binding the contract event 0x5359a4d931e34226aa199102a165152f231215dd06fd7900d8739037e764e2d1.
//
// Solidity: event DepartureRemoved(address indexed local, uint256 indexed targetChainId)
func (_Router *RouterFilterer) FilterDepartureRemoved(opts *bind.FilterOpts, local []common.Address, targetChainId []*big.Int) (*RouterDepartureRemovedIterator, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "DepartureRemoved", localRule, targetChainIdRule)
	if err != nil {
		return nil, err
	}
	return &RouterDepartureRemovedIterator{contract: _Router.contract, event: "DepartureRemoved", logs: logs, sub: sub}, nil
}

// WatchDepartureRemoved is a free log subscription operation binding the contract event 0x5359a4d931e34226aa199102a165152f231215dd06fd7900d8739037e764e2d1.
//
// Solidity: event DepartureRemoved(address indexed local, uint256 indexed targetChainId)
func (_Router *RouterFilterer) WatchDepartureRemoved(opts *bind.WatchOpts, sink chan<- *RouterDepartureRemoved, local []common.Address, targetChainId []*big.Int) (event.Subscription, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "DepartureRemoved", localRule, targetChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterDepartureRemoved)
				if err := _Router.contract.UnpackLog(event, "DepartureRemoved", log); err != nil {
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

// ParseDepartureRemoved is a log parse operation binding the contract event 0x5359a4d931e34226aa199102a165152f231215dd06fd7900d8739037e764e2d1.
//
// Solidity: event DepartureRemoved(address indexed local, uint256 indexed targetChainId)
func (_Router *RouterFilterer) ParseDepartureRemoved(log types.Log) (*RouterDepartureRemoved, error) {
	event := new(RouterDepartureRemoved)
	if err := _Router.contract.UnpackLog(event, "DepartureRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Router contract.
type RouterRoleAdminChangedIterator struct {
	Event *RouterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RouterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRoleAdminChanged)
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
		it.Event = new(RouterRoleAdminChanged)
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
func (it *RouterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRoleAdminChanged represents a RoleAdminChanged event raised by the Router contract.
type RouterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Router *RouterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RouterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Router.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RouterRoleAdminChangedIterator{contract: _Router.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Router *RouterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RouterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Router.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRoleAdminChanged)
				if err := _Router.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Router *RouterFilterer) ParseRoleAdminChanged(log types.Log) (*RouterRoleAdminChanged, error) {
	event := new(RouterRoleAdminChanged)
	if err := _Router.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Router contract.
type RouterRoleGrantedIterator struct {
	Event *RouterRoleGranted // Event containing the contract specifics and raw log

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
func (it *RouterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRoleGranted)
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
		it.Event = new(RouterRoleGranted)
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
func (it *RouterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRoleGranted represents a RoleGranted event raised by the Router contract.
type RouterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Router *RouterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RouterRoleGrantedIterator, error) {

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

	logs, sub, err := _Router.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RouterRoleGrantedIterator{contract: _Router.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Router *RouterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RouterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Router.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRoleGranted)
				if err := _Router.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Router *RouterFilterer) ParseRoleGranted(log types.Log) (*RouterRoleGranted, error) {
	event := new(RouterRoleGranted)
	if err := _Router.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Router contract.
type RouterRoleRevokedIterator struct {
	Event *RouterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RouterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRoleRevoked)
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
		it.Event = new(RouterRoleRevoked)
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
func (it *RouterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRoleRevoked represents a RoleRevoked event raised by the Router contract.
type RouterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Router *RouterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RouterRoleRevokedIterator, error) {

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

	logs, sub, err := _Router.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RouterRoleRevokedIterator{contract: _Router.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Router *RouterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RouterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Router.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRoleRevoked)
				if err := _Router.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Router *RouterFilterer) ParseRoleRevoked(log types.Log) (*RouterRoleRevoked, error) {
	event := new(RouterRoleRevoked)
	if err := _Router.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
