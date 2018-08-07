package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LookupBackend(rpcUrl string) (*ethclient.Client, chan bool, error) {
	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, nil, err
	}

	block, err := ethClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Latest known block is: ", block.NumberU64())

	progress, err := ethClient.SyncProgress(context.Background())
	if err != nil {
		return nil, nil, err
	}
	completed := make(chan bool)
	if progress != nil {
		fmt.Println("Client is in syncing state - any operations will be delayed until finished")
		go trackGethProgress(ethClient, progress, completed)
	} else {
		fmt.Println("Geth process fully synced")
		close(completed)
	}

	return ethClient, completed, nil
}
func trackGethProgress(client *ethclient.Client, lastProgress *ethereum.SyncProgress, completed chan<- bool) {
	bar := pb.New64(int64(lastProgress.HighestBlock)).
		SetTotal(int64(lastProgress.CurrentBlock)).
		Start()
	defer bar.Finish()
	defer close(completed)
	for {
		progress, err := client.SyncProgress(context.Background())
		if err != nil {
			fmt.Println("Error querying client progress: " + err.Error())
			return
		}
		if progress == nil {
			bar.Finish()
			fmt.Println("Client in fully synced state. Proceeding...")
			return
		}
		bar.SetCurrent(int64(progress.CurrentBlock))
		bar.SetTotal(int64(progress.HighestBlock))
		time.Sleep(10 * time.Second)
	}
}
