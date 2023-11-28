package oracle

import (
	"context"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/assert/v2"
)

func TestCalculateRating(t *testing.T) {
	commission := uint64(100)
	// Prepare test data
	validator := YieldedValidator{
		Account:           common.HexToAddress("0x0DC0d841F962759DA25547c686fa440cF6C28C61"),
		Yield:             "0.55",
		Commisson:         &commission,
		Rank:              uint16(1),
		RegistrationBlock: uint64(1),
		PbftCount:         uint64(10000),
	}
	client, err := ethclient.Dial("ws://localhost:8777")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	currentBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	// Validate the rating
	rating, _, _ := validator.calculateRating(client)
	commission_float := float64(*validator.Commisson)
	yield_float, err := strconv.ParseFloat(validator.Yield, 64)
	if err != nil {
		t.Fatalf("Failed to parse yield: %v", err)
	}
	commission_percentage := commission_float / float64(100000)
	adjusted_apy := (1 - commission_percentage) * yield_float * 100
	blocksSinceRegistration := currentBlock.NumberU64() - validator.RegistrationBlock
	continuity := float64(blocksSinceRegistration) / float64(currentBlock.NumberU64()-validator.RegistrationBlock)

	expected_rating := float64(0.4)*adjusted_apy - float64(0.1)*commission_float + float64(0.5)*continuity
	assert.Equal(t, int64(expected_rating*1000), rating)
}
