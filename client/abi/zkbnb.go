// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ZkBNBMetaData contains all meta data concerning the ZkBNB contract.
var ZkBNBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBEP20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_nftL1Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_pubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_pubKeyY\",\"type\":\"bytes32\"}],\"name\":\"registerZNS\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"requestFullExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_nftIndex\",\"type\":\"uint32\"}],\"name\":\"requestFullExitNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZkBNBABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkBNBMetaData.ABI instead.
var ZkBNBABI = ZkBNBMetaData.ABI

// ZkBNB is an auto generated Go binding around an Ethereum contract.
type ZkBNB struct {
	ZkBNBCaller     // Read-only binding to the contract
	ZkBNBTransactor // Write-only binding to the contract
	ZkBNBFilterer   // Log filterer for contract events
}

// ZkBNBCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkBNBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBNBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkBNBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBNBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkBNBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBNBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkBNBSession struct {
	Contract     *ZkBNB            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkBNBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkBNBCallerSession struct {
	Contract *ZkBNBCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZkBNBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkBNBTransactorSession struct {
	Contract     *ZkBNBTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkBNBRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkBNBRaw struct {
	Contract *ZkBNB // Generic contract binding to access the raw methods on
}

// ZkBNBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkBNBCallerRaw struct {
	Contract *ZkBNBCaller // Generic read-only contract binding to access the raw methods on
}

// ZkBNBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkBNBTransactorRaw struct {
	Contract *ZkBNBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkBNB creates a new instance of ZkBNB, bound to a specific deployed contract.
func NewZkBNB(address common.Address, backend bind.ContractBackend) (*ZkBNB, error) {
	contract, err := bindZkBNB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkBNB{ZkBNBCaller: ZkBNBCaller{contract: contract}, ZkBNBTransactor: ZkBNBTransactor{contract: contract}, ZkBNBFilterer: ZkBNBFilterer{contract: contract}}, nil
}

// NewZkBNBCaller creates a new read-only instance of ZkBNB, bound to a specific deployed contract.
func NewZkBNBCaller(address common.Address, caller bind.ContractCaller) (*ZkBNBCaller, error) {
	contract, err := bindZkBNB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkBNBCaller{contract: contract}, nil
}

// NewZkBNBTransactor creates a new write-only instance of ZkBNB, bound to a specific deployed contract.
func NewZkBNBTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkBNBTransactor, error) {
	contract, err := bindZkBNB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkBNBTransactor{contract: contract}, nil
}

// NewZkBNBFilterer creates a new log filterer instance of ZkBNB, bound to a specific deployed contract.
func NewZkBNBFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkBNBFilterer, error) {
	contract, err := bindZkBNB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkBNBFilterer{contract: contract}, nil
}

// bindZkBNB binds a generic wrapper to an already deployed contract.
func bindZkBNB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZkBNBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkBNB *ZkBNBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkBNB.Contract.ZkBNBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkBNB *ZkBNBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.Contract.ZkBNBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkBNB *ZkBNBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkBNB.Contract.ZkBNBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkBNB *ZkBNBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkBNB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkBNB *ZkBNBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkBNB *ZkBNBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkBNB.Contract.contract.Transact(opts, method, params...)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_ZkBNB *ZkBNBTransactor) DepositBEP20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositBEP20", _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_ZkBNB *ZkBNBSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBEP20(&_ZkBNB.TransactOpts, _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBEP20(&_ZkBNB.TransactOpts, _token, _amount, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_ZkBNB *ZkBNBTransactor) DepositBNB(opts *bind.TransactOpts, _accountName string) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositBNB", _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_ZkBNB *ZkBNBSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBNB(&_ZkBNB.TransactOpts, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBNB(&_ZkBNB.TransactOpts, _accountName)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBTransactor) DepositNft(opts *bind.TransactOpts, _accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositNft", _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositNft(&_ZkBNB.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositNft(&_ZkBNB.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _pubKeyX, bytes32 _pubKeyY) payable returns()
func (_ZkBNB *ZkBNBTransactor) RegisterZNS(opts *bind.TransactOpts, _name string, _owner common.Address, _pubKeyX [32]byte, _pubKeyY [32]byte) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "registerZNS", _name, _owner, _pubKeyX, _pubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _pubKeyX, bytes32 _pubKeyY) payable returns()
func (_ZkBNB *ZkBNBSession) RegisterZNS(_name string, _owner common.Address, _pubKeyX [32]byte, _pubKeyY [32]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.RegisterZNS(&_ZkBNB.TransactOpts, _name, _owner, _pubKeyX, _pubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _pubKeyX, bytes32 _pubKeyY) payable returns()
func (_ZkBNB *ZkBNBTransactorSession) RegisterZNS(_name string, _owner common.Address, _pubKeyX [32]byte, _pubKeyY [32]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.RegisterZNS(&_ZkBNB.TransactOpts, _name, _owner, _pubKeyX, _pubKeyY)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_ZkBNB *ZkBNBTransactor) RequestFullExit(opts *bind.TransactOpts, _accountName string, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "requestFullExit", _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_ZkBNB *ZkBNBSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExit(&_ZkBNB.TransactOpts, _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_ZkBNB *ZkBNBTransactorSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExit(&_ZkBNB.TransactOpts, _accountName, _asset)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_ZkBNB *ZkBNBTransactor) RequestFullExitNft(opts *bind.TransactOpts, _accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "requestFullExitNft", _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_ZkBNB *ZkBNBSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExitNft(&_ZkBNB.TransactOpts, _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_ZkBNB *ZkBNBTransactorSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExitNft(&_ZkBNB.TransactOpts, _accountName, _nftIndex)
}
