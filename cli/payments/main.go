package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/cli/helpers"
	"github.com/mysteriumnetwork/payments/promises"
)

var paymentsContract = flag.String("payments.contract", "", "Address of payments contract")
var identity = flag.String("payments.identity", "", "Identity for balance checking")
var gethUrl = flag.String("geth.url", "", "URL value of started geth to connect")

func main() {
	flag.Parse()

	client, syncCompleted, err := helpers.LookupBackend(*gethUrl)
	checkError(err)
	<-syncCompleted

	contract, err := promises.NewIdentityRegistryCaller(common.HexToAddress(*paymentsContract), client)
	checkError(err)

	paymentsSession := promises.IdentityRegistryCallerSession{
		Contract: contract,
		CallOpts: bind.CallOpts{},
	}

	registered, err := paymentsSession.IsRegistered(common.HexToAddress(*identity))
	checkError(err)
	fmt.Println("Identity: ", *identity, "registration status: ", registered)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}
