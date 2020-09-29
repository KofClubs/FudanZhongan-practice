// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package view

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

// ViewABI is the input ABI used to generate the binding from.
const ViewABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getStr\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// View is an auto generated Go binding around an Ethereum contract.
type View struct {
	ViewCaller     // Read-only binding to the contract
	ViewTransactor // Write-only binding to the contract
	ViewFilterer   // Log filterer for contract events
}

// ViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type ViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ViewSession struct {
	Contract     *View             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ViewCallerSession struct {
	Contract *ViewCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ViewTransactorSession struct {
	Contract     *ViewTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type ViewRaw struct {
	Contract *View // Generic contract binding to access the raw methods on
}

// ViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ViewCallerRaw struct {
	Contract *ViewCaller // Generic read-only contract binding to access the raw methods on
}

// ViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ViewTransactorRaw struct {
	Contract *ViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewView creates a new instance of View, bound to a specific deployed contract.
func NewView(address common.Address, backend bind.ContractBackend) (*View, error) {
	contract, err := bindView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &View{ViewCaller: ViewCaller{contract: contract}, ViewTransactor: ViewTransactor{contract: contract}, ViewFilterer: ViewFilterer{contract: contract}}, nil
}

// NewViewCaller creates a new read-only instance of View, bound to a specific deployed contract.
func NewViewCaller(address common.Address, caller bind.ContractCaller) (*ViewCaller, error) {
	contract, err := bindView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ViewCaller{contract: contract}, nil
}

// NewViewTransactor creates a new write-only instance of View, bound to a specific deployed contract.
func NewViewTransactor(address common.Address, transactor bind.ContractTransactor) (*ViewTransactor, error) {
	contract, err := bindView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ViewTransactor{contract: contract}, nil
}

// NewViewFilterer creates a new log filterer instance of View, bound to a specific deployed contract.
func NewViewFilterer(address common.Address, filterer bind.ContractFilterer) (*ViewFilterer, error) {
	contract, err := bindView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ViewFilterer{contract: contract}, nil
}

// bindView binds a generic wrapper to an already deployed contract.
func bindView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ViewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_View *ViewRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _View.Contract.ViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_View *ViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _View.Contract.ViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_View *ViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _View.Contract.ViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_View *ViewCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _View.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_View *ViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _View.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_View *ViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _View.Contract.contract.Transact(opts, method, params...)
}

// GetStr is a free data retrieval call binding the contract method 0xb8c9e4ed.
//
// Solidity: function getStr() view returns(string)
func (_View *ViewCaller) GetStr(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _View.contract.Call(opts, out, "getStr")
	return *ret0, err
}

// GetStr is a free data retrieval call binding the contract method 0xb8c9e4ed.
//
// Solidity: function getStr() view returns(string)
func (_View *ViewSession) GetStr() (string, error) {
	return _View.Contract.GetStr(&_View.CallOpts)
}

// GetStr is a free data retrieval call binding the contract method 0xb8c9e4ed.
//
// Solidity: function getStr() view returns(string)
func (_View *ViewCallerSession) GetStr() (string, error) {
	return _View.Contract.GetStr(&_View.CallOpts)
}
