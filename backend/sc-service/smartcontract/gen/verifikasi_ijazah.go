// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifikasi_ijazah

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

// VerifikasiIjazahMetaData contains all meta data concerning the VerifikasiIjazah contract.
var VerifikasiIjazahMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"degreeHash\",\"type\":\"bytes32\"}],\"name\":\"DegreeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"degreeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sekolah\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"issueDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsUrl\",\"type\":\"string\"}],\"name\":\"DegreeIssued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"degrees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"degreeHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sekolah\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"issueDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsUrl\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_degreeHash\",\"type\":\"bytes32\"}],\"name\":\"deleteDegree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_degreeHash\",\"type\":\"bytes32\"}],\"name\":\"getDegree\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_degreeHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_sekolah\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_issueDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ipfsUrl\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_mataPelajaran\",\"type\":\"string[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_nilai\",\"type\":\"uint8[]\"}],\"name\":\"issueDegree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listDegrees\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611be4806100606000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063578d30da146100675780638da5cb5b1461008357806391f2bceb146100a1578063bc50b3cd146100bd578063d04a419b146100f0578063f266c79114610124575b600080fd5b610081600480360381019061007c9190610ff7565b610142565b005b61008b6103fc565b6040516100989190611135565b60405180910390f35b6100bb60048036038101906100b69190611150565b610420565b005b6100d760048036038101906100d29190611150565b61059a565b6040516100e7949392919061121a565b60405180910390f35b61010a60048036038101906101059190611150565b6106da565b60405161011b959493929190611437565b60405180910390f35b61012c610a75565b6040516101399190611564565b60405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101c7906115f8565b60405180910390fd5b6000801b60016000888152602001908152602001600020600001541461022b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022290611664565b60405180910390fd5b805182511461026f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610266906116f6565b60405180910390fd5b6000600160008881526020019081526020016000209050868160000181905550858160010190816102a09190611922565b50848160020181905550838160030190816102bb9190611922565b5060005b835181101561038d578160040160405180604001604052808684815181106102ea576102e96119f4565b5b6020026020010151815260200185848151811061030a576103096119f4565b5b602002602001015160ff16815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190816103569190611922565b5060208201518160010160006101000a81548160ff021916908360ff1602179055505050808061038590611a52565b9150506102bf565b506002879080600181540180825580915050600190039060005260206000200160009091909190915055867fb604f6e596d89a3657daccdbb69c509b6039e34c5cff6500a0ff22efec8ac4328787876040516103eb93929190611a9a565b60405180910390a250505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146104ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a5906115f8565b60405180910390fd5b6000801b600160008381526020019081526020016000206000015403610509576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050090611b2b565b60405180910390fd5b600160008281526020019081526020016000206000808201600090556001820160006105359190610b93565b600282016000905560038201600061054d9190610b93565b60048201600061055d9190610bd3565b505061056a600282610acd565b807f2378393e5e45ca60b5cc050d80fa00b05dc03948e54f0260efb5af5229b9ca3e60405160405180910390a250565b60016020528060005260406000206000915090508060000154908060010180546105c390611745565b80601f01602080910402602001604051908101604052809291908181526020018280546105ef90611745565b801561063c5780601f106106115761010080835404028352916020019161063c565b820191906000526020600020905b81548152906001019060200180831161061f57829003601f168201915b50505050509080600201549080600301805461065790611745565b80601f016020809104026020016040519081016040528092919081815260200182805461068390611745565b80156106d05780601f106106a5576101008083540402835291602001916106d0565b820191906000526020600020905b8154815290600101906020018083116106b357829003601f168201915b5050505050905084565b60606000606080606060006001600088815260200190815260200160002090506000801b816000015403610743576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073a90611b2b565b60405180910390fd5b60008160040180549050905060008167ffffffffffffffff81111561076b5761076a610cb4565b5b60405190808252806020026020018201604052801561079e57816020015b60608152602001906001900390816107895790505b50905060008267ffffffffffffffff8111156107bd576107bc610cb4565b5b6040519080825280602002602001820160405280156107eb5781602001602082028036833780820191505090505b50905060005b838110156109365784600401818154811061080f5761080e6119f4565b5b9060005260206000209060020201600001805461082b90611745565b80601f016020809104026020016040519081016040528092919081815260200182805461085790611745565b80156108a45780601f10610879576101008083540402835291602001916108a4565b820191906000526020600020905b81548152906001019060200180831161088757829003601f168201915b50505050508382815181106108bc576108bb6119f4565b5b60200260200101819052508460040181815481106108dd576108dc6119f4565b5b906000526020600020906002020160010160009054906101000a900460ff1682828151811061090f5761090e6119f4565b5b602002602001019060ff16908160ff1681525050808061092e90611a52565b9150506107f1565b5083600101846002015485600301848484805461095290611745565b80601f016020809104026020016040519081016040528092919081815260200182805461097e90611745565b80156109cb5780601f106109a0576101008083540402835291602001916109cb565b820191906000526020600020905b8154815290600101906020018083116109ae57829003601f168201915b505050505094508280546109de90611745565b80601f0160208091040260200160405190810160405280929190818152602001828054610a0a90611745565b8015610a575780601f10610a2c57610100808354040283529160200191610a57565b820191906000526020600020905b815481529060010190602001808311610a3a57829003601f168201915b50505050509250985098509850985098505050505091939590929450565b60606002805480602002602001604051908101604052809291908181526020018280548015610ac357602002820191906000526020600020905b815481526020019060010190808311610aaf575b5050505050905090565b60005b8280549050811015610b8e5781838281548110610af057610aef6119f4565b5b906000526020600020015403610b7b578260018480549050610b129190611b4b565b81548110610b2357610b226119f4565b5b9060005260206000200154838281548110610b4157610b406119f4565b5b906000526020600020018190555082805480610b6057610b5f611b7f565b5b60019003818190600052602060002001600090559055610b8e565b8080610b8690611a52565b915050610ad0565b505050565b508054610b9f90611745565b6000825580601f10610bb15750610bd0565b601f016020900490600052602060002090810190610bcf9190610bf7565b5b50565b5080546000825560020290600052602060002090810190610bf49190610c14565b50565b5b80821115610c10576000816000905550600101610bf8565b5090565b5b80821115610c4b5760008082016000610c2e9190610b93565b6001820160006101000a81549060ff021916905550600201610c15565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610c7681610c63565b8114610c8157600080fd5b50565b600081359050610c9381610c6d565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610cec82610ca3565b810181811067ffffffffffffffff82111715610d0b57610d0a610cb4565b5b80604052505050565b6000610d1e610c4f565b9050610d2a8282610ce3565b919050565b600067ffffffffffffffff821115610d4a57610d49610cb4565b5b610d5382610ca3565b9050602081019050919050565b82818337600083830152505050565b6000610d82610d7d84610d2f565b610d14565b905082815260208101848484011115610d9e57610d9d610c9e565b5b610da9848285610d60565b509392505050565b600082601f830112610dc657610dc5610c99565b5b8135610dd6848260208601610d6f565b91505092915050565b6000819050919050565b610df281610ddf565b8114610dfd57600080fd5b50565b600081359050610e0f81610de9565b92915050565b600067ffffffffffffffff821115610e3057610e2f610cb4565b5b602082029050602081019050919050565b600080fd5b6000610e59610e5484610e15565b610d14565b90508083825260208201905060208402830185811115610e7c57610e7b610e41565b5b835b81811015610ec357803567ffffffffffffffff811115610ea157610ea0610c99565b5b808601610eae8982610db1565b85526020850194505050602081019050610e7e565b5050509392505050565b600082601f830112610ee257610ee1610c99565b5b8135610ef2848260208601610e46565b91505092915050565b600067ffffffffffffffff821115610f1657610f15610cb4565b5b602082029050602081019050919050565b600060ff82169050919050565b610f3d81610f27565b8114610f4857600080fd5b50565b600081359050610f5a81610f34565b92915050565b6000610f73610f6e84610efb565b610d14565b90508083825260208201905060208402830185811115610f9657610f95610e41565b5b835b81811015610fbf5780610fab8882610f4b565b845260208401935050602081019050610f98565b5050509392505050565b600082601f830112610fde57610fdd610c99565b5b8135610fee848260208601610f60565b91505092915050565b60008060008060008060c0878903121561101457611013610c59565b5b600061102289828a01610c84565b965050602087013567ffffffffffffffff81111561104357611042610c5e565b5b61104f89828a01610db1565b955050604061106089828a01610e00565b945050606087013567ffffffffffffffff81111561108157611080610c5e565b5b61108d89828a01610db1565b935050608087013567ffffffffffffffff8111156110ae576110ad610c5e565b5b6110ba89828a01610ecd565b92505060a087013567ffffffffffffffff8111156110db576110da610c5e565b5b6110e789828a01610fc9565b9150509295509295509295565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061111f826110f4565b9050919050565b61112f81611114565b82525050565b600060208201905061114a6000830184611126565b92915050565b60006020828403121561116657611165610c59565b5b600061117484828501610c84565b91505092915050565b61118681610c63565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156111c65780820151818401526020810190506111ab565b60008484015250505050565b60006111dd8261118c565b6111e78185611197565b93506111f78185602086016111a8565b61120081610ca3565b840191505092915050565b61121481610ddf565b82525050565b600060808201905061122f600083018761117d565b818103602083015261124181866111d2565b9050611250604083018561120b565b818103606083015261126281846111d2565b905095945050505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b60006112b58261118c565b6112bf8185611299565b93506112cf8185602086016111a8565b6112d881610ca3565b840191505092915050565b60006112ef83836112aa565b905092915050565b6000602082019050919050565b600061130f8261126d565b6113198185611278565b93508360208202850161132b85611289565b8060005b85811015611367578484038952815161134885826112e3565b9450611353836112f7565b925060208a0199505060018101905061132f565b50829750879550505050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6113ae81610f27565b82525050565b60006113c083836113a5565b60208301905092915050565b6000602082019050919050565b60006113e482611379565b6113ee8185611384565b93506113f983611395565b8060005b8381101561142a57815161141188826113b4565b975061141c836113cc565b9250506001810190506113fd565b5085935050505092915050565b600060a082019050818103600083015261145181886111d2565b9050611460602083018761120b565b818103604083015261147281866111d2565b905081810360608301526114868185611304565b9050818103608083015261149a81846113d9565b90509695505050505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6114db81610c63565b82525050565b60006114ed83836114d2565b60208301905092915050565b6000602082019050919050565b6000611511826114a6565b61151b81856114b1565b9350611526836114c2565b8060005b8381101561155757815161153e88826114e1565b9750611549836114f9565b92505060018101905061152a565b5085935050505092915050565b6000602082019050818103600083015261157e8184611506565b905092915050565b7f48616e79612070656d696c696b2079616e672062697361206d656c616b756b6160008201527f6e20696e692e0000000000000000000000000000000000000000000000000000602082015250565b60006115e2602683611197565b91506115ed82611586565b604082019050919050565b60006020820190508181036000830152611611816115d5565b9050919050565b7f496a617a6168207375646168207465726461667461722e000000000000000000600082015250565b600061164e601783611197565b915061165982611618565b602082019050919050565b6000602082019050818103600083015261167d81611641565b9050919050565b7f4a756d6c6168206d6174612070656c616a6172616e2064616e206e696c61692060008201527f68617275732073616d612e000000000000000000000000000000000000000000602082015250565b60006116e0602b83611197565b91506116eb82611684565b604082019050919050565b6000602082019050818103600083015261170f816116d3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061175d57607f821691505b6020821081036117705761176f611716565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026117d87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261179b565b6117e2868361179b565b95508019841693508086168417925050509392505050565b6000819050919050565b600061181f61181a61181584610ddf565b6117fa565b610ddf565b9050919050565b6000819050919050565b61183983611804565b61184d61184582611826565b8484546117a8565b825550505050565b600090565b611862611855565b61186d818484611830565b505050565b5b818110156118915761188660008261185a565b600181019050611873565b5050565b601f8211156118d6576118a781611776565b6118b08461178b565b810160208510156118bf578190505b6118d36118cb8561178b565b830182611872565b50505b505050565b600082821c905092915050565b60006118f9600019846008026118db565b1980831691505092915050565b600061191283836118e8565b9150826002028217905092915050565b61192b8261118c565b67ffffffffffffffff81111561194457611943610cb4565b5b61194e8254611745565b611959828285611895565b600060209050601f83116001811461198c576000841561197a578287015190505b6119848582611906565b8655506119ec565b601f19841661199a86611776565b60005b828110156119c25784890151825560018201915060208501945060208101905061199d565b868310156119df57848901516119db601f8916826118e8565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611a5d82610ddf565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a8f57611a8e611a23565b5b600182019050919050565b60006060820190508181036000830152611ab481866111d2565b9050611ac3602083018561120b565b8181036040830152611ad581846111d2565b9050949350505050565b7f496a617a616820746964616b20646974656d756b616e2e000000000000000000600082015250565b6000611b15601783611197565b9150611b2082611adf565b602082019050919050565b60006020820190508181036000830152611b4481611b08565b9050919050565b6000611b5682610ddf565b9150611b6183610ddf565b9250828203905081811115611b7957611b78611a23565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220252296abe5de146584859c1d3ca0514cad48a72a2d06ea8770cf6e7dd075573b64736f6c63430008130033",
}

// VerifikasiIjazahABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifikasiIjazahMetaData.ABI instead.
var VerifikasiIjazahABI = VerifikasiIjazahMetaData.ABI

// VerifikasiIjazahBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifikasiIjazahMetaData.Bin instead.
var VerifikasiIjazahBin = VerifikasiIjazahMetaData.Bin

// DeployVerifikasiIjazah deploys a new Ethereum contract, binding an instance of VerifikasiIjazah to it.
func DeployVerifikasiIjazah(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerifikasiIjazah, error) {
	parsed, err := VerifikasiIjazahMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifikasiIjazahBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifikasiIjazah{VerifikasiIjazahCaller: VerifikasiIjazahCaller{contract: contract}, VerifikasiIjazahTransactor: VerifikasiIjazahTransactor{contract: contract}, VerifikasiIjazahFilterer: VerifikasiIjazahFilterer{contract: contract}}, nil
}

// VerifikasiIjazah is an auto generated Go binding around an Ethereum contract.
type VerifikasiIjazah struct {
	VerifikasiIjazahCaller     // Read-only binding to the contract
	VerifikasiIjazahTransactor // Write-only binding to the contract
	VerifikasiIjazahFilterer   // Log filterer for contract events
}

// VerifikasiIjazahCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifikasiIjazahCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifikasiIjazahTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifikasiIjazahTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifikasiIjazahFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifikasiIjazahFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifikasiIjazahSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifikasiIjazahSession struct {
	Contract     *VerifikasiIjazah // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifikasiIjazahCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifikasiIjazahCallerSession struct {
	Contract *VerifikasiIjazahCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VerifikasiIjazahTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifikasiIjazahTransactorSession struct {
	Contract     *VerifikasiIjazahTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VerifikasiIjazahRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifikasiIjazahRaw struct {
	Contract *VerifikasiIjazah // Generic contract binding to access the raw methods on
}

// VerifikasiIjazahCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifikasiIjazahCallerRaw struct {
	Contract *VerifikasiIjazahCaller // Generic read-only contract binding to access the raw methods on
}

// VerifikasiIjazahTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifikasiIjazahTransactorRaw struct {
	Contract *VerifikasiIjazahTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifikasiIjazah creates a new instance of VerifikasiIjazah, bound to a specific deployed contract.
func NewVerifikasiIjazah(address common.Address, backend bind.ContractBackend) (*VerifikasiIjazah, error) {
	contract, err := bindVerifikasiIjazah(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifikasiIjazah{VerifikasiIjazahCaller: VerifikasiIjazahCaller{contract: contract}, VerifikasiIjazahTransactor: VerifikasiIjazahTransactor{contract: contract}, VerifikasiIjazahFilterer: VerifikasiIjazahFilterer{contract: contract}}, nil
}

// NewVerifikasiIjazahCaller creates a new read-only instance of VerifikasiIjazah, bound to a specific deployed contract.
func NewVerifikasiIjazahCaller(address common.Address, caller bind.ContractCaller) (*VerifikasiIjazahCaller, error) {
	contract, err := bindVerifikasiIjazah(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifikasiIjazahCaller{contract: contract}, nil
}

// NewVerifikasiIjazahTransactor creates a new write-only instance of VerifikasiIjazah, bound to a specific deployed contract.
func NewVerifikasiIjazahTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifikasiIjazahTransactor, error) {
	contract, err := bindVerifikasiIjazah(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifikasiIjazahTransactor{contract: contract}, nil
}

// NewVerifikasiIjazahFilterer creates a new log filterer instance of VerifikasiIjazah, bound to a specific deployed contract.
func NewVerifikasiIjazahFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifikasiIjazahFilterer, error) {
	contract, err := bindVerifikasiIjazah(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifikasiIjazahFilterer{contract: contract}, nil
}

// bindVerifikasiIjazah binds a generic wrapper to an already deployed contract.
func bindVerifikasiIjazah(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifikasiIjazahMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifikasiIjazah *VerifikasiIjazahRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifikasiIjazah.Contract.VerifikasiIjazahCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifikasiIjazah *VerifikasiIjazahRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifikasiIjazah.Contract.VerifikasiIjazahTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifikasiIjazah *VerifikasiIjazahRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifikasiIjazah.Contract.VerifikasiIjazahTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifikasiIjazah *VerifikasiIjazahCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifikasiIjazah.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifikasiIjazah *VerifikasiIjazahTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifikasiIjazah.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifikasiIjazah *VerifikasiIjazahTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifikasiIjazah.Contract.contract.Transact(opts, method, params...)
}

// Degrees is a free data retrieval call binding the contract method 0xbc50b3cd.
//
// Solidity: function degrees(bytes32 ) view returns(bytes32 degreeHash, string sekolah, uint256 issueDate, string ipfsUrl)
func (_VerifikasiIjazah *VerifikasiIjazahCaller) Degrees(opts *bind.CallOpts, arg0 [32]byte) (struct {
	DegreeHash [32]byte
	Sekolah    string
	IssueDate  *big.Int
	IpfsUrl    string
}, error) {
	var out []interface{}
	err := _VerifikasiIjazah.contract.Call(opts, &out, "degrees", arg0)

	outstruct := new(struct {
		DegreeHash [32]byte
		Sekolah    string
		IssueDate  *big.Int
		IpfsUrl    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DegreeHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Sekolah = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.IssueDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IpfsUrl = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Degrees is a free data retrieval call binding the contract method 0xbc50b3cd.
//
// Solidity: function degrees(bytes32 ) view returns(bytes32 degreeHash, string sekolah, uint256 issueDate, string ipfsUrl)
func (_VerifikasiIjazah *VerifikasiIjazahSession) Degrees(arg0 [32]byte) (struct {
	DegreeHash [32]byte
	Sekolah    string
	IssueDate  *big.Int
	IpfsUrl    string
}, error) {
	return _VerifikasiIjazah.Contract.Degrees(&_VerifikasiIjazah.CallOpts, arg0)
}

// Degrees is a free data retrieval call binding the contract method 0xbc50b3cd.
//
// Solidity: function degrees(bytes32 ) view returns(bytes32 degreeHash, string sekolah, uint256 issueDate, string ipfsUrl)
func (_VerifikasiIjazah *VerifikasiIjazahCallerSession) Degrees(arg0 [32]byte) (struct {
	DegreeHash [32]byte
	Sekolah    string
	IssueDate  *big.Int
	IpfsUrl    string
}, error) {
	return _VerifikasiIjazah.Contract.Degrees(&_VerifikasiIjazah.CallOpts, arg0)
}

// GetDegree is a free data retrieval call binding the contract method 0xd04a419b.
//
// Solidity: function getDegree(bytes32 _degreeHash) view returns(string, uint256, string, string[], uint8[])
func (_VerifikasiIjazah *VerifikasiIjazahCaller) GetDegree(opts *bind.CallOpts, _degreeHash [32]byte) (string, *big.Int, string, []string, []uint8, error) {
	var out []interface{}
	err := _VerifikasiIjazah.contract.Call(opts, &out, "getDegree", _degreeHash)

	if err != nil {
		return *new(string), *new(*big.Int), *new(string), *new([]string), *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new([]string)).(*[]string)
	out4 := *abi.ConvertType(out[4], new([]uint8)).(*[]uint8)

	return out0, out1, out2, out3, out4, err

}

// GetDegree is a free data retrieval call binding the contract method 0xd04a419b.
//
// Solidity: function getDegree(bytes32 _degreeHash) view returns(string, uint256, string, string[], uint8[])
func (_VerifikasiIjazah *VerifikasiIjazahSession) GetDegree(_degreeHash [32]byte) (string, *big.Int, string, []string, []uint8, error) {
	return _VerifikasiIjazah.Contract.GetDegree(&_VerifikasiIjazah.CallOpts, _degreeHash)
}

// GetDegree is a free data retrieval call binding the contract method 0xd04a419b.
//
// Solidity: function getDegree(bytes32 _degreeHash) view returns(string, uint256, string, string[], uint8[])
func (_VerifikasiIjazah *VerifikasiIjazahCallerSession) GetDegree(_degreeHash [32]byte) (string, *big.Int, string, []string, []uint8, error) {
	return _VerifikasiIjazah.Contract.GetDegree(&_VerifikasiIjazah.CallOpts, _degreeHash)
}

// ListDegrees is a free data retrieval call binding the contract method 0xf266c791.
//
// Solidity: function listDegrees() view returns(bytes32[])
func (_VerifikasiIjazah *VerifikasiIjazahCaller) ListDegrees(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _VerifikasiIjazah.contract.Call(opts, &out, "listDegrees")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ListDegrees is a free data retrieval call binding the contract method 0xf266c791.
//
// Solidity: function listDegrees() view returns(bytes32[])
func (_VerifikasiIjazah *VerifikasiIjazahSession) ListDegrees() ([][32]byte, error) {
	return _VerifikasiIjazah.Contract.ListDegrees(&_VerifikasiIjazah.CallOpts)
}

// ListDegrees is a free data retrieval call binding the contract method 0xf266c791.
//
// Solidity: function listDegrees() view returns(bytes32[])
func (_VerifikasiIjazah *VerifikasiIjazahCallerSession) ListDegrees() ([][32]byte, error) {
	return _VerifikasiIjazah.Contract.ListDegrees(&_VerifikasiIjazah.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifikasiIjazah *VerifikasiIjazahCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifikasiIjazah.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifikasiIjazah *VerifikasiIjazahSession) Owner() (common.Address, error) {
	return _VerifikasiIjazah.Contract.Owner(&_VerifikasiIjazah.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifikasiIjazah *VerifikasiIjazahCallerSession) Owner() (common.Address, error) {
	return _VerifikasiIjazah.Contract.Owner(&_VerifikasiIjazah.CallOpts)
}

// DeleteDegree is a paid mutator transaction binding the contract method 0x91f2bceb.
//
// Solidity: function deleteDegree(bytes32 _degreeHash) returns()
func (_VerifikasiIjazah *VerifikasiIjazahTransactor) DeleteDegree(opts *bind.TransactOpts, _degreeHash [32]byte) (*types.Transaction, error) {
	return _VerifikasiIjazah.contract.Transact(opts, "deleteDegree", _degreeHash)
}

// DeleteDegree is a paid mutator transaction binding the contract method 0x91f2bceb.
//
// Solidity: function deleteDegree(bytes32 _degreeHash) returns()
func (_VerifikasiIjazah *VerifikasiIjazahSession) DeleteDegree(_degreeHash [32]byte) (*types.Transaction, error) {
	return _VerifikasiIjazah.Contract.DeleteDegree(&_VerifikasiIjazah.TransactOpts, _degreeHash)
}

// DeleteDegree is a paid mutator transaction binding the contract method 0x91f2bceb.
//
// Solidity: function deleteDegree(bytes32 _degreeHash) returns()
func (_VerifikasiIjazah *VerifikasiIjazahTransactorSession) DeleteDegree(_degreeHash [32]byte) (*types.Transaction, error) {
	return _VerifikasiIjazah.Contract.DeleteDegree(&_VerifikasiIjazah.TransactOpts, _degreeHash)
}

// IssueDegree is a paid mutator transaction binding the contract method 0x578d30da.
//
// Solidity: function issueDegree(bytes32 _degreeHash, string _sekolah, uint256 _issueDate, string _ipfsUrl, string[] _mataPelajaran, uint8[] _nilai) returns()
func (_VerifikasiIjazah *VerifikasiIjazahTransactor) IssueDegree(opts *bind.TransactOpts, _degreeHash [32]byte, _sekolah string, _issueDate *big.Int, _ipfsUrl string, _mataPelajaran []string, _nilai []uint8) (*types.Transaction, error) {
	return _VerifikasiIjazah.contract.Transact(opts, "issueDegree", _degreeHash, _sekolah, _issueDate, _ipfsUrl, _mataPelajaran, _nilai)
}

// IssueDegree is a paid mutator transaction binding the contract method 0x578d30da.
//
// Solidity: function issueDegree(bytes32 _degreeHash, string _sekolah, uint256 _issueDate, string _ipfsUrl, string[] _mataPelajaran, uint8[] _nilai) returns()
func (_VerifikasiIjazah *VerifikasiIjazahSession) IssueDegree(_degreeHash [32]byte, _sekolah string, _issueDate *big.Int, _ipfsUrl string, _mataPelajaran []string, _nilai []uint8) (*types.Transaction, error) {
	return _VerifikasiIjazah.Contract.IssueDegree(&_VerifikasiIjazah.TransactOpts, _degreeHash, _sekolah, _issueDate, _ipfsUrl, _mataPelajaran, _nilai)
}

// IssueDegree is a paid mutator transaction binding the contract method 0x578d30da.
//
// Solidity: function issueDegree(bytes32 _degreeHash, string _sekolah, uint256 _issueDate, string _ipfsUrl, string[] _mataPelajaran, uint8[] _nilai) returns()
func (_VerifikasiIjazah *VerifikasiIjazahTransactorSession) IssueDegree(_degreeHash [32]byte, _sekolah string, _issueDate *big.Int, _ipfsUrl string, _mataPelajaran []string, _nilai []uint8) (*types.Transaction, error) {
	return _VerifikasiIjazah.Contract.IssueDegree(&_VerifikasiIjazah.TransactOpts, _degreeHash, _sekolah, _issueDate, _ipfsUrl, _mataPelajaran, _nilai)
}

// VerifikasiIjazahDegreeDeletedIterator is returned from FilterDegreeDeleted and is used to iterate over the raw logs and unpacked data for DegreeDeleted events raised by the VerifikasiIjazah contract.
type VerifikasiIjazahDegreeDeletedIterator struct {
	Event *VerifikasiIjazahDegreeDeleted // Event containing the contract specifics and raw log

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
func (it *VerifikasiIjazahDegreeDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifikasiIjazahDegreeDeleted)
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
		it.Event = new(VerifikasiIjazahDegreeDeleted)
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
func (it *VerifikasiIjazahDegreeDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifikasiIjazahDegreeDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifikasiIjazahDegreeDeleted represents a DegreeDeleted event raised by the VerifikasiIjazah contract.
type VerifikasiIjazahDegreeDeleted struct {
	DegreeHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDegreeDeleted is a free log retrieval operation binding the contract event 0x2378393e5e45ca60b5cc050d80fa00b05dc03948e54f0260efb5af5229b9ca3e.
//
// Solidity: event DegreeDeleted(bytes32 indexed degreeHash)
func (_VerifikasiIjazah *VerifikasiIjazahFilterer) FilterDegreeDeleted(opts *bind.FilterOpts, degreeHash [][32]byte) (*VerifikasiIjazahDegreeDeletedIterator, error) {

	var degreeHashRule []interface{}
	for _, degreeHashItem := range degreeHash {
		degreeHashRule = append(degreeHashRule, degreeHashItem)
	}

	logs, sub, err := _VerifikasiIjazah.contract.FilterLogs(opts, "DegreeDeleted", degreeHashRule)
	if err != nil {
		return nil, err
	}
	return &VerifikasiIjazahDegreeDeletedIterator{contract: _VerifikasiIjazah.contract, event: "DegreeDeleted", logs: logs, sub: sub}, nil
}

// WatchDegreeDeleted is a free log subscription operation binding the contract event 0x2378393e5e45ca60b5cc050d80fa00b05dc03948e54f0260efb5af5229b9ca3e.
//
// Solidity: event DegreeDeleted(bytes32 indexed degreeHash)
func (_VerifikasiIjazah *VerifikasiIjazahFilterer) WatchDegreeDeleted(opts *bind.WatchOpts, sink chan<- *VerifikasiIjazahDegreeDeleted, degreeHash [][32]byte) (event.Subscription, error) {

	var degreeHashRule []interface{}
	for _, degreeHashItem := range degreeHash {
		degreeHashRule = append(degreeHashRule, degreeHashItem)
	}

	logs, sub, err := _VerifikasiIjazah.contract.WatchLogs(opts, "DegreeDeleted", degreeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifikasiIjazahDegreeDeleted)
				if err := _VerifikasiIjazah.contract.UnpackLog(event, "DegreeDeleted", log); err != nil {
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

// ParseDegreeDeleted is a log parse operation binding the contract event 0x2378393e5e45ca60b5cc050d80fa00b05dc03948e54f0260efb5af5229b9ca3e.
//
// Solidity: event DegreeDeleted(bytes32 indexed degreeHash)
func (_VerifikasiIjazah *VerifikasiIjazahFilterer) ParseDegreeDeleted(log types.Log) (*VerifikasiIjazahDegreeDeleted, error) {
	event := new(VerifikasiIjazahDegreeDeleted)
	if err := _VerifikasiIjazah.contract.UnpackLog(event, "DegreeDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifikasiIjazahDegreeIssuedIterator is returned from FilterDegreeIssued and is used to iterate over the raw logs and unpacked data for DegreeIssued events raised by the VerifikasiIjazah contract.
type VerifikasiIjazahDegreeIssuedIterator struct {
	Event *VerifikasiIjazahDegreeIssued // Event containing the contract specifics and raw log

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
func (it *VerifikasiIjazahDegreeIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifikasiIjazahDegreeIssued)
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
		it.Event = new(VerifikasiIjazahDegreeIssued)
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
func (it *VerifikasiIjazahDegreeIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifikasiIjazahDegreeIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifikasiIjazahDegreeIssued represents a DegreeIssued event raised by the VerifikasiIjazah contract.
type VerifikasiIjazahDegreeIssued struct {
	DegreeHash [32]byte
	Sekolah    string
	IssueDate  *big.Int
	IpfsUrl    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDegreeIssued is a free log retrieval operation binding the contract event 0xb604f6e596d89a3657daccdbb69c509b6039e34c5cff6500a0ff22efec8ac432.
//
// Solidity: event DegreeIssued(bytes32 indexed degreeHash, string sekolah, uint256 issueDate, string ipfsUrl)
func (_VerifikasiIjazah *VerifikasiIjazahFilterer) FilterDegreeIssued(opts *bind.FilterOpts, degreeHash [][32]byte) (*VerifikasiIjazahDegreeIssuedIterator, error) {

	var degreeHashRule []interface{}
	for _, degreeHashItem := range degreeHash {
		degreeHashRule = append(degreeHashRule, degreeHashItem)
	}

	logs, sub, err := _VerifikasiIjazah.contract.FilterLogs(opts, "DegreeIssued", degreeHashRule)
	if err != nil {
		return nil, err
	}
	return &VerifikasiIjazahDegreeIssuedIterator{contract: _VerifikasiIjazah.contract, event: "DegreeIssued", logs: logs, sub: sub}, nil
}

// WatchDegreeIssued is a free log subscription operation binding the contract event 0xb604f6e596d89a3657daccdbb69c509b6039e34c5cff6500a0ff22efec8ac432.
//
// Solidity: event DegreeIssued(bytes32 indexed degreeHash, string sekolah, uint256 issueDate, string ipfsUrl)
func (_VerifikasiIjazah *VerifikasiIjazahFilterer) WatchDegreeIssued(opts *bind.WatchOpts, sink chan<- *VerifikasiIjazahDegreeIssued, degreeHash [][32]byte) (event.Subscription, error) {

	var degreeHashRule []interface{}
	for _, degreeHashItem := range degreeHash {
		degreeHashRule = append(degreeHashRule, degreeHashItem)
	}

	logs, sub, err := _VerifikasiIjazah.contract.WatchLogs(opts, "DegreeIssued", degreeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifikasiIjazahDegreeIssued)
				if err := _VerifikasiIjazah.contract.UnpackLog(event, "DegreeIssued", log); err != nil {
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

// ParseDegreeIssued is a log parse operation binding the contract event 0xb604f6e596d89a3657daccdbb69c509b6039e34c5cff6500a0ff22efec8ac432.
//
// Solidity: event DegreeIssued(bytes32 indexed degreeHash, string sekolah, uint256 issueDate, string ipfsUrl)
func (_VerifikasiIjazah *VerifikasiIjazahFilterer) ParseDegreeIssued(log types.Log) (*VerifikasiIjazahDegreeIssued, error) {
	event := new(VerifikasiIjazahDegreeIssued)
	if err := _VerifikasiIjazah.contract.UnpackLog(event, "DegreeIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
