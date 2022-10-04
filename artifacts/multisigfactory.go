// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisigfactory

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

// MultisigfactoryMetaData contains all meta data concerning the Multisigfactory contract.
var MultisigfactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"instantiation\",\"type\":\"address\"}],\"name\":\"ContractInstantiation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getInstantiationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instantiations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isInstantiation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MultisigfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MultisigfactoryMetaData.ABI instead.
var MultisigfactoryABI = MultisigfactoryMetaData.ABI

// Multisigfactory is an auto generated Go binding around an Ethereum contract.
type Multisigfactory struct {
	MultisigfactoryCaller     // Read-only binding to the contract
	MultisigfactoryTransactor // Write-only binding to the contract
	MultisigfactoryFilterer   // Log filterer for contract events
}

// MultisigfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisigfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisigfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisigfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisigfactorySession struct {
	Contract     *Multisigfactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisigfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisigfactoryCallerSession struct {
	Contract *MultisigfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MultisigfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisigfactoryTransactorSession struct {
	Contract     *MultisigfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MultisigfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisigfactoryRaw struct {
	Contract *Multisigfactory // Generic contract binding to access the raw methods on
}

// MultisigfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisigfactoryCallerRaw struct {
	Contract *MultisigfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MultisigfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisigfactoryTransactorRaw struct {
	Contract *MultisigfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisigfactory creates a new instance of Multisigfactory, bound to a specific deployed contract.
func NewMultisigfactory(address common.Address, backend bind.ContractBackend) (*Multisigfactory, error) {
	contract, err := bindMultisigfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisigfactory{MultisigfactoryCaller: MultisigfactoryCaller{contract: contract}, MultisigfactoryTransactor: MultisigfactoryTransactor{contract: contract}, MultisigfactoryFilterer: MultisigfactoryFilterer{contract: contract}}, nil
}

// NewMultisigfactoryCaller creates a new read-only instance of Multisigfactory, bound to a specific deployed contract.
func NewMultisigfactoryCaller(address common.Address, caller bind.ContractCaller) (*MultisigfactoryCaller, error) {
	contract, err := bindMultisigfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigfactoryCaller{contract: contract}, nil
}

// NewMultisigfactoryTransactor creates a new write-only instance of Multisigfactory, bound to a specific deployed contract.
func NewMultisigfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigfactoryTransactor, error) {
	contract, err := bindMultisigfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigfactoryTransactor{contract: contract}, nil
}

// NewMultisigfactoryFilterer creates a new log filterer instance of Multisigfactory, bound to a specific deployed contract.
func NewMultisigfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigfactoryFilterer, error) {
	contract, err := bindMultisigfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisigfactoryFilterer{contract: contract}, nil
}

// bindMultisigfactory binds a generic wrapper to an already deployed contract.
func bindMultisigfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisigfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisigfactory *MultisigfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisigfactory.Contract.MultisigfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisigfactory *MultisigfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisigfactory.Contract.MultisigfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisigfactory *MultisigfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisigfactory.Contract.MultisigfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisigfactory *MultisigfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisigfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisigfactory *MultisigfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisigfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisigfactory *MultisigfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisigfactory.Contract.contract.Transact(opts, method, params...)
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Multisigfactory *MultisigfactoryCaller) GetInstantiationCount(opts *bind.CallOpts, creator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multisigfactory.contract.Call(opts, &out, "getInstantiationCount", creator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Multisigfactory *MultisigfactorySession) GetInstantiationCount(creator common.Address) (*big.Int, error) {
	return _Multisigfactory.Contract.GetInstantiationCount(&_Multisigfactory.CallOpts, creator)
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Multisigfactory *MultisigfactoryCallerSession) GetInstantiationCount(creator common.Address) (*big.Int, error) {
	return _Multisigfactory.Contract.GetInstantiationCount(&_Multisigfactory.CallOpts, creator)
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Multisigfactory *MultisigfactoryCaller) Instantiations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multisigfactory.contract.Call(opts, &out, "instantiations", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Multisigfactory *MultisigfactorySession) Instantiations(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Multisigfactory.Contract.Instantiations(&_Multisigfactory.CallOpts, arg0, arg1)
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Multisigfactory *MultisigfactoryCallerSession) Instantiations(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Multisigfactory.Contract.Instantiations(&_Multisigfactory.CallOpts, arg0, arg1)
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Multisigfactory *MultisigfactoryCaller) IsInstantiation(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Multisigfactory.contract.Call(opts, &out, "isInstantiation", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Multisigfactory *MultisigfactorySession) IsInstantiation(arg0 common.Address) (bool, error) {
	return _Multisigfactory.Contract.IsInstantiation(&_Multisigfactory.CallOpts, arg0)
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Multisigfactory *MultisigfactoryCallerSession) IsInstantiation(arg0 common.Address) (bool, error) {
	return _Multisigfactory.Contract.IsInstantiation(&_Multisigfactory.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0x53d9d910.
//
// Solidity: function create(address[] _owners, uint256 _required, uint256 _dailyLimit) returns(address)
func (_Multisigfactory *MultisigfactoryTransactor) Create(opts *bind.TransactOpts, _owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _Multisigfactory.contract.Transact(opts, "create", _owners, _required, _dailyLimit)
}

// Create is a paid mutator transaction binding the contract method 0x53d9d910.
//
// Solidity: function create(address[] _owners, uint256 _required, uint256 _dailyLimit) returns(address)
func (_Multisigfactory *MultisigfactorySession) Create(_owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _Multisigfactory.Contract.Create(&_Multisigfactory.TransactOpts, _owners, _required, _dailyLimit)
}

// Create is a paid mutator transaction binding the contract method 0x53d9d910.
//
// Solidity: function create(address[] _owners, uint256 _required, uint256 _dailyLimit) returns(address)
func (_Multisigfactory *MultisigfactoryTransactorSession) Create(_owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _Multisigfactory.Contract.Create(&_Multisigfactory.TransactOpts, _owners, _required, _dailyLimit)
}

// MultisigfactoryContractInstantiationIterator is returned from FilterContractInstantiation and is used to iterate over the raw logs and unpacked data for ContractInstantiation events raised by the Multisigfactory contract.
type MultisigfactoryContractInstantiationIterator struct {
	Event *MultisigfactoryContractInstantiation // Event containing the contract specifics and raw log

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
func (it *MultisigfactoryContractInstantiationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigfactoryContractInstantiation)
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
		it.Event = new(MultisigfactoryContractInstantiation)
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
func (it *MultisigfactoryContractInstantiationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigfactoryContractInstantiationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigfactoryContractInstantiation represents a ContractInstantiation event raised by the Multisigfactory contract.
type MultisigfactoryContractInstantiation struct {
	Sender        common.Address
	Instantiation common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContractInstantiation is a free log retrieval operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_Multisigfactory *MultisigfactoryFilterer) FilterContractInstantiation(opts *bind.FilterOpts) (*MultisigfactoryContractInstantiationIterator, error) {

	logs, sub, err := _Multisigfactory.contract.FilterLogs(opts, "ContractInstantiation")
	if err != nil {
		return nil, err
	}
	return &MultisigfactoryContractInstantiationIterator{contract: _Multisigfactory.contract, event: "ContractInstantiation", logs: logs, sub: sub}, nil
}

// WatchContractInstantiation is a free log subscription operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_Multisigfactory *MultisigfactoryFilterer) WatchContractInstantiation(opts *bind.WatchOpts, sink chan<- *MultisigfactoryContractInstantiation) (event.Subscription, error) {

	logs, sub, err := _Multisigfactory.contract.WatchLogs(opts, "ContractInstantiation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigfactoryContractInstantiation)
				if err := _Multisigfactory.contract.UnpackLog(event, "ContractInstantiation", log); err != nil {
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
func (_Multisigfactory *MultisigfactoryFilterer) ParseContractInstantiation(log types.Log) (*MultisigfactoryContractInstantiation, error) {
	event := new(MultisigfactoryContractInstantiation)
	if err := _Multisigfactory.contract.UnpackLog(event, "ContractInstantiation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
