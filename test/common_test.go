package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/magefile/mage/sh"
)

var client *ethclient.Client

const privateKey0 = "45bb96530f3d1972fdcd2005c1987a371d0b6d378b77561c6beeaca27498f46b"

func TestMain(m *testing.M) {
	once := sync.Once{}
	defer once.Do(func() { teardownCompose() })
	err := startCompose()
	if err != nil {
		log.Fatalf("error starting docker compose: %s", err)
	}
	waitForGanache()
	code := m.Run()
	once.Do(func() { teardownCompose() })
	os.Exit(code)
}

func startCompose() error {
	_, err := sh.Output("docker-compose", "-f", "docker-compose.yml", "up", "-d")
	if err != nil {
		return err
	}
	return nil
}

func teardownCompose() (string, error) {
	return sh.Output("docker-compose", "-f", "docker-compose.yml", "down")
}

func waitForGanache() error {
	timeout := time.After(15 * time.Second)
	for {
		select {
		case <-timeout:
			return fmt.Errorf("timed out waiting for ganache to start")
		case <-time.After(200 * time.Millisecond):
			err := connectToGanache()
			if err == nil {
				return nil
			}
		}
	}
}

func connectToGanache() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	client, err = ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		return fmt.Errorf("error connecting to ganache: %s", err)
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	_, err = client.ChainID(ctx2)
	if err != nil {
		return fmt.Errorf("error getting chain ID: %s", err)
	}
	return nil
}
