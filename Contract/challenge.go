// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challenge

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChallengeABI is the input ABI used to generate the binding from.
const ChallengeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"setB\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"a\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setA\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"b\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChallengeBin is the compiled bytecode used for deploying new contracts.
const ChallengeBin = `608060405234801561001057600080fd5b506102de806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806309cdcf9b146100675780630dbe671f146100ad5780631fc376f7146100f75780634df7e3d014610119578063a1c5191514610137578063d46300fd14610155575b600080fd5b6100936004803603602081101561007d57600080fd5b810190808035906020019092919050505061019f565b604051808215151515815260200191505060405180910390f35b6100b56101b4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100ff6101d9565b604051808215151515815260200191505060405180910390f35b610121610270565b6040518082815260200191505060405180910390f35b61013f610276565b6040518082815260200191505060405180910390f35b61015d610280565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008160018190555081600154149050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905090565b60015481565b6000600154905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690509056fea265627a7a72305820b83ec25149d58e94fa43391f27ad64cd54dd768d05ed51b9800960d2c54fd55d64736f6c63430005090032`

// DeployChallenge deploys a new Ethereum contract, binding an instance of Challenge to it.
func DeployChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Challenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// Challenge is an auto generated Go binding around an Ethereum contract.
type Challenge struct {
	ChallengeCaller     // Read-only binding to the contract
	ChallengeTransactor // Write-only binding to the contract
	ChallengeFilterer   // Log filterer for contract events
}

// ChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeSession struct {
	Contract     *Challenge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeCallerSession struct {
	Contract *ChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTransactorSession struct {
	Contract     *ChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeRaw struct {
	Contract *Challenge // Generic contract binding to access the raw methods on
}

// ChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeCallerRaw struct {
	Contract *ChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTransactorRaw struct {
	Contract *ChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallenge creates a new instance of Challenge, bound to a specific deployed contract.
func NewChallenge(address common.Address, backend bind.ContractBackend) (*Challenge, error) {
	contract, err := bindChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// NewChallengeCaller creates a new read-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeCaller, error) {
	contract, err := bindChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeCaller{contract: contract}, nil
}

// NewChallengeTransactor creates a new write-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTransactor, error) {
	contract, err := bindChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTransactor{contract: contract}, nil
}

// NewChallengeFilterer creates a new log filterer instance of Challenge, bound to a specific deployed contract.
func NewChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFilterer, error) {
	contract, err := bindChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFilterer{contract: contract}, nil
}

// bindChallenge binds a generic wrapper to an already deployed contract.
func bindChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.ChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() constant returns(address)
func (_Challenge *ChallengeCaller) A(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Challenge.contract.Call(opts, out, "a")
	return *ret0, err
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() constant returns(address)
func (_Challenge *ChallengeSession) A() (common.Address, error) {
	return _Challenge.Contract.A(&_Challenge.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() constant returns(address)
func (_Challenge *ChallengeCallerSession) A() (common.Address, error) {
	return _Challenge.Contract.A(&_Challenge.CallOpts)
}

// B is a free data retrieval call binding the contract method 0x4df7e3d0.
//
// Solidity: function b() constant returns(uint256)
func (_Challenge *ChallengeCaller) B(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Challenge.contract.Call(opts, out, "b")
	return *ret0, err
}

// B is a free data retrieval call binding the contract method 0x4df7e3d0.
//
// Solidity: function b() constant returns(uint256)
func (_Challenge *ChallengeSession) B() (*big.Int, error) {
	return _Challenge.Contract.B(&_Challenge.CallOpts)
}

// B is a free data retrieval call binding the contract method 0x4df7e3d0.
//
// Solidity: function b() constant returns(uint256)
func (_Challenge *ChallengeCallerSession) B() (*big.Int, error) {
	return _Challenge.Contract.B(&_Challenge.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() constant returns(address)
func (_Challenge *ChallengeCaller) GetA(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Challenge.contract.Call(opts, out, "getA")
	return *ret0, err
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() constant returns(address)
func (_Challenge *ChallengeSession) GetA() (common.Address, error) {
	return _Challenge.Contract.GetA(&_Challenge.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() constant returns(address)
func (_Challenge *ChallengeCallerSession) GetA() (common.Address, error) {
	return _Challenge.Contract.GetA(&_Challenge.CallOpts)
}

// GetB is a free data retrieval call binding the contract method 0xa1c51915.
//
// Solidity: function getB() constant returns(uint256)
func (_Challenge *ChallengeCaller) GetB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Challenge.contract.Call(opts, out, "getB")
	return *ret0, err
}

// GetB is a free data retrieval call binding the contract method 0xa1c51915.
//
// Solidity: function getB() constant returns(uint256)
func (_Challenge *ChallengeSession) GetB() (*big.Int, error) {
	return _Challenge.Contract.GetB(&_Challenge.CallOpts)
}

// GetB is a free data retrieval call binding the contract method 0xa1c51915.
//
// Solidity: function getB() constant returns(uint256)
func (_Challenge *ChallengeCallerSession) GetB() (*big.Int, error) {
	return _Challenge.Contract.GetB(&_Challenge.CallOpts)
}

// SetA is a paid mutator transaction binding the contract method 0x1fc376f7.
//
// Solidity: function setA() returns(bool)
func (_Challenge *ChallengeTransactor) SetA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "setA")
}

// SetA is a paid mutator transaction binding the contract method 0x1fc376f7.
//
// Solidity: function setA() returns(bool)
func (_Challenge *ChallengeSession) SetA() (*types.Transaction, error) {
	return _Challenge.Contract.SetA(&_Challenge.TransactOpts)
}

// SetA is a paid mutator transaction binding the contract method 0x1fc376f7.
//
// Solidity: function setA() returns(bool)
func (_Challenge *ChallengeTransactorSession) SetA() (*types.Transaction, error) {
	return _Challenge.Contract.SetA(&_Challenge.TransactOpts)
}

// SetB is a paid mutator transaction binding the contract method 0x09cdcf9b.
//
// Solidity: function setB(uint256 _b) returns(bool)
func (_Challenge *ChallengeTransactor) SetB(opts *bind.TransactOpts, _b *big.Int) (*types.Transaction, error) {
	return _Challenge.contract.Transact(opts, "setB", _b)
}

// SetB is a paid mutator transaction binding the contract method 0x09cdcf9b.
//
// Solidity: function setB(uint256 _b) returns(bool)
func (_Challenge *ChallengeSession) SetB(_b *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetB(&_Challenge.TransactOpts, _b)
}

// SetB is a paid mutator transaction binding the contract method 0x09cdcf9b.
//
// Solidity: function setB(uint256 _b) returns(bool)
func (_Challenge *ChallengeTransactorSession) SetB(_b *big.Int) (*types.Transaction, error) {
	return _Challenge.Contract.SetB(&_Challenge.TransactOpts, _b)
}
