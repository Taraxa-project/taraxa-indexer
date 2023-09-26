package events

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

const testEthereumURL = "https://rpc.mainnet.taraxa.io"

func TestGetValidatorsRegisteredInBlock(t *testing.T) {
	_, err := ethclient.Dial(testEthereumURL)
	if err != nil {
		t.Fatalf("Failed to create Ethereum client: %v", err)
	}

	tests := []struct {
		from             uint64
		to               uint64
		expectRegistries bool
		expectError      bool
	}{
		{1, 1, false, false},            // Test with default from and to
		{1, 100, false, false},          // Test with specific block range
		{101, 100, false, true},         // Test with invalid block range
		{4084000, 4084049, true, false}, // Test with a block range that has validator registrations
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("From %d To %d", test.from, test.to), func(t *testing.T) {
			validators, err := GetValidatorsRegisteredInBlock(testEthereumURL, test.from, test.to)

			if test.expectError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if len(validators) == 0 && test.expectRegistries {
					t.Errorf("Expected at least one validator registration event, but got none")
				}

			}
		})
	}
}
