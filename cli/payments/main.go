package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MysteriumNetwork/payments/cli/helpers"
	"github.com/MysteriumNetwork/payments/promises/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var paymentsContract = flag.String("payments.contract", "", "Address of payments contract")
var identity = flag.String("payments.identity", "", "Identity for balance checking")

func main() {
	flag.Parse()

	client, syncCompleted, err := helpers.LookupBackend()
	checkError(err)
	<-syncCompleted

	contract, err := generated.NewIdentityRegistryCaller(common.HexToAddress(*paymentsContract), client)
	checkError(err)

	paymentsSession := generated.IdentityRegistryCallerSession{
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
