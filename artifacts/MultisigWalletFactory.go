// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mypackage

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

// MypackageMetaData contains all meta data concerning the Mypackage contract.
var MypackageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"instantiation\",\"type\":\"address\"}],\"name\":\"ContractInstantiation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getInstantiationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instantiations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isInstantiation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MypackageABI is the input ABI used to generate the binding from.
// Deprecated: Use MypackageMetaData.ABI instead.
var MypackageABI = MypackageMetaData.ABI

// Mypackage is an auto generated Go binding around an Ethereum contract.
type Mypackage struct {
	MypackageCaller     // Read-only binding to the contract
	MypackageTransactor // Write-only binding to the contract
	MypackageFilterer   // Log filterer for contract events
}

// MypackageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MypackageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MypackageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MypackageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MypackageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MypackageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MypackageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MypackageSession struct {
	Contract     *Mypackage        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MypackageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MypackageCallerSession struct {
	Contract *MypackageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MypackageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MypackageTransactorSession struct {
	Contract     *MypackageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MypackageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MypackageRaw struct {
	Contract *Mypackage // Generic contract binding to access the raw methods on
}

// MypackageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MypackageCallerRaw struct {
	Contract *MypackageCaller // Generic read-only contract binding to access the raw methods on
}

// MypackageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MypackageTransactorRaw struct {
	Contract *MypackageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMypackage creates a new instance of Mypackage, bound to a specific deployed contract.
func NewMypackage(address common.Address, backend bind.ContractBackend) (*Mypackage, error) {
	contract, err := bindMypackage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mypackage{MypackageCaller: MypackageCaller{contract: contract}, MypackageTransactor: MypackageTransactor{contract: contract}, MypackageFilterer: MypackageFilterer{contract: contract}}, nil
}

// NewMypackageCaller creates a new read-only instance of Mypackage, bound to a specific deployed contract.
func NewMypackageCaller(address common.Address, caller bind.ContractCaller) (*MypackageCaller, error) {
	contract, err := bindMypackage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MypackageCaller{contract: contract}, nil
}

// NewMypackageTransactor creates a new write-only instance of Mypackage, bound to a specific deployed contract.
func NewMypackageTransactor(address common.Address, transactor bind.ContractTransactor) (*MypackageTransactor, error) {
	contract, err := bindMypackage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MypackageTransactor{contract: contract}, nil
}

// NewMypackageFilterer creates a new log filterer instance of Mypackage, bound to a specific deployed contract.
func NewMypackageFilterer(address common.Address, filterer bind.ContractFilterer) (*MypackageFilterer, error) {
	contract, err := bindMypackage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MypackageFilterer{contract: contract}, nil
}

// bindMypackage binds a generic wrapper to an already deployed contract.
func bindMypackage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MypackageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mypackage *MypackageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mypackage.Contract.MypackageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mypackage *MypackageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mypackage.Contract.MypackageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mypackage *MypackageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mypackage.Contract.MypackageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mypackage *MypackageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mypackage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mypackage *MypackageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mypackage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mypackage *MypackageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mypackage.Contract.contract.Transact(opts, method, params...)
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Mypackage *MypackageCaller) GetInstantiationCount(opts *bind.CallOpts, creator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mypackage.contract.Call(opts, &out, "getInstantiationCount", creator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Mypackage *MypackageSession) GetInstantiationCount(creator common.Address) (*big.Int, error) {
	return _Mypackage.Contract.GetInstantiationCount(&_Mypackage.CallOpts, creator)
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Mypackage *MypackageCallerSession) GetInstantiationCount(creator common.Address) (*big.Int, error) {
	return _Mypackage.Contract.GetInstantiationCount(&_Mypackage.CallOpts, creator)
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Mypackage *MypackageCaller) Instantiations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mypackage.contract.Call(opts, &out, "instantiations", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Mypackage *MypackageSession) Instantiations(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Mypackage.Contract.Instantiations(&_Mypackage.CallOpts, arg0, arg1)
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Mypackage *MypackageCallerSession) Instantiations(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Mypackage.Contract.Instantiations(&_Mypackage.CallOpts, arg0, arg1)
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Mypackage *MypackageCaller) IsInstantiation(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Mypackage.contract.Call(opts, &out, "isInstantiation", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Mypackage *MypackageSession) IsInstantiation(arg0 common.Address) (bool, error) {
	return _Mypackage.Contract.IsInstantiation(&_Mypackage.CallOpts, arg0)
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Mypackage *MypackageCallerSession) IsInstantiation(arg0 common.Address) (bool, error) {
	return _Mypackage.Contract.IsInstantiation(&_Mypackage.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0x53d9d910.
//
// Solidity: function create(address[] _owners, uint256 _required, uint256 _dailyLimit) returns(address)
func (_Mypackage *MypackageTransactor) Create(opts *bind.TransactOpts, _owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _Mypackage.contract.Transact(opts, "create", _owners, _required, _dailyLimit)
}

// Create is a paid mutator transaction binding the contract method 0x53d9d910.
//
// Solidity: function create(address[] _owners, uint256 _required, uint256 _dailyLimit) returns(address)
func (_Mypackage *MypackageSession) Create(_owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _Mypackage.Contract.Create(&_Mypackage.TransactOpts, _owners, _required, _dailyLimit)
}

// Create is a paid mutator transaction binding the contract method 0x53d9d910.
//
// Solidity: function create(address[] _owners, uint256 _required, uint256 _dailyLimit) returns(address)
func (_Mypackage *MypackageTransactorSession) Create(_owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _Mypackage.Contract.Create(&_Mypackage.TransactOpts, _owners, _required, _dailyLimit)
}

// MypackageContractInstantiationIterator is returned from FilterContractInstantiation and is used to iterate over the raw logs and unpacked data for ContractInstantiation events raised by the Mypackage contract.
type MypackageContractInstantiationIterator struct {
	Event *MypackageContractInstantiation // Event containing the contract specifics and raw log

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
func (it *MypackageContractInstantiationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MypackageContractInstantiation)
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
		it.Event = new(MypackageContractInstantiation)
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
func (it *MypackageContractInstantiationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MypackageContractInstantiationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MypackageContractInstantiation represents a ContractInstantiation event raised by the Mypackage contract.
type MypackageContractInstantiation struct {
	Sender        common.Address
	Instantiation common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContractInstantiation is a free log retrieval operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_Mypackage *MypackageFilterer) FilterContractInstantiation(opts *bind.FilterOpts) (*MypackageContractInstantiationIterator, error) {

	logs, sub, err := _Mypackage.contract.FilterLogs(opts, "ContractInstantiation")
	if err != nil {
		return nil, err
	}
	return &MypackageContractInstantiationIterator{contract: _Mypackage.contract, event: "ContractInstantiation", logs: logs, sub: sub}, nil
}

// WatchContractInstantiation is a free log subscription operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_Mypackage *MypackageFilterer) WatchContractInstantiation(opts *bind.WatchOpts, sink chan<- *MypackageContractInstantiation) (event.Subscription, error) {

	logs, sub, err := _Mypackage.contract.WatchLogs(opts, "ContractInstantiation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MypackageContractInstantiation)
				if err := _Mypackage.contract.UnpackLog(event, "ContractInstantiation", log); err != nil {
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

// ParseContractInstantiation is a log parse operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_Mypackage *MypackageFilterer) ParseContractInstantiation(log types.Log) (*MypackageContractInstantiation, error) {
	event := new(MypackageContractInstantiation)
	if err := _Mypackage.contract.UnpackLog(event, "ContractInstantiation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
