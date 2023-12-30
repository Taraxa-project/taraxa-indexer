GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
BUILD_DIR=build
EXECUTABLE=taraxa-indexer

help:
	@echo "This is a helper makefile for taraxa-indexer"
	@echo "Targets:"
	@echo "    lint:        run lint"
	@echo "    generate:    regenerate all api generated files"
	@echo "    check:       run tests"
	@echo "    tidy         tidy go mod"
	@echo "    make         builds executable"
	@echo "	   abigen       generates typings for solidity contracts"

$(GOBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.55.2

.PHONY: tools
tools: $(GOBIN)/golangci-lint

lint: tools
	$(GOBIN)/golangci-lint run ./...

generate:
	go generate ./...

check:
	go test ./...

tidy:
	@echo "tidy..."
	go mod tidy

build: clean
	@mkdir -p $(BUILD_DIR)/linux_amd64
	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/linux_amd64/$(EXECUTABLE)

clean:
	rm -rf $(BUILD_DIR)

abigen:
	abigen --abi=/Users/vargaelod/lara/taraxa-indexer/abi/dpos/Dpos.abi --pkg=dpos_contract --out=/Users/vargaelod/lara/taraxa-indexer/abi/dpos/Dpos.go
	abigen --abi=/Users/vargaelod/lara/taraxa-indexer/abi/oracle/Oracle.abi --pkg=apy_oracle --out=/Users/vargaelod/lara/taraxa-indexer/abi/oracle/Apy_oracle.go
	abigen --abi=/Users/vargaelod/lara/taraxa-indexer/abi/lara/Lara.abi --pkg=lara_contract --out=/Users/vargaelod/lara/taraxa-indexer/abi/lara/Lara.go