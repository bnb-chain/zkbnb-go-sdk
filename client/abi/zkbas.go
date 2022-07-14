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

// ZkbasPairInfo is an auto generated low-level Go binding around an user-defined struct.
type ZkbasPairInfo struct {
	TokenA               common.Address
	TokenB               common.Address
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
}

// ZkbasMetaData contains all meta data concerning the Zkbas contract.
var ZkbasMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBEP20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_nftL1Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_pubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_pubKeyY\",\"type\":\"bytes32\"}],\"name\":\"registerZNS\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"requestFullExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_nftIndex\",\"type\":\"uint32\"}],\"name\":\"requestFullExitNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"internalType\":\"structZkbas.PairInfo\",\"name\":\"_pairInfo\",\"type\":\"tuple\"}],\"name\":\"updatePairRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZkbasABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkbasMetaData.ABI instead.
var ZkbasABI = ZkbasMetaData.ABI

// Zkbas is an auto generated Go binding around an Ethereum contract.
type Zkbas struct {
	ZkbasCaller     // Read-only binding to the contract
	ZkbasTransactor // Write-only binding to the contract
	ZkbasFilterer   // Log filterer for contract events
}

// ZkbasCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkbasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkbasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkbasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkbasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkbasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkbasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkbasSession struct {
	Contract     *Zkbas            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkbasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkbasCallerSession struct {
	Contract *ZkbasCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZkbasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkbasTransactorSession struct {
	Contract     *ZkbasTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkbasRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkbasRaw struct {
	Contract *Zkbas // Generic contract binding to access the raw methods on
}

// ZkbasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkbasCallerRaw struct {
	Contract *ZkbasCaller // Generic read-only contract binding to access the raw methods on
}

// ZkbasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkbasTransactorRaw struct {
	Contract *ZkbasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkbas creates a new instance of Zkbas, bound to a specific deployed contract.
func NewZkbas(address common.Address, backend bind.ContractBackend) (*Zkbas, error) {
	contract, err := bindZkbas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zkbas{ZkbasCaller: ZkbasCaller{contract: contract}, ZkbasTransactor: ZkbasTransactor{contract: contract}, ZkbasFilterer: ZkbasFilterer{contract: contract}}, nil
}

// NewZkbasCaller creates a new read-only instance of Zkbas, bound to a specific deployed contract.
func NewZkbasCaller(address common.Address, caller bind.ContractCaller) (*ZkbasCaller, error) {
	contract, err := bindZkbas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkbasCaller{contract: contract}, nil
}

// NewZkbasTransactor creates a new write-only instance of Zkbas, bound to a specific deployed contract.
func NewZkbasTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkbasTransactor, error) {
	contract, err := bindZkbas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkbasTransactor{contract: contract}, nil
}

// NewZkbasFilterer creates a new log filterer instance of Zkbas, bound to a specific deployed contract.
func NewZkbasFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkbasFilterer, error) {
	contract, err := bindZkbas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkbasFilterer{contract: contract}, nil
}

// bindZkbas binds a generic wrapper to an already deployed contract.
func bindZkbas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZkbasABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zkbas *ZkbasRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zkbas.Contract.ZkbasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zkbas *ZkbasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.Contract.ZkbasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zkbas *ZkbasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zkbas.Contract.ZkbasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zkbas *ZkbasCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zkbas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zkbas *ZkbasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zkbas *ZkbasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zkbas.Contract.contract.Transact(opts, method, params...)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zkbas *ZkbasTransactor) CreatePair(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "createPair", _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zkbas *ZkbasSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.CreatePair(&_Zkbas.TransactOpts, _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zkbas *ZkbasTransactorSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.CreatePair(&_Zkbas.TransactOpts, _tokenA, _tokenB)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zkbas *ZkbasTransactor) DepositBEP20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "depositBEP20", _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zkbas *ZkbasSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositBEP20(&_Zkbas.TransactOpts, _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zkbas *ZkbasTransactorSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositBEP20(&_Zkbas.TransactOpts, _token, _amount, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zkbas *ZkbasTransactor) DepositBNB(opts *bind.TransactOpts, _accountName string) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "depositBNB", _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zkbas *ZkbasSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositBNB(&_Zkbas.TransactOpts, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zkbas *ZkbasTransactorSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositBNB(&_Zkbas.TransactOpts, _accountName)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zkbas *ZkbasTransactor) DepositNft(opts *bind.TransactOpts, _accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "depositNft", _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zkbas *ZkbasSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositNft(&_Zkbas.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zkbas *ZkbasTransactorSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositNft(&_Zkbas.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _pubKeyX, bytes32 _pubKeyY) payable returns()
func (_Zkbas *ZkbasTransactor) RegisterZNS(opts *bind.TransactOpts, _name string, _owner common.Address, _pubKeyX [32]byte, _pubKeyY [32]byte) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "registerZNS", _name, _owner, _pubKeyX, _pubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _pubKeyX, bytes32 _pubKeyY) payable returns()
func (_Zkbas *ZkbasSession) RegisterZNS(_name string, _owner common.Address, _pubKeyX [32]byte, _pubKeyY [32]byte) (*types.Transaction, error) {
	return _Zkbas.Contract.RegisterZNS(&_Zkbas.TransactOpts, _name, _owner, _pubKeyX, _pubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _pubKeyX, bytes32 _pubKeyY) payable returns()
func (_Zkbas *ZkbasTransactorSession) RegisterZNS(_name string, _owner common.Address, _pubKeyX [32]byte, _pubKeyY [32]byte) (*types.Transaction, error) {
	return _Zkbas.Contract.RegisterZNS(&_Zkbas.TransactOpts, _name, _owner, _pubKeyX, _pubKeyY)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zkbas *ZkbasTransactor) RequestFullExit(opts *bind.TransactOpts, _accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "requestFullExit", _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zkbas *ZkbasSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.RequestFullExit(&_Zkbas.TransactOpts, _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zkbas *ZkbasTransactorSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.RequestFullExit(&_Zkbas.TransactOpts, _accountName, _asset)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zkbas *ZkbasTransactor) RequestFullExitNft(opts *bind.TransactOpts, _accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "requestFullExitNft", _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zkbas *ZkbasSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zkbas.Contract.RequestFullExitNft(&_Zkbas.TransactOpts, _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zkbas *ZkbasTransactorSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zkbas.Contract.RequestFullExitNft(&_Zkbas.TransactOpts, _accountName, _nftIndex)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zkbas *ZkbasTransactor) UpdatePairRate(opts *bind.TransactOpts, _pairInfo ZkbasPairInfo) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "updatePairRate", _pairInfo)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zkbas *ZkbasSession) UpdatePairRate(_pairInfo ZkbasPairInfo) (*types.Transaction, error) {
	return _Zkbas.Contract.UpdatePairRate(&_Zkbas.TransactOpts, _pairInfo)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zkbas *ZkbasTransactorSession) UpdatePairRate(_pairInfo ZkbasPairInfo) (*types.Transaction, error) {
	return _Zkbas.Contract.UpdatePairRate(&_Zkbas.TransactOpts, _pairInfo)
}
