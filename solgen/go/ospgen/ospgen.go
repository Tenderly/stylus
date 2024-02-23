// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ospgen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/tenderly/stylus/go-ethereum"
	"github.com/tenderly/stylus/go-ethereum/accounts/abi"
	"github.com/tenderly/stylus/go-ethereum/accounts/abi/bind"
	"github.com/tenderly/stylus/go-ethereum/common"
	"github.com/tenderly/stylus/go-ethereum/core/types"
	"github.com/tenderly/stylus/go-ethereum/event"
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

// ErrorGuard is an auto generated low-level Go binding around an user-defined struct.
type ErrorGuard struct {
	FrameStack [32]byte
	ValueStack [32]byte
	InterStack [32]byte
	OnErrorPc  Value
}

// ExecutionContext is an auto generated low-level Go binding around an user-defined struct.
type ExecutionContext struct {
	MaxInboxMessagesRead *big.Int
	Bridge               common.Address
}

// GuardStack is an auto generated low-level Go binding around an user-defined struct.
type GuardStack struct {
	Proved        []ErrorGuard
	RemainingHash [32]byte
	Enabled       bool
}

// Instruction is an auto generated low-level Go binding around an user-defined struct.
type Instruction struct {
	Opcode       uint16
	ArgumentData *big.Int
}

// Machine is an auto generated low-level Go binding around an user-defined struct.
type Machine struct {
	Status          uint8
	ValueStack      ValueStack
	InternalStack   ValueStack
	FrameStack      StackFrameWindow
	GuardStack      GuardStack
	GlobalStateHash [32]byte
	ModuleIdx       uint32
	FunctionIdx     uint32
	FunctionPc      uint32
	ModulesRoot     [32]byte
}

// Module is an auto generated low-level Go binding around an user-defined struct.
type Module struct {
	GlobalsMerkleRoot   [32]byte
	ModuleMemory        ModuleMemory
	TablesMerkleRoot    [32]byte
	FunctionsMerkleRoot [32]byte
	InternalsOffset     uint32
}

// ModuleMemory is an auto generated low-level Go binding around an user-defined struct.
type ModuleMemory struct {
	Size       uint64
	MaxSize    uint64
	MerkleRoot [32]byte
}

// StackFrame is an auto generated low-level Go binding around an user-defined struct.
type StackFrame struct {
	ReturnPc              Value
	LocalsMerkleRoot      [32]byte
	CallerModule          uint32
	CallerModuleInternals uint32
}

// StackFrameWindow is an auto generated low-level Go binding around an user-defined struct.
type StackFrameWindow struct {
	Proved        []StackFrame
	RemainingHash [32]byte
}

// Value is an auto generated low-level Go binding around an user-defined struct.
type Value struct {
	ValueType uint8
	Contents  *big.Int
}

// ValueArray is an auto generated low-level Go binding around an user-defined struct.
type ValueArray struct {
	Inner []Value
}

// ValueStack is an auto generated low-level Go binding around an user-defined struct.
type ValueStack struct {
	Proved        ValueArray
	RemainingHash [32]byte
}

// HashProofHelperMetaData contains all meta data concerning the HashProofHelper contract.
var HashProofHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"NotProven\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"}],\"name\":\"PreimagePartProven\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clearSplitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"getPreimagePart\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keccakStates\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"proveWithFullPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"proveWithSplitPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611cea806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063740085d71461005c57806379754cba14610085578063ae364ac2146100a6578063b7465799146100b0578063d4e5dd2b146100d2575b600080fd5b61006f61006a3660046118e0565b6100e5565b60405161007c9190611959565b60405180910390f35b6100986100933660046119bb565b6101d8565b60405190815260200161007c565b6100ae6106cb565b005b6100c36100be366004611a16565b610713565b60405161007c93929190611a3f565b6100986100e0366004611a71565b6107c9565b6000828152602081815260408083206001600160401b0385168452909152902080546060919060ff16610142576040516309cb23c960e11b8152600481018590526001600160401b03841660248201526044015b60405180910390fd5b80600101805461015190611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461017d90611ac4565b80156101ca5780601f1061019f576101008083540402835291602001916101ca565b820191906000526020600020905b8154815290600101906020018083116101ad57829003601f168201915b505050505091505092915050565b60006001821615156002831615610230573360009081526001602081905260408220805467ffffffffffffffff191681559190610217908301826117a2565b6102256002830160006117df565b600982016000905550505b80806102445750610242608886611b15565b155b6102845760405162461bcd60e51b81526020600482015260116024820152701393d517d09313d0d2d7d0531251d39151607a1b6044820152606401610139565b3360009081526001602052604090206009810154806102bc57815467ffffffffffffffff19166001600160401b038716178255610306565b81546001600160401b038781169116146103065760405162461bcd60e51b815260206004820152600b60248201526a1112519197d3d19194d15560aa1b6044820152606401610139565b61031282898986610920565b8061032760206001600160401b038916611b3f565b11801561034057508160090154866001600160401b0316105b1561045a57600081876001600160401b0316111561036e5761036b826001600160401b038916611b57565b90505b60008261038560206001600160401b038b16611b3f565b61038f9190611b57565b90508881111561039c5750875b815b8181101561045657846001018b8b838181106103bc576103bc611b6e565b9050013560f81c60f81b90808054806103d490611ac4565b80601f81146103e2576103f8565b83600052602060002060ff1984168155603f9350505b506002919091019091558154600116156104215790600052602060002090602091828204019190065b909190919091601f036101000a81548160ff02191690600160f81b84040217905550808061044e90611b84565b91505061039e565b5050505b8261046c5750600092506106c3915050565b60005b602081101561053c576000610485600883611b9f565b9050610492600582611b9f565b61049d600583611b15565b6104a8906005611bb3565b6104b29190611b3f565b905060006104c1600884611b15565b6104cc906008611bb3565b8560020183601981106104e1576104e1611b6e565b60048104909101546001600160401b036008600390931683026101000a9091041690911c9150610512908490611bb3565b61051d9060f8611b57565b60ff909116901b9590951794508061053481611b84565b91505061046f565b50604051806040016040528060011515815260200183600101805461056090611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461058c90611ac4565b80156105d95780601f106105ae576101008083540402835291602001916105d9565b820191906000526020600020905b8154815290600101906020018083116105bc57829003601f168201915b50505091909252505060008581526020818152604080832086546001600160401b0316845282529091208251815460ff1916901515178155828201518051919261062b926001850192909101906117ee565b505082546040516001600160401b03909116915085907ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c90610671906001870190611bd2565b60405180910390a33360009081526001602081905260408220805467ffffffffffffffff1916815591906106a7908301826117a2565b6106b56002830160006117df565b600982016000905550505050505b949350505050565b3360009081526001602081905260408220805467ffffffffffffffff1916815591906106f9908301826117a2565b6107076002830160006117df565b60098201600090555050565b6001602081905260009182526040909120805491810180546001600160401b039093169261074090611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461076c90611ac4565b80156107b95780601f1061078e576101008083540402835291602001916107b9565b820191906000526020600020905b81548152906001019060200180831161079c57829003601f168201915b5050505050908060090154905083565b600083836040516107db929190611c7a565b604051908190039020905060606001600160401b03831684111561087957600061080e6001600160401b03851686611b57565b9050602081111561081d575060205b856001600160401b038516866108338483611b3f565b9261084093929190611c8a565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929450505050505b6040805180820182526001808252602080830185815260008781528083528581206001600160401b038a1682528352949094208351815460ff1916901515178155935180519394936108d29385019291909101906117ee565b50905050826001600160401b0316827ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c836040516109109190611959565b60405180910390a3509392505050565b828290508460090160008282546109379190611b3f565b90915550505b81158015610949575080155b1561095357610ba2565b60005b6088811015610a7a5760008382101561098c5784848381811061097b5761097b611b6e565b919091013560f81c91506109af9050565b81841415610998576001175b6109a460016088611b57565b8214156109af576080175b60006109bc600884611b9f565b90506109c9600582611b9f565b6109d4600583611b15565b6109df906005611bb3565b6109e99190611b3f565b90506109f6600884611b15565b610a01906008611bb3565b6001600160401b03168260ff166001600160401b0316901b876002018260198110610a2e57610a2e611b6e565b6004810490910180546001600160401b0360086003909416939093026101000a808204841690941883168402929093021990921617905550819050610a7281611b84565b915050610956565b50610a83611872565b60005b6019811015610af557856002018160198110610aa457610aa4611b6e565b600491828204019190066008029054906101000a90046001600160401b03166001600160401b0316828260198110610ade57610ade611b6e565b602002015280610aed81611b84565b915050610a86565b50610aff81610ba8565b905060005b6019811015610b7b57818160198110610b1f57610b1f611b6e565b6020020151866002018260198110610b3957610b39611b6e565b600491828204019190066008026101000a8154816001600160401b0302191690836001600160401b031602179055508080610b7390611b84565b915050610b04565b506088831015610b8b5750610ba2565b610b988360888187611c8a565b935093505061093d565b50505050565b610bb0611872565b610bb8611891565b610bc0611891565b610bc8611872565b600060405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015611797576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e0151181818189288018390526001600160401b0360028202166001603f1b91829004179092188652510485600260200201516002026001600160401b03161785600060200201511884600160200201526001603f1b856003602002015181610e1957610e19611aff565b0485600360200201516002026001600160401b03161785600160200201511884600260200201526001603f1b856004602002015181610e5a57610e5a611aff565b0485600460200201516002026001600160401b03161785600260058110610e8357610e83611b6e565b602002015118606085015284516001603f1b9086516060808901519390920460029091026001600160401b031617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a01805190911890529084525163100000009060208901516001600160401b03641000000000909102169190041761010084015260408701516001603d1b9060408901516001600160401b03600890910216919004176101608401526060870151628000009060608901516001600160401b036502000000000090910216919004176102608401526080870151654000000000009060808901516001600160401b036204000090910216919004176102c084015260a08701516001603f1b900487600560200201516002026001600160401b031617836002601981106110f3576110f3611b6e565b602002015260c08701516210000081046001602c1b9091026001600160401b039081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166001603e1b909104176103008501526101408801516101408901516001600160401b036001603e1b909102169190041760808401526101608701516001603a1b906101608901516001600160401b036040909102169190041760e084015261018087015162200000906101808901516001600160401b036001602b1b90910216919004176101408401526101a08701516602000000000000906101a08901516001600160401b0361800090910216919004176102408401526101c08701516008906101c08901516001600160401b036001603d1b90910216919004176102a08401526101e0870151641000000000906101e08901516001600160401b03631000000090910216919004176020840152610200808801516102008901516001600160401b0366800000000000009091021691900417610120840152610220870151648000000000906102208901516001600160401b03630200000090910216919004176101808401526102408701516001602b1b906102408901516001600160401b036220000090910216919004176101e0840152610260870151610100906102608901516001600160401b03600160381b90910216919004176102e0840152610280870151642000000000906102808901516001600160401b036308000000909102169190041760608401526102a08701516001602c1b906102a08901516001600160401b0362100000909102169190041760c08401526102c08701516302000000906102c08901516001600160401b0364800000000090910216919004176101c08401526102e0870151600160381b906102e08901516001600160401b036101009091021691900417610220840152610300870151660400000000000090048760186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f602002015282600160200201518360156020020151191683601060200201511887601060200201528260026020020151836016602002015119168360116020020151188760116020020152826003602002015183601760200201511916836012602002015118876012602002015282600460200201518360186020020151191683601360200201511887601360200201528260056020020151836000602002015119168360146020020151188760146020020152826006602002015183600160200201511916836015602002015118876015602002015282600760200201518360026020020151191683601660200201511887601660200201528260086020020151836003602002015119168360176020020151188760176020020152826009602002015183600460200201511916836018602002015118876018602002015281816018811061178557611785611b6e565b60200201518751188752600101610cee565b509495945050505050565b5080546117ae90611ac4565b6000825580601f106117be575050565b601f0160209004906000526020600020908101906117dc91906118af565b50565b506117dc9060078101906118af565b8280546117fa90611ac4565b90600052602060002090601f01602090048101928261181c5760008555611862565b82601f1061183557805160ff1916838001178555611862565b82800160010185558215611862579182015b82811115611862578251825591602001919060010190611847565b5061186e9291506118af565b5090565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b5b8082111561186e57600081556001016118b0565b80356001600160401b03811681146118db57600080fd5b919050565b600080604083850312156118f357600080fd5b82359150611903602084016118c4565b90509250929050565b6000815180845260005b8181101561193257602081850181015186830182015201611916565b81811115611944576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061196c602083018461190c565b9392505050565b60008083601f84011261198557600080fd5b5081356001600160401b0381111561199c57600080fd5b6020830191508360208285010111156119b457600080fd5b9250929050565b600080600080606085870312156119d157600080fd5b84356001600160401b038111156119e757600080fd5b6119f387828801611973565b9095509350611a069050602086016118c4565b9396929550929360400135925050565b600060208284031215611a2857600080fd5b81356001600160a01b038116811461196c57600080fd5b6001600160401b0384168152606060208201526000611a61606083018561190c565b9050826040830152949350505050565b600080600060408486031215611a8657600080fd5b83356001600160401b03811115611a9c57600080fd5b611aa886828701611973565b9094509250611abb9050602085016118c4565b90509250925092565b600181811c90821680611ad857607f821691505b60208210811415611af957634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601260045260246000fd5b600082611b2457611b24611aff565b500690565b634e487b7160e01b600052601160045260246000fd5b60008219821115611b5257611b52611b29565b500190565b600082821015611b6957611b69611b29565b500390565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611b9857611b98611b29565b5060010190565b600082611bae57611bae611aff565b500490565b6000816000190483118215151615611bcd57611bcd611b29565b500290565b600060208083526000845481600182811c915080831680611bf457607f831692505b858310811415611c1257634e487b7160e01b85526022600452602485fd5b878601838152602001818015611c2f5760018114611c4057611c6b565b60ff19861682528782019650611c6b565b60008b81526020902060005b86811015611c6557815484820152908501908901611c4c565b83019750505b50949998505050505050505050565b8183823760009101908152919050565b60008085851115611c9a57600080fd5b83861115611ca757600080fd5b505082019391909203915056fea26469706673582212202f2dcb3af29934bcb41f7861970718df9ad65661f9f6cb26c88c963224a1e4c264736f6c63430008090033",
}

// HashProofHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use HashProofHelperMetaData.ABI instead.
var HashProofHelperABI = HashProofHelperMetaData.ABI

// HashProofHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashProofHelperMetaData.Bin instead.
var HashProofHelperBin = HashProofHelperMetaData.Bin

// DeployHashProofHelper deploys a new Ethereum contract, binding an instance of HashProofHelper to it.
func DeployHashProofHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashProofHelper, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashProofHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// HashProofHelper is an auto generated Go binding around an Ethereum contract.
type HashProofHelper struct {
	HashProofHelperCaller     // Read-only binding to the contract
	HashProofHelperTransactor // Write-only binding to the contract
	HashProofHelperFilterer   // Log filterer for contract events
}

// HashProofHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashProofHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashProofHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashProofHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashProofHelperSession struct {
	Contract     *HashProofHelper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashProofHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashProofHelperCallerSession struct {
	Contract *HashProofHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HashProofHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashProofHelperTransactorSession struct {
	Contract     *HashProofHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HashProofHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashProofHelperRaw struct {
	Contract *HashProofHelper // Generic contract binding to access the raw methods on
}

// HashProofHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashProofHelperCallerRaw struct {
	Contract *HashProofHelperCaller // Generic read-only contract binding to access the raw methods on
}

// HashProofHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashProofHelperTransactorRaw struct {
	Contract *HashProofHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashProofHelper creates a new instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelper(address common.Address, backend bind.ContractBackend) (*HashProofHelper, error) {
	contract, err := bindHashProofHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// NewHashProofHelperCaller creates a new read-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperCaller(address common.Address, caller bind.ContractCaller) (*HashProofHelperCaller, error) {
	contract, err := bindHashProofHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperCaller{contract: contract}, nil
}

// NewHashProofHelperTransactor creates a new write-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*HashProofHelperTransactor, error) {
	contract, err := bindHashProofHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperTransactor{contract: contract}, nil
}

// NewHashProofHelperFilterer creates a new log filterer instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*HashProofHelperFilterer, error) {
	contract, err := bindHashProofHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperFilterer{contract: contract}, nil
}

// bindHashProofHelper binds a generic wrapper to an already deployed contract.
func bindHashProofHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.HashProofHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transact(opts, method, params...)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCaller) GetPreimagePart(opts *bind.CallOpts, fullHash [32]byte, offset uint64) ([]byte, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "getPreimagePart", fullHash, offset)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCallerSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCaller) KeccakStates(opts *bind.CallOpts, arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "keccakStates", arg0)

	outstruct := new(struct {
		Offset uint64
		Part   []byte
		Length *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Offset = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Part = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Length = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCallerSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactor) ClearSplitProof(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "clearSplitProof")
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactorSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithFullPreimage(opts *bind.TransactOpts, data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithFullPreimage", data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithSplitPreimage(opts *bind.TransactOpts, data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithSplitPreimage", data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// HashProofHelperPreimagePartProvenIterator is returned from FilterPreimagePartProven and is used to iterate over the raw logs and unpacked data for PreimagePartProven events raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProvenIterator struct {
	Event *HashProofHelperPreimagePartProven // Event containing the contract specifics and raw log

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
func (it *HashProofHelperPreimagePartProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashProofHelperPreimagePartProven)
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
		it.Event = new(HashProofHelperPreimagePartProven)
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
func (it *HashProofHelperPreimagePartProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashProofHelperPreimagePartProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashProofHelperPreimagePartProven represents a PreimagePartProven event raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProven struct {
	FullHash [32]byte
	Offset   uint64
	Part     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreimagePartProven is a free log retrieval operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) FilterPreimagePartProven(opts *bind.FilterOpts, fullHash [][32]byte, offset []uint64) (*HashProofHelperPreimagePartProvenIterator, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.FilterLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperPreimagePartProvenIterator{contract: _HashProofHelper.contract, event: "PreimagePartProven", logs: logs, sub: sub}, nil
}

// WatchPreimagePartProven is a free log subscription operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) WatchPreimagePartProven(opts *bind.WatchOpts, sink chan<- *HashProofHelperPreimagePartProven, fullHash [][32]byte, offset []uint64) (event.Subscription, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.WatchLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashProofHelperPreimagePartProven)
				if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
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

// ParsePreimagePartProven is a log parse operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) ParsePreimagePartProven(log types.Log) (*HashProofHelperPreimagePartProven, error) {
	event := new(HashProofHelperPreimagePartProven)
	if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOneStepProofEntryMetaData contains all meta data concerning the IOneStepProofEntry contract.
var IOneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProofEntryMetaData.ABI instead.
var IOneStepProofEntryABI = IOneStepProofEntryMetaData.ABI

// IOneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type IOneStepProofEntry struct {
	IOneStepProofEntryCaller     // Read-only binding to the contract
	IOneStepProofEntryTransactor // Write-only binding to the contract
	IOneStepProofEntryFilterer   // Log filterer for contract events
}

// IOneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProofEntrySession struct {
	Contract     *IOneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IOneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProofEntryCallerSession struct {
	Contract *IOneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IOneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProofEntryTransactorSession struct {
	Contract     *IOneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IOneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProofEntryRaw struct {
	Contract *IOneStepProofEntry // Generic contract binding to access the raw methods on
}

// IOneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCallerRaw struct {
	Contract *IOneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactorRaw struct {
	Contract *IOneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProofEntry creates a new instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*IOneStepProofEntry, error) {
	contract, err := bindIOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntry{IOneStepProofEntryCaller: IOneStepProofEntryCaller{contract: contract}, IOneStepProofEntryTransactor: IOneStepProofEntryTransactor{contract: contract}, IOneStepProofEntryFilterer: IOneStepProofEntryFilterer{contract: contract}}, nil
}

// NewIOneStepProofEntryCaller creates a new read-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProofEntryCaller, error) {
	contract, err := bindIOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryCaller{contract: contract}, nil
}

// NewIOneStepProofEntryTransactor creates a new write-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProofEntryTransactor, error) {
	contract, err := bindIOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryTransactor{contract: contract}, nil
}

// NewIOneStepProofEntryFilterer creates a new log filterer instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProofEntryFilterer, error) {
	contract, err := bindIOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryFilterer{contract: contract}, nil
}

// bindIOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindIOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// IOneStepProverMetaData contains all meta data concerning the IOneStepProver contract.
var IOneStepProverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"instruction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"result\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"resultMod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProverABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProverMetaData.ABI instead.
var IOneStepProverABI = IOneStepProverMetaData.ABI

// IOneStepProver is an auto generated Go binding around an Ethereum contract.
type IOneStepProver struct {
	IOneStepProverCaller     // Read-only binding to the contract
	IOneStepProverTransactor // Write-only binding to the contract
	IOneStepProverFilterer   // Log filterer for contract events
}

// IOneStepProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProverSession struct {
	Contract     *IOneStepProver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOneStepProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProverCallerSession struct {
	Contract *IOneStepProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IOneStepProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProverTransactorSession struct {
	Contract     *IOneStepProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOneStepProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProverRaw struct {
	Contract *IOneStepProver // Generic contract binding to access the raw methods on
}

// IOneStepProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProverCallerRaw struct {
	Contract *IOneStepProverCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProverTransactorRaw struct {
	Contract *IOneStepProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProver creates a new instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProver(address common.Address, backend bind.ContractBackend) (*IOneStepProver, error) {
	contract, err := bindIOneStepProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProver{IOneStepProverCaller: IOneStepProverCaller{contract: contract}, IOneStepProverTransactor: IOneStepProverTransactor{contract: contract}, IOneStepProverFilterer: IOneStepProverFilterer{contract: contract}}, nil
}

// NewIOneStepProverCaller creates a new read-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProverCaller, error) {
	contract, err := bindIOneStepProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverCaller{contract: contract}, nil
}

// NewIOneStepProverTransactor creates a new write-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProverTransactor, error) {
	contract, err := bindIOneStepProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverTransactor{contract: contract}, nil
}

// NewIOneStepProverFilterer creates a new log filterer instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProverFilterer, error) {
	contract, err := bindIOneStepProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverFilterer{contract: contract}, nil
}

// bindIOneStepProver binds a generic wrapper to an already deployed contract.
func bindIOneStepProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.IOneStepProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	var out []interface{}
	err := _IOneStepProver.contract.Call(opts, &out, "executeOneStep", execCtx, mach, mod, instruction, proof)

	outstruct := new(struct {
		Result    Machine
		ResultMod Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.ResultMod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCallerSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// OneStepProofEntryMetaData contains all meta data concerning the OneStepProofEntry contract.
var OneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"prover0_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMem_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMath_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverHostIo_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prover0\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverHostIo\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMath\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMem\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002daa38038062002daa8339810160408190526200003491620000a5565b600080546001600160a01b039586166001600160a01b031991821617909155600180549486169482169490941790935560028054928516928416929092179091556003805491909316911617905562000102565b80516001600160a01b0381168114620000a057600080fd5b919050565b60008060008060808587031215620000bc57600080fd5b620000c78562000088565b9350620000d76020860162000088565b9250620000e76040860162000088565b9150620000f76060860162000088565b905092959194509250565b612c9880620001126000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631f128bc01461005c57806330a5509f1461008c5780635d3adcfb1461009f5780635f52fd7c146100c057806366e5d9c3146100d3575b600080fd5b60015461006f906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b60005461006f906001600160a01b031681565b6100b26100ad366004611ffd565b6100e6565b604051908152602001610083565b60035461006f906001600160a01b031681565b60025461006f906001600160a01b031681565b60006100f0611ea5565b6100f8611f37565b604080516020810190915260608152604080518082019091526000808252602082015260006101288888836107a3565b909550905088610137866109a3565b1461017f5760405162461bcd60e51b815260206004820152601360248201527209a828690929c8abe848a8c9ea48abe9082a69606b1b60448201526064015b60405180910390fd5b60008551600381111561019457610194612098565b146101ae576101a2856109a3565b9550505050505061079a565b650800000000006101c08b60016120c4565b14156101d357600285526101a2856109a3565b6101de888883610c2c565b90945090506101ee888883610cd8565b809250819450505084610120015161021b8660c0015163ffffffff168686610db29092919063ffffffff16565b146102575760405162461bcd60e51b815260206004820152600c60248201526b1353d115531154d7d493d3d560a21b6044820152606401610176565b6040805160208101909152606081526040805160208101909152606081526102808a8a85610dfb565b90945092506102908a8a85610cd8565b9350915061029f8a8a85610cd8565b809450819250505060006102c988610100015163ffffffff168685610e559092919063ffffffff16565b905060006102ec8960e0015163ffffffff168385610e9b9092919063ffffffff16565b9050876060015181146103365760405162461bcd60e51b815260206004820152601260248201527110905117d1955390d51253d394d7d493d3d560721b6044820152606401610176565b506103499250899150839050818b6120dc565b975097505060008460c0015163ffffffff169050600185610100018181516103719190612106565b63ffffffff1690525081516000602861ffff8316108015906103985750603561ffff831611155b806103b85750603661ffff8316108015906103b85750603e61ffff831611155b806103c7575061ffff8216603f145b806103d6575061ffff82166040145b156103ed57506001546001600160a01b03166105e2565b61ffff821660451480610404575061ffff82166050145b806104325750604661ffff83161080159061043257506104266009604661212e565b61ffff168261ffff1611155b806104605750606761ffff83161080159061046057506104546002606761212e565b61ffff168261ffff1611155b806104805750606a61ffff8316108015906104805750607861ffff831611155b806104ae5750605161ffff8316108015906104ae57506104a26009605161212e565b61ffff168261ffff1611155b806104dc5750607961ffff8316108015906104dc57506104d06002607961212e565b61ffff168261ffff1611155b806104fc5750607c61ffff8316108015906104fc5750608a61ffff831611155b8061050b575061ffff821660a7145b80610528575061ffff821660ac1480610528575061ffff821660ad145b80610548575060c061ffff831610801590610548575060c461ffff831611155b80610568575060bc61ffff831610801590610568575060bf61ffff831611155b1561057f57506002546001600160a01b03166105e2565b61801061ffff83161080159061059b575061801361ffff831611155b806105bd575061802061ffff8316108015906105bd575061802761ffff831611155b156105d457506003546001600160a01b03166105e2565b506000546001600160a01b03165b806001600160a01b031663a173659b8e8989888f8f6040518763ffffffff1660e01b815260040161061896959493929190612319565b60006040518083038186803b15801561063057600080fd5b505afa158015610644573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261066c91908101906129f8565b9097509550600061ffff8316618023148061068c575061ffff8316618024145b15905080156106a7576106a0868589610db2565b6101208901525b6002885160038111156106bc576106bc612098565b1480156106d157506106d18860800151610f10565b156107865760006106e58960800151610f2f565b805160608b01519192506106ff9190602082015260609052565b6020818101518a8201518083019190915260408051928301905260608252526040818101518a8201516020808201929092528251918201909252606081529052606081015161074f908a90610ff3565b604080518082018252600080825260209182018190528251808401909352808352828201528a0151610780916110df565b50600088525b61078f886109a3565b985050505050505050505b95945050505050565b6107ab611ea5565b816000806107ba8787856110ef565b9350905060ff81166107cf576000915061084c565b8060ff16600114156107e4576001915061084c565b8060ff16600214156107f9576002915061084c565b8060ff166003141561080e576003915061084c565b60405162461bcd60e51b8152602060048201526013602482015272554e4b4e4f574e5f4d4143485f53544154555360681b6044820152606401610176565b50610855611f88565b61085d611f88565b60008060008061087e60408051808201909152606081526000602082015290565b60408051606080820183528152600060208201819052918101829052906108a68f8f8d611125565b9b5098506108b58f8f8d611125565b9b5097506108c48f8f8d611224565b9b5092506108d38f8f8d61134c565b9b5091506108e28f8f8d611467565b9b5096506108f18f8f8d611483565b9b5095506109008f8f8d611483565b9b50945061090f8f8f8d611483565b9b50935061091e8f8f8d611467565b809c5081925050506040518061014001604052808b600381111561094457610944612098565b81526020018a81526020018981526020018481526020018381526020018881526020018763ffffffff1681526020018663ffffffff1681526020018563ffffffff168152602001828152509b5050505050505050505050935093915050565b600080825160038111156109b9576109b9612098565b1415610b095760006109ce83602001516114e7565b6109db84604001516114e7565b6109e8856060015161156c565b60a086015160c087015160e0808901516101008a01516101208b01516040516f26b0b1b434b73290393ab73734b7339d60811b602082015260308101999099526050890197909752607088019590955260908701939093526001600160e01b031991811b821660b087015291821b811660b486015291901b1660b883015260bc82015260dc016040516020818303038152906040529050610a8c8360800151611605565b8015610a9e5750826080015160400151155b15610ab157805160209091012092915050565b60808301516040015160009015610ac95750600160f81b5b8181610ad8866080015161161c565b604051602001610aea93929190612b6e565b6040516020818303038152906040528051906020012092505050919050565b600182516003811115610b1e57610b1e612098565b1415610b6f5760a08201516040517026b0b1b434b732903334b734b9b432b21d60791b602082015260318101919091526051015b604051602081830303815290604052805190602001209050919050565b600282516003811115610b8457610b84612098565b1415610bae576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001610b52565b600382516003811115610bc357610bc3612098565b1415610bed576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001610b52565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606401610176565b919050565b610c34611f37565b604080516060810182526000808252602082018190529181018290528391906000806000610c638a8a88611467565b96509450610c728a8a886116a9565b96509350610c818a8a88611467565b96509250610c908a8a88611467565b96509150610c9f8a8a88611483565b6040805160a08101825297885260208801969096529486019390935250606084015263ffffffff16608083015290969095509350505050565b604080516020810190915260608152816000610cf58686846110ef565b92509050600060ff82166001600160401b03811115610d1657610d166124cb565b604051908082528060200260200182016040528015610d3f578160200160208202803683370190505b50905060005b8260ff168160ff161015610d9657610d5e888886611467565b838360ff1681518110610d7357610d73612bb5565b602002602001018196508281525050508080610d8e90612bcb565b915050610d45565b5060405180602001604052808281525093505050935093915050565b6000610df38484610dc285611724565b6040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b815250611791565b949350505050565b604080518082019091526000808252602082015281600080610e1e87878561189b565b93509150610e2d8787856118f4565b6040805180820190915261ffff90941684526020840191909152919791965090945050505050565b6000610df38484610e6585611949565b6040518060400160405280601881526020017724b739ba393ab1ba34b7b71036b2b935b632903a3932b29d60411b815250611791565b60405168233ab731ba34b7b71d60b91b602082015260298101829052600090819060490160405160208183030381529060405280519060200120905061079a85858360405180604001604052806015815260200174233ab731ba34b7b71036b2b935b632903a3932b29d60591b815250611791565b600081604001518015610f295750610f2782611605565b155b92915050565b610f37611fa6565b815151600114610f7d5760405162461bcd60e51b81526020600482015260116024820152700848288be8eaa82a488a6be988a9c8ea89607b1b6044820152606401610176565b81518051600090610f9057610f90612bb5565b6020026020010151905060006001600160401b03811115610fb357610fb36124cb565b604051908082528060200260200182016040528015610fec57816020015b610fd9611fa6565b815260200190600190039081610fd15790505b5090915290565b60048151600681111561100857611008612098565b1415611015575060029052565b602081015160068251600681111561102f5761102f612098565b1461106e5760405162461bcd60e51b815260206004820152600f60248201526e494e56414c49445f50435f5459504560881b6044820152606401610176565b606081901c156110b25760405162461bcd60e51b815260206004820152600f60248201526e494e56414c49445f50435f4441544160881b6044820152606401610176565b63ffffffff808216610100850152602082901c811660e085015260409190911c1660c09092019190915250565b81516110eb9082611993565b5050565b60008184848281811061110457611104612bb5565b919091013560f81c925081905061111a81612beb565b915050935093915050565b61112d611f88565b81600061113b868684611467565b92509050600061114c8787856118f4565b935090506000816001600160401b0381111561116a5761116a6124cb565b6040519080825280602002602001820160405280156111af57816020015b60408051808201909152600080825260208201528152602001906001900390816111885790505b50905060005b81518110156111fd576111c9898987611a86565b8383815181106111db576111db612bb5565b60200260200101819750829052505080806111f590612beb565b9150506111b5565b50604080516060810182529081019182529081526020810192909252509590945092505050565b604080518082019091526060815260006020820152816000611247868684611467565b92509050606086868481811061125f5761125f612bb5565b909101356001600160f81b0319161590506112e7578261127e81612beb565b604080516001808252818301909252919550909150816020015b6112a0611fda565b8152602001906001900390816112985790505090506112c0878785611b82565b826000815181106112d3576112d3612bb5565b60200260200101819550829052505061132b565b826112f181612beb565b60408051600080825260208201909252919550909150611327565b611314611fda565b81526020019060019003908161130c5790505b5090505b60405180604001604052808281526020018381525093505050935093915050565b60408051606080820183528152600060208201819052918101919091528160008080611379888886611c1b565b94509150611388888886611c1b565b94509050611397888886611467565b94509250606081156113dc5760408051600080825260208201909252906113d4565b6113c1611fa6565b8152602001906001900390816113b95790505b50905061143c565b60408051600180825281830190925290816020015b6113f9611fa6565b8152602001906001900390816113f1579050509050611419898987611c48565b8260008151811061142c5761142c612bb5565b6020026020010181975082905250505b6040518060600160405280828152602001858152602001841515815250955050505050935093915050565b600081816114768686846118f4565b9097909650945050505050565b600081815b60048110156114de5760088363ffffffff16901b92508585838181106114b0576114b0612bb5565b919091013560f81c939093179250816114c881612beb565b92505080806114d690612beb565b915050611488565b50935093915050565b60208101518151515160005b818110156115655783516115109061150b9083611cd6565b611d0e565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c01604051602081830303815290604052805190602001209250808061155d90612beb565b9150506114f3565b5050919050565b602081015160005b8251518110156115ff576115a48360000151828151811061159757611597612bb5565b6020026020010151611d2b565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b602082015260328101919091526052810183905260720160405160208183030381529060405280519060200120915080806115f790612beb565b915050611574565b50919050565b805151600090158015610f29575050602001511590565b602081015160005b8251518110156115ff576116548360000151828151811061164757611647612bb5565b6020026020010151611d9b565b6040516b23bab0b9321039ba30b1b59d60a11b6020820152602c810191909152604c8101839052606c0160405160208183030381529060405280519060200120915080806116a190612beb565b915050611624565b604080516060810182526000808252602082018190529181019190915281600080806116d6888886611df3565b945092506116e5888886611df3565b945091506116f4888886611467565b604080516060810182526001600160401b0396871681529490951660208501529383015250969095509350505050565b600081600001516117388360200151611e51565b6040848101516060860151608087015192516626b7b23ab6329d60c91b6020820152602781019590955260478501939093526067840152608783019190915260e01b6001600160e01b03191660a782015260ab01610b52565b8160005b85515181101561185a57600185166117f6578282876000015183815181106117bf576117bf612bb5565b60200260200101516040516020016117d993929190612c06565b604051602081830303815290604052805190602001209150611841565b828660000151828151811061180d5761180d612bb5565b60200260200101518360405160200161182893929190612c06565b6040516020818303038152906040528051906020012091505b60019490941c938061185281612beb565b915050611795565b508315610df35760405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606401610176565b600081815b60028110156114de5760088361ffff16901b92508585838181106118c6576118c6612bb5565b919091013560f81c939093179250816118de81612beb565b92505080806118ec90612beb565b9150506118a0565b600081815b60208110156114de57600883901b925085858381811061191b5761191b612bb5565b919091013560f81c9390931792508161193381612beb565b925050808061194190612beb565b9150506118f9565b600081600001518260200151604051602001610b529291906b24b739ba393ab1ba34b7b71d60a11b815260f09290921b6001600160f01b031916600c830152600e820152602e0190565b8151516000906119a49060016120c4565b6001600160401b038111156119bb576119bb6124cb565b604051908082528060200260200182016040528015611a0057816020015b60408051808201909152600080825260208201528152602001906001900390816119d95790505b50905060005b835151811015611a5c578351805182908110611a2457611a24612bb5565b6020026020010151828281518110611a3e57611a3e612bb5565b60200260200101819052508080611a5490612beb565b915050611a06565b50818184600001515181518110611a7557611a75612bb5565b602090810291909101015290915250565b6040805180820190915260008082526020820152816000858583818110611aaf57611aaf612bb5565b919091013560f81c9150829050611ac581612beb565b925050611ad0600690565b6006811115611ae157611ae1612098565b60ff168160ff161115611b275760405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b6044820152606401610176565b6000611b348787856118f4565b809450819250505060405180604001604052808360ff166006811115611b5c57611b5c612098565b6006811115611b6d57611b6d612098565b81526020018281525093505050935093915050565b611b8a611fda565b81611ba5604080518082019091526000808252602082015290565b6000806000611bb5898987611a86565b95509350611bc4898987611467565b95509250611bd3898987611483565b95509150611be2898987611483565b60408051608081018252968752602087019590955263ffffffff9384169486019490945290911660608401525090969095509350505050565b600081848482818110611c3057611c30612bb5565b919091013560f81c1515925081905061111a81612beb565b611c50611fa6565b81611c6b604080518082019091526000808252602082015290565b6000806000611c7b898987611467565b95509250611c8a898987611467565b95509150611c99898987611467565b95509050611ca8898987611a86565b6040805160808101825295865260208601949094529284019190915260608301529097909650945050505050565b60408051808201909152600080825260208201528251805183908110611cfe57611cfe612bb5565b6020026020010151905092915050565b600081600001518260200151604051602001610b52929190612c2d565b6000611d3a8260000151611d0e565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401610b52565b6000816000015182602001518360400151611db98560600151611d0e565b6040516b22b93937b91033bab0b9321d60a11b6020820152602c810194909452604c840192909252606c830152608c82015260ac01610b52565b600081815b60088110156114de576008836001600160401b0316901b9250858583818110611e2357611e23612bb5565b919091013560f81c93909317925081611e3b81612beb565b9250508080611e4990612beb565b915050611df8565b805160208083015160408085015190516626b2b6b7b93c9d60c91b938101939093526001600160c01b031960c094851b811660278501529190931b16602f8201526037810191909152600090605701610b52565b6040805161014081019091528060008152602001611ec1611f88565b8152602001611ece611f88565b8152602001611eee60408051808201909152606081526000602082015290565b8152604080516060808201835281526000602082810182905292820152910190815260006020820181905260408201819052606082018190526080820181905260a09091015290565b6040805160a081019091526000815260208101611f6d604080516060810182526000808252602082018190529181019190915290565b81526000602082018190526040820181905260609091015290565b60408051606080820183529181019182529081526000602082015290565b6040805160808101825260008082526020808301829052828401829052835180850190945281845283015290606082015290565b6040805160c0810190915260006080820181815260a08301919091528190611f6d565b600080600080600085870360a081121561201657600080fd5b604081121561202457600080fd5b50859450604086013593506060860135925060808601356001600160401b038082111561205057600080fd5b818801915088601f83011261206457600080fd5b81358181111561207357600080fd5b89602082850101111561208557600080fd5b9699959850939650602001949392505050565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156120d7576120d76120ae565b500190565b600080858511156120ec57600080fd5b838611156120f957600080fd5b5050820193919092039150565b600063ffffffff808316818516808303821115612125576121256120ae565b01949350505050565b600061ffff808316818516808303821115612125576121256120ae565b6004811061215b5761215b612098565b9052565b80516007811061217157612171612098565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b808410156121ce576121ba82865161215f565b9382019360019390930192908501906121a7565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b8281101561224d57845161221985825161215f565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101612204565b509687015197909601969096525093949350505050565b805160608084528151848201819052600092602091908201906080870190855b818110156122c7578351805184528581015186850152604080820151908501528601516122b38785018261215f565b509284019260a09290920191600101612284565b50508286015183880152604086015193506122e6604088018515159052565b9695505050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8635815260006101a060208901356001600160a01b03811680821461233d57600080fd5b8060208601525050806040840152612358818401895161214b565b506020870151610140806101c08501526123766102e085018361217e565b9150604089015161019f1980868503016101e0870152612396848361217e565b935060608b0151915080868503016102008701526123b484836121e2565b935060808b015191508086850301610220870152506123d38382612264565b92505060a089015161024085015260c08901516123f961026086018263ffffffff169052565b5060e089015163ffffffff81166102808601525061010089015163ffffffff81166102a0860152506101208901516102c085015261248f60608501898051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b6124a981850188805161ffff168252602090810151910152565b508281036101808401526124be8185876122f0565b9998505050505050505050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715612503576125036124cb565b60405290565b604051602081016001600160401b0381118282101715612503576125036124cb565b604051608081016001600160401b0381118282101715612503576125036124cb565b604051606081016001600160401b0381118282101715612503576125036124cb565b60405160a081016001600160401b0381118282101715612503576125036124cb565b60405161014081016001600160401b0381118282101715612503576125036124cb565b604051601f8201601f191681016001600160401b03811182821017156125dc576125dc6124cb565b604052919050565b805160048110610c2757600080fd5b60006001600160401b0382111561260c5761260c6124cb565b5060051b60200190565b60006040828403121561262857600080fd5b6126306124e1565b905081516007811061264157600080fd5b808252506020820151602082015292915050565b6000604080838503121561266857600080fd5b6126706124e1565b915082516001600160401b038082111561268957600080fd5b8185019150602080838803121561269f57600080fd5b6126a7612509565b8351838111156126b657600080fd5b80850194505087601f8501126126cb57600080fd5b835192506126e06126db846125f3565b6125b4565b83815260069390931b840182019282810190898511156126ff57600080fd5b948301945b84861015612725576127168a87612616565b82529486019490830190612704565b8252508552948501519484019490945250909392505050565b805163ffffffff81168114610c2757600080fd5b6000604080838503121561276557600080fd5b61276d6124e1565b915082516001600160401b0381111561278557600080fd5b8301601f8101851361279657600080fd5b805160206127a66126db836125f3565b82815260a092830284018201928282019190898511156127c557600080fd5b948301945b8486101561282e5780868b0312156127e25760008081fd5b6127ea61252b565b6127f48b88612616565b81528787015185820152606061280b81890161273e565b8983015261281b6080890161273e565b90820152835294850194918301916127ca565b50808752505080860151818601525050505092915050565b80518015158114610c2757600080fd5b6000606080838503121561286957600080fd5b61287161254d565b915082516001600160401b0381111561288957600080fd5b8301601f8101851361289a57600080fd5b805160206128aa6126db836125f3565b82815260a092830284018201928282019190898511156128c957600080fd5b948301945b848610156129235780868b0312156128e65760008081fd5b6128ee61252b565b865181528487015185820152604080880151908201526129108b898901612616565b81890152835294850194918301916128ce565b5086525085810151908501525061293f91505060408301612846565b604082015292915050565b80516001600160401b0381168114610c2757600080fd5b600081830360e081121561297457600080fd5b61297c61256f565b8351815291506060601f198201121561299457600080fd5b5061299d61254d565b6129a96020840161294a565b81526129b76040840161294a565b602082015260608301516040820152806020830152506080820151604082015260a082015160608201526129ed60c0830161273e565b608082015292915050565b600080610100808486031215612a0d57600080fd5b83516001600160401b0380821115612a2457600080fd5b908501906101408288031215612a3957600080fd5b612a41612591565b612a4a836125e4565b8152602083015182811115612a5e57600080fd5b612a6a89828601612655565b602083015250604083015182811115612a8257600080fd5b612a8e89828601612655565b604083015250606083015182811115612aa657600080fd5b612ab289828601612752565b606083015250608083015182811115612aca57600080fd5b612ad689828601612856565b60808301525060a083015160a0820152612af260c0840161273e565b60c0820152612b0360e0840161273e565b60e0820152612b1384840161273e565b938101939093525061012090810151908201529150612b358460208501612961565b90509250929050565b60005b83811015612b59578181015183820152602001612b41565b83811115612b68576000848401525b50505050565b60008451612b80818460208901612b3e565b6b2bb4ba341033bab0b932399d60a11b9201918252506001600160f81b031992909216600c830152600d820152602d01919050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff811415612be257612be26120ae565b60010192915050565b6000600019821415612bff57612bff6120ae565b5060010190565b60008451612c18818460208901612b3e565b91909101928352506020820152604001919050565b652b30b63ab29d60d11b8152600060078410612c4b57612c4b612098565b5060f89290921b600683015260078201526027019056fea26469706673582212202fcf5a237807ba1b3d2a30f3b2fb58a4c33089207700753a2ed7d79cd3f1e2f464736f6c63430008090033",
}

// OneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryMetaData.ABI instead.
var OneStepProofEntryABI = OneStepProofEntryMetaData.ABI

// OneStepProofEntryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryMetaData.Bin instead.
var OneStepProofEntryBin = OneStepProofEntryMetaData.Bin

// DeployOneStepProofEntry deploys a new Ethereum contract, binding an instance of OneStepProofEntry to it.
func DeployOneStepProofEntry(auth *bind.TransactOpts, backend bind.ContractBackend, prover0_ common.Address, proverMem_ common.Address, proverMath_ common.Address, proverHostIo_ common.Address) (common.Address, *types.Transaction, *OneStepProofEntry, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryBin), backend, prover0_, proverMem_, proverMath_, proverHostIo_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// OneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntry struct {
	OneStepProofEntryCaller     // Read-only binding to the contract
	OneStepProofEntryTransactor // Write-only binding to the contract
	OneStepProofEntryFilterer   // Log filterer for contract events
}

// OneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntrySession struct {
	Contract     *OneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryCallerSession struct {
	Contract *OneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryTransactorSession struct {
	Contract     *OneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryRaw struct {
	Contract *OneStepProofEntry // Generic contract binding to access the raw methods on
}

// OneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryCallerRaw struct {
	Contract *OneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactorRaw struct {
	Contract *OneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntry creates a new instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*OneStepProofEntry, error) {
	contract, err := bindOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryCaller creates a new read-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryCaller, error) {
	contract, err := bindOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryCaller{contract: contract}, nil
}

// NewOneStepProofEntryTransactor creates a new write-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryTransactor, error) {
	contract, err := bindOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryTransactor{contract: contract}, nil
}

// NewOneStepProofEntryFilterer creates a new log filterer instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryFilterer, error) {
	contract, err := bindOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryFilterer{contract: contract}, nil
}

// bindOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.OneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) Prover0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "prover0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverHostIo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverHostIo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMath(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMath")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// OneStepProofEntryLibMetaData contains all meta data concerning the OneStepProofEntryLib contract.
var OneStepProofEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122020b6f549982a5090a8f6f1021556e3fcbbc3008af51a1cdeb979a35e2476492164736f6c63430008090033",
}

// OneStepProofEntryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryLibMetaData.ABI instead.
var OneStepProofEntryLibABI = OneStepProofEntryLibMetaData.ABI

// OneStepProofEntryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryLibMetaData.Bin instead.
var OneStepProofEntryLibBin = OneStepProofEntryLibMetaData.Bin

// DeployOneStepProofEntryLib deploys a new Ethereum contract, binding an instance of OneStepProofEntryLib to it.
func DeployOneStepProofEntryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProofEntryLib, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// OneStepProofEntryLib is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntryLib struct {
	OneStepProofEntryLibCaller     // Read-only binding to the contract
	OneStepProofEntryLibTransactor // Write-only binding to the contract
	OneStepProofEntryLibFilterer   // Log filterer for contract events
}

// OneStepProofEntryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntryLibSession struct {
	Contract     *OneStepProofEntryLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryLibCallerSession struct {
	Contract *OneStepProofEntryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OneStepProofEntryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryLibTransactorSession struct {
	Contract     *OneStepProofEntryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryLibRaw struct {
	Contract *OneStepProofEntryLib // Generic contract binding to access the raw methods on
}

// OneStepProofEntryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCallerRaw struct {
	Contract *OneStepProofEntryLibCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactorRaw struct {
	Contract *OneStepProofEntryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntryLib creates a new instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLib(address common.Address, backend bind.ContractBackend) (*OneStepProofEntryLib, error) {
	contract, err := bindOneStepProofEntryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryLibCaller creates a new read-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryLibCaller, error) {
	contract, err := bindOneStepProofEntryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibCaller{contract: contract}, nil
}

// NewOneStepProofEntryLibTransactor creates a new write-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryLibTransactor, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibTransactor{contract: contract}, nil
}

// NewOneStepProofEntryLibFilterer creates a new log filterer instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryLibFilterer, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibFilterer{contract: contract}, nil
}

// bindOneStepProofEntryLib binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transact(opts, method, params...)
}

// OneStepProver0MetaData contains all meta data concerning the OneStepProver0 contract.
var OneStepProver0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612cf8806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a173659b14610030575b600080fd5b61004361003e36600461206d565b61005a565b60405161005192919061230c565b60405180910390f35b610062611f0e565b61006a611fd7565b610073876128db565b915061008436879003870187612a04565b905060006100956020870187612a9b565b905061202861ffff82166100ac57506102eb6102cd565b61ffff8216600114156100c257506102f66102cd565b61ffff8216600f14156100d857506102fd6102cd565b61ffff8216601014156100ee57506103246102cd565b61ffff8216618009141561010557506103bf6102cd565b61ffff821661800b141561011c575061043c6102cd565b61ffff821661800c141561013357506104cb6102cd565b61ffff821661800a141561014a575061060d6102cd565b61ffff82166011141561016057506106fd6102cd565b61ffff821661800314156101775750610ada6102cd565b61ffff8216618004141561018e5750610b1a6102cd565b61ffff8216602014156101a45750610b786102cd565b61ffff8216602114156101ba5750610bba6102cd565b61ffff8216602314156101d05750610bff6102cd565b61ffff8216602414156101e65750610c276102cd565b61ffff821661800214156101fd5750610c576102cd565b61ffff8216601a14156102135750610cf46102cd565b61ffff8216601b14156102295750610d016102cd565b604161ffff8316108015906102435750604461ffff831611155b156102515750610d706102cd565b61ffff8216618005148061026a575061ffff8216618006145b156102785750610e646102cd565b61ffff8216618008141561028f5750610f376102cd565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b60448201526064015b60405180910390fd5b6102de84848989898663ffffffff16565b5050965096945050505050565b505060029092525050565b5050505050565b600061030c8660600151610f46565b805190915061031c908790610fe6565b505050505050565b61033b610330866110d2565b60208701519061114e565b600061034a866060015161115e565b905061036761035c82604001516111aa565b60208801519061114e565b61037761035c82606001516111aa565b602084013563ffffffff811681146103a15760405162461bcd60e51b81526004016102c490612abf565b63ffffffff1660e08701525050600061010090940193909352505050565b6103cb610330866110d2565b6103db6103308660c001516111aa565b6103eb61033085608001516111aa565b6020808401359081901c604082901c156104175760405162461bcd60e51b81526004016102c490612ae6565b63ffffffff90811660c08801521660e086015250506000610100909301929092525050565b610448610330866110d2565b6000610457866060015161115e565b905061046961035c82604001516111aa565b61047961035c82606001516111aa565b6020808501359081901c604082901c156104a55760405162461bcd60e51b81526004016102c490612ae6565b63ffffffff90811660c08901521660e08701525050600061010090940193909352505050565b60008360200135905060006104eb6104e688602001516111dd565b6111fc565b90506104f5611fd7565b604080516020810190915260608152600061051187878361128d565b9093509050610521878783611339565b6101208c015191935091506105418363ffffffff80881690879061141316565b146105995760405162461bcd60e51b815260206004820152602260248201527f43524f53535f4d4f44554c455f494e5445524e414c5f4d4f44554c45535f524f60448201526113d560f21b60648201526084016102c4565b6105b06105a58b6110d2565b60208c01519061114e565b6105c06105a58b60c001516111aa565b6105d06105a58a608001516111aa565b63ffffffff841660c08b015260808301516105eb9086612b33565b63ffffffff1660e08b0152505060006101009098019790975250505050505050565b610619610330866110d2565b6106296103308660c001516111aa565b61063961033085608001516111aa565b6000610648866060015161115e565b9050806060015163ffffffff16600014156106675750600285526102f6565b602084013563ffffffff811681146106c15760405162461bcd60e51b815260206004820152601d60248201527f4241445f43414c4c45525f494e5445524e414c5f43414c4c5f4441544100000060448201526064016102c4565b604082015163ffffffff1660c088015260608201516106e1908290612b33565b63ffffffff1660e0880152505060006101008601525050505050565b6000806107106104e688602001516111dd565b90506000806000808060006107316040518060200160405280606081525090565b61073c8b8b8761145e565b9550935061074b8b8b876114c5565b909650945061075b8b8b876114e1565b9550925061076a8b8b8761145e565b955091506107798b8b876114c5565b90975094506107898b8b87611339565b6040516d21b0b6361034b73234b932b1ba1d60911b60208201526001600160c01b031960c088901b16602e8201526036810189905290965090915060009060560160408051601f19818403018152919052805160209182012091508d0135811461082e5760405162461bcd60e51b81526020600482015260166024820152754241445f43414c4c5f494e4449524543545f4441544160501b60448201526064016102c4565b610844826001600160401b03871686868c611517565b90508d60400151811461088b5760405162461bcd60e51b815260206004820152600f60248201526e10905117d51050931154d7d493d3d5608a1b60448201526064016102c4565b826001600160401b03168963ffffffff16106108b557505060028d52506102f69650505050505050565b505050505060006108d6604080518082019091526000808252602082015290565b6040805160208101909152606081526108f08a8a866114c5565b945092506108ff8a8a866115b9565b9450915061090e8a8a86611339565b94509050600061092b8263ffffffff808b1690879087906116b516565b90508681146109705760405162461bcd60e51b815260206004820152601160248201527010905117d153115351539514d7d493d3d5607a1b60448201526064016102c4565b8584146109a0578d60025b9081600381111561098e5761098e612151565b815250505050505050505050506102f6565b6004835160068111156109b5576109b5612151565b14156109c3578d600261097b565b6005835160068111156109d8576109d8612151565b1415610a37576020830151985063ffffffff89168914610a325760405162461bcd60e51b81526020600482015260156024820152744241445f46554e435f5245465f434f4e54454e545360581b60448201526064016102c4565b610a6f565b60405162461bcd60e51b815260206004820152600d60248201526c4241445f454c454d5f5459504560981b60448201526064016102c4565b5050505050505050610a8361035c876110d2565b6000610a92876060015161115e565b9050610aaf610aa482604001516111aa565b60208901519061114e565b610abf610aa482606001516111aa565b5063ffffffff1660e086015260006101008601525050505050565b602083013563ffffffff81168114610b045760405162461bcd60e51b81526004016102c490612abf565b63ffffffff166101009095019490945250505050565b6000610b2c6104e687602001516111dd565b905063ffffffff81161561031c57602084013563ffffffff81168114610b645760405162461bcd60e51b81526004016102c490612abf565b63ffffffff16610100870152505050505050565b6000610b87866060015161115e565b90506000610b9f826020015186602001358686611751565b6020880151909150610bb1908261114e565b50505050505050565b6000610bc986602001516111dd565b90506000610bda876060015161115e565b9050610bf1816020015186602001358487876117e9565b602090910152505050505050565b6000610c15856000015185602001358585611751565b602087015190915061031c908261114e565b6000610c3686602001516111dd565b9050610c4d856000015185602001358386866117e9565b9094525050505050565b6000610c6686602001516111dd565b90506000610c7787602001516111dd565b90506000610c8888602001516111dd565b905060006040518060800160405280838152602001886020013560001b8152602001610cb3856111fc565b63ffffffff168152602001610cc7866111fc565b63ffffffff168152509050610ce9818a6060015161188390919063ffffffff16565b505050505050505050565b61031c85602001516111dd565b6000610d136104e687602001516111dd565b90506000610d2487602001516111dd565b90506000610d3588602001516111dd565b905063ffffffff831615610d57576020880151610d52908261114e565b610d66565b6020880151610d66908361114e565b5050505050505050565b6000610d7f6020850185612a9b565b9050600061ffff821660411415610d9857506000610e1b565b61ffff821660421415610dad57506001610e1b565b61ffff821660431415610dc257506002610e1b565b61ffff821660441415610dd757506003610e1b565b60405162461bcd60e51b8152602060048201526019602482015278434f4e53545f505553485f494e56414c49445f4f50434f444560381b60448201526064016102c4565b610bb16040518060400160405280836006811115610e3b57610e3b612151565b815260200187602001356001600160401b0316815250886020015161114e90919063ffffffff16565b6040805180820190915260008082526020820152618005610e886020860186612a9b565b61ffff161415610eb657610e9f86602001516111dd565b6040870151909150610eb1908261114e565b61031c565b618006610ec66020860186612a9b565b61ffff161415610eef57610edd86604001516111dd565b6020870151909150610eb1908261114e565b60405162461bcd60e51b815260206004820152601c60248201527f4d4f56455f494e5445524e414c5f494e56414c49445f4f50434f44450000000060448201526064016102c4565b6000610c15866020015161196a565b610f4e612032565b815151600114610f705760405162461bcd60e51b81526004016102c490612b5b565b81518051600090610f8357610f83612b86565b6020026020010151905060006001600160401b03811115610fa657610fa6612457565b604051908082528060200260200182016040528015610fdf57816020015b610fcc612032565b815260200190600190039081610fc45790505b5090915290565b600481516006811115610ffb57610ffb612151565b1415611008575060029052565b602081015160068251600681111561102257611022612151565b146110615760405162461bcd60e51b815260206004820152600f60248201526e494e56414c49445f50435f5459504560881b60448201526064016102c4565b606081901c156110a55760405162461bcd60e51b815260206004820152600f60248201526e494e56414c49445f50435f4441544160881b60448201526064016102c4565b63ffffffff808216610100850152602082901c811660e085015260409190911c1660c09092019190915250565b60408051808201909152600080825260208201526111488261010001518360e001518460c001516040805180820190915260008082526020820152506040805180820182526006815263ffffffff94909416602093841b67ffffffff00000000161791901b63ffffffff60401b16179082015290565b92915050565b815161115a9082611998565b5050565b611166612032565b8151516001146111885760405162461bcd60e51b81526004016102c490612b5b565b8151805160009061119b5761119b612b86565b60200260200101519050919050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6040805180820190915260008082526020820152815161114890611a61565b6020810151600090818351600681111561121857611218612151565b1461124f5760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b60448201526064016102c4565b64010000000081106111485760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b60448201526064016102c4565b611295611fd7565b6040805160608101825260008082526020820181905291810182905283919060008060006112c48a8a886114c5565b965094506112d38a8a88611b6a565b965093506112e28a8a886114c5565b965092506112f18a8a886114c5565b965091506113008a8a88611be5565b6040805160a08101825297885260208801969096529486019390935250606084015263ffffffff16608083015290969095509350505050565b6040805160208101909152606081528160006113568686846114e1565b92509050600060ff82166001600160401b0381111561137757611377612457565b6040519080825280602002602001820160405280156113a0578160200160208202803683370190505b50905060005b8260ff168160ff1610156113f7576113bf8888866114c5565b838360ff16815181106113d4576113d4612b86565b6020026020010181965082815250505080806113ef90612b9c565b9150506113a6565b5060405180602001604052808281525093505050935093915050565b6000611454848461142385611c40565b6040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b815250611cc6565b90505b9392505050565b600081815b60088110156114bc576008836001600160401b0316901b925085858381811061148e5761148e612b86565b919091013560f81c939093179250816114a681612bbc565b92505080806114b490612bbc565b915050611463565b50935093915050565b600081816114d4868684611dd0565b9097909650945050505050565b6000818484828181106114f6576114f6612b86565b919091013560f81c925081905061150c81612bbc565b915050935093915050565b604051652a30b136329d60d11b60208201526001600160f81b031960f885901b1660268201526001600160c01b031960c084901b166027820152602f81018290526000908190604f016040516020818303038152906040528051906020012090506115ae878783604051806040016040528060128152602001712a30b136329036b2b935b632903a3932b29d60711b815250611cc6565b979650505050505050565b60408051808201909152600080825260208201528160008585838181106115e2576115e2612b86565b919091013560f81c91508290506115f881612bbc565b925050611603600690565b600681111561161457611614612151565b60ff168160ff16111561165a5760405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b60448201526064016102c4565b6000611667878785611dd0565b809450819250505060405180604001604052808360ff16600681111561168f5761168f612151565b60068111156116a0576116a0612151565b81526020018281525093505050935093915050565b600080836116c284611e25565b6040516d2a30b136329032b632b6b2b73a1d60911b6020820152602e810192909252604e820152606e016040516020818303038152906040528051906020012090506117458686836040518060400160405280601a81526020017f5461626c6520656c656d656e74206d65726b6c6520747265653a000000000000815250611cc6565b9150505b949350505050565b60408051808201909152600080825260208201526000611781604080518082019091526000808252602082015290565b60408051602081019091526060815261179b8686856115b9565b935091506117aa868685611339565b9350905060006117bb828985611e42565b90508881146117dc5760405162461bcd60e51b81526004016102c490612bd7565b5090979650505050505050565b6000611805604080518082019091526000808252602082015290565b600061181d6040518060200160405280606081525090565b6118288686846115b9565b9093509150611838868684611339565b925090506000611849828a86611e42565b905089811461186a5760405162461bcd60e51b81526004016102c490612bd7565b611875828a8a611e42565b9a9950505050505050505050565b815151600090611894906001612c02565b6001600160401b038111156118ab576118ab612457565b6040519080825280602002602001820160405280156118e457816020015b6118d1612032565b8152602001906001900390816118c95790505b50905060005b83515181101561194057835180518290811061190857611908612b86565b602002602001015182828151811061192257611922612b86565b6020026020010181905250808061193890612bbc565b9150506118ea565b5081818460000151518151811061195957611959612b86565b602090810291909101015290915250565b604080518082019091526000808252602082015281515151611457611990600183612c1a565b845190611e82565b8151516000906119a9906001612c02565b6001600160401b038111156119c0576119c0612457565b604051908082528060200260200182016040528015611a0557816020015b60408051808201909152600080825260208201528152602001906001900390816119de5790505b50905060005b835151811015611940578351805182908110611a2957611a29612b86565b6020026020010151828281518110611a4357611a43612b86565b60200260200101819052508080611a5990612bbc565b915050611a0b565b604080518082019091526000808252602082015281518051611a8590600190612c1a565b81518110611a9557611a95612b86565b6020026020010151905060006001836000015151611ab39190612c1a565b6001600160401b03811115611aca57611aca612457565b604051908082528060200260200182016040528015611b0f57816020015b6040805180820190915260008082526020820152815260200190600190039081611ae85790505b50905060005b8151811015610fdf578351805182908110611b3257611b32612b86565b6020026020010151828281518110611b4c57611b4c612b86565b60200260200101819052508080611b6290612bbc565b915050611b15565b60408051606081018252600080825260208201819052918101919091528160008080611b9788888661145e565b94509250611ba688888661145e565b94509150611bb58888866114c5565b604080516060810182526001600160401b0396871681529490951660208501529383015250969095509350505050565b600081815b60048110156114bc5760088363ffffffff16901b9250858583818110611c1257611c12612b86565b919091013560f81c93909317925081611c2a81612bbc565b9250508080611c3890612bbc565b915050611bea565b60008160000151611c548360200151611eba565b6040848101516060860151608087015192516626b7b23ab6329d60c91b6020820152602781019590955260478501939093526067840152608783019190915260e01b6001600160e01b03191660a782015260ab015b604051602081830303815290604052805190602001209050919050565b8160005b855151811015611d8f5760018516611d2b57828287600001518381518110611cf457611cf4612b86565b6020026020010151604051602001611d0e93929190612c31565b604051602081830303815290604052805190602001209150611d76565b8286600001518281518110611d4257611d42612b86565b602002602001015183604051602001611d5d93929190612c31565b6040516020818303038152906040528051906020012091505b60019490941c9380611d8781612bbc565b915050611cca565b5083156117495760405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b60448201526064016102c4565b600081815b60208110156114bc57600883901b9250858583818110611df757611df7612b86565b919091013560f81c93909317925081611e0f81612bbc565b9250508080611e1d90612bbc565b915050611dd5565b600081600001518260200151604051602001611ca9929190612c77565b60006114548484611e5285611e25565b604051806040016040528060128152602001712b30b63ab29036b2b935b632903a3932b29d60711b815250611cc6565b60408051808201909152600080825260208201528251805183908110611eaa57611eaa612b86565b6020026020010151905092915050565b805160208083015160408085015190516626b2b6b7b93c9d60c91b938101939093526001600160c01b031960c094851b811660278501529190931b16602f8201526037810191909152600090605701611ca9565b6040805161014081019091528060008152602001611f4360408051606080820183529181019182529081526000602082015290565b8152602001611f6960408051606080820183529181019182529081526000602082015290565b8152602001611f8e604051806040016040528060608152602001600080191681525090565b8152604080516060808201835281526000602082810182905292820152910190815260006020820181905260408201819052606082018190526080820181905260a09091015290565b6040805160a08101909152600081526020810161200d604080516060810182526000808252602082018190529181019190915290565b81526000602082018190526040820181905260609091015290565b612030612cac565b565b6040805160c0810190915260006080820181815260a0830191909152819061200d565b60006040828403121561206757600080fd5b50919050565b6000806000806000808688036101a081121561208857600080fd5b6120928989612055565b965060408801356001600160401b03808211156120ae57600080fd5b818a01915061014080838d0312156120c557600080fd5b82985060e0605f19850112156120da57600080fd5b60608b0197506120ec8c828d01612055565b9650506101808a013592508083111561210457600080fd5b828a0192508a601f84011261211857600080fd5b823591508082111561212957600080fd5b5089602082840101111561213c57600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b6004811061217757612177612151565b9052565b80516007811061218d5761218d612151565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b808410156121ea576121d682865161217b565b9382019360019390930192908501906121c3565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b8281101561226957845161223585825161217b565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101612220565b509687015197909601969096525093949350505050565b805160608084528151848201819052600092602091908201906080870190855b818110156122e3578351805184528581015186850152604080820151908501528601516122cf8785018261217b565b509284019260a092909201916001016122a0565b5050828601518388015260408601519350612302604088018515159052565b9695505050505050565b60006101008083526123218184018651612167565b6020850151610140610120818187015261233f61024087018461219a565b9250604088015160ff1980888603018489015261235c858361219a565b945060608a01519350808886030161016089015261237a85856121fe565b945060808a015193508088860301610180890152505061239a8383612280565b925060a08801516101a087015260c088015191506123c16101c087018363ffffffff169052565b60e088015163ffffffff81166101e088015291509287015163ffffffff811661020087015292870151610220860152509150611457905060208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561248f5761248f612457565b60405290565b604051602081016001600160401b038111828210171561248f5761248f612457565b604051608081016001600160401b038111828210171561248f5761248f612457565b604051606081016001600160401b038111828210171561248f5761248f612457565b60405161014081016001600160401b038111828210171561248f5761248f612457565b60405160a081016001600160401b038111828210171561248f5761248f612457565b604051601f8201601f191681016001600160401b038111828210171561256857612568612457565b604052919050565b80356004811061257f57600080fd5b919050565b60006001600160401b0382111561259d5761259d612457565b5060051b60200190565b6000604082840312156125b957600080fd5b6125c161246d565b90508135600781106125d257600080fd5b808252506020820135602082015292915050565b600060408083850312156125f957600080fd5b61260161246d565b915082356001600160401b038082111561261a57600080fd5b8185019150602080838803121561263057600080fd5b612638612495565b83358381111561264757600080fd5b80850194505087601f85011261265c57600080fd5b8335925061267161266c84612584565b612540565b83815260069390931b8401820192828101908985111561269057600080fd5b948301945b848610156126b6576126a78a876125a7565b82529486019490830190612695565b8252508552948501359484019490945250909392505050565b803563ffffffff8116811461257f57600080fd5b600060408083850312156126f657600080fd5b6126fe61246d565b915082356001600160401b0381111561271657600080fd5b8301601f8101851361272757600080fd5b8035602061273761266c83612584565b82815260a0928302840182019282820191908985111561275657600080fd5b948301945b848610156127bf5780868b0312156127735760008081fd5b61277b6124b7565b6127858b886125a7565b81528787013585820152606061279c8189016126cf565b898301526127ac608089016126cf565b908201528352948501949183019161275b565b50808752505080860135818601525050505092915050565b8035801515811461257f57600080fd5b600060608083850312156127fa57600080fd5b6128026124d9565b915082356001600160401b0381111561281a57600080fd5b8301601f8101851361282b57600080fd5b8035602061283b61266c83612584565b82815260a0928302840182019282820191908985111561285a57600080fd5b948301945b848610156128b45780868b0312156128775760008081fd5b61287f6124b7565b863581528487013585820152604080880135908201526128a18b8989016125a7565b818901528352948501949183019161285f565b508652508581013590850152506128d0915050604083016127d7565b604082015292915050565b600061014082360312156128ee57600080fd5b6128f66124fb565b6128ff83612570565b815260208301356001600160401b038082111561291b57600080fd5b612927368387016125e6565b6020840152604085013591508082111561294057600080fd5b61294c368387016125e6565b6040840152606085013591508082111561296557600080fd5b612971368387016126e3565b6060840152608085013591508082111561298a57600080fd5b50612997368286016127e7565b60808301525060a083013560a08201526129b360c084016126cf565b60c08201526129c460e084016126cf565b60e08201526101006129d78185016126cf565b9082015261012092830135928101929092525090565b80356001600160401b038116811461257f57600080fd5b600081830360e0811215612a1757600080fd5b612a1f61251e565b833581526060601f1983011215612a3557600080fd5b612a3d6124d9565b9150612a4b602085016129ed565b8252612a59604085016129ed565b6020830152606084013560408301528160208201526080840135604082015260a08401356060820152612a8e60c085016126cf565b6080820152949350505050565b600060208284031215612aad57600080fd5b813561ffff8116811461145757600080fd5b6020808252600d908201526c4241445f43414c4c5f4441544160981b604082015260600190565b6020808252601a908201527f4241445f43524f53535f4d4f44554c455f43414c4c5f44415441000000000000604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff808316818516808303821115612b5257612b52612b1d565b01949350505050565b6020808252601190820152700848288beae929c889eaebe988a9c8ea89607b1b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff811415612bb357612bb3612b1d565b60010192915050565b6000600019821415612bd057612bd0612b1d565b5060010190565b60208082526011908201527015d493d391d7d3515492d31157d493d3d5607a1b604082015260600190565b60008219821115612c1557612c15612b1d565b500190565b600082821015612c2c57612c2c612b1d565b500390565b6000845160005b81811015612c525760208188018101518583015201612c38565b81811115612c61576000828501525b5091909101928352506020820152604001919050565b652b30b63ab29d60d11b8152600060078410612c9557612c95612151565b5060f89290921b6006830152600782015260270190565b634e487b7160e01b600052605160045260246000fdfea26469706673582212205b3d8ec923394c66d272bc5cab03600693df0c13f7aec94f7417ff23c5de9cf564736f6c63430008090033",
}

// OneStepProver0ABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProver0MetaData.ABI instead.
var OneStepProver0ABI = OneStepProver0MetaData.ABI

// OneStepProver0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProver0MetaData.Bin instead.
var OneStepProver0Bin = OneStepProver0MetaData.Bin

// DeployOneStepProver0 deploys a new Ethereum contract, binding an instance of OneStepProver0 to it.
func DeployOneStepProver0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProver0, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProver0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// OneStepProver0 is an auto generated Go binding around an Ethereum contract.
type OneStepProver0 struct {
	OneStepProver0Caller     // Read-only binding to the contract
	OneStepProver0Transactor // Write-only binding to the contract
	OneStepProver0Filterer   // Log filterer for contract events
}

// OneStepProver0Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProver0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProver0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProver0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProver0Session struct {
	Contract     *OneStepProver0   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProver0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProver0CallerSession struct {
	Contract *OneStepProver0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OneStepProver0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProver0TransactorSession struct {
	Contract     *OneStepProver0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OneStepProver0Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProver0Raw struct {
	Contract *OneStepProver0 // Generic contract binding to access the raw methods on
}

// OneStepProver0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProver0CallerRaw struct {
	Contract *OneStepProver0Caller // Generic read-only contract binding to access the raw methods on
}

// OneStepProver0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProver0TransactorRaw struct {
	Contract *OneStepProver0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProver0 creates a new instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0(address common.Address, backend bind.ContractBackend) (*OneStepProver0, error) {
	contract, err := bindOneStepProver0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// NewOneStepProver0Caller creates a new read-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Caller(address common.Address, caller bind.ContractCaller) (*OneStepProver0Caller, error) {
	contract, err := bindOneStepProver0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Caller{contract: contract}, nil
}

// NewOneStepProver0Transactor creates a new write-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Transactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProver0Transactor, error) {
	contract, err := bindOneStepProver0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Transactor{contract: contract}, nil
}

// NewOneStepProver0Filterer creates a new log filterer instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Filterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProver0Filterer, error) {
	contract, err := bindOneStepProver0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Filterer{contract: contract}, nil
}

// bindOneStepProver0 binds a generic wrapper to an already deployed contract.
func bindOneStepProver0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.OneStepProver0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Caller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProver0.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Session) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0CallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverHostIoMetaData contains all meta data concerning the OneStepProverHostIo contract.
var OneStepProverHostIoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50613534806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a173659b14610030575b600080fd5b61004361003e366004612821565b61005a565b604051610051929190612ac0565b60405180910390f35b610062612674565b61006a61273d565b6100738761308f565b9150610084368790038701876131b8565b90506000610095602087018761324f565b905061278e61801061ffff8316108015906100b6575061801361ffff831611155b156100c457506101e06101c1565b61ffff821661802014156100db57506103286101c1565b61ffff821661802114156100f257506105c86101c1565b61ffff8216618022141561010957506108ed6101c1565b61ffff8216618023141561012057506108f96101c1565b61ffff821661802414156101375750610a1f6101c1565b61ffff8216618025141561014e5750610ac76101c1565b61ffff821661802614156101655750610b926101c1565b61ffff8216618027141561017c5750610ba86101c1565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b60448201526064015b60405180910390fd5b6101d38a85858a8a8a8763ffffffff16565b5050965096945050505050565b60006101ef602085018561324f565b90506101f9612798565b6000610206858583610bdb565b60a08a0151919350915061021983610cb6565b146102595760405162461bcd60e51b815260206004820152601060248201526f4241445f474c4f42414c5f535441544560801b60448201526064016101b8565b61ffff83166180101480610272575061ffff8316618011145b156102945761028f8888848961028a8987818d613273565b610d37565b61030c565b61ffff831661801214156102ac5761028f8883610ebc565b61ffff831661801314156102c45761028f8883610f53565b60405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f474c4f42414c53544154455f4f50434f444500000000000060448201526064016101b8565b61031582610cb6565b60a0909801979097525050505050505050565b600061033f61033a8760200151610fc9565b610fee565b63ffffffff169050600061035961033a8860200151610fc9565b63ffffffff1690508560200151600001516001600160401b031681602061038091906132b3565b118061039557506103926020826132e1565b15155b156103bc578660025b908160038111156103b1576103b1612905565b8152505050506105c0565b60006103c96020836132f5565b90506000806103e46040518060200160405280606081525090565b60208a01516103f690858a8a8761107f565b90945090925090506060600089898681811061041457610414613309565b919091013560f81c915085905061042a8161331f565b95505060ff8116610503573660006104448b88818f613273565b9150915085828260405161045992919061333a565b60405180910390201461049d5760405162461bcd60e51b815260206004820152600c60248201526b4241445f505245494d41474560a01b60448201526064016101b8565b60006104aa8b60206132b3565b9050818111156104b75750805b6104c3818c8486613273565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092975061054495505050505050565b60405162461bcd60e51b81526020600482015260166024820152752aa725a727aba72fa82922a4a6a0a3a2afa82927a7a360511b60448201526064016101b8565b60005b825181101561058857610574858285848151811061056757610567613309565b016020015160f81c611119565b9450806105808161331f565b915050610547565b5061059483878661119f565b60208d01516040015281516105b7906105ac9061121e565b60208f015190611251565b50505050505050505b505050505050565b60006105da61033a8760200151610fc9565b63ffffffff16905060006105f461033a8860200151610fc9565b63ffffffff169050600061061361060e8960200151610fc9565b611261565b6001600160401b031690506020860135158015610631575088358110155b15610659578760035b9081600381111561064d5761064d612905565b815250505050506105c0565b602080880151516001600160401b0316906106759084906132b3565b118061068a57506106876020836132e1565b15155b156106975787600261063a565b60006106a46020846132f5565b90506000806106bf6040518060200160405280606081525090565b60208b01516106d190858b8b8761107f565b90945090925090508888848181106106eb576106eb613309565b909101356001600160f81b03191615905061073e5760405162461bcd60e51b81526020600482015260136024820152722aa725a727aba72fa4a72127ac2fa82927a7a360691b60448201526064016101b8565b826107488161331f565b93505061278e6000808c602001351415610766576112f291506107a6565b60018c60200135141561077d576115f391506107a6565b8d60025b9081600381111561079457610794612905565b815250505050505050505050506105c0565b6107c68f888d8d899080926107bd93929190613273565b8663ffffffff16565b9050806107d5578d6002610781565b50508288101561081b5760405162461bcd60e51b81526020600482015260116024820152702120a22fa6a2a9a9a0a3a2afa82927a7a360791b60448201526064016101b8565b6000610827848a61334a565b905060005b60208163ffffffff1610801561085057508161084e63ffffffff83168b6132b3565b105b156108a9576108958463ffffffff83168d8d8261086d8f8c6132b3565b61087791906132b3565b81811061088657610886613309565b919091013560f81c9050611119565b9350806108a181613361565b91505061082c565b6108b483878661119f565b60208e0151604001526108dc6108c98261121e565b8f6020015161125190919063ffffffff16565b505050505050505050505050505050565b50506001909252505050565b60006040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b815250905060008661012001519050600061094561033a8960200151610fc9565b63ffffffff16905061096481886020015161189590919063ffffffff16565b6109705787600261063a565b6000806109906109816020856132f5565b60208b0151908989600061107f565b50915091506000806109a48c848b8b6118ca565b925050915060006109c08360016109bb91906132b3565b611aaa565b905080156109eb576109e0876109d78560016132b3565b8760008c611aca565b6101208e0152610a09565b610a026109f98460016132b3565b8390878b611b74565b6101208e01525b6105b76105ac610a1a8560016132b3565b61121e565b60408051808201909152601381527226b7b23ab6329036b2b935b632903a3932b29d60691b6020820152600080610a58888287876118ca565b50915091506000610a6883611aaa565b90508015610aa75781518051610a809060019061334a565b81518110610a9057610a90613309565b602002602001015189610120018181525050610abb565b610ab48284600087611b74565b6101208a01525b50505050505050505050565b6000610ad68660600151611c7e565b90506000610ae78760200151611d17565b90506000610af88860400151611d17565b90506000610b5e8961010001518a60e001518b60c001516040805180820190915260008082526020820152506040805180820182526006815263ffffffff94909416602093841b67ffffffff00000000161791901b63ffffffff60401b16179082015290565b9050610b7a610b6f85858585611d9c565b60808b015190611dc6565b610abb610b87600161121e565b60208b015190611251565b610b9f8560800151611ead565b50505050505050565b6000610bba61033a8760200151610fc9565b60809096015163ffffffff9096161515604090960195909552505050505050565b610be3612798565b81610bec6127bd565b610bf46127bd565b60005b600260ff82161015610c3f57610c0e888886611f71565b848360ff1660028110610c2357610c23613309565b6020020191909152935080610c3781613385565b915050610bf7565b5060005b600260ff82161015610c9957610c5a888886611f8d565b838360ff1660028110610c6f57610c6f613309565b6001600160401b039093166020939093020191909152935080610c9181613385565b915050610c43565b506040805180820190915291825260208201529590945092505050565b80518051602091820151828401518051908401516040516c23b637b130b61039ba30ba329d60991b95810195909552602d850193909352604d8401919091526001600160c01b031960c091821b8116606d85015291901b166075820152600090607d015b604051602081830303815290604052805190602001209050919050565b6000610d4961033a8860200151610fc9565b63ffffffff1690506000610d6361033a8960200151610fc9565b9050600263ffffffff821610610d7b5787600261039e565b6020870151610d8a9083611895565b610d965787600261039e565b6000610da36020846132f5565b9050600080610dbe6040518060200160405280606081525090565b60208b0151610dd090858a8a8761107f565b9094509092509050618010610de860208b018b61324f565b61ffff161415610e2d57610e1f848b600001518763ffffffff1660028110610e1257610e12613309565b602002015183919061119f565b60208c015160400152610eae565b618011610e3d60208b018b61324f565b61ffff161415610e6c578951829063ffffffff871660028110610e6257610e62613309565b6020020152610eae565b60405162461bcd60e51b81526020600482015260176024820152764241445f474c4f42414c5f53544154455f4f50434f444560481b60448201526064016101b8565b505050505050505050505050565b6000610ece61033a8460200151610fc9565b9050600263ffffffff821610610ee657505060029052565b610f4e610f4383602001518363ffffffff1660028110610f0857610f08613309565b602002015160408051808201909152600080825260208201525060408051808201909152600181526001600160401b03909116602082015290565b602085015190611251565b505050565b6000610f6561060e8460200151610fc9565b90506000610f7961033a8560200151610fc9565b9050600263ffffffff821610610f93575050600290915250565b8183602001518263ffffffff1660028110610fb057610fb0613309565b6001600160401b03909216602092909202015250505050565b60408051808201909152600080825260208201528151610fe890611ff4565b92915050565b6020810151600090818351600681111561100a5761100a612905565b146110415760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b60448201526064016101b8565b6401000000008110610fe85760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b60448201526064016101b8565b6000806110986040518060200160405280606081525090565b8391506110a6868684611f71565b90935091506110b68686846120fd565b9250905060006110c782898661119f565b90508860400151811461110d5760405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b60448201526064016101b8565b50955095509592505050565b6000602083106111635760405162461bcd60e51b81526020600482015260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b60448201526064016101b8565b6000836111726001602061334a565b61117c919061334a565b6111879060086133a5565b60ff848116821b911b198616179150505b9392505050565b6040516b26b2b6b7b93c903632b0b31d60a11b6020820152602c81018290526000908190604c016040516020818303038152906040528051906020012090506112158585836040518060400160405280601381526020017226b2b6b7b93c9036b2b935b632903a3932b29d60691b815250611b74565b95945050505050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b815161125d90826121d7565b5050565b602081015160009060018351600681111561127e5761127e612905565b146112b55760405162461bcd60e51b81526020600482015260076024820152661393d517d24d8d60ca1b60448201526064016101b8565b600160401b8110610fe85760405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b60448201526064016101b8565b6000602882101561133a5760405162461bcd60e51b81526020600482015260126024820152712120a22fa9a2a8a4a72127ac2fa82927a7a360711b60448201526064016101b8565b600061134884846020611f8d565b50809150506000848460405161135f92919061333a565b60405190819003902090506000806001600160401b0388161561141f5761138c60408a0160208b016133c4565b6001600160a01b03166316bf55796113a560018b6133ed565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b1580156113e457600080fd5b505afa1580156113f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061141c9190613415565b91505b6001600160401b038416156114d15761143e60408a0160208b016133c4565b6001600160a01b031663d5719dc26114576001876133ed565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b15801561149657600080fd5b505afa1580156114aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ce9190613415565b90505b60408051602081018490529081018490526060810182905260009060800160405160208183030381529060405280519060200120905089602001602081019061151a91906133c4565b6040516316bf557960e01b81526001600160401b038b1660048201526001600160a01b0391909116906316bf55799060240160206040518083038186803b15801561156457600080fd5b505afa158015611578573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159c9190613415565b81146115e15760405162461bcd60e51b81526020600482015260146024820152734241445f534551494e424f585f4d45535341474560601b60448201526064016101b8565b6001955050505050505b949350505050565b6000607182101561163a5760405162461bcd60e51b81526020600482015260116024820152702120a22fa222a620aca2a22fa82927a7a360791b60448201526064016101b8565b60006001600160401b038516156116ee5761165b60408701602088016133c4565b6001600160a01b031663d5719dc26116746001886133ed565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b1580156116b357600080fd5b505afa1580156116c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116eb9190613415565b90505b60006116fd8460718188613273565b60405161170b92919061333a565b6040518091039020905060008585600081811061172a5761172a613309565b9050013560f81c60f81b90506000611744878760016122a0565b50905060008282611759607160218b8d613273565b8760405160200161176e95949392919061342e565b60408051601f198184030181528282528051602091820120838201899052838301819052825180850384018152606090940190925282519201919091209091506117be60408c0160208d016133c4565b604051636ab8cee160e11b81526001600160401b038c1660048201526001600160a01b03919091169063d5719dc29060240160206040518083038186803b15801561180857600080fd5b505afa15801561181c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118409190613415565b81146118845760405162461bcd60e51b81526020600482015260136024820152724241445f44454c415945445f4d45535341474560681b60448201526064016101b8565b5060019a9950505050505050505050565b81516000906001600160401b03166118ae8360206132b3565b1115801561119857506118c26020836132e1565b159392505050565b60006118e26040518060200160405280606081525090565b60408051602081019091526060815260408051808201909152601381527226b7b23ab6329036b2b935b632903a3932b29d60691b602082015261012088015161192961273d565b600061193689898c6122f5565b9a50915061194589898c6123a1565b9a50905061195489898c6120fd565b9a5063ffffffff80831698509096506000906119769088908a9086906123fc16565b90508381146119bd5760405162461bcd60e51b81526020600482015260136024820152722ba927a723afa927a7aa2fa327a92fa622a0a360691b60448201526064016101b8565b50505060006119d28660016109bb91906132b3565b90508015611a2b576119e58660016132b3565b8551516001901b14611a265760405162461bcd60e51b815260206004820152600a6024820152692ba927a723afa622a0a360b11b60448201526064016101b8565b611a9d565b611a3688888b6120fd565b995093506000611a54611a4a8860016132b3565b8690600087611b74565b9050828114611a9b5760405162461bcd60e51b815260206004820152601360248201527257524f4e475f524f4f545f464f525f5a45524f60681b60448201526064016101b8565b505b5050509450945094915050565b60008115801590610fe85750611ac160018361334a565b82161592915050565b600083855b6001811115611b3c57838286604051602001611aed9392919061346d565b604051602081830303815290604052805190602001209150838586604051602001611b1a9392919061346d565b60408051601f198184030181529190528051602090910120945060011c611acf565b838883604051602001611b519392919061346d565b604051602081830303815290604052805190602001209250505095945050505050565b8160005b855151811015611c3d5760018516611bd957828287600001518381518110611ba257611ba2613309565b6020026020010151604051602001611bbc9392919061346d565b604051602081830303815290604052805190602001209150611c24565b8286600001518281518110611bf057611bf0613309565b602002602001015183604051602001611c0b9392919061346d565b6040516020818303038152906040528051906020012091505b60019490941c9380611c358161331f565b915050611b78565b5083156115eb5760405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b60448201526064016101b8565b602081015160005b825151811015611d1157611cb683600001518281518110611ca957611ca9613309565b602002602001015161243d565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b60208201526032810191909152605281018390526072016040516020818303038152906040528051906020012091508080611d099061331f565b915050611c86565b50919050565b60208101518151515160005b81811015611d95578351611d4090611d3b90836124ad565b6124e5565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c016040516020818303038152906040528051906020012092508080611d8d9061331f565b915050611d23565b5050919050565b611da46127db565b5060408051608081018252948552602085019390935291830152606082015290565b815151600090611dd79060016132b3565b6001600160401b03811115611dee57611dee612c0b565b604051908082528060200260200182016040528015611e2757816020015b611e146127db565b815260200190600190039081611e0c5790505b50905060005b835151811015611e83578351805182908110611e4b57611e4b613309565b6020026020010151828281518110611e6557611e65613309565b60200260200101819052508080611e7b9061331f565b915050611e2d565b50818184600001515181518110611e9c57611e9c613309565b602090810291909101015290915250565b611eb56127db565b815151600114611efb5760405162461bcd60e51b81526020600482015260116024820152700848288be8eaa82a488a6be988a9c8ea89607b1b60448201526064016101b8565b81518051600090611f0e57611f0e613309565b6020026020010151905060006001600160401b03811115611f3157611f31612c0b565b604051908082528060200260200182016040528015611f6a57816020015b611f576127db565b815260200190600190039081611f4f5790505b5090915290565b60008181611f808686846122a0565b9097909650945050505050565b600081815b6008811015611feb576008836001600160401b0316901b9250858583818110611fbd57611fbd613309565b919091013560f81c93909317925081611fd58161331f565b9250508080611fe39061331f565b915050611f92565b50935093915050565b6040805180820190915260008082526020820152815180516120189060019061334a565b8151811061202857612028613309565b6020026020010151905060006001836000015151612046919061334a565b6001600160401b0381111561205d5761205d612c0b565b6040519080825280602002602001820160405280156120a257816020015b604080518082019091526000808252602082015281526020019060019003908161207b5790505b50905060005b8151811015611f6a5783518051829081106120c5576120c5613309565b60200260200101518282815181106120df576120df613309565b602002602001018190525080806120f59061331f565b9150506120a8565b60408051602081019091526060815281600061211a868684612502565b92509050600060ff82166001600160401b0381111561213b5761213b612c0b565b604051908082528060200260200182016040528015612164578160200160208202803683370190505b50905060005b8260ff168160ff1610156121bb57612183888886611f71565b838360ff168151811061219857612198613309565b6020026020010181965082815250505080806121b390613385565b91505061216a565b5060405180602001604052808281525093505050935093915050565b8151516000906121e89060016132b3565b6001600160401b038111156121ff576121ff612c0b565b60405190808252806020026020018201604052801561224457816020015b604080518082019091526000808252602082015281526020019060019003908161221d5790505b50905060005b835151811015611e8357835180518290811061226857612268613309565b602002602001015182828151811061228257612282613309565b602002602001018190525080806122989061331f565b91505061224a565b600081815b6020811015611feb57600883901b92508585838181106122c7576122c7613309565b919091013560f81c939093179250816122df8161331f565b92505080806122ed9061331f565b9150506122a5565b6122fd61273d565b60408051606081018252600080825260208201819052918101829052839190600080600061232c8a8a88611f71565b9650945061233b8a8a88612538565b9650935061234a8a8a88611f71565b965092506123598a8a88611f71565b965091506123688a8a886123a1565b6040805160a08101825297885260208801969096529486019390935250606084015263ffffffff16608083015290969095509350505050565b600081815b6004811015611feb5760088363ffffffff16901b92508585838181106123ce576123ce613309565b919091013560f81c939093179250816123e68161331f565b92505080806123f49061331f565b9150506123a6565b60006115eb848461240c856125b3565b6040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b815250611b74565b600061244c82600001516124e5565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401610d1a565b604080518082019091526000808252602082015282518051839081106124d5576124d5613309565b6020026020010151905092915050565b600081600001518260200151604051602001610d1a9291906134b3565b60008184848281811061251757612517613309565b919091013560f81c925081905061252d8161331f565b915050935093915050565b60408051606081018252600080825260208201819052918101919091528160008080612565888886611f8d565b94509250612574888886611f8d565b94509150612583888886611f71565b604080516060810182526001600160401b0396871681529490951660208501529383015250969095509350505050565b600081600001516125c78360200151612620565b6040848101516060860151608087015192516626b7b23ab6329d60c91b6020820152602781019590955260478501939093526067840152608783019190915260e01b6001600160e01b03191660a782015260ab01610d1a565b805160208083015160408085015190516626b2b6b7b93c9d60c91b938101939093526001600160c01b031960c094851b811660278501529190931b16602f8201526037810191909152600090605701610d1a565b60408051610140810190915280600081526020016126a960408051606080820183529181019182529081526000602082015290565b81526020016126cf60408051606080820183529181019182529081526000602082015290565b81526020016126f4604051806040016040528060608152602001600080191681525090565b8152604080516060808201835281526000602082810182905292820152910190815260006020820181905260408201819052606082018190526080820181905260a09091015290565b6040805160a081019091526000815260208101612773604080516060810182526000808252602082018190529181019190915290565b81526000602082018190526040820181905260609091015290565b6127966134e8565b565b60405180604001604052806127ab6127bd565b81526020016127b86127bd565b905290565b60405180604001604052806002906020820280368337509192915050565b6040805160808101825260008082526020808301829052828401829052835180850190945281845283015290606082015290565b600060408284031215611d1157600080fd5b6000806000806000808688036101a081121561283c57600080fd5b612846898961280f565b965060408801356001600160401b038082111561286257600080fd5b818a01915061014080838d03121561287957600080fd5b82985060e0605f198501121561288e57600080fd5b60608b0197506128a08c828d0161280f565b9650506101808a01359250808311156128b857600080fd5b828a0192508a601f8401126128cc57600080fd5b82359150808211156128dd57600080fd5b508960208284010111156128f057600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b6004811061292b5761292b612905565b9052565b80516007811061294157612941612905565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b8084101561299e5761298a82865161292f565b938201936001939093019290850190612977565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015612a1d5784516129e985825161292f565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a0909301926001016129d4565b509687015197909601969096525093949350505050565b805160608084528151848201819052600092602091908201906080870190855b81811015612a9757835180518452858101518685015260408082015190850152860151612a838785018261292f565b509284019260a09290920191600101612a54565b5050828601518388015260408601519350612ab6604088018515159052565b9695505050505050565b6000610100808352612ad5818401865161291b565b60208501516101406101208181870152612af361024087018461294e565b9250604088015160ff19808886030184890152612b10858361294e565b945060608a015193508088860301610160890152612b2e85856129b2565b945060808a0151935080888603016101808901525050612b4e8383612a34565b925060a08801516101a087015260c08801519150612b756101c087018363ffffffff169052565b60e088015163ffffffff81166101e088015291509287015163ffffffff811661020087015292870151610220860152509150611198905060208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715612c4357612c43612c0b565b60405290565b604051602081016001600160401b0381118282101715612c4357612c43612c0b565b604051608081016001600160401b0381118282101715612c4357612c43612c0b565b604051606081016001600160401b0381118282101715612c4357612c43612c0b565b60405161014081016001600160401b0381118282101715612c4357612c43612c0b565b60405160a081016001600160401b0381118282101715612c4357612c43612c0b565b604051601f8201601f191681016001600160401b0381118282101715612d1c57612d1c612c0b565b604052919050565b803560048110612d3357600080fd5b919050565b60006001600160401b03821115612d5157612d51612c0b565b5060051b60200190565b600060408284031215612d6d57600080fd5b612d75612c21565b9050813560078110612d8657600080fd5b808252506020820135602082015292915050565b60006040808385031215612dad57600080fd5b612db5612c21565b915082356001600160401b0380821115612dce57600080fd5b81850191506020808388031215612de457600080fd5b612dec612c49565b833583811115612dfb57600080fd5b80850194505087601f850112612e1057600080fd5b83359250612e25612e2084612d38565b612cf4565b83815260069390931b84018201928281019089851115612e4457600080fd5b948301945b84861015612e6a57612e5b8a87612d5b565b82529486019490830190612e49565b8252508552948501359484019490945250909392505050565b803563ffffffff81168114612d3357600080fd5b60006040808385031215612eaa57600080fd5b612eb2612c21565b915082356001600160401b03811115612eca57600080fd5b8301601f81018513612edb57600080fd5b80356020612eeb612e2083612d38565b82815260a09283028401820192828201919089851115612f0a57600080fd5b948301945b84861015612f735780868b031215612f275760008081fd5b612f2f612c6b565b612f398b88612d5b565b815287870135858201526060612f50818901612e83565b89830152612f6060808901612e83565b9082015283529485019491830191612f0f565b50808752505080860135818601525050505092915050565b80358015158114612d3357600080fd5b60006060808385031215612fae57600080fd5b612fb6612c8d565b915082356001600160401b03811115612fce57600080fd5b8301601f81018513612fdf57600080fd5b80356020612fef612e2083612d38565b82815260a0928302840182019282820191908985111561300e57600080fd5b948301945b848610156130685780868b03121561302b5760008081fd5b613033612c6b565b863581528487013585820152604080880135908201526130558b898901612d5b565b8189015283529485019491830191613013565b5086525085810135908501525061308491505060408301612f8b565b604082015292915050565b600061014082360312156130a257600080fd5b6130aa612caf565b6130b383612d24565b815260208301356001600160401b03808211156130cf57600080fd5b6130db36838701612d9a565b602084015260408501359150808211156130f457600080fd5b61310036838701612d9a565b6040840152606085013591508082111561311957600080fd5b61312536838701612e97565b6060840152608085013591508082111561313e57600080fd5b5061314b36828601612f9b565b60808301525060a083013560a082015261316760c08401612e83565b60c082015261317860e08401612e83565b60e082015261010061318b818501612e83565b9082015261012092830135928101929092525090565b80356001600160401b0381168114612d3357600080fd5b600081830360e08112156131cb57600080fd5b6131d3612cd2565b833581526060601f19830112156131e957600080fd5b6131f1612c8d565b91506131ff602085016131a1565b825261320d604085016131a1565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015261324260c08501612e83565b6080820152949350505050565b60006020828403121561326157600080fd5b813561ffff8116811461119857600080fd5b6000808585111561328357600080fd5b8386111561329057600080fd5b5050820193919092039150565b634e487b7160e01b600052601160045260246000fd5b600082198211156132c6576132c661329d565b500190565b634e487b7160e01b600052601260045260246000fd5b6000826132f0576132f06132cb565b500690565b600082613304576133046132cb565b500490565b634e487b7160e01b600052603260045260246000fd5b60006000198214156133335761333361329d565b5060010190565b8183823760009101908152919050565b60008282101561335c5761335c61329d565b500390565b600063ffffffff8083168181141561337b5761337b61329d565b6001019392505050565b600060ff821660ff81141561339c5761339c61329d565b60010192915050565b60008160001904831182151516156133bf576133bf61329d565b500290565b6000602082840312156133d657600080fd5b81356001600160a01b038116811461119857600080fd5b60006001600160401b038381169083168181101561340d5761340d61329d565b039392505050565b60006020828403121561342757600080fd5b5051919050565b6001600160f81b031986168152606085901b6bffffffffffffffffffffffff191660018201528284601583013760159201918201526035019392505050565b6000845160005b8181101561348e5760208188018101518583015201613474565b8181111561349d576000828501525b5091909101928352506020820152604001919050565b652b30b63ab29d60d11b81526000600784106134d1576134d1612905565b5060f89290921b6006830152600782015260270190565b634e487b7160e01b600052605160045260246000fdfea2646970667358221220fa4712a494805c4fb40c17250d33a1b31ac8a8bd8e52ccc3629f00402a8c754e64736f6c63430008090033",
}

// OneStepProverHostIoABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverHostIoMetaData.ABI instead.
var OneStepProverHostIoABI = OneStepProverHostIoMetaData.ABI

// OneStepProverHostIoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverHostIoMetaData.Bin instead.
var OneStepProverHostIoBin = OneStepProverHostIoMetaData.Bin

// DeployOneStepProverHostIo deploys a new Ethereum contract, binding an instance of OneStepProverHostIo to it.
func DeployOneStepProverHostIo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverHostIo, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverHostIoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// OneStepProverHostIo is an auto generated Go binding around an Ethereum contract.
type OneStepProverHostIo struct {
	OneStepProverHostIoCaller     // Read-only binding to the contract
	OneStepProverHostIoTransactor // Write-only binding to the contract
	OneStepProverHostIoFilterer   // Log filterer for contract events
}

// OneStepProverHostIoCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverHostIoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverHostIoSession struct {
	Contract     *OneStepProverHostIo // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverHostIoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverHostIoCallerSession struct {
	Contract *OneStepProverHostIoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverHostIoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverHostIoTransactorSession struct {
	Contract     *OneStepProverHostIoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverHostIoRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverHostIoRaw struct {
	Contract *OneStepProverHostIo // Generic contract binding to access the raw methods on
}

// OneStepProverHostIoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCallerRaw struct {
	Contract *OneStepProverHostIoCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverHostIoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactorRaw struct {
	Contract *OneStepProverHostIoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverHostIo creates a new instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIo(address common.Address, backend bind.ContractBackend) (*OneStepProverHostIo, error) {
	contract, err := bindOneStepProverHostIo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// NewOneStepProverHostIoCaller creates a new read-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverHostIoCaller, error) {
	contract, err := bindOneStepProverHostIo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoCaller{contract: contract}, nil
}

// NewOneStepProverHostIoTransactor creates a new write-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverHostIoTransactor, error) {
	contract, err := bindOneStepProverHostIo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoTransactor{contract: contract}, nil
}

// NewOneStepProverHostIoFilterer creates a new log filterer instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverHostIoFilterer, error) {
	contract, err := bindOneStepProverHostIo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoFilterer{contract: contract}, nil
}

// bindOneStepProverHostIo binds a generic wrapper to an already deployed contract.
func bindOneStepProverHostIo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverHostIo.contract.Call(opts, &out, "executeOneStep", execCtx, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// OneStepProverMathMetaData contains all meta data concerning the OneStepProverMath contract.
var OneStepProverMathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061255e806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a173659b14610030575b600080fd5b61004361003e3660046118ce565b61005a565b604051610051929190611b6d565b60405180910390f35b6100626117e3565b6040805160a0810182526000808252825160608082018552828252602080830184905282860184905284019190915292820181905291810182905260808101919091526100ae87612137565b91506100bf36879003870187612260565b905060006100d060208701876122f7565b90506118ac61ffff8216604514806100ec575061ffff82166050145b156100fa57506103096102eb565b604661ffff831610801590610122575061011660096046612331565b61ffff168261ffff1611155b1561013057506104226102eb565b606761ffff831610801590610158575061014c60026067612331565b61ffff168261ffff1611155b1561016657506105056102eb565b606a61ffff8316108015906101805750607861ffff831611155b1561018e575061056d6102eb565b605161ffff8316108015906101b657506101aa60096051612331565b61ffff168261ffff1611155b156101c4575061075a6102eb565b607961ffff8316108015906101ec57506101e060026079612331565b61ffff168261ffff1611155b156101fa57506107bf6102eb565b607c61ffff8316108015906102145750608a61ffff831611155b1561022257506108126102eb565b61ffff821660a7141561023857506109cd6102eb565b61ffff821660ac148061024f575061ffff821660ad145b1561025d57506109ee6102eb565b60c061ffff831610801590610277575060c461ffff831611155b156102855750610a426102eb565b60bc61ffff83161080159061029f575060bf61ffff831611155b156102ad5750610c576102eb565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b60448201526064015b60405180910390fd5b6102fc84848989898663ffffffff16565b5050965096945050505050565b60006103188660200151610de2565b9050604561032960208601866122f7565b61ffff16141561036a57600081516006811115610348576103486119b2565b146103655760405162461bcd60e51b81526004016102e290612357565b6103e7565b605061037960208601866122f7565b61ffff1614156103b557600181516006811115610398576103986119b2565b146103655760405162461bcd60e51b81526004016102e290612378565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa2a8ad60c91b60448201526064016102e2565b60008160200151600014156103fe57506001610402565b5060005b61041961040e82610e07565b602089015190610e3a565b50505050505050565b60006104396104348760200151610de2565b610e4a565b9050600061044d6104348860200151610de2565b90506000604661046060208801886122f7565b61046a9190612399565b905060008061ffff831660021480610486575061ffff83166004145b80610495575061ffff83166006145b806104a4575061ffff83166008145b156104c4576104b284610ec1565b91506104bd85610ec1565b90506104d2565b505063ffffffff8083169084165b60006104df838386610eed565b90506104f86104ed82611096565b60208d015190610e3a565b5050505050505050505050565b60006105176104348760200151610de2565b90506000606761052a60208701876122f7565b6105349190612399565b9050600061054a8363ffffffff168360206110c9565b905061056361055882610e07565b60208a015190610e3a565b5050505050505050565b600061057f6104348760200151610de2565b905060006105936104348860200151610de2565b9050600080606a6105a760208901896122f7565b6105b19190612399565b90508061ffff166003141561062f5763ffffffff841615806105e957508260030b637fffffff191480156105e957508360030b600019145b15610612578860025b90816003811115610605576106056119b2565b8152505050505050610753565b8360030b8360030b81610627576106276123bc565b059150610737565b8061ffff166005141561066c5763ffffffff841661064f578860026105f2565b8360030b8360030b81610664576106646123bc565b079150610737565b8061ffff16600a141561068c5763ffffffff8316601f85161b9150610737565b8061ffff16600c14156106ac5763ffffffff8316601f85161c9150610737565b8061ffff16600b14156106ca57600383900b601f85161d9150610737565b8061ffff16600d14156106e8576106e1838561128d565b9150610737565b8061ffff16600e14156106ff576106e183856112cf565b6000806107198563ffffffff168763ffffffff1685611311565b915091508015610733575050600289525061075392505050565b5091505b61074e61074383610e07565b60208b015190610e3a565b505050505b5050505050565b600061077161076c8760200151610de2565b611497565b9050600061078561076c8860200151610de2565b90506000605161079860208801886122f7565b6107a29190612399565b905060006107b1838584610eed565b905061074e61074382611096565b60006107d161076c8760200151610de2565b9050600060796107e460208701876122f7565b6107ee9190612399565b905060006107fe838360406110c9565b63ffffffff1690506105636105588261150e565b600061082461076c8760200151610de2565b9050600061083861076c8860200151610de2565b9050600080607c61084c60208901896122f7565b6108569190612399565b90508061ffff16600314156108bf576001600160401b038416158061089557508260070b677fffffffffffffff1914801561089557508360070b600019145b156108a2578860026105f2565b8360070b8360070b816108b7576108b76123bc565b0591506109c1565b8061ffff16600514156108ff576001600160401b0384166108e2578860026105f2565b8360070b8360070b816108f7576108f76123bc565b0791506109c1565b8061ffff16600a1415610922576001600160401b038316603f85161b91506109c1565b8061ffff16600c1415610945576001600160401b038316603f85161c91506109c1565b8061ffff16600b141561096357600783900b603f85161d91506109c1565b8061ffff16600d14156109815761097a8385611544565b91506109c1565b8061ffff16600e14156109985761097a8385611592565b60006109a5848684611311565b909350905080156109bf5750506002885250610753915050565b505b61074e6107438361150e565b60006109df61076c8760200151610de2565b90508061041961040e82610e07565b6000610a006104348760200151610de2565b9050600060ac610a1360208701876122f7565b61ffff161415610a2d57610a2682610ec1565b9050610a36565b5063ffffffff81165b61041961040e8261150e565b60008060c0610a5460208701876122f7565b61ffff161415610a6a5750600090506008610b41565b60c1610a7960208701876122f7565b61ffff161415610a8f5750600090506010610b41565b60c2610a9e60208701876122f7565b61ffff161415610ab45750600190506008610b41565b60c3610ac360208701876122f7565b61ffff161415610ad95750600190506010610b41565b60c4610ae860208701876122f7565b61ffff161415610afe5750600190506020610b41565b60405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f455854454e445f53414d455f5459504560401b60448201526064016102e2565b600080836006811115610b5657610b566119b2565b1415610b67575063ffffffff610b71565b506001600160401b035b6000610b808960200151610de2565b9050836006811115610b9457610b946119b2565b81516006811115610ba757610ba76119b2565b14610bf05760405162461bcd60e51b81526020600482015260196024820152784241445f455854454e445f53414d455f545950455f5459504560381b60448201526064016102e2565b6000610c03600160ff861681901b6123d2565b602083018051821690529050610c1a6001856123e9565b60ff166001901b826020015116600014610c3c57602082018051821985161790525b60208a0151610c4b9083610e3a565b50505050505050505050565b60008060bc610c6960208701876122f7565b61ffff161415610c7f5750600090506002610d2c565b60bd610c8e60208701876122f7565b61ffff161415610ca45750600190506003610d2c565b60be610cb360208701876122f7565b61ffff161415610cc95750600290506000610d2c565b60bf610cd860208701876122f7565b61ffff161415610cee5750600390506001610d2c565b60405162461bcd60e51b81526020600482015260136024820152721253959053125117d491525395115494149155606a1b60448201526064016102e2565b6000610d3b8860200151610de2565b9050816006811115610d4f57610d4f6119b2565b81516006811115610d6257610d626119b2565b14610daa5760405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f5245494e544552505245545f5459504560401b60448201526064016102e2565b80836006811115610dbd57610dbd6119b2565b90816006811115610dd057610dd06119b2565b90525060208801516105639082610e3a565b60408051808201909152600080825260208201528151610e01906115e0565b92915050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b8151610e4690826116f0565b5050565b60208101516000908183516006811115610e6657610e666119b2565b14610e835760405162461bcd60e51b81526004016102e290612357565b6401000000008110610e015760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b60448201526064016102e2565b60006380000000821615610ee3575063ffffffff1667ffffffff000000001790565b5063ffffffff1690565b600061ffff8216610f1457826001600160401b0316846001600160401b031614905061108f565b61ffff821660011415610f3e57826001600160401b0316846001600160401b03161415905061108f565b61ffff821660021415610f5b578260070b8460070b12905061108f565b61ffff821660031415610f8457826001600160401b0316846001600160401b031610905061108f565b61ffff821660041415610fa1578260070b8460070b13905061108f565b61ffff821660051415610fca57826001600160401b0316846001600160401b031611905061108f565b61ffff821660061415610fe8578260070b8460070b1315905061108f565b61ffff82166007141561101257826001600160401b0316846001600160401b03161115905061108f565b61ffff821660081415611030578260070b8460070b1215905061108f565b61ffff82166009141561105a57826001600160401b0316846001600160401b03161015905061108f565b60405162461bcd60e51b815260206004820152600a6024820152690424144204952454c4f560b41b60448201526064016102e2565b9392505050565b604080518082019091526000808252602082015281156110ba57610e016001610e07565b610e016000610e07565b919050565b60008161ffff16602014806110e257508161ffff166040145b6111295760405162461bcd60e51b8152602060048201526018602482015277057524f4e4720555345204f462067656e65726963556e4f760441b60448201526064016102e2565b61ffff831661119a5761ffff82165b60008163ffffffff1611801561116d575061115460018261240c565b63ffffffff166001901b856001600160401b0316166000145b156111845761117d60018261240c565b9050611138565b6111928161ffff851661240c565b91505061108f565b61ffff8316600114156111f35760005b8261ffff168163ffffffff161080156111d55750600163ffffffff82161b85166001600160401b0316155b156111ec576111e5600182612429565b90506111aa565b905061108f565b61ffff831660021415611259576000805b8361ffff168263ffffffff16101561125057600163ffffffff83161b86166001600160401b03161561123e5761123b600182612429565b90505b8161124881612448565b925050611204565b915061108f9050565b60405162461bcd60e51b815260206004820152600960248201526804241442049556e4f760bc1b60448201526064016102e2565b600061129a60208361246c565b91506112a782602061240c565b63ffffffff168363ffffffff16901c8263ffffffff168463ffffffff16901b17905092915050565b60006112dc60208361246c565b91506112e982602061240c565b63ffffffff168363ffffffff16901b8263ffffffff168463ffffffff16901c17905092915050565b60008061ffff8316611329575050828201600061148f565b8261ffff1660011415611342575050818303600061148f565b8261ffff166002141561135b575050828202600061148f565b8261ffff16600414156113af576001600160401b038416611382575060009050600161148f565b836001600160401b0316856001600160401b0316816113a3576113a36123bc565b0460009150915061148f565b8261ffff1660061415611403576001600160401b0384166113d6575060009050600161148f565b836001600160401b0316856001600160401b0316816113f7576113f76123bc565b0660009150915061148f565b8261ffff166007141561141c575050828216600061148f565b8261ffff1660081415611435575050828217600061148f565b8261ffff166009141561144e575050828218600061148f565b60405162461bcd60e51b81526020600482015260166024820152750494e56414c49445f47454e455249435f42494e5f4f560541b60448201526064016102e2565b935093915050565b60208101516000906001835160068111156114b4576114b46119b2565b146114d15760405162461bcd60e51b81526004016102e290612378565b600160401b8110610e015760405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b60448201526064016102e2565b60408051808201909152600080825260208201525060408051808201909152600181526001600160401b03909116602082015290565b600061155160408361248f565b915061155e8260406124a9565b6001600160401b0316836001600160401b0316901c826001600160401b0316846001600160401b0316901b17905092915050565b600061159f60408361248f565b91506115ac8260406124a9565b6001600160401b0316836001600160401b0316901b826001600160401b0316846001600160401b0316901c17905092915050565b604080518082019091526000808252602082015281518051611604906001906123d2565b81518110611614576116146124c9565b602002602001015190506000600183600001515161163291906123d2565b6001600160401b0381111561164957611649611cb8565b60405190808252806020026020018201604052801561168e57816020015b60408051808201909152600080825260208201528152602001906001900390816116675790505b50905060005b81518110156116e95783518051829081106116b1576116b16124c9565b60200260200101518282815181106116cb576116cb6124c9565b602002602001018190525080806116e1906124df565b915050611694565b5090915290565b8151516000906117019060016124fa565b6001600160401b0381111561171857611718611cb8565b60405190808252806020026020018201604052801561175d57816020015b60408051808201909152600080825260208201528152602001906001900390816117365790505b50905060005b8351518110156117b9578351805182908110611781576117816124c9565b602002602001015182828151811061179b5761179b6124c9565b602002602001018190525080806117b1906124df565b915050611763565b508181846000015151815181106117d2576117d26124c9565b602090810291909101015290915250565b604080516101408101909152806000815260200161181860408051606080820183529181019182529081526000602082015290565b815260200161183e60408051606080820183529181019182529081526000602082015290565b8152602001611863604051806040016040528060608152602001600080191681525090565b8152604080516060808201835281526000602082810182905292820152910190815260006020820181905260408201819052606082018190526080820181905260a09091015290565b6118b4612512565b565b6000604082840312156118c857600080fd5b50919050565b6000806000806000808688036101a08112156118e957600080fd5b6118f389896118b6565b965060408801356001600160401b038082111561190f57600080fd5b818a01915061014080838d03121561192657600080fd5b82985060e0605f198501121561193b57600080fd5b60608b01975061194d8c828d016118b6565b9650506101808a013592508083111561196557600080fd5b828a0192508a601f84011261197957600080fd5b823591508082111561198a57600080fd5b5089602082840101111561199d57600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b600481106119d8576119d86119b2565b9052565b8051600781106119ee576119ee6119b2565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015611a4b57611a378286516119dc565b938201936001939093019290850190611a24565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015611aca578451611a968582516119dc565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101611a81565b509687015197909601969096525093949350505050565b805160608084528151848201819052600092602091908201906080870190855b81811015611b4457835180518452858101518685015260408082015190850152860151611b30878501826119dc565b509284019260a09290920191600101611b01565b5050828601518388015260408601519350611b63604088018515159052565b9695505050505050565b6000610100808352611b8281840186516119c8565b60208501516101406101208181870152611ba06102408701846119fb565b9250604088015160ff19808886030184890152611bbd85836119fb565b945060608a015193508088860301610160890152611bdb8585611a5f565b945060808a0151935080888603016101808901525050611bfb8383611ae1565b925060a08801516101a087015260c08801519150611c226101c087018363ffffffff169052565b60e088015163ffffffff81166101e088015291509287015163ffffffff81166102008701529287015161022086015250915061108f905060208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715611cf057611cf0611cb8565b60405290565b604051602081016001600160401b0381118282101715611cf057611cf0611cb8565b604051608081016001600160401b0381118282101715611cf057611cf0611cb8565b604051606081016001600160401b0381118282101715611cf057611cf0611cb8565b60405161014081016001600160401b0381118282101715611cf057611cf0611cb8565b60405160a081016001600160401b0381118282101715611cf057611cf0611cb8565b604051601f8201601f191681016001600160401b0381118282101715611dc957611dc9611cb8565b604052919050565b8035600481106110c457600080fd5b60006001600160401b03821115611df957611df9611cb8565b5060051b60200190565b600060408284031215611e1557600080fd5b611e1d611cce565b9050813560078110611e2e57600080fd5b808252506020820135602082015292915050565b60006040808385031215611e5557600080fd5b611e5d611cce565b915082356001600160401b0380821115611e7657600080fd5b81850191506020808388031215611e8c57600080fd5b611e94611cf6565b833583811115611ea357600080fd5b80850194505087601f850112611eb857600080fd5b83359250611ecd611ec884611de0565b611da1565b83815260069390931b84018201928281019089851115611eec57600080fd5b948301945b84861015611f1257611f038a87611e03565b82529486019490830190611ef1565b8252508552948501359484019490945250909392505050565b803563ffffffff811681146110c457600080fd5b60006040808385031215611f5257600080fd5b611f5a611cce565b915082356001600160401b03811115611f7257600080fd5b8301601f81018513611f8357600080fd5b80356020611f93611ec883611de0565b82815260a09283028401820192828201919089851115611fb257600080fd5b948301945b8486101561201b5780868b031215611fcf5760008081fd5b611fd7611d18565b611fe18b88611e03565b815287870135858201526060611ff8818901611f2b565b8983015261200860808901611f2b565b9082015283529485019491830191611fb7565b50808752505080860135818601525050505092915050565b803580151581146110c457600080fd5b6000606080838503121561205657600080fd5b61205e611d3a565b915082356001600160401b0381111561207657600080fd5b8301601f8101851361208757600080fd5b80356020612097611ec883611de0565b82815260a092830284018201928282019190898511156120b657600080fd5b948301945b848610156121105780868b0312156120d35760008081fd5b6120db611d18565b863581528487013585820152604080880135908201526120fd8b898901611e03565b81890152835294850194918301916120bb565b5086525085810135908501525061212c91505060408301612033565b604082015292915050565b6000610140823603121561214a57600080fd5b612152611d5c565b61215b83611dd1565b815260208301356001600160401b038082111561217757600080fd5b61218336838701611e42565b6020840152604085013591508082111561219c57600080fd5b6121a836838701611e42565b604084015260608501359150808211156121c157600080fd5b6121cd36838701611f3f565b606084015260808501359150808211156121e657600080fd5b506121f336828601612043565b60808301525060a083013560a082015261220f60c08401611f2b565b60c082015261222060e08401611f2b565b60e0820152610100612233818501611f2b565b9082015261012092830135928101929092525090565b80356001600160401b03811681146110c457600080fd5b600081830360e081121561227357600080fd5b61227b611d7f565b833581526060601f198301121561229157600080fd5b612299611d3a565b91506122a760208501612249565b82526122b560408501612249565b6020830152606084013560408301528160208201526080840135604082015260a084013560608201526122ea60c08501611f2b565b6080820152949350505050565b60006020828403121561230957600080fd5b813561ffff8116811461108f57600080fd5b634e487b7160e01b600052601160045260246000fd5b600061ffff80831681851680830382111561234e5761234e61231b565b01949350505050565b6020808252600790820152662727aa2fa4999960c91b604082015260600190565b6020808252600790820152661393d517d24d8d60ca1b604082015260600190565b600061ffff838116908316818110156123b4576123b461231b565b039392505050565b634e487b7160e01b600052601260045260246000fd5b6000828210156123e4576123e461231b565b500390565b600060ff821660ff8416808210156124035761240361231b565b90039392505050565b600063ffffffff838116908316818110156123b4576123b461231b565b600063ffffffff80831681851680830382111561234e5761234e61231b565b600063ffffffff808316818114156124625761246261231b565b6001019392505050565b600063ffffffff80841680612483576124836123bc565b92169190910692915050565b60006001600160401b0380841680612483576124836123bc565b60006001600160401b03838116908316818110156123b4576123b461231b565b634e487b7160e01b600052603260045260246000fd5b60006000198214156124f3576124f361231b565b5060010190565b6000821982111561250d5761250d61231b565b500190565b634e487b7160e01b600052605160045260246000fdfea26469706673582212205731ce8cc33e286c2ec2949a54264fb52f4f50a33d8e5a721065ed4c792af64764736f6c63430008090033",
}

// OneStepProverMathABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMathMetaData.ABI instead.
var OneStepProverMathABI = OneStepProverMathMetaData.ABI

// OneStepProverMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMathMetaData.Bin instead.
var OneStepProverMathBin = OneStepProverMathMetaData.Bin

// DeployOneStepProverMath deploys a new Ethereum contract, binding an instance of OneStepProverMath to it.
func DeployOneStepProverMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMath, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// OneStepProverMath is an auto generated Go binding around an Ethereum contract.
type OneStepProverMath struct {
	OneStepProverMathCaller     // Read-only binding to the contract
	OneStepProverMathTransactor // Write-only binding to the contract
	OneStepProverMathFilterer   // Log filterer for contract events
}

// OneStepProverMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMathSession struct {
	Contract     *OneStepProverMath // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProverMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMathCallerSession struct {
	Contract *OneStepProverMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProverMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMathTransactorSession struct {
	Contract     *OneStepProverMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProverMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMathRaw struct {
	Contract *OneStepProverMath // Generic contract binding to access the raw methods on
}

// OneStepProverMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMathCallerRaw struct {
	Contract *OneStepProverMathCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactorRaw struct {
	Contract *OneStepProverMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMath creates a new instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMath(address common.Address, backend bind.ContractBackend) (*OneStepProverMath, error) {
	contract, err := bindOneStepProverMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// NewOneStepProverMathCaller creates a new read-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMathCaller, error) {
	contract, err := bindOneStepProverMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathCaller{contract: contract}, nil
}

// NewOneStepProverMathTransactor creates a new write-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMathTransactor, error) {
	contract, err := bindOneStepProverMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathTransactor{contract: contract}, nil
}

// NewOneStepProverMathFilterer creates a new log filterer instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMathFilterer, error) {
	contract, err := bindOneStepProverMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathFilterer{contract: contract}, nil
}

// bindOneStepProverMath binds a generic wrapper to an already deployed contract.
func bindOneStepProverMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.OneStepProverMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMath.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverMemoryMetaData contains all meta data concerning the OneStepProverMemory contract.
var OneStepProverMemoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"frameStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"valueStack\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"interStack\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"onErrorPc\",\"type\":\"tuple\"}],\"internalType\":\"structErrorGuard[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structGuardStack\",\"name\":\"guardStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612079806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a173659b14610030575b600080fd5b61004361003e366004611425565b61005a565b6040516100519291906116c4565b60405180910390f35b61006261133a565b6040805160a0810182526000808252825160608082018552828252602080830184905282860184905284019190915292820181905291810182905260808101919091526100ae87611c93565b91506100bf36879003870187611dbc565b905060006100d06020870187611e53565b9050611403602861ffff8316108015906100ef5750603561ffff831611155b156100fd57506101b4610196565b603661ffff8316108015906101175750603e61ffff831611155b15610125575061061f610196565b61ffff8216603f141561013b57506109c9610196565b61ffff8216604014156101515750610a01610196565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b60448201526064015b60405180910390fd5b6101a784848989898663ffffffff16565b5050965096945050505050565b6000808060286101c76020880188611e53565b61ffff1614156101e05750600091506004905081610435565b60296101ef6020880188611e53565b61ffff161415610209575060019150600890506000610435565b602a6102186020880188611e53565b61ffff161415610232575060029150600490506000610435565b602b6102416020880188611e53565b61ffff16141561025b575060039150600890506000610435565b602c61026a6020880188611e53565b61ffff1614156102835750600091506001905080610435565b602d6102926020880188611e53565b61ffff1614156102ab5750600091506001905081610435565b602e6102ba6020880188611e53565b61ffff1614156102d4575060009150600290506001610435565b602f6102e36020880188611e53565b61ffff1614156102fc5750600091506002905081610435565b603061030b6020880188611e53565b61ffff16141561032357506001915081905080610435565b60316103326020880188611e53565b61ffff16141561034b5750600191508190506000610435565b603261035a6020880188611e53565b61ffff1614156103735750600191506002905081610435565b60336103826020880188611e53565b61ffff16141561039c575060019150600290506000610435565b60346103ab6020880188611e53565b61ffff1614156103c45750600191506004905081610435565b60356103d36020880188611e53565b61ffff1614156103ed575060019150600490506000610435565b60405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f4d454d4f52595f4c4f41445f4f50434f4445000000000000604482015260640161018d565b600061044c6104478a60200151610ab0565b610ad5565b6104609063ffffffff166020890135611e8d565b6020890151909150600090819061047b9084878b8b86610b66565b50915091508115610496575050600289525061061892505050565b8084156105d5578560011480156104be575060008760068111156104bc576104bc611509565b145b156104d4578060000b63ffffffff1690506105d5565b8560011480156104f5575060018760068111156104f3576104f3611509565b145b156105025760000b6105d5565b8560021480156105235750600087600681111561052157610521611509565b145b15610539578060010b63ffffffff1690506105d5565b85600214801561055a5750600187600681111561055857610558611509565b145b156105675760010b6105d5565b8560041480156105885750600187600681111561058657610586611509565b145b156105955760030b6105d5565b60405162461bcd60e51b815260206004820152601560248201527410905117d491505117d096551154d7d4d251d39151605a1b604482015260640161018d565b61061060405180604001604052808960068111156105f5576105f5611509565b81526001600160401b0384166020918201528e015190610c3e565b505050505050505b5050505050565b6000808060366106326020880188611e53565b61ffff16141561064857506004915060006107b7565b60376106576020880188611e53565b61ffff16141561066d57506008915060016107b7565b603861067c6020880188611e53565b61ffff16141561069257506004915060026107b7565b60396106a16020880188611e53565b61ffff1614156106b757506008915060036107b7565b603a6106c66020880188611e53565b61ffff1614156106dc57506001915060006107b7565b603b6106eb6020880188611e53565b61ffff16141561070157506002915060006107b7565b603c6107106020880188611e53565b61ffff161415610725575060019150816107b7565b603d6107346020880188611e53565b61ffff16141561074a57506002915060016107b7565b603e6107596020880188611e53565b61ffff16141561076f57506004915060016107b7565b60405162461bcd60e51b815260206004820152601b60248201527f494e56414c49445f4d454d4f52595f53544f52455f4f50434f44450000000000604482015260640161018d565b60006107c68960200151610ab0565b90508160068111156107da576107da611509565b815160068111156107ed576107ed611509565b1461082b5760405162461bcd60e51b815260206004820152600e60248201526d4241445f53544f52455f5459504560901b604482015260640161018d565b806020015192506008846001600160401b03161015610876576001610851856008611ea5565b6001600160401b031660016001600160401b0316901b6108719190611ed4565b831692505b5050600061088a6104478960200151610ab0565b61089e9063ffffffff166020880135611e8d565b90508660200151600001516001600160401b0316836001600160401b0316826108c79190611e8d565b11156108d95750506002865250610618565b604080516020810190915260608152600090600019906000805b876001600160401b03168110156109a65760006109108288611e8d565b9050600061091f602083611f12565b90508581146109645760001986146109465761093c858786610c4e565b60208f0151604001525b6109578e60200151828e8e8b610ccf565b9098509196509094509250845b6000610971602084611f26565b905061097e85828c610d69565b945060088a6001600160401b0316901c9950505050808061099e90611f3a565b9150506108f3565b506109b2828483610c4e565b60208c015160400152505050505050505050505050565b6020840151516000906109e0906201000090611f55565b90506109f96109ee82610dee565b602088015190610c3e565b505050505050565b602084015151600090610a18906201000090611f55565b90506000610a2c6104478860200151610ab0565b90506000610a4363ffffffff808416908516611e8d565b90508660200151602001516001600160401b03168111610a9857610a6a6201000082611f7b565b60208801516001600160401b039091169052610a93610a8884610dee565b60208a015190610c3e565b610aa6565b610aa6610a88600019610dee565b5050505050505050565b60408051808201909152600080825260208201528151610acf90610e21565b92915050565b60208101516000908183516006811115610af157610af1611509565b14610b285760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b604482015260640161018d565b6401000000008110610acf5760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b604482015260640161018d565b8551600090819081906001600160401b0316610b82888a611e8d565b1115610b975750600191506000905082610c32565b600019600080805b8a811015610c25576000610bb3828e611e8d565b90506000610bc2602083611f12565b9050858114610be257610bd88f828e8e8e610ccf565b509a509095509350845b6000610bef602084611f26565b9050610bfc846008611f7b565b610c068783610f31565b60ff16901b851794505050508080610c1d90611f3a565b915050610b9f565b5060009550935085925050505b96509650969350505050565b8151610c4a9082610fab565b5050565b6040516b26b2b6b7b93c903632b0b31d60a11b6020820152602c81018290526000908190604c01604051602081830303815290604052805190602001209050610cc48585836040518060400160405280601381526020017226b2b6b7b93c9036b2b935b632903a3932b29d60691b81525061109e565b9150505b9392505050565b600080610ce86040518060200160405280606081525090565b839150610cf68686846111b0565b9093509150610d068686846111cc565b925090506000610d17828986610c4e565b905088604001518114610d5d5760405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b604482015260640161018d565b50955095509592505050565b600060208310610db35760405162461bcd60e51b81526020600482015260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b604482015260640161018d565b600083610dc260016020611f9a565b610dcc9190611f9a565b610dd7906008611f7b565b60ff848116821b911b198616179150509392505050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b604080518082019091526000808252602082015281518051610e4590600190611f9a565b81518110610e5557610e55611fb1565b6020026020010151905060006001836000015151610e739190611f9a565b6001600160401b03811115610e8a57610e8a61180f565b604051908082528060200260200182016040528015610ecf57816020015b6040805180820190915260008082526020820152815260200190600190039081610ea85790505b50905060005b8151811015610f2a578351805182908110610ef257610ef2611fb1565b6020026020010151828281518110610f0c57610f0c611fb1565b60200260200101819052508080610f2290611f3a565b915050610ed5565b5090915290565b600060208210610f7c5760405162461bcd60e51b81526020600482015260166024820152750848288bea0aa9898be988a828cbe84b2a88abe9288b60531b604482015260640161018d565b600082610f8b60016020611f9a565b610f959190611f9a565b610fa0906008611f7b565b9390931c9392505050565b815151600090610fbc906001611e8d565b6001600160401b03811115610fd357610fd361180f565b60405190808252806020026020018201604052801561101857816020015b6040805180820190915260008082526020820152815260200190600190039081610ff15790505b50905060005b83515181101561107457835180518290811061103c5761103c611fb1565b602002602001015182828151811061105657611056611fb1565b6020026020010181905250808061106c90611f3a565b91505061101e565b5081818460000151518151811061108d5761108d611fb1565b602090810291909101015290915250565b8160005b8551518110156111675760018516611103578282876000015183815181106110cc576110cc611fb1565b60200260200101516040516020016110e693929190611fc7565b60405160208183030381529060405280519060200120915061114e565b828660000151828151811061111a5761111a611fb1565b60200260200101518360405160200161113593929190611fc7565b6040516020818303038152906040528051906020012091505b60019490941c938061115f81611f3a565b9150506110a2565b5083156111a85760405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b604482015260640161018d565b949350505050565b600081816111bf8686846112a6565b9097909650945050505050565b6040805160208101909152606081528160006111e9868684611304565b92509050600060ff82166001600160401b0381111561120a5761120a61180f565b604051908082528060200260200182016040528015611233578160200160208202803683370190505b50905060005b8260ff168160ff16101561128a576112528888866111b0565b838360ff168151811061126757611267611fb1565b6020026020010181965082815250505080806112829061200d565b915050611239565b5060405180602001604052808281525093505050935093915050565b600081815b60208110156112fb57600883901b92508585838181106112cd576112cd611fb1565b919091013560f81c939093179250816112e581611f3a565b92505080806112f390611f3a565b9150506112ab565b50935093915050565b60008184848281811061131957611319611fb1565b919091013560f81c925081905061132f81611f3a565b915050935093915050565b604080516101408101909152806000815260200161136f60408051606080820183529181019182529081526000602082015290565b815260200161139560408051606080820183529181019182529081526000602082015290565b81526020016113ba604051806040016040528060608152602001600080191681525090565b8152604080516060808201835281526000602082810182905292820152910190815260006020820181905260408201819052606082018190526080820181905260a09091015290565b61140b61202d565b565b60006040828403121561141f57600080fd5b50919050565b6000806000806000808688036101a081121561144057600080fd5b61144a898961140d565b965060408801356001600160401b038082111561146657600080fd5b818a01915061014080838d03121561147d57600080fd5b82985060e0605f198501121561149257600080fd5b60608b0197506114a48c828d0161140d565b9650506101808a01359250808311156114bc57600080fd5b828a0192508a601f8401126114d057600080fd5b82359150808211156114e157600080fd5b508960208284010111156114f457600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b6004811061152f5761152f611509565b9052565b80516007811061154557611545611509565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b808410156115a25761158e828651611533565b93820193600193909301929085019061157b565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b828110156116215784516115ed858251611533565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a0909301926001016115d8565b509687015197909601969096525093949350505050565b805160608084528151848201819052600092602091908201906080870190855b8181101561169b5783518051845285810151868501526040808201519085015286015161168787850182611533565b509284019260a09290920191600101611658565b50508286015183880152604086015193506116ba604088018515159052565b9695505050505050565b60006101008083526116d9818401865161151f565b602085015161014061012081818701526116f7610240870184611552565b9250604088015160ff198088860301848901526117148583611552565b945060608a01519350808886030161016089015261173285856115b6565b945060808a01519350808886030161018089015250506117528383611638565b925060a08801516101a087015260c088015191506117796101c087018363ffffffff169052565b60e088015163ffffffff81166101e088015291509287015163ffffffff811661020087015292870151610220860152509150610cc8905060208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a083015263ffffffff60808201511660c08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156118475761184761180f565b60405290565b604051602081016001600160401b03811182821017156118475761184761180f565b604051608081016001600160401b03811182821017156118475761184761180f565b604051606081016001600160401b03811182821017156118475761184761180f565b60405161014081016001600160401b03811182821017156118475761184761180f565b60405160a081016001600160401b03811182821017156118475761184761180f565b604051601f8201601f191681016001600160401b03811182821017156119205761192061180f565b604052919050565b80356004811061193757600080fd5b919050565b60006001600160401b038211156119555761195561180f565b5060051b60200190565b60006040828403121561197157600080fd5b611979611825565b905081356007811061198a57600080fd5b808252506020820135602082015292915050565b600060408083850312156119b157600080fd5b6119b9611825565b915082356001600160401b03808211156119d257600080fd5b818501915060208083880312156119e857600080fd5b6119f061184d565b8335838111156119ff57600080fd5b80850194505087601f850112611a1457600080fd5b83359250611a29611a248461193c565b6118f8565b83815260069390931b84018201928281019089851115611a4857600080fd5b948301945b84861015611a6e57611a5f8a8761195f565b82529486019490830190611a4d565b8252508552948501359484019490945250909392505050565b803563ffffffff8116811461193757600080fd5b60006040808385031215611aae57600080fd5b611ab6611825565b915082356001600160401b03811115611ace57600080fd5b8301601f81018513611adf57600080fd5b80356020611aef611a248361193c565b82815260a09283028401820192828201919089851115611b0e57600080fd5b948301945b84861015611b775780868b031215611b2b5760008081fd5b611b3361186f565b611b3d8b8861195f565b815287870135858201526060611b54818901611a87565b89830152611b6460808901611a87565b9082015283529485019491830191611b13565b50808752505080860135818601525050505092915050565b8035801515811461193757600080fd5b60006060808385031215611bb257600080fd5b611bba611891565b915082356001600160401b03811115611bd257600080fd5b8301601f81018513611be357600080fd5b80356020611bf3611a248361193c565b82815260a09283028401820192828201919089851115611c1257600080fd5b948301945b84861015611c6c5780868b031215611c2f5760008081fd5b611c3761186f565b86358152848701358582015260408088013590820152611c598b89890161195f565b8189015283529485019491830191611c17565b50865250858101359085015250611c8891505060408301611b8f565b604082015292915050565b60006101408236031215611ca657600080fd5b611cae6118b3565b611cb783611928565b815260208301356001600160401b0380821115611cd357600080fd5b611cdf3683870161199e565b60208401526040850135915080821115611cf857600080fd5b611d043683870161199e565b60408401526060850135915080821115611d1d57600080fd5b611d2936838701611a9b565b60608401526080850135915080821115611d4257600080fd5b50611d4f36828601611b9f565b60808301525060a083013560a0820152611d6b60c08401611a87565b60c0820152611d7c60e08401611a87565b60e0820152610100611d8f818501611a87565b9082015261012092830135928101929092525090565b80356001600160401b038116811461193757600080fd5b600081830360e0811215611dcf57600080fd5b611dd76118d6565b833581526060601f1983011215611ded57600080fd5b611df5611891565b9150611e0360208501611da5565b8252611e1160408501611da5565b6020830152606084013560408301528160208201526080840135604082015260a08401356060820152611e4660c08501611a87565b6080820152949350505050565b600060208284031215611e6557600080fd5b813561ffff81168114610cc857600080fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115611ea057611ea0611e77565b500190565b60006001600160401b0380831681851681830481118215151615611ecb57611ecb611e77565b02949350505050565b60006001600160401b0383811690831681811015611ef457611ef4611e77565b039392505050565b634e487b7160e01b600052601260045260246000fd5b600082611f2157611f21611efc565b500490565b600082611f3557611f35611efc565b500690565b6000600019821415611f4e57611f4e611e77565b5060010190565b60006001600160401b0380841680611f6f57611f6f611efc565b92169190910492915050565b6000816000190483118215151615611f9557611f95611e77565b500290565b600082821015611fac57611fac611e77565b500390565b634e487b7160e01b600052603260045260246000fd5b6000845160005b81811015611fe85760208188018101518583015201611fce565b81811115611ff7576000828501525b5091909101928352506020820152604001919050565b600060ff821660ff81141561202457612024611e77565b60010192915050565b634e487b7160e01b600052605160045260246000fdfea264697066735822122005ebd2447a5e0f9075aaae57bf53a6ea108208c873db53f75c6f131ddd0acdcd64736f6c63430008090033",
}

// OneStepProverMemoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMemoryMetaData.ABI instead.
var OneStepProverMemoryABI = OneStepProverMemoryMetaData.ABI

// OneStepProverMemoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMemoryMetaData.Bin instead.
var OneStepProverMemoryBin = OneStepProverMemoryMetaData.Bin

// DeployOneStepProverMemory deploys a new Ethereum contract, binding an instance of OneStepProverMemory to it.
func DeployOneStepProverMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMemory, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// OneStepProverMemory is an auto generated Go binding around an Ethereum contract.
type OneStepProverMemory struct {
	OneStepProverMemoryCaller     // Read-only binding to the contract
	OneStepProverMemoryTransactor // Write-only binding to the contract
	OneStepProverMemoryFilterer   // Log filterer for contract events
}

// OneStepProverMemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMemorySession struct {
	Contract     *OneStepProverMemory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverMemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMemoryCallerSession struct {
	Contract *OneStepProverMemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverMemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMemoryTransactorSession struct {
	Contract     *OneStepProverMemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverMemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMemoryRaw struct {
	Contract *OneStepProverMemory // Generic contract binding to access the raw methods on
}

// OneStepProverMemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCallerRaw struct {
	Contract *OneStepProverMemoryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactorRaw struct {
	Contract *OneStepProverMemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMemory creates a new instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemory(address common.Address, backend bind.ContractBackend) (*OneStepProverMemory, error) {
	contract, err := bindOneStepProverMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// NewOneStepProverMemoryCaller creates a new read-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMemoryCaller, error) {
	contract, err := bindOneStepProverMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryCaller{contract: contract}, nil
}

// NewOneStepProverMemoryTransactor creates a new write-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMemoryTransactor, error) {
	contract, err := bindOneStepProverMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryTransactor{contract: contract}, nil
}

// NewOneStepProverMemoryFilterer creates a new log filterer instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMemoryFilterer, error) {
	contract, err := bindOneStepProverMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryFilterer{contract: contract}, nil
}

// bindOneStepProverMemory binds a generic wrapper to an already deployed contract.
func bindOneStepProverMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.OneStepProverMemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMemory.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemorySession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa173659b.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),((bytes32,bytes32,bytes32,(uint8,uint256))[],bytes32,bool),bytes32,uint32,uint32,uint32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}
