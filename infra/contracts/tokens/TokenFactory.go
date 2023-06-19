// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"

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

// TokenFactoryABI is the input ABI used to generate the binding from.
const TokenFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"artifact\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ERC1155_CREATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"artifact\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ERC20_CREATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"artifact\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ERC721_CREATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"deployERC1155\",\"outputs\":[{\"internalType\":\"contractPeggedERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"deployERC20\",\"outputs\":[{\"internalType\":\"contractPeggedERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"deployERC721\",\"outputs\":[{\"internalType\":\"contractPeggedERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beacon20_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beacon721_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beacon1155_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenFactoryBin is the compiled bytecode used for deploying new contracts.
var TokenFactoryBin = "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61215c8061010d6000396000f3fe60806040523480156200001157600080fd5b5060043610620000885760003560e01c80638da5cb5b11620000635780638da5cb5b1462000105578063c0c53b8b1462000127578063e36111ea1462000147578063f2fde38b146200017d5762000088565b80631f7bbfb4146200008d578063715018a614620000c35780637d23ae9514620000cf575b600080fd5b620000ab6004803603810190620000a5919062000b7d565b6200019d565b604051620000ba919062000c8d565b60405180910390f35b620000cd62000341565b005b620000ed6004803603810190620000e7919062000caa565b62000359565b604051620000fc919062000d88565b60405180910390f35b6200010f62000500565b6040516200011e919062000dca565b60405180910390f35b6200014560048036038101906200013f919062000e18565b62000529565b005b6200016560048036038101906200015f919062000caa565b62000677565b60405162000174919062000e99565b60405180910390f35b6200019b600480360381019062000195919062000eb6565b6200081e565b005b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620001d190620009f9565b620001dd919062000f23565b604051809103906000f080158015620001fa573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff16634cd88b7685856040518363ffffffff1660e01b81526004016200023a92919062000fde565b600060405180830381600087803b1580156200025557600080fd5b505af11580156200026a573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b8152600401620002a9919062000dca565b600060405180830381600087803b158015620002c457600080fd5b505af1158015620002d9573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f520f4142a2e54493df9c3829e1a6ecba49fec74835a9f267f2341c006584ac9e60405160405180910390a38091505092915050565b6200034b620008a8565b6200035760006200092d565b565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516200038d90620009f9565b62000399919062000f23565b604051809103906000f080158015620003b6573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff1663a6487c538686866040518463ffffffff1660e01b8152600401620003f89392919062001019565b600060405180830381600087803b1580156200041357600080fd5b505af115801562000428573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b815260040162000467919062000dca565b600060405180830381600087803b1580156200048257600080fd5b505af115801562000497573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f5d8ff0bd03029fa85dc30fbd6ee73cd6b2cfc3d5eecd5d6c8099c0fc67fb31a160405160405180910390a3809150509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600073ffffffffffffffffffffffffffffffffffffffff166200054b62000500565b73ffffffffffffffffffffffffffffffffffffffff1614620005a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200059b90620010bb565b60405180910390fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000672336200092d565b505050565b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620006ab90620009f9565b620006b7919062000f23565b604051809103906000f080158015620006d4573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff1663a6487c538686866040518463ffffffff1660e01b8152600401620007169392919062001019565b600060405180830381600087803b1580156200073157600080fd5b505af115801562000746573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b815260040162000785919062000dca565b600060405180830381600087803b158015620007a057600080fd5b505af1158015620007b5573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8b305b1e96e31e4f8c5853053bc99ae2e88e18f4d29eef36aa37edbb6637493560405160405180910390a3809150509392505050565b62000828620008a8565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200089a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620008919062001153565b60405180910390fd5b620008a5816200092d565b50565b620008b2620009f1565b73ffffffffffffffffffffffffffffffffffffffff16620008d262000500565b73ffffffffffffffffffffffffffffffffffffffff16146200092b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200092290620011c5565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b610f3f80620011e883390190565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000a708262000a25565b810181811067ffffffffffffffff8211171562000a925762000a9162000a36565b5b80604052505050565b600062000aa762000a07565b905062000ab5828262000a65565b919050565b600067ffffffffffffffff82111562000ad85762000ad762000a36565b5b62000ae38262000a25565b9050602081019050919050565b82818337600083830152505050565b600062000b1662000b108462000aba565b62000a9b565b90508281526020810184848401111562000b355762000b3462000a20565b5b62000b4284828562000af0565b509392505050565b600082601f83011262000b625762000b6162000a1b565b5b813562000b7484826020860162000aff565b91505092915050565b6000806040838503121562000b975762000b9662000a11565b5b600083013567ffffffffffffffff81111562000bb85762000bb762000a16565b5b62000bc68582860162000b4a565b925050602083013567ffffffffffffffff81111562000bea5762000be962000a16565b5b62000bf88582860162000b4a565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600062000c4d62000c4762000c418462000c02565b62000c22565b62000c02565b9050919050565b600062000c618262000c2c565b9050919050565b600062000c758262000c54565b9050919050565b62000c878162000c68565b82525050565b600060208201905062000ca4600083018462000c7c565b92915050565b60008060006060848603121562000cc65762000cc562000a11565b5b600084013567ffffffffffffffff81111562000ce75762000ce662000a16565b5b62000cf58682870162000b4a565b935050602084013567ffffffffffffffff81111562000d195762000d1862000a16565b5b62000d278682870162000b4a565b925050604084013567ffffffffffffffff81111562000d4b5762000d4a62000a16565b5b62000d598682870162000b4a565b9150509250925092565b600062000d708262000c54565b9050919050565b62000d828162000d63565b82525050565b600060208201905062000d9f600083018462000d77565b92915050565b600062000db28262000c02565b9050919050565b62000dc48162000da5565b82525050565b600060208201905062000de1600083018462000db9565b92915050565b62000df28162000da5565b811462000dfe57600080fd5b50565b60008135905062000e128162000de7565b92915050565b60008060006060848603121562000e345762000e3362000a11565b5b600062000e448682870162000e01565b935050602062000e578682870162000e01565b925050604062000e6a8682870162000e01565b9150509250925092565b600062000e818262000c54565b9050919050565b62000e938162000e74565b82525050565b600060208201905062000eb0600083018462000e88565b92915050565b60006020828403121562000ecf5762000ece62000a11565b5b600062000edf8482850162000e01565b91505092915050565b600082825260208201905092915050565b50565b600062000f0b60008362000ee8565b915062000f188262000ef9565b600082019050919050565b600060408201905062000f3a600083018462000db9565b818103602083015262000f4d8162000efc565b905092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101562000f9157808201518184015260208101905062000f74565b60008484015250505050565b600062000faa8262000f55565b62000fb6818562000f60565b935062000fc881856020860162000f71565b62000fd38162000a25565b840191505092915050565b6000604082019050818103600083015262000ffa818562000f9d565b9050818103602083015262001010818462000f9d565b90509392505050565b6000606082019050818103600083015262001035818662000f9d565b905081810360208301526200104b818562000f9d565b9050818103604083015262001061818462000f9d565b9050949350505050565b7f696e697469616c697a6564000000000000000000000000000000000000000000600082015250565b6000620010a3600b8362000f60565b9150620010b0826200106b565b602082019050919050565b60006020820190508181036000830152620010d68162001094565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006200113b60268362000f60565b91506200114882620010dd565b604082019050919050565b600060208201905081810360008301526200116e816200112c565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000620011ad60208362000f60565b9150620011ba8262001175565b602082019050919050565b60006020820190508181036000830152620011e0816200119e565b905091905056fe608060405260405162000f3f38038062000f3f8339818101604052810190620000299190620006a6565b6200003d828260006200004560201b60201c565b5050620009c7565b62000056836200013c60201b60201c565b8273ffffffffffffffffffffffffffffffffffffffff167f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e60405160405180910390a2600082511180620000a75750805b156200013757620001358373ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200012391906200070c565b83620002db60201b620000371760201c565b505b505050565b62000152816200031160201b620000641760201c565b62000194576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200018b90620007c5565b60405180910390fd5b6200021b8173ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200020a91906200070c565b6200031160201b620000641760201c565b6200025d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000254906200085d565b60405180910390fd5b80620002977fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6200033460201b620000871760201c565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b606062000309838360405180606001604052806027815260200162000f18602791396200033e60201b60201c565b905092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000819050919050565b60606000808573ffffffffffffffffffffffffffffffffffffffff16856040516200036a9190620008cc565b600060405180830381855af49150503d8060008114620003a7576040519150601f19603f3d011682016040523d82523d6000602084013e620003ac565b606091505b5091509150620003c586838387620003d060201b60201c565b925050509392505050565b60608315620004405760008351036200043757620003f4856200031160201b60201c565b62000436576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200042d9062000935565b60405180910390fd5b5b82905062000453565b6200045283836200045b60201b60201c565b5b949350505050565b6000825111156200046f5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620004a59190620009a3565b60405180910390fd5b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620004ef82620004c2565b9050919050565b6200050181620004e2565b81146200050d57600080fd5b50565b6000815190506200052181620004f6565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200057c8262000531565b810181811067ffffffffffffffff821117156200059e576200059d62000542565b5b80604052505050565b6000620005b3620004ae565b9050620005c1828262000571565b919050565b600067ffffffffffffffff821115620005e457620005e362000542565b5b620005ef8262000531565b9050602081019050919050565b60005b838110156200061c578082015181840152602081019050620005ff565b60008484015250505050565b60006200063f6200063984620005c6565b620005a7565b9050828152602081018484840111156200065e576200065d6200052c565b5b6200066b848285620005fc565b509392505050565b600082601f8301126200068b576200068a62000527565b5b81516200069d84826020860162000628565b91505092915050565b60008060408385031215620006c057620006bf620004b8565b5b6000620006d08582860162000510565b925050602083015167ffffffffffffffff811115620006f457620006f3620004bd565b5b620007028582860162000673565b9150509250929050565b600060208284031215620007255762000724620004b8565b5b6000620007358482850162000510565b91505092915050565b600082825260208201905092915050565b7f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e60008201527f7472616374000000000000000000000000000000000000000000000000000000602082015250565b6000620007ad6025836200073e565b9150620007ba826200074f565b604082019050919050565b60006020820190508181036000830152620007e0816200079e565b9050919050565b7f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960008201527f73206e6f74206120636f6e747261637400000000000000000000000000000000602082015250565b6000620008456030836200073e565b91506200085282620007e7565b604082019050919050565b60006020820190508181036000830152620008788162000836565b9050919050565b600081519050919050565b600081905092915050565b6000620008a2826200087f565b620008ae81856200088a565b9350620008c0818560208601620005fc565b80840191505092915050565b6000620008da828462000895565b915081905092915050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b60006200091d601d836200073e565b91506200092a82620008e5565b602082019050919050565b6000602082019050818103600083015262000950816200090e565b9050919050565b600081519050919050565b60006200096f8262000957565b6200097b81856200073e565b93506200098d818560208601620005fc565b620009988162000531565b840191505092915050565b60006020820190508181036000830152620009bf818462000962565b905092915050565b61054180620009d76000396000f3fe6080604052366100135761001161001d565b005b61001b61001d565b005b610025610091565b610035610030610093565b610110565b565b606061005c83836040518060600160405280602781526020016104e560279139610136565b905092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000819050919050565b565b600061009d6101bc565b73ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061010b919061033b565b905090565b3660008037600080366000845af43d6000803e8060008114610131573d6000f35b3d6000fd5b60606000808573ffffffffffffffffffffffffffffffffffffffff168560405161016091906103d9565b600060405180830381855af49150503d806000811461019b576040519150601f19603f3d011682016040523d82523d6000602084013e6101a0565b606091505b50915091506101b186838387610213565b925050509392505050565b60006101ea7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b610087565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060831561027557600083510361026d5761022d85610064565b61026c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102639061044d565b60405180910390fd5b5b829050610280565b61027f8383610288565b5b949350505050565b60008251111561029b5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102cf91906104c2565b60405180910390fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610308826102dd565b9050919050565b610318816102fd565b811461032357600080fd5b50565b6000815190506103358161030f565b92915050565b600060208284031215610351576103506102d8565b5b600061035f84828501610326565b91505092915050565b600081519050919050565b600081905092915050565b60005b8381101561039c578082015181840152602081019050610381565b60008484015250505050565b60006103b382610368565b6103bd8185610373565b93506103cd81856020860161037e565b80840191505092915050565b60006103e582846103a8565b915081905092915050565b600082825260208201905092915050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b6000610437601d836103f0565b915061044282610401565b602082019050919050565b600060208201905081810360008301526104668161042a565b9050919050565b600081519050919050565b6000601f19601f8301169050919050565b60006104948261046d565b61049e81856103f0565b93506104ae81856020860161037e565b6104b781610478565b840191505092915050565b600060208201905081810360008301526104dc8184610489565b90509291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e2d3045857e1cab9624680014fb7a7a0406fa90030c0970e6d8486d66eb4fcd764736f6c63430008120033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212209bb09da1474c312ab14caced2df2acee4c36636db24a323119b5dfd97e2bd55764736f6c63430008120033"

// DeployTokenFactory deploys a new Conflux contract, binding an instance of TokenFactory to it.
func DeployTokenFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.UnsignedTransaction, *types.Hash, *TokenFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFactoryABI))
	if err != nil {
		return nil, nil, nil, err
	}

	tx, hash, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenFactoryBin), backend)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, hash, &TokenFactory{TokenFactoryCaller: TokenFactoryCaller{contract: contract}, TokenFactoryTransactor: TokenFactoryTransactor{contract: contract}, TokenFactoryFilterer: TokenFactoryFilterer{contract: contract}}, nil
}

// TokenFactory is an auto generated Go binding around an Conflux contract.
type TokenFactory struct {
	TokenFactoryCaller     // Read-only binding to the contract
	TokenFactoryTransactor // Write-only binding to the contract
	TokenFactoryFilterer   // Log filterer for contract events
}

// TokenFactoryCaller is an auto generated read-only Go binding around an Conflux contract.
type TokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type TokenFactoryBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryTransactor is an auto generated write-only Go binding around an Conflux contract.
type TokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type TokenFactoryBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type TokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactorySession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type TokenFactorySession struct {
	Contract     *TokenFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenFactoryCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type TokenFactoryCallerSession struct {
	Contract *TokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenFactoryTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type TokenFactoryTransactorSession struct {
	Contract     *TokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenFactoryRaw is an auto generated low-level Go binding around an Conflux contract.
type TokenFactoryRaw struct {
	Contract *TokenFactory // Generic contract binding to access the raw methods on
}

// TokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type TokenFactoryCallerRaw struct {
	Contract *TokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type TokenFactoryTransactorRaw struct {
	Contract *TokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFactory creates a new instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactory(address types.Address, backend bind.ContractBackend) (*TokenFactory, error) {
	contract, err := bindTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFactory{TokenFactoryCaller: TokenFactoryCaller{contract: contract}, TokenFactoryTransactor: TokenFactoryTransactor{contract: contract}, TokenFactoryFilterer: TokenFactoryFilterer{contract: contract}}, nil
}

// NewTokenFactoryCaller creates a new read-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryCaller(address types.Address, caller bind.ContractCaller) (*TokenFactoryCaller, error) {
	contract, err := bindTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryCaller{contract: contract}, nil
}

// NewTokenFactoryTransactor creates a new write-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryTransactor(address types.Address, transactor bind.ContractTransactor) (*TokenFactoryTransactor, error) {
	contract, err := bindTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTransactor{contract: contract}, nil
}

// NewTokenFactoryFilterer creates a new log filterer instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryFilterer(address types.Address, filterer bind.ContractFilterer) (*TokenFactoryFilterer, error) {
	contract, err := bindTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryFilterer{contract: contract}, nil
}

// NewTokenFactoryCaller creates a new read-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryBulkCaller(address types.Address, caller bind.ContractCaller) (*TokenFactoryBulkCaller, error) {
	contract, err := bindTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryBulkCaller{contract: contract}, nil
}

// NewTokenFactoryBulkTransactor creates a new write-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*TokenFactoryBulkTransactor, error) {
	contract, err := bindTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryBulkTransactor{contract: contract}, nil
}

// bindTokenFactory binds a generic wrapper to an already deployed contract.
func bindTokenFactory(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.TokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactory *TokenFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _TokenFactory.contract.Call(opts, &out, "owner")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactory *TokenFactoryBulkCaller) Owner(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _TokenFactory.contract.GenRequest(opts, "owner")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _TokenFactory.contract.DecodeOutput(&out, rawOut, "owner")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactory *TokenFactorySession) Owner() (common.Address, error) {
	return _TokenFactory.Contract.Owner(&_TokenFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) Owner() (common.Address, error) {
	return _TokenFactory.Contract.Owner(&_TokenFactory.CallOpts)
}

// DeployERC1155 is a paid mutator transaction binding the contract method 0xe36111ea.
//
// Solidity: function deployERC1155(string name_, string symbol_, string uri_) returns(address)
func (_TokenFactory *TokenFactoryTransactor) DeployERC1155(opts *bind.TransactOpts, name_ string, symbol_ string, uri_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.contract.Transact(opts, "deployERC1155", name_, symbol_, uri_)
}

// DeployERC1155 is a paid mutator transaction binding the contract method 0xe36111ea.
//
// Solidity: function deployERC1155(string name_, string symbol_, string uri_) returns(address)
func (_TokenFactory *TokenFactoryBulkTransactor) DeployERC1155(opts *bind.TransactOpts, name_ string, symbol_ string, uri_ string) types.UnsignedTransaction {
	return _TokenFactory.contract.GenUnsignedTransaction(opts, "deployERC1155", name_, symbol_, uri_)
}

// DeployERC1155 is a paid mutator transaction binding the contract method 0xe36111ea.
//
// Solidity: function deployERC1155(string name_, string symbol_, string uri_) returns(address)
func (_TokenFactory *TokenFactorySession) DeployERC1155(name_ string, symbol_ string, uri_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.DeployERC1155(&_TokenFactory.TransactOpts, name_, symbol_, uri_)
}

// DeployERC1155 is a paid mutator transaction binding the contract method 0xe36111ea.
//
// Solidity: function deployERC1155(string name_, string symbol_, string uri_) returns(address)
func (_TokenFactory *TokenFactoryTransactorSession) DeployERC1155(name_ string, symbol_ string, uri_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.DeployERC1155(&_TokenFactory.TransactOpts, name_, symbol_, uri_)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0x1f7bbfb4.
//
// Solidity: function deployERC20(string name_, string symbol_) returns(address)
func (_TokenFactory *TokenFactoryTransactor) DeployERC20(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.contract.Transact(opts, "deployERC20", name_, symbol_)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0x1f7bbfb4.
//
// Solidity: function deployERC20(string name_, string symbol_) returns(address)
func (_TokenFactory *TokenFactoryBulkTransactor) DeployERC20(opts *bind.TransactOpts, name_ string, symbol_ string) types.UnsignedTransaction {
	return _TokenFactory.contract.GenUnsignedTransaction(opts, "deployERC20", name_, symbol_)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0x1f7bbfb4.
//
// Solidity: function deployERC20(string name_, string symbol_) returns(address)
func (_TokenFactory *TokenFactorySession) DeployERC20(name_ string, symbol_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.DeployERC20(&_TokenFactory.TransactOpts, name_, symbol_)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0x1f7bbfb4.
//
// Solidity: function deployERC20(string name_, string symbol_) returns(address)
func (_TokenFactory *TokenFactoryTransactorSession) DeployERC20(name_ string, symbol_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.DeployERC20(&_TokenFactory.TransactOpts, name_, symbol_)
}

// DeployERC721 is a paid mutator transaction binding the contract method 0x7d23ae95.
//
// Solidity: function deployERC721(string name_, string symbol_, string uri_) returns(address)
func (_TokenFactory *TokenFactoryTransactor) DeployERC721(opts *bind.TransactOpts, name_ string, symbol_ string, uri_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.contract.Transact(opts, "deployERC721", name_, symbol_, uri_)
}

// DeployERC721 is a paid mutator transaction binding the contract method 0x7d23ae95.
//
// Solidity: function deployERC721(string name_, string symbol_, string uri_) returns(address)
func (_TokenFactory *TokenFactoryBulkTransactor) DeployERC721(opts *bind.TransactOpts, name_ string, symbol_ string, uri_ string) types.UnsignedTransaction {
	return _TokenFactory.contract.GenUnsignedTransaction(opts, "deployERC721", name_, symbol_, uri_)
}

// DeployERC721 is a paid mutator transaction binding the contract method 0x7d23ae95.
//
// Solidity: function deployERC721(string name_, string symbol_, string uri_) returns(address)
func (_TokenFactory *TokenFactorySession) DeployERC721(name_ string, symbol_ string, uri_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.DeployERC721(&_TokenFactory.TransactOpts, name_, symbol_, uri_)
}

// DeployERC721 is a paid mutator transaction binding the contract method 0x7d23ae95.
//
// Solidity: function deployERC721(string name_, string symbol_, string uri_) returns(address)
func (_TokenFactory *TokenFactoryTransactorSession) DeployERC721(name_ string, symbol_ string, uri_ string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.DeployERC721(&_TokenFactory.TransactOpts, name_, symbol_, uri_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address beacon20_, address beacon721_, address beacon1155_) returns()
func (_TokenFactory *TokenFactoryTransactor) Initialize(opts *bind.TransactOpts, beacon20_ common.Address, beacon721_ common.Address, beacon1155_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.contract.Transact(opts, "initialize", beacon20_, beacon721_, beacon1155_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address beacon20_, address beacon721_, address beacon1155_) returns()
func (_TokenFactory *TokenFactoryBulkTransactor) Initialize(opts *bind.TransactOpts, beacon20_ common.Address, beacon721_ common.Address, beacon1155_ common.Address) types.UnsignedTransaction {
	return _TokenFactory.contract.GenUnsignedTransaction(opts, "initialize", beacon20_, beacon721_, beacon1155_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address beacon20_, address beacon721_, address beacon1155_) returns()
func (_TokenFactory *TokenFactorySession) Initialize(beacon20_ common.Address, beacon721_ common.Address, beacon1155_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.Initialize(&_TokenFactory.TransactOpts, beacon20_, beacon721_, beacon1155_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address beacon20_, address beacon721_, address beacon1155_) returns()
func (_TokenFactory *TokenFactoryTransactorSession) Initialize(beacon20_ common.Address, beacon721_ common.Address, beacon1155_ common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.Initialize(&_TokenFactory.TransactOpts, beacon20_, beacon721_, beacon1155_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactory *TokenFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactory *TokenFactoryBulkTransactor) RenounceOwnership(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _TokenFactory.contract.GenUnsignedTransaction(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactory *TokenFactorySession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.RenounceOwnership(&_TokenFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactory *TokenFactoryTransactorSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.RenounceOwnership(&_TokenFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactory *TokenFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactory *TokenFactoryBulkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) types.UnsignedTransaction {
	return _TokenFactory.contract.GenUnsignedTransaction(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactory *TokenFactorySession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.TransferOwnership(&_TokenFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactory *TokenFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TokenFactory.Contract.TransferOwnership(&_TokenFactory.TransactOpts, newOwner)
}

// TokenFactoryERC1155CREATEDIterator is returned from FilterERC1155CREATED and is used to iterate over the raw logs and unpacked data for ERC1155CREATED events raised by the TokenFactory contract.
type TokenFactoryERC1155CREATEDIterator struct {
	Event *TokenFactoryERC1155CREATED // Event containing the contract specifics and raw log

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
func (it *TokenFactoryERC1155CREATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryERC1155CREATED)
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
		it.Event = new(TokenFactoryERC1155CREATED)
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
func (it *TokenFactoryERC1155CREATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryERC1155CREATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryERC1155CREATED represents a ERC1155CREATED event raised by the TokenFactory contract.
type TokenFactoryERC1155CREATED struct {
	Artifact common.Address
	Creator  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// TokenFactoryERC1155CREATEDOrChainReorg represents a ERC1155CREATED subscription event raised by the TokenFactory contract.
type TokenFactoryERC1155CREATEDOrChainReorg struct {
	Event      *TokenFactoryERC1155CREATED
	ChainReorg *types.ChainReorg
}

// FilterERC1155CREATED is a free log retrieval operation binding the contract event 0x8b305b1e96e31e4f8c5853053bc99ae2e88e18f4d29eef36aa37edbb66374935.
//
// Solidity: event ERC1155_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) FilterERC1155CREATED(opts *bind.FilterOpts, artifact []common.Address, creator []common.Address) (*TokenFactoryERC1155CREATEDIterator, error) {

	var artifactRule []interface{}
	for _, artifactItem := range artifact {
		artifactRule = append(artifactRule, artifactItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, err := _TokenFactory.contract.FilterLogs(opts, "ERC1155_CREATED", artifactRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryERC1155CREATEDIterator{contract: _TokenFactory.contract, event: "ERC1155_CREATED", logs: logs}, nil
}

// WatchERC1155CREATED is a free log subscription operation binding the contract event 0x8b305b1e96e31e4f8c5853053bc99ae2e88e18f4d29eef36aa37edbb66374935.
//
// Solidity: event ERC1155_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) WatchERC1155CREATED(opts *bind.WatchOpts, sink chan<- *TokenFactoryERC1155CREATEDOrChainReorg, artifact []common.Address, creator []common.Address) (event.Subscription, error) {

	var artifactRule []interface{}
	for _, artifactItem := range artifact {
		artifactRule = append(artifactRule, artifactItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "ERC1155_CREATED", artifactRule, creatorRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryERC1155CREATEDOrChainReorg)
				event.Event = new(TokenFactoryERC1155CREATED)

				if log.ChainReorg == nil {
					if err := _TokenFactory.contract.UnpackLog(event.Event, "ERC1155_CREATED", *log.Log); err != nil {
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

// ParseERC1155CREATED is a log parse operation binding the contract event 0x8b305b1e96e31e4f8c5853053bc99ae2e88e18f4d29eef36aa37edbb66374935.
//
// Solidity: event ERC1155_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) ParseERC1155CREATED(log types.Log) (*TokenFactoryERC1155CREATED, error) {
	event := new(TokenFactoryERC1155CREATED)
	if err := _TokenFactory.contract.UnpackLog(event, "ERC1155_CREATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryERC20CREATEDIterator is returned from FilterERC20CREATED and is used to iterate over the raw logs and unpacked data for ERC20CREATED events raised by the TokenFactory contract.
type TokenFactoryERC20CREATEDIterator struct {
	Event *TokenFactoryERC20CREATED // Event containing the contract specifics and raw log

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
func (it *TokenFactoryERC20CREATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryERC20CREATED)
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
		it.Event = new(TokenFactoryERC20CREATED)
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
func (it *TokenFactoryERC20CREATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryERC20CREATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryERC20CREATED represents a ERC20CREATED event raised by the TokenFactory contract.
type TokenFactoryERC20CREATED struct {
	Artifact common.Address
	Creator  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// TokenFactoryERC20CREATEDOrChainReorg represents a ERC20CREATED subscription event raised by the TokenFactory contract.
type TokenFactoryERC20CREATEDOrChainReorg struct {
	Event      *TokenFactoryERC20CREATED
	ChainReorg *types.ChainReorg
}

// FilterERC20CREATED is a free log retrieval operation binding the contract event 0x520f4142a2e54493df9c3829e1a6ecba49fec74835a9f267f2341c006584ac9e.
//
// Solidity: event ERC20_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) FilterERC20CREATED(opts *bind.FilterOpts, artifact []common.Address, creator []common.Address) (*TokenFactoryERC20CREATEDIterator, error) {

	var artifactRule []interface{}
	for _, artifactItem := range artifact {
		artifactRule = append(artifactRule, artifactItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, err := _TokenFactory.contract.FilterLogs(opts, "ERC20_CREATED", artifactRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryERC20CREATEDIterator{contract: _TokenFactory.contract, event: "ERC20_CREATED", logs: logs}, nil
}

// WatchERC20CREATED is a free log subscription operation binding the contract event 0x520f4142a2e54493df9c3829e1a6ecba49fec74835a9f267f2341c006584ac9e.
//
// Solidity: event ERC20_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) WatchERC20CREATED(opts *bind.WatchOpts, sink chan<- *TokenFactoryERC20CREATEDOrChainReorg, artifact []common.Address, creator []common.Address) (event.Subscription, error) {

	var artifactRule []interface{}
	for _, artifactItem := range artifact {
		artifactRule = append(artifactRule, artifactItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "ERC20_CREATED", artifactRule, creatorRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryERC20CREATEDOrChainReorg)
				event.Event = new(TokenFactoryERC20CREATED)

				if log.ChainReorg == nil {
					if err := _TokenFactory.contract.UnpackLog(event.Event, "ERC20_CREATED", *log.Log); err != nil {
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

// ParseERC20CREATED is a log parse operation binding the contract event 0x520f4142a2e54493df9c3829e1a6ecba49fec74835a9f267f2341c006584ac9e.
//
// Solidity: event ERC20_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) ParseERC20CREATED(log types.Log) (*TokenFactoryERC20CREATED, error) {
	event := new(TokenFactoryERC20CREATED)
	if err := _TokenFactory.contract.UnpackLog(event, "ERC20_CREATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryERC721CREATEDIterator is returned from FilterERC721CREATED and is used to iterate over the raw logs and unpacked data for ERC721CREATED events raised by the TokenFactory contract.
type TokenFactoryERC721CREATEDIterator struct {
	Event *TokenFactoryERC721CREATED // Event containing the contract specifics and raw log

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
func (it *TokenFactoryERC721CREATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryERC721CREATED)
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
		it.Event = new(TokenFactoryERC721CREATED)
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
func (it *TokenFactoryERC721CREATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryERC721CREATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryERC721CREATED represents a ERC721CREATED event raised by the TokenFactory contract.
type TokenFactoryERC721CREATED struct {
	Artifact common.Address
	Creator  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// TokenFactoryERC721CREATEDOrChainReorg represents a ERC721CREATED subscription event raised by the TokenFactory contract.
type TokenFactoryERC721CREATEDOrChainReorg struct {
	Event      *TokenFactoryERC721CREATED
	ChainReorg *types.ChainReorg
}

// FilterERC721CREATED is a free log retrieval operation binding the contract event 0x5d8ff0bd03029fa85dc30fbd6ee73cd6b2cfc3d5eecd5d6c8099c0fc67fb31a1.
//
// Solidity: event ERC721_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) FilterERC721CREATED(opts *bind.FilterOpts, artifact []common.Address, creator []common.Address) (*TokenFactoryERC721CREATEDIterator, error) {

	var artifactRule []interface{}
	for _, artifactItem := range artifact {
		artifactRule = append(artifactRule, artifactItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, err := _TokenFactory.contract.FilterLogs(opts, "ERC721_CREATED", artifactRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryERC721CREATEDIterator{contract: _TokenFactory.contract, event: "ERC721_CREATED", logs: logs}, nil
}

// WatchERC721CREATED is a free log subscription operation binding the contract event 0x5d8ff0bd03029fa85dc30fbd6ee73cd6b2cfc3d5eecd5d6c8099c0fc67fb31a1.
//
// Solidity: event ERC721_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) WatchERC721CREATED(opts *bind.WatchOpts, sink chan<- *TokenFactoryERC721CREATEDOrChainReorg, artifact []common.Address, creator []common.Address) (event.Subscription, error) {

	var artifactRule []interface{}
	for _, artifactItem := range artifact {
		artifactRule = append(artifactRule, artifactItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "ERC721_CREATED", artifactRule, creatorRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryERC721CREATEDOrChainReorg)
				event.Event = new(TokenFactoryERC721CREATED)

				if log.ChainReorg == nil {
					if err := _TokenFactory.contract.UnpackLog(event.Event, "ERC721_CREATED", *log.Log); err != nil {
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

// ParseERC721CREATED is a log parse operation binding the contract event 0x5d8ff0bd03029fa85dc30fbd6ee73cd6b2cfc3d5eecd5d6c8099c0fc67fb31a1.
//
// Solidity: event ERC721_CREATED(address indexed artifact, address indexed creator)
func (_TokenFactory *TokenFactoryFilterer) ParseERC721CREATED(log types.Log) (*TokenFactoryERC721CREATED, error) {
	event := new(TokenFactoryERC721CREATED)
	if err := _TokenFactory.contract.UnpackLog(event, "ERC721_CREATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenFactory contract.
type TokenFactoryOwnershipTransferredIterator struct {
	Event *TokenFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryOwnershipTransferred)
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
		it.Event = new(TokenFactoryOwnershipTransferred)
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
func (it *TokenFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the TokenFactory contract.
type TokenFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// TokenFactoryOwnershipTransferredOrChainReorg represents a OwnershipTransferred subscription event raised by the TokenFactory contract.
type TokenFactoryOwnershipTransferredOrChainReorg struct {
	Event      *TokenFactoryOwnershipTransferred
	ChainReorg *types.ChainReorg
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFactory *TokenFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, err := _TokenFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryOwnershipTransferredIterator{contract: _TokenFactory.contract, event: "OwnershipTransferred", logs: logs}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFactory *TokenFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenFactoryOwnershipTransferredOrChainReorg, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryOwnershipTransferredOrChainReorg)
				event.Event = new(TokenFactoryOwnershipTransferred)

				if log.ChainReorg == nil {
					if err := _TokenFactory.contract.UnpackLog(event.Event, "OwnershipTransferred", *log.Log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFactory *TokenFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*TokenFactoryOwnershipTransferred, error) {
	event := new(TokenFactoryOwnershipTransferred)
	if err := _TokenFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
