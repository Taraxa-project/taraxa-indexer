package indexer

import (
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/utils"
	"github.com/Taraxa-project/taraxa-indexer/models"
)

func TestUpdateBalancesInternal(t *testing.T) {
	// Prepare test data
	accounts := []models.Account{
		{
			Address: "0x1111111111111111111111111111111111111111",
			Balance: "100",
		},
		{
			Address: "0x0DC0d841F962759DA25547c686fa440cF6C28C61",
			Balance: "50",
		},
	}
	trx := models.Transaction{
		From:  "0x1111111111111111111111111111111111111111",
		To:    "0x0DC0d841F962759DA25547c686fa440cF6C28C61",
		Value: "20",
	}

	// Invoke the method
	err := UpdateBalancesInternal(&accounts, trx)

	// Validate the result
	if err != nil {
		t.Errorf("UpdateBalancesInternal failed with error: %v", err)
	}

	// Validate the updated balances
	i := utils.FindBalance(&accounts, "0x1111111111111111111111111111111111111111")
	t.Log(accounts[i])
	if i == -1 || accounts[i].Balance != "80" {
		t.Error("UpdateBalancesInternal failed to update 'from' balance correctly")
	}

	j := utils.FindBalance(&accounts, "0x0DC0d841F962759DA25547c686fa440cF6C28C61")
	t.Log(accounts[j])
	if j == -1 || accounts[j].Balance != "70" {
		t.Error("UpdateBalancesInternal failed to update 'to' balance correctly")
	}
}

func TestUpdateBalances(t *testing.T) {
	// Prepare test data
	accounts := []models.Account{
		{
			Address: "0x1111111111111111111111111111111111111111",
			Balance: "100",
		},
		{
			Address: "0x0DC0d841F962759DA25547c686fa440cF6C28C61",
			Balance: "50",
		},
	}
	trx := &chain.Transaction{
		Logs: []chain.EventLog{{

			Address:          "0x00000000000000000000000000000000000000fe",
			Data:             "0x000000000000000000000000000000000000000000000005d9da3b556bb3aa86",
			LogIndex:         "0",
			Removed:          false,
			Topics:           []string{"0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7", "0x00000000000000000000000021db400dcb1ef3bc3aee4f3d028ec1939b7fadd6", "0x0000000000000000000000004beaf4ce3c239ac7195a1e422725c0465271fb42"},
			TransactionHash:  "0xd8c9296770c696b313128f1cc913b1a5e90ddc62b049ceb8a476b1125d65d3a4",
			TransactionIndex: "1",
		},
			{
				Address:  "0x00000000000000000000000000000000000000fe",
				Data:     "0x000000000000000000000000000000000000000000000000a7a44a964be1f30a",
				LogIndex: "0",
				Removed:  false,
				Topics: []string{
					"0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0",
					"0x0000000000000000000000000dc0d841f962759da25547c686fa440cf6c28c61",
					"0x000000000000000000000000ed4d5f4f3641cbc056e466d15dbe2403e38056f8",
				},
				TransactionHash:  "0xe667503bfec2ade69c5e03398aa29a88e035931cadd2caf265c0c85345f3f40e",
				TransactionIndex: "105",
			},
		},
		Transaction: models.Transaction{
			Value: "30",
			From:  "0x1111111111111111111111111111111111111111",
			To:    "0x0DC0d841F962759DA25547c686fa440cF6C28C61",
		},
	}

	// Invoke the method
	err := UpdateBalances(&accounts, trx)

	// Validate the result
	if err != nil {
		t.Errorf("UpdateBalances failed with error: %v", err)
	}

	// Validate the updated balances
	i := utils.FindBalance(&accounts, "0x1111111111111111111111111111111111111111")
	if i == -1 || accounts[i].Balance != "70" {
		t.Error("UpdateBalances failed to update 'from' balance correctly. Sould be 70 but is ", accounts[i].Balance)
	}

	j := utils.FindBalance(&accounts, "0x0DC0d841F962759DA25547c686fa440cF6C28C61")
	if j == -1 || accounts[j].Balance != "12079862109893161818" {
		t.Error("UpdateBalances failed to update 'to' balance correctly. Should be 12079862109893161818 but is ", accounts[j].Balance)
	}
}
