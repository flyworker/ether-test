// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ubi

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
	_ = abi.ConvertType
)

// CpaccountMetaData contains all meta data concerning the Cpaccount contract.
var CpaccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newExpiration\",\"type\":\"uint256\"}],\"name\":\"changeBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"newMultiaddrs\",\"type\":\"string[]\"}],\"name\":\"changeMultiaddrs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taskId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"}],\"name\":\"submitUBIProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_windowPoStProofType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_peerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_multiAddresses\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"multiAddresses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peerId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isSubmitted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"windowPoStProofType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CpaccountABI is the input ABI used to generate the binding from.
// Deprecated: Use CpaccountMetaData.ABI instead.
var CpaccountABI = CpaccountMetaData.ABI

// Cpaccount is an auto generated Go binding around an Ethereum contract.
type Cpaccount struct {
	CpaccountCaller     // Read-only binding to the contract
	CpaccountTransactor // Write-only binding to the contract
	CpaccountFilterer   // Log filterer for contract events
}

// CpaccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type CpaccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CpaccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CpaccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CpaccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CpaccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CpaccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CpaccountSession struct {
	Contract     *Cpaccount        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CpaccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CpaccountCallerSession struct {
	Contract *CpaccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CpaccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CpaccountTransactorSession struct {
	Contract     *CpaccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CpaccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type CpaccountRaw struct {
	Contract *Cpaccount // Generic contract binding to access the raw methods on
}

// CpaccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CpaccountCallerRaw struct {
	Contract *CpaccountCaller // Generic read-only contract binding to access the raw methods on
}

// CpaccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CpaccountTransactorRaw struct {
	Contract *CpaccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCpaccount creates a new instance of Cpaccount, bound to a specific deployed contract.
func NewCpaccount(address common.Address, backend bind.ContractBackend) (*Cpaccount, error) {
	contract, err := bindCpaccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cpaccount{CpaccountCaller: CpaccountCaller{contract: contract}, CpaccountTransactor: CpaccountTransactor{contract: contract}, CpaccountFilterer: CpaccountFilterer{contract: contract}}, nil
}

// NewCpaccountCaller creates a new read-only instance of Cpaccount, bound to a specific deployed contract.
func NewCpaccountCaller(address common.Address, caller bind.ContractCaller) (*CpaccountCaller, error) {
	contract, err := bindCpaccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CpaccountCaller{contract: contract}, nil
}

// NewCpaccountTransactor creates a new write-only instance of Cpaccount, bound to a specific deployed contract.
func NewCpaccountTransactor(address common.Address, transactor bind.ContractTransactor) (*CpaccountTransactor, error) {
	contract, err := bindCpaccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CpaccountTransactor{contract: contract}, nil
}

// NewCpaccountFilterer creates a new log filterer instance of Cpaccount, bound to a specific deployed contract.
func NewCpaccountFilterer(address common.Address, filterer bind.ContractFilterer) (*CpaccountFilterer, error) {
	contract, err := bindCpaccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CpaccountFilterer{contract: contract}, nil
}

// bindCpaccount binds a generic wrapper to an already deployed contract.
func bindCpaccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CpaccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cpaccount *CpaccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cpaccount.Contract.CpaccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cpaccount *CpaccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cpaccount.Contract.CpaccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cpaccount *CpaccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cpaccount.Contract.CpaccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cpaccount *CpaccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cpaccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cpaccount *CpaccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cpaccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cpaccount *CpaccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cpaccount.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Cpaccount *CpaccountCaller) Beneficiary(opts *bind.CallOpts) (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	var out []interface{}
	err := _Cpaccount.contract.Call(opts, &out, "beneficiary")

	outstruct := new(struct {
		BeneficiaryAddress common.Address
		Quota              *big.Int
		Expiration         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BeneficiaryAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quota = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Cpaccount *CpaccountSession) Beneficiary() (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	return _Cpaccount.Contract.Beneficiary(&_Cpaccount.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Cpaccount *CpaccountCallerSession) Beneficiary() (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	return _Cpaccount.Contract.Beneficiary(&_Cpaccount.CallOpts)
}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Cpaccount *CpaccountCaller) MultiAddresses(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cpaccount.contract.Call(opts, &out, "multiAddresses", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Cpaccount *CpaccountSession) MultiAddresses(arg0 *big.Int) (string, error) {
	return _Cpaccount.Contract.MultiAddresses(&_Cpaccount.CallOpts, arg0)
}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Cpaccount *CpaccountCallerSession) MultiAddresses(arg0 *big.Int) (string, error) {
	return _Cpaccount.Contract.MultiAddresses(&_Cpaccount.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cpaccount *CpaccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cpaccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cpaccount *CpaccountSession) Owner() (common.Address, error) {
	return _Cpaccount.Contract.Owner(&_Cpaccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cpaccount *CpaccountCallerSession) Owner() (common.Address, error) {
	return _Cpaccount.Contract.Owner(&_Cpaccount.CallOpts)
}

// PeerId is a free data retrieval call binding the contract method 0x60531df4.
//
// Solidity: function peerId() view returns(string)
func (_Cpaccount *CpaccountCaller) PeerId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cpaccount.contract.Call(opts, &out, "peerId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PeerId is a free data retrieval call binding the contract method 0x60531df4.
//
// Solidity: function peerId() view returns(string)
func (_Cpaccount *CpaccountSession) PeerId() (string, error) {
	return _Cpaccount.Contract.PeerId(&_Cpaccount.CallOpts)
}

// PeerId is a free data retrieval call binding the contract method 0x60531df4.
//
// Solidity: function peerId() view returns(string)
func (_Cpaccount *CpaccountCallerSession) PeerId() (string, error) {
	return _Cpaccount.Contract.PeerId(&_Cpaccount.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string proof, bool isSubmitted)
func (_Cpaccount *CpaccountCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	Proof       string
	IsSubmitted bool
}, error) {
	var out []interface{}
	err := _Cpaccount.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		TaskId      string
		TaskType    uint8
		Proof       string
		IsSubmitted bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TaskId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TaskType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Proof = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IsSubmitted = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string proof, bool isSubmitted)
func (_Cpaccount *CpaccountSession) Tasks(arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	Proof       string
	IsSubmitted bool
}, error) {
	return _Cpaccount.Contract.Tasks(&_Cpaccount.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string proof, bool isSubmitted)
func (_Cpaccount *CpaccountCallerSession) Tasks(arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	Proof       string
	IsSubmitted bool
}, error) {
	return _Cpaccount.Contract.Tasks(&_Cpaccount.CallOpts, arg0)
}

// WindowPoStProofType is a free data retrieval call binding the contract method 0xdc68f1ff.
//
// Solidity: function windowPoStProofType() view returns(uint8)
func (_Cpaccount *CpaccountCaller) WindowPoStProofType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cpaccount.contract.Call(opts, &out, "windowPoStProofType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// WindowPoStProofType is a free data retrieval call binding the contract method 0xdc68f1ff.
//
// Solidity: function windowPoStProofType() view returns(uint8)
func (_Cpaccount *CpaccountSession) WindowPoStProofType() (uint8, error) {
	return _Cpaccount.Contract.WindowPoStProofType(&_Cpaccount.CallOpts)
}

// WindowPoStProofType is a free data retrieval call binding the contract method 0xdc68f1ff.
//
// Solidity: function windowPoStProofType() view returns(uint8)
func (_Cpaccount *CpaccountCallerSession) WindowPoStProofType() (uint8, error) {
	return _Cpaccount.Contract.WindowPoStProofType(&_Cpaccount.CallOpts)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Cpaccount *CpaccountTransactor) ChangeBeneficiary(opts *bind.TransactOpts, newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Cpaccount.contract.Transact(opts, "changeBeneficiary", newBeneficiary, newQuota, newExpiration)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Cpaccount *CpaccountSession) ChangeBeneficiary(newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Cpaccount.Contract.ChangeBeneficiary(&_Cpaccount.TransactOpts, newBeneficiary, newQuota, newExpiration)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Cpaccount *CpaccountTransactorSession) ChangeBeneficiary(newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Cpaccount.Contract.ChangeBeneficiary(&_Cpaccount.TransactOpts, newBeneficiary, newQuota, newExpiration)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Cpaccount *CpaccountTransactor) ChangeMultiaddrs(opts *bind.TransactOpts, newMultiaddrs []string) (*types.Transaction, error) {
	return _Cpaccount.contract.Transact(opts, "changeMultiaddrs", newMultiaddrs)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Cpaccount *CpaccountSession) ChangeMultiaddrs(newMultiaddrs []string) (*types.Transaction, error) {
	return _Cpaccount.Contract.ChangeMultiaddrs(&_Cpaccount.TransactOpts, newMultiaddrs)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Cpaccount *CpaccountTransactorSession) ChangeMultiaddrs(newMultiaddrs []string) (*types.Transaction, error) {
	return _Cpaccount.Contract.ChangeMultiaddrs(&_Cpaccount.TransactOpts, newMultiaddrs)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Cpaccount *CpaccountTransactor) ChangeOwnerAddress(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cpaccount.contract.Transact(opts, "changeOwnerAddress", newOwner)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Cpaccount *CpaccountSession) ChangeOwnerAddress(newOwner common.Address) (*types.Transaction, error) {
	return _Cpaccount.Contract.ChangeOwnerAddress(&_Cpaccount.TransactOpts, newOwner)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Cpaccount *CpaccountTransactorSession) ChangeOwnerAddress(newOwner common.Address) (*types.Transaction, error) {
	return _Cpaccount.Contract.ChangeOwnerAddress(&_Cpaccount.TransactOpts, newOwner)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x5ded591b.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _proof) returns()
func (_Cpaccount *CpaccountTransactor) SubmitUBIProof(opts *bind.TransactOpts, _taskId string, _taskType uint8, _proof string) (*types.Transaction, error) {
	return _Cpaccount.contract.Transact(opts, "submitUBIProof", _taskId, _taskType, _proof)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x5ded591b.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _proof) returns()
func (_Cpaccount *CpaccountSession) SubmitUBIProof(_taskId string, _taskType uint8, _proof string) (*types.Transaction, error) {
	return _Cpaccount.Contract.SubmitUBIProof(&_Cpaccount.TransactOpts, _taskId, _taskType, _proof)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x5ded591b.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _proof) returns()
func (_Cpaccount *CpaccountTransactorSession) SubmitUBIProof(_taskId string, _taskType uint8, _proof string) (*types.Transaction, error) {
	return _Cpaccount.Contract.SubmitUBIProof(&_Cpaccount.TransactOpts, _taskId, _taskType, _proof)
}
