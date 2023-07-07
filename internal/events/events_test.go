package events

import (
	"math/big"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

func TestDecodeRewardsTopics(t *testing.T) {
	// Test case with valid logs
	logs := []models.EventLog{{

		Address:  "0x00000000000000000000000000000000000000fe",
		Data:     "0x000000000000000000000000000000000000000000000005d9da3b556bb3aa86",
		LogIndex: 0,
		Removed:  false,
		Topics: []string{"0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7",
			"0x00000000000000000000000021db400dcb1ef3bc3aee4f3d028ec1939b7fadd6",
			"0x0000000000000000000000004beaf4ce3c239ac7195a1e422725c0465271fb42"},
		TransactionHash:  "0xd8c9296770c696b313128f1cc913b1a5e90ddc62b049ceb8a476b1125d65d3a4",
		TransactionIndex: 1,
	},
		{
			Address:  "0x00000000000000000000000000000000000000fe",
			Data:     "0x000000000000000000000000000000000000000000000000a7a44a964be1f30a",
			LogIndex: 0,
			Removed:  false,
			Topics: []string{
				"0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0",
				"0x0000000000000000000000000dc0d841f962759da25547c686fa440cf6c28c61",
				"0x000000000000000000000000ed4d5f4f3641cbc056e466d15dbe2403e38056f8",
			},
			TransactionHash:  "0xe667503bfec2ade69c5e03398aa29a88e035931cadd2caf265c0c85345f3f40e",
			TransactionIndex: 105,
		},
	}

	decodedEvents, err := DecodeRewardsTopics(logs)
	t.Log(decodedEvents)
	if err != nil {
		t.Errorf("Failed to decode rewards topics: %v", err)
	}

	valueOne, _ := new(big.Int).SetString("107931645057766238854", 10)
	valueTwo, _ := new(big.Int).SetString("12079862109893161738", 10)
	expectedEvents := []LogReward{
		{
			EventName: "RewardsClaimed",
			Account:   "0x21DB400dCB1eF3bC3AEe4f3d028ec1939b7FadD6",
			Validator: "0x4BEAf4ce3c239Ac7195a1e422725c0465271fb42",
			Value:     valueOne,
		},
		{
			EventName: "CommissionRewardsClaimed",
			Account:   "0x0DC0d841F962759DA25547c686fa440cF6C28C61",
			Validator: "0xEd4d5f4F3641cbC056E466d15DBE2403e38056f8",
			Value:     valueTwo,
		},
	}

	if len(decodedEvents) != len(expectedEvents) {
		t.Errorf("Unexpected number of decoded events. Got %d, expected %d", len(decodedEvents), len(expectedEvents))
	}

	for i, event := range decodedEvents {
		expectedEvent := expectedEvents[i]

		if event.EventName != expectedEvent.EventName {
			t.Errorf("Mismatched event name at index %d. Got %s, expected %s", i, event.EventName, expectedEvent.EventName)
		}

		if event.Account != expectedEvent.Account {
			t.Errorf("Mismatched account at index %d. Got %s, expected %s", i, event.Account, expectedEvent.Account)
		}

		if event.Validator != expectedEvent.Validator {
			t.Errorf("Mismatched validator at index %d. Got %s, expected %s", i, event.Validator, expectedEvent.Validator)
		}

		if event.Value.Cmp(expectedEvent.Value) != 0 {
			t.Errorf("Mismatched value at index %d. Got %s, expected %s", i, event.Value.String(), expectedEvent.Value.String())
		}
	}
}
