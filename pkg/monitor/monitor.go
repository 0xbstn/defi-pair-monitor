package monitor

// CONST
const (
	Arbitrum = "Arbitrum"
	Ethereum = "Ethereum"
)

//VARIABLES

var factoryListArbitrum = []Factory{
	{
		Name:           "UniswapV3",
		FactoryAddress: [32]byte{0x1F98431c8aD98523631AE4a59f267346ea31F984}, // Factory address for UniswapV3 on Arbitrum
	},
	{
		Name:           "SushiSwap",
		FactoryAddress: [32]byte{0xc35DADB65012eC5796536bD9864eD8773aBc74C4}, // Factory address for SushiSwap on Arbitrum
	},
}

var factoryListEthereum = []Factory{
	{
		Name:           "UniswapV3",
		FactoryAddress: [32]byte{0x1F98431c8aD98523631AE4a59f267346ea31F984}, // Factory address for UniswapV3 on Ethereum
	},
	{
		Name:           "UniswapV2",
		FactoryAddress: [32]byte{0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f}, // Factory address for UniswapV2 on Ethereum
	},
}

var ArbitrumMonitor = Monitor{
	Name:        Arbitrum,
	ChainID:     42161,
	RPC:         "https://arb1.arbitrum.io/rpc",
	FactoryList: factoryListArbitrum,
}
var EthereumMonitor = Monitor{
	Name:        Ethereum,
	ChainID:     1,
	RPC:         "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
	FactoryList: factoryListEthereum,
}
