package monitor

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func listenForPairCreatedEvents(monitor *Monitor, factory Factory) {
	client, err := ethclient.Dial(monitor.RPC)
	if err != nil {
		log.Fatalf("Failed to connect to %s RPC: %v", monitor.Name, err)
	}
	defer client.Close()

	// Define the PairCreated event signature
	pairCreatedSig := []byte("PairCreated(address,address,address,uint256)")

	// Set up a query to filter logs for the PairCreated event signature
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			common.BytesToAddress(factory.FactoryAddress[:]),
		},
		Topics: [][]common.Hash{
			{
				common.BytesToHash(pairCreatedSig),
			},
		},
	}

	logs := make(chan types.Log)

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to get block number for %s: %v", monitor.Name, err)
	} else {
		log.Printf("Connected to %s RPC. Current block number: %d", monitor.Name, blockNumber)
	}

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to %s logs: %v", monitor.Name, err)
	}
	defer sub.Unsubscribe()
	log.Printf("Successfully subscribed to %s logs", monitor.Name)

	// Listen for new PairCreated events
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logs:
			log.Printf("New log received for %s", monitor.Name)
			token0 := common.BytesToAddress(vLog.Topics[1].Bytes())
			token1 := common.BytesToAddress(vLog.Topics[2].Bytes())
			pair := common.BytesToAddress(vLog.Topics[3].Bytes())
			pairNumber := new(big.Int).SetBytes(vLog.Data)

			var logMessage string
			if monitor.Name == Ethereum {
				logMessage = fmt.Sprintf("New pair created on %s: token0: %s, token1: %s, pair: %s, pairNumber: %s", Ethereum, token0.Hex(), token1.Hex(), pair.Hex(), pairNumber.String())
			} else if monitor.Name == Arbitrum {
				logMessage = fmt.Sprintf("New pair created on %s: token0: %s, token1: %s, pair: %s, pairNumber: %s", Arbitrum, token0.Hex(), token1.Hex(), pair.Hex(), pairNumber.String())
			}
			fmt.Println(logMessage)
		}
	}
}
