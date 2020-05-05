/* Generated by ts-generator ver. 0.0.8 */
/* tslint:disable */

import { Contract, ContractFactory, Signer } from 'ethers'
import { Provider } from 'ethers/providers'
import { UnsignedTransaction } from 'ethers/utils/transaction'

import { GlobalInbox } from './GlobalInbox'

export class GlobalInboxFactory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(): Promise<GlobalInbox> {
    return super.deploy() as Promise<GlobalInbox>
  }
  getDeployTransaction(): UnsignedTransaction {
    return super.getDeployTransaction()
  }
  attach(address: string): GlobalInbox {
    return super.attach(address) as GlobalInbox
  }
  connect(signer: Signer): GlobalInboxFactory {
    return super.connect(signer) as GlobalInboxFactory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): GlobalInbox {
    return new Contract(address, _abi, signerOrProvider) as GlobalInbox
  }
}

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'chain',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'value',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'bytes',
        name: 'data',
        type: 'bytes',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'messageNum',
        type: 'uint256',
      },
    ],
    name: 'ContractTransactionMessageDelivered',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'chain',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'address',
        name: 'erc20',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'value',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'messageNum',
        type: 'uint256',
      },
    ],
    name: 'ERC20DepositMessageDelivered',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'chain',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'address',
        name: 'erc721',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'id',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'messageNum',
        type: 'uint256',
      },
    ],
    name: 'ERC721DepositMessageDelivered',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'chain',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'value',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'messageNum',
        type: 'uint256',
      },
    ],
    name: 'EthDepositMessageDelivered',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'chain',
        type: 'address',
      },
    ],
    name: 'TransactionMessageBatchDelivered',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'chain',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'seqNumber',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'value',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'bytes',
        name: 'data',
        type: 'bytes',
      },
    ],
    name: 'TransactionMessageDelivered',
    type: 'event',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: '_tokenContract',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_owner',
        type: 'address',
      },
    ],
    name: 'getERC20Balance',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: '_erc721',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_owner',
        type: 'address',
      },
    ],
    name: 'getERC721Tokens',
    outputs: [
      {
        internalType: 'uint256[]',
        name: '',
        type: 'uint256[]',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: '_owner',
        type: 'address',
      },
    ],
    name: 'getEthBalance',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: '_erc721',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_owner',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_tokenId',
        type: 'uint256',
      },
    ],
    name: 'hasERC721',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: '_owner',
        type: 'address',
      },
    ],
    name: 'ownedERC20s',
    outputs: [
      {
        internalType: 'address[]',
        name: '',
        type: 'address[]',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: '_owner',
        type: 'address',
      },
    ],
    name: 'ownedERC721s',
    outputs: [
      {
        internalType: 'address[]',
        name: '',
        type: 'address[]',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: '_tokenContract',
        type: 'address',
      },
    ],
    name: 'withdrawERC20',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: '_erc721',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_tokenId',
        type: 'uint256',
      },
    ],
    name: 'withdrawERC721',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [],
    name: 'withdrawEth',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: 'account',
        type: 'address',
      },
    ],
    name: 'getInbox',
    outputs: [
      {
        internalType: 'bytes32',
        name: '',
        type: 'bytes32',
      },
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'bytes',
        name: '_messages',
        type: 'bytes',
      },
    ],
    name: 'sendMessages',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: '_chain',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_seqNumber',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: '_value',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'sendTransactionMessage',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: '_chain',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
    ],
    name: 'depositEthMessage',
    outputs: [],
    payable: true,
    stateMutability: 'payable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: '_chain',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_erc20',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_value',
        type: 'uint256',
      },
    ],
    name: 'depositERC20Message',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: '_chain',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_erc721',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_id',
        type: 'uint256',
      },
    ],
    name: 'depositERC721Message',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_from',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_value',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'forwardContractTransactionMessage',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_from',
        type: 'address',
      },
    ],
    name: 'forwardEthMessage',
    outputs: [],
    payable: true,
    stateMutability: 'payable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: 'chain',
        type: 'address',
      },
      {
        internalType: 'bytes',
        name: 'transactions',
        type: 'bytes',
      },
    ],
    name: 'deliverTransactionBatch',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
]

const _bytecode =
  '0x608060405234801561001057600080fd5b5061221a806100206000396000f3fe6080604052600436106101095760003560e01c80638f5ed73e11610095578063bca22b7611610064578063bca22b7614610569578063c3a8962c146105b2578063e4eb8c63146105ed578063f3e414f814610668578063f4f3b200146106a157610109565b80638f5ed73e146103fb57806396588a271461049b578063a0ef91df146104c9578063b7dabd7a146104de57610109565b80634d2301cc116100dc5780634d2301cc1461026f5780635bd21290146102b45780636e2b89c5146102e457806384cb7997146103175780638b7010aa146103b257610109565b8063022016811461010e5780630758fb0a1461015a57806333f2ac42146101e557806345a53f0914610218575b600080fd5b34801561011a57600080fd5b506101416004803603602081101561013157600080fd5b50356001600160a01b03166106d4565b6040805192835260208301919091528051918290030190f35b34801561016657600080fd5b506101956004803603604081101561017d57600080fd5b506001600160a01b03813581169160200135166106f7565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101d15781810151838201526020016101b9565b505050509050019250505060405180910390f35b3480156101f157600080fd5b506101956004803603602081101561020857600080fd5b50356001600160a01b03166107bd565b34801561022457600080fd5b5061025b6004803603606081101561023b57600080fd5b506001600160a01b03813581169160208101359091169060400135610880565b604080519115158252519081900360200190f35b34801561027b57600080fd5b506102a26004803603602081101561029257600080fd5b50356001600160a01b0316610900565b60408051918252519081900360200190f35b6102e2600480360360408110156102ca57600080fd5b506001600160a01b038135811691602001351661091b565b005b3480156102f057600080fd5b506101956004803603602081101561030757600080fd5b50356001600160a01b0316610934565b34801561032357600080fd5b506102e26004803603608081101561033a57600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b81111561037457600080fd5b82018360208201111561038657600080fd5b803590602001918460018302840111600160201b831117156103a757600080fd5b5090925090506109eb565b3480156103be57600080fd5b506102e2600480360360808110156103d557600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610a35565b34801561040757600080fd5b506102e2600480360360a081101561041e57600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359181019060a081016080820135600160201b81111561045d57600080fd5b82018360208201111561046f57600080fd5b803590602001918460018302840111600160201b8311171561049057600080fd5b509092509050610a53565b6102e2600480360360408110156104b157600080fd5b506001600160a01b0381358116916020013516610a9f565b3480156104d557600080fd5b506102e2610ab4565b3480156104ea57600080fd5b506102e26004803603604081101561050157600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561052b57600080fd5b82018360208201111561053d57600080fd5b803590602001918460018302840111600160201b8311171561055e57600080fd5b509092509050610aff565b34801561057557600080fd5b506102e26004803603608081101561058c57600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610bdb565b3480156105be57600080fd5b506102a2600480360360408110156105d557600080fd5b506001600160a01b0381358116916020013516610bf3565b3480156105f957600080fd5b506102e26004803603602081101561061057600080fd5b810190602081018135600160201b81111561062a57600080fd5b82018360208201111561063c57600080fd5b803590602001918460018302840111600160201b8311171561065d57600080fd5b509092509050610c5c565b34801561067457600080fd5b506102e26004803603604081101561068b57600080fd5b506001600160a01b038135169060200135610d20565b3480156106ad57600080fd5b506102e2600480360360208110156106c457600080fd5b50356001600160a01b0316610de4565b6001600160a01b0316600090815260036020526040902080546001909101549091565b6001600160a01b038082166000908152600260209081526040808320938616835290839052902054606091908061074057505060408051600081526020810190915290506107b7565b81600101600182038154811061075257fe5b90600052602060002090600302016002018054806020026020016040519081016040528092919081815260200182805480156107ad57602002820191906000526020600020905b815481526020019060010190808311610799575b5050505050925050505b92915050565b6001600160a01b03811660009081526002602090815260409182902060018101548351818152818402810190930190935260609290918391801561080b578160200160208202803883390190505b50805190915060005b818110156108765783600101818154811061082b57fe5b600091825260209091206003909102015483516001600160a01b039091169084908390811061085657fe5b6001600160a01b0390921660209283029190910190910152600101610814565b5090949350505050565b6001600160a01b038083166000908152600260209081526040808320938716835290839052812054909190806108bb576000925050506108f9565b8160010160018203815481106108cd57fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b6001600160a01b031660009081526020819052604090205490565b61092482610eb1565b61093082823334610ed0565b5050565b6001600160a01b03811660009081526001602081815260409283902091820154835181815281830281019092019093526060928391908015610980578160200160208202803883390190505b50805190915060005b81811015610876578360010181815481106109a057fe5b600091825260209091206002909102015483516001600160a01b03909116908490839081106109cb57fe5b6001600160a01b0390921660209283029190910190910152600101610989565b610a2e3386868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f6f92505050565b5050505050565b610a4082858361107a565b610a4d84843385856110f6565b50505050565b610a97868633878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061119092505050565b505050505050565b610aa833610eb1565b61093033838334610ed0565b6000610abf33610900565b3360008181526020819052604080822082905551929350909183156108fc0291849190818181858888f19350505050158015610930573d6000803e3d6000fd5b333214610b41576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b6000600683834342604051602001808660ff1660ff1660f81b81526001018585808284378083019250505083815260200182815260200195505050505050604051602081830303815290604052805190602001209050610ba1848261123f565b6040516001600160a01b038516907f9cd591b0e52bcf1c506475ee03776192ea3d99f35150ef6651b339333b7372c490600090a250505050565b610be6828583611275565b610a4d8484338585611302565b6001600160a01b03808216600090815260016020908152604080832093861683529083905281205490919080610c2e576000925050506107b7565b816001016001820381548110610c4057fe5b9060005260206000209060020201600101549250505092915050565b6000808080845b80841015610d1757610cac87878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525088925061139c915050565b9297509095509350915084610cc057610d17565b610d0387878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525088925087915061145e9050565b909550935084610d1257610d17565b610c63565b50505050505050565b610d2b338383611562565b610d7c576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b158015610dd057600080fd5b505af1158015610a97573d6000803e3d6000fd5b6000610df08233610bf3565b9050610dfd3383836117ca565b610e385760405162461bcd60e51b815260040180806020018281038252602e8152602001806121b8602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b158015610e8757600080fd5b505af1158015610e9b573d6000803e3d6000fd5b505050506040513d6020811015610a4d57600080fd5b6001600160a01b03166000908152602081905260409020805434019055565b6001600160a01b03841660009081526003602052604081206001908101540190610efe85858543428761195d565b9050610f0a868261123f565b336001600160a01b0316856001600160a01b0316876001600160a01b03167ffd0d0553177fec183128f048fbde54554a3a67302f7ebd7f735215a3582907053486604051808381526020018281526020019250505060405180910390a4505050505050565b6001600160a01b03851660009081526003602052604081206001908101540190610f9e868686864342886119c7565b9050610faa878261123f565b846001600160a01b0316866001600160a01b0316886001600160a01b03167f362b3acbdbf0277aefa83754ea8d39fc1c55d01d9351cf78969923f8cfee612c8787876040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561103557818101518382015260200161101d565b50505050905090810190601f1680156110625780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a450505050505050565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd91606480830192600092919082900301818387803b1580156110ce57600080fd5b505af11580156110e2573d6000803e3d6000fd5b505050506110f1828483611aa6565b505050565b6001600160a01b0385166000908152600360205260408120600190810154019061112586868686434288611c2a565b9050611131878261123f565b604080516001600160a01b0386811682526020820186905281830185905291518288169289811692908b16917f40baf11a4a4a4be2a155dbf303fbaec6fabd52e267268bd7e3de4b4ed8a2e0959181900360600190a450505050505050565b60006111a9878787878787805190602001204342611c49565b90506111b5878261123f565b846001600160a01b0316866001600160a01b0316886001600160a01b03167fcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b38787876040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360008381101561103557818101518382015260200161101d565b6001600160a01b038216600090815260036020526040902080546112639083611cc0565b81556001908101805490910190555050565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd9160648083019260209291908290030181600087803b1580156112ca57600080fd5b505af11580156112de573d6000803e3d6000fd5b505050506040513d60208110156112f457600080fd5b506110f19050828483611cec565b6001600160a01b0385166000908152600360205260408120600190810154019061133186868686434288611dc3565b905061133d878261123f565b604080516001600160a01b0386811682526020820186905281830185905291518288169289811692908b16917fb13d04085b4a9f87fecfccf9b72081bb8a273498d6b08b4bccf2940d555b5e609181900360600190a450505050505050565b60008060008060008060008088905060008a82815181106113b957fe5b016020015160019092019160f81c9050600681146113e95750600097508896508795508594506114559350505050565b6113f38b83611dd6565b9196509094509150846114185750600097508896508795508594506114559350505050565b6114228b83611dd6565b9196509093509150846114475750600097508896508795508594506114559350505050565b506001975095509093509150505b92959194509250565b60008060018314156114b35760008060008061147a8989611e4e565b93509350935093508361149757600088955095505050505061155a565b6114a2338383611e9b565b50600183955095505050505061155a565b600283141561150c5760008060008060006114ce8a8a611ef9565b94509450945094509450846114ee5760008996509650505050505061155a565b6114fa33838584612000565b5060018496509650505050505061155a565b60038314156115535760008060008060006115278a8a611ef9565b94509450945094509450846115475760008996509650505050505061155a565b6114fa33838584612030565b5060009050825b935093915050565b6001600160a01b0380841660009081526002602090815260408083209386168352908390528120549091908061159d576000925050506108f9565b60008260010160018303815481106115b157fe5b6000918252602080832088845260016003909302019182019052604090912054909150806115e65760009450505050506108f9565b6002820180548291600185019160009190600019810190811061160557fe5b60009182526020808320909101548352820192909252604001902055600282018054600019810190811061163557fe5b906000526020600020015482600201600183038154811061165257fe5b60009182526020808320909101929092558781526001840190915260408120556002820180548061167f57fe5b60008281526020812082016000199081019190915501905560028201546117bc57600184018054849186916000919060001981019081106116bc57fe5b600091825260208083206003909202909101546001600160a01b0316835282019290925260400190205560018401805460001981019081106116fa57fe5b906000526020600020906003020184600101600185038154811061171a57fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b039092169190911781556002808301805461175d92840191906120ee565b5050506001600160a01b0387166000908152602085905260408120556001840180548061178657fe5b60008281526020812060036000199093019283020180546001600160a01b0319168155906117b7600283018261213e565b505090555b506001979650505050505050565b6000816117d9575060016108f9565b6001600160a01b03808516600090815260016020908152604080832093871683529083905290205480611811576000925050506108f9565b600082600101600183038154811061182557fe5b90600052602060002090600202019050806001015485111561184d57600093505050506108f9565b60018101805486900390819055611950576001830180548391859160009190600019810190811061187a57fe5b600091825260208083206002909202909101546001600160a01b0316835282019290925260400190205560018301805460001981019081106118b857fe5b90600052602060002090600202018360010160018403815481106118d857fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061192657fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b5060019695505050505050565b60408051600160f81b6020808301919091526001600160601b03196060998a1b811660218401529790981b909616603587015260498601949094526069850192909252608984015260a9808401919091528151808403909101815260c99092019052805191012090565b6000600488888888888888604051602001808960ff1660ff1660f81b8152600101886001600160a01b03166001600160a01b031660601b8152601401876001600160a01b03166001600160a01b031660601b815260140186815260200185805190602001908083835b60208310611a4f5780518252601f199092019160209182019101611a30565b51815160209384036101000a600019018019909216911617905292019586525084810193909352506040808401919091528051808403820181526060909301905281519101209d9c50505050505050505050505050565b6001600160a01b03808416600090815260026020908152604080832093861683529083905290205480611b66576040805180820182526001600160a01b0386811682528251600080825260208083019095528484019182526001878101805491820180825590835291869020855160039092020180546001600160a01b03191691909416178355905180519194611b459260028501929091019061215f565b5050506001600160a01b038516600090815260208490526040902081905590505b6000826001016001830381548110611b7a57fe5b9060005260206000209060030201905080600101600085815260200190815260200160002054600014611bf4576040805162461bcd60e51b815260206004820152601d60248201527f63616e27742061646420616c7265616479206f776e656420746f6b656e000000604482015290519081900360640190fd5b60028101805460018181018355600083815260208082209093018890559254968352909201909152604090209290925550505050565b6000611c3d600389898989898989612054565b98975050505050505050565b6040805160006020808301919091526001600160601b031960609b8c1b81166021840152998b1b8a1660358301529790991b9097166049890152605d880194909452607d870192909252609d86015260bd85015260dd808501919091528251808503909101815260fd909301909152815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b80611cf6576110f1565b6001600160a01b03808416600090815260016020908152604080832093861683529083905290205480611d8f57506040805180820182526001600160a01b0385811680835260006020808501828152600188810180548083018083559186528486209851600290910290980180546001600160a01b03191698909716979097178655905194019390935590815290849052919091208190555b82826001016001830381548110611da257fe5b60009182526020909120600160029092020101805490910190555050505050565b6000611c3d600289898989898989612054565b6000806000808551905084811080611df057506021858203105b80611e125750600060ff16868681518110611e0757fe5b016020015160f81c14155b15611e27575060009250839150829050611e47565b600160218601611e3f8888840163ffffffff6120d216565b935093509350505b9250925092565b60008060008060008060008088905060008a8281518110611e6b57fe5b016020015160019092019160f81c9050600581146113e95750600097508896508795508594506114559350505050565b6001600160a01b038316600090815260208190526040812054821115611ec3575060006108f9565b506001600160a01b0392831660009081526020819052604080822080548490039055929093168352912080549091019055600190565b6000806000806000806000806000808a905060008c8281518110611f1957fe5b016020015160019092019160f81c905060068114611f4d5750600099508a9850899750879650869550611ff6945050505050565b611f578d83611dd6565b919750909550915085611f805750600099508a9850899750879650869550611ff6945050505050565b611f8a8d83611dd6565b919750909450915085611fb35750600099508a9850899750879650869550611ff6945050505050565b611fbd8d83611dd6565b919750909350915085611fe65750600099508a9850899750879650869550611ff6945050505050565b5060019950975091955093509150505b9295509295909350565b600061200d8584846117ca565b61201957506000612028565b612024848484611cec565b5060015b949350505050565b600061203d858484611562565b61204957506000612028565b612024848484611aa6565b6040805160f89990991b6001600160f81b0319166020808b0191909152606098891b6001600160601b031990811660218c015297891b881660358b01529590971b9095166049880152605d870192909252607d860152609d85015260bd808501929092528251808503909201825260dd909301909152805191012090565b600081602001835110156120e557600080fd5b50016020015190565b82805482825590600052602060002090810192821561212e5760005260206000209182015b8281111561212e578254825591600101919060010190612113565b5061213a92915061219a565b5090565b508054600082559060005260206000209081019061215c919061219a565b50565b82805482825590600052602060002090810192821561212e579160200282015b8281111561212e57825182559160200191906001019061217f565b6121b491905b8082111561213a57600081556001016121a0565b9056fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a72315820642bcc81a696cfdbaa13504fead56eb71ba0f9f435490a900267fad1c976fb1d64736f6c634300050f0032'
