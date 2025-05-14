.PHONY: check-copyright
check-copyright:
	go run cmd/copyright/main.go

.PHONY: generate
generate:
	go generate

.PHONY: generate-matic
generate-matic:
	go run bindings/abi/abigen.go --localdir=./bindings/matic/abi --contracts=ChildContract.json,RootChainManager.json,RootChain.json --out=bindings/matic --pkg=matic

.PHONY: test
test:
	go test --short -race -cover ./...
