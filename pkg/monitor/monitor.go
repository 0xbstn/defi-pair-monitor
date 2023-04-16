package monitor

import (
	"fmt"
	"github.com/0xbstn/defi-pair-monitor/config"
	"github.com/ethereum/go-ethereum/common"
)

// CONST
const (
	Arbitrum = "Arbitrum"
	Ethereum = "Ethereum"
)

// VARIABLES
var configInstance = config.GetConfigInstance("config/config.json")

var factoryListArbitrum = []Factory{
	{
		Name:           "UniswapV3",
		FactoryAddress: common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"), // Factory address for UniswapV3 on Arbitrum
	},
	{
		Name:           "SushiSwap",
		FactoryAddress: common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"), // Factory address for SushiSwap on Arbitrum
	},
}

var factoryListEthereum = []Factory{
	{
		Name:           "UniswapV3",
		FactoryAddress: common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"), // Factory address for UniswapV3 on Ethereum
	},
	{
		Name:           "UniswapV2",
		FactoryAddress: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), // Factory address for UniswapV2 on Ethereum
	},
}

var ArbitrumMonitor = Monitor{
	Name:        Arbitrum,
	ChainID:     42161,
	RPC:         configInstance.Arbitrum,
	FactoryList: factoryListArbitrum,
}
var EthereumMonitor = Monitor{
	Name:        Ethereum,
	ChainID:     1,
	RPC:         configInstance.Ethereum,
	FactoryList: factoryListEthereum,
}

func (m *Monitor) StartListeningForPairCreatedEvents() {
	for _, factory := range m.FactoryList {
		go listenForPairCreatedEvents(m, factory)
	}
}

func StartPairTracking() {
	// Start tracking pairs on Ethereum and Arbitrum
	fmt.Println("Starting pair tracking Arbitrum...")
	ArbitrumMonitor.StartListeningForPairCreatedEvents()
	fmt.Println("Starting pair tracking Ethereum...")
	EthereumMonitor.StartListeningForPairCreatedEvents()

	// Block the main goroutine forever
	select {}
}
