// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamond_3

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

// DiamondDiamondArgs is an auto generated low-level Go binding around an user-defined struct.
type DiamondDiamondArgs struct {
	Owner common.Address
}

// IDiamondCutFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// Diamond3MetaData contains all meta data concerning the Diamond3 contract.
var Diamond3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structDiamond.DiamondArgs\",\"name\":\"_args\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405260405161177438038061177483398101604081905261002291611201565b6040805160008082526020820190925261003d91849161012b565b805161004890610353565b50507fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f6020527f699d9daa71b280d05a152715774afa0a81a312594b2d731d6b0b2552b7d6f69f8054600160ff1991821681179092557ff97e938d8af42f52387bb74b8b526fda8f184cc2aa534340a8d75a88fbecc77580548216831790557f65d510a5d8f7ef134ec444f7f34ee808c8eeb5177cdfd16be0c40fe1ab43369580548216831790556307f5828d60e41b6000527f5622121b47b8cd0120c4efe45dd5483242f54a3d49bd7679be565d47694918c3805490911690911790556115c0565b60005b835181101561030857600084828151811061014b5761014b61138f565b60200260200101516020015190506000600281111561016c5761016c6113a5565b81600281111561017e5761017e6113a5565b036101d2576101cd8583815181106101985761019861138f565b6020026020010151600001518684815181106101b6576101b661138f565b6020026020010151604001516103d660201b60201c565b6102f5565b60018160028111156101e6576101e66113a5565b03610235576101cd8583815181106102005761020061138f565b60200260200101516000015186848151811061021e5761021e61138f565b6020026020010151604001516106ab60201b60201c565b6002816002811115610249576102496113a5565b03610298576101cd8583815181106102635761026361138f565b6020026020010151600001518684815181106102815761028161138f565b60200260200101516040015161099a60201b60201c565b60405162461bcd60e51b815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f727265637420466163657443756044820152663a20b1ba34b7b760c91b60648201526084015b60405180910390fd5b5080610300816113d1565b91505061012e565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb67383838360405161033c9392919061143a565b60405180910390a161034e8282610af0565b505050565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c132080546001600160a01b031981166001600160a01b038481169182179093556040516000805160206116c8833981519152939092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b60008151116104295760405162461bcd60e51b815260206004820152602b602482015260008051602061175483398151915260448201526a1858d95d081d1bc818dd5d60aa1b60648201526084016102ec565b6000805160206116c88339815191526001600160a01b0383166104915760405162461bcd60e51b815260206004820152602c602482015260008051602061171083398151915260448201526b65206164647265737328302960a01b60648201526084016102ec565b6001600160a01b03831660009081526001820160205260408120549061ffff82169003610536576104da8460405180606001604052806024815260200161173060249139610cfd565b6002820180546001600160a01b038616600081815260018087016020908152604083208201805461ffff191661ffff90961695909517909455845490810185559381529190912090910180546001600160a01b03191690911790555b60005b83518110156106a45760008482815181106105565761055661138f565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b031680156105fc5760405162461bcd60e51b815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f60448201527f6e207468617420616c726561647920657869737473000000000000000000000060648201526084016102ec565b6001600160a01b03871660008181526001878101602090815260408084208054938401815584528184206008840401805463ffffffff60079095166004026101000a948502191660e089901c94909402939093179092556001600160e01b031986168352889052902080546001600160b01b031916909117600160a01b61ffff8716021790558361068c8161153a565b9450505050808061069c906113d1565b915050610539565b5050505050565b60008151116106fe5760405162461bcd60e51b815260206004820152602b602482015260008051602061175483398151915260448201526a1858d95d081d1bc818dd5d60aa1b60648201526084016102ec565b6000805160206116c88339815191526001600160a01b0383166107665760405162461bcd60e51b815260206004820152602c602482015260008051602061171083398151915260448201526b65206164647265737328302960a01b60648201526084016102ec565b6001600160a01b03831660009081526001820160205260408120549061ffff8216900361080b576107af8460405180606001604052806024815260200161173060249139610cfd565b6002820180546001600160a01b038616600081815260018087016020908152604083208201805461ffff191661ffff90961695909517909455845490810185559381529190912090910180546001600160a01b03191690911790555b60005b83518110156106a457600084828151811061082b5761082b61138f565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b0390811690871681036108d65760405162461bcd60e51b815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e20776974682073616d652066756e6374696f6e000000000000000060648201526084016102ec565b6108e08183610d1e565b6001600160e01b03198216600081815260208781526040808320805461ffff60a01b1916600160a01b61ffff8b16021781556001600160a01b038c168085526001808c0185529285208054938401815585528385206008840401805463ffffffff60079095166004026101000a948502191660e08a901c94909402939093179092559390925287905281546001600160a01b031916179055836109828161153a565b94505050508080610992906113d1565b91505061080e565b60008151116109ed5760405162461bcd60e51b815260206004820152602b602482015260008051602061175483398151915260448201526a1858d95d081d1bc818dd5d60aa1b60648201526084016102ec565b6000805160206116c88339815191526001600160a01b03831615610a795760405162461bcd60e51b815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f7665206661636574206164647260448201527f657373206d75737420626520616464726573732830290000000000000000000060648201526084016102ec565b60005b8251811015610aea576000838281518110610a9957610a9961138f565b6020908102919091018101516001600160e01b031981166000908152918590526040909120549091506001600160a01b0316610ad58183610d1e565b50508080610ae2906113d1565b915050610a7c565b50505050565b6001600160a01b038216610b7757805115610b735760405162461bcd60e51b815260206004820152603c60248201527f4c69624469616d6f6e644375743a205f696e697420697320616464726573732860448201527f3029206275745f63616c6c64617461206973206e6f7420656d7074790000000060648201526084016102ec565b5050565b6000815111610bee5760405162461bcd60e51b815260206004820152603d60248201527f4c69624469616d6f6e644375743a205f63616c6c6461746120697320656d707460448201527f7920627574205f696e6974206973206e6f74206164647265737328302900000060648201526084016102ec565b6001600160a01b0382163014610c2057610c20826040518060600160405280602881526020016116e860289139610cfd565b600080836001600160a01b031683604051610c3b919061155b565b600060405180830381855af49150503d8060008114610c76576040519150601f19603f3d011682016040523d82523d6000602084013e610c7b565b606091505b509150915081610aea57805115610ca6578060405162461bcd60e51b81526004016102ec9190611577565b60405162461bcd60e51b815260206004820152602660248201527f4c69624469616d6f6e644375743a205f696e69742066756e6374696f6e2072656044820152651d995c9d195960d21b60648201526084016102ec565b813b8181610aea5760405162461bcd60e51b81526004016102ec9190611577565b6000805160206116c88339815191526001600160a01b038316610da95760405162461bcd60e51b815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e277420657869737400000000000000000060648201526084016102ec565b306001600160a01b03841603610e185760405162461bcd60e51b815260206004820152602e60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201526d3a30b1363290333ab731ba34b7b760911b60648201526084016102ec565b6001600160e01b03198216600090815260208281526040808320546001600160a01b0387168452600180860190935290832054600160a01b90910461ffff169291610e6291611591565b9050808214610f4e576001600160a01b03851660009081526001840160205260408120805483908110610e9757610e9761138f565b600091825260208083206008830401546001600160a01b038a168452600188019091526040909220805460079092166004026101000a90920460e01b925082919085908110610ee857610ee861138f565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c929092029390931790556001600160e01b031992909216825284905260409020805461ffff60a01b1916600160a01b61ffff8516021790555b6001600160a01b03851660009081526001840160205260409020805480610f7757610f776115aa565b60008281526020808220600860001990940193840401805463ffffffff600460078716026101000a0219169055919092556001600160e01b0319861682528490526040812080546001600160b01b03191690558190036106a4576002830154600090610fe590600190611591565b6001600160a01b038716600090815260018087016020526040909120015490915061ffff168082146110a45760008560020183815481106110285761102861138f565b6000918252602090912001546002870180546001600160a01b0390921692508291849081106110595761105961138f565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394851617905592909116815260018781019092526040902001805461ffff191661ffff83161790555b846002018054806110b7576110b76115aa565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03891682526001878101909152604090912001805461ffff1916905550505050505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b038111828210171561114157611141611109565b60405290565b604051601f8201601f191681016001600160401b038111828210171561116f5761116f611109565b604052919050565b60006001600160401b0382111561119057611190611109565b5060051b60200190565b80516001600160a01b03811681146111b157600080fd5b919050565b6000602082840312156111c857600080fd5b604051602081016001600160401b03811182821017156111ea576111ea611109565b6040529050806111f98361119a565b905292915050565b6000806040838503121561121457600080fd5b82516001600160401b038082111561122b57600080fd5b818501915085601f83011261123f57600080fd5b8151602061125461124f83611177565b611147565b82815260059290921b8401810191818101908984111561127357600080fd5b8286015b848110156113715780518681111561128e57600080fd5b87016060818d03601f190112156112a457600080fd5b6112ac61111f565b6112b786830161119a565b81526040820151600381106112cb57600080fd5b818701526060820151888111156112e157600080fd5b8083019250508c603f8301126112f657600080fd5b8582015161130661124f82611177565b81815260059190911b830160400190878101908f83111561132657600080fd5b6040850194505b8285101561135c5784516001600160e01b03198116811461134d57600080fd5b8252938801939088019061132d565b60408401525050845250918301918301611277565b5096506113829050888883016111b6565b9450505050509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016113e3576113e36113bb565b5060010190565b60005b838110156114055781810151838201526020016113ed565b50506000910152565b600081518084526114268160208601602086016113ea565b601f01601f19169290920160200192915050565b60006060808301818452808751808352608092508286019150828160051b8701016020808b0160005b8481101561150a57898403607f19018652815180516001600160a01b031685528381015189860190600381106114a957634e487b7160e01b600052602160045260246000fd5b868601526040918201519186018a905281519081905290840190600090898701905b808310156114f55783516001600160e01b03191682529286019260019290920191908601906114cb565b50978501979550505090820190600101611463565b50506001600160a01b038a1690880152868103604088015261152c818961140e565b9a9950505050505050505050565b600061ffff808316818103611551576115516113bb565b6001019392505050565b6000825161156d8184602087016113ea565b9190910192915050565b60208152600061158a602083018461140e565b9392505050565b818103818111156115a4576115a46113bb565b92915050565b634e487b7160e01b600052603160045260246000fd5b60fa806115ce6000396000f3fe608060405236600a57005b600080356001600160e01b03191681527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c602081905260409091205481906001600160a01b03168060a15760405162461bcd60e51b815260206004820181905260248201527f4469616d6f6e643a2046756e6374696f6e20646f6573206e6f74206578697374604482015260640160405180910390fd5b3660008037600080366000845af43d6000803e80801560bf573d6000f35b3d6000fdfea2646970667358221220538f13fe9ba2996135ecaec357bf229d28151f1367da42b72973cc7f034f28f864736f6c63430008130033c8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c4c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f64654c69624469616d6f6e644375743a204164642066616365742063616e277420624c69624469616d6f6e644375743a204e657720666163657420686173206e6f20636f64654c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e2066",
}

// Diamond3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Diamond3MetaData.ABI instead.
var Diamond3ABI = Diamond3MetaData.ABI

// Diamond3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Diamond3MetaData.Bin instead.
var Diamond3Bin = Diamond3MetaData.Bin

// DeployDiamond3 deploys a new Ethereum contract, binding an instance of Diamond3 to it.
func DeployDiamond3(auth *bind.TransactOpts, backend bind.ContractBackend, _diamondCut []IDiamondCutFacetCut, _args DiamondDiamondArgs) (common.Address, *types.Transaction, *Diamond3, error) {
	parsed, err := Diamond3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Diamond3Bin), backend, _diamondCut, _args)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Diamond3{Diamond3Caller: Diamond3Caller{contract: contract}, Diamond3Transactor: Diamond3Transactor{contract: contract}, Diamond3Filterer: Diamond3Filterer{contract: contract}}, nil
}

// Diamond3 is an auto generated Go binding around an Ethereum contract.
type Diamond3 struct {
	Diamond3Caller     // Read-only binding to the contract
	Diamond3Transactor // Write-only binding to the contract
	Diamond3Filterer   // Log filterer for contract events
}

// Diamond3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Diamond3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Diamond3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Diamond3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Diamond3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Diamond3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Diamond3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Diamond3Session struct {
	Contract     *Diamond3         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Diamond3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Diamond3CallerSession struct {
	Contract *Diamond3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Diamond3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Diamond3TransactorSession struct {
	Contract     *Diamond3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Diamond3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Diamond3Raw struct {
	Contract *Diamond3 // Generic contract binding to access the raw methods on
}

// Diamond3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Diamond3CallerRaw struct {
	Contract *Diamond3Caller // Generic read-only contract binding to access the raw methods on
}

// Diamond3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Diamond3TransactorRaw struct {
	Contract *Diamond3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamond3 creates a new instance of Diamond3, bound to a specific deployed contract.
func NewDiamond3(address common.Address, backend bind.ContractBackend) (*Diamond3, error) {
	contract, err := bindDiamond3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Diamond3{Diamond3Caller: Diamond3Caller{contract: contract}, Diamond3Transactor: Diamond3Transactor{contract: contract}, Diamond3Filterer: Diamond3Filterer{contract: contract}}, nil
}

// NewDiamond3Caller creates a new read-only instance of Diamond3, bound to a specific deployed contract.
func NewDiamond3Caller(address common.Address, caller bind.ContractCaller) (*Diamond3Caller, error) {
	contract, err := bindDiamond3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Diamond3Caller{contract: contract}, nil
}

// NewDiamond3Transactor creates a new write-only instance of Diamond3, bound to a specific deployed contract.
func NewDiamond3Transactor(address common.Address, transactor bind.ContractTransactor) (*Diamond3Transactor, error) {
	contract, err := bindDiamond3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Diamond3Transactor{contract: contract}, nil
}

// NewDiamond3Filterer creates a new log filterer instance of Diamond3, bound to a specific deployed contract.
func NewDiamond3Filterer(address common.Address, filterer bind.ContractFilterer) (*Diamond3Filterer, error) {
	contract, err := bindDiamond3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Diamond3Filterer{contract: contract}, nil
}

// bindDiamond3 binds a generic wrapper to an already deployed contract.
func bindDiamond3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Diamond3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond3 *Diamond3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond3.Contract.Diamond3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond3 *Diamond3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond3.Contract.Diamond3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond3 *Diamond3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond3.Contract.Diamond3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond3 *Diamond3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond3 *Diamond3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond3 *Diamond3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond3.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond3 *Diamond3Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Diamond3.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond3 *Diamond3Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond3.Contract.Fallback(&_Diamond3.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond3 *Diamond3TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond3.Contract.Fallback(&_Diamond3.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond3 *Diamond3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond3 *Diamond3Session) Receive() (*types.Transaction, error) {
	return _Diamond3.Contract.Receive(&_Diamond3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond3 *Diamond3TransactorSession) Receive() (*types.Transaction, error) {
	return _Diamond3.Contract.Receive(&_Diamond3.TransactOpts)
}