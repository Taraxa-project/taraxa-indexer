package events

import (
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/stretchr/testify/assert"
)

func TestDecodeEventDynamicDpos(t *testing.T) {
	// Test RewardsClaimedEvent
	rewardsClaimedTopic := "0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7" // Example topic for RewardsClaimed event
	rewardsClaimedData := "0x000000000000000000000000000000000000000000000000095e271c526c181b"  // Example data for RewardsClaimed event
	rewardsClaimedLog := models.EventLog{
		Address:          "0x00000000000000000000000000000000000000fe",
		LogIndex:         0,
		Removed:          false,
		Topics:           []string{rewardsClaimedTopic, "0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec", "0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec"},
		Data:             rewardsClaimedData,
		TransactionHash:  "0xecd3243842f3fe4a0e8419a0ac85e4a6098bde7635de53bcad18e57f80cc6463",
		TransactionIndex: 1,
	}

	{
		name, decoded, err := DecodeEventDynamic(rewardsClaimedLog)
		t.Log(name)
		t.Log(decoded)
		if err != nil {
			t.Errorf("Failed to decode RewardsClaimedEvent: %v", err)
		}
	}
	// Test CommissionRewardsClaimedEvent
	commissionRewardsClaimedTopic := "0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0" // Example topic for CommissionRewardsClaimed event
	commissionRewardsClaimedData := "0x000000000000000000000000000000000000000000000000a7a44a964be1f30a"  // Example data for CommissionRewardsClaimed event
	commissionRewardsClaimedLog := models.EventLog{
		Address:          "0x00000000000000000000000000000000000000fe",
		LogIndex:         0,
		Removed:          false,
		Topics:           []string{commissionRewardsClaimedTopic, "0x0000000000000000000000000dc0d841f962759da25547c686fa440cf6c28c61", "0x000000000000000000000000ed4d5f4f3641cbc056e466d15dbe2403e38056f8"},
		Data:             commissionRewardsClaimedData,
		TransactionHash:  "0xe667503bfec2ade69c5e03398aa29a88e035931cadd2caf265c0c85345f3f40e",
		TransactionIndex: 1,
	}

	{
		name, decoded, err := DecodeEventDynamic(commissionRewardsClaimedLog)
		if err != nil {
			t.Errorf("Failed to decode CommissionRewardsClaimedEvent: %v", err)
		}
		t.Log(name)
		t.Log(decoded)
	}
}

func TestDecodeEventDynamicUnknownAddress(t *testing.T) {
	// Test RewardsClaimedEvent
	rewardsClaimedTopic := "0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7" // Example topic for RewardsClaimed event
	rewardsClaimedData := "0x000000000000000000000000000000000000000000000000095e271c526c181b"  // Example data for RewardsClaimed event
	rewardsClaimedLog := models.EventLog{
		Address:          "0xca67d0c50C4f5363aA47a11C54486aD8F28A7F2b",
		LogIndex:         0,
		Removed:          false,
		Topics:           []string{rewardsClaimedTopic, "0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec", "0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec"},
		Data:             rewardsClaimedData,
		TransactionHash:  "0xecd3243842f3fe4a0e8419a0ac85e4a6098bde7635de53bcad18e57f80cc6463",
		TransactionIndex: 1,
	}

	{
		name, decoded, err := DecodeEventDynamic(rewardsClaimedLog)
		assert.Equal(t, name, "")
		assert.Nil(t, decoded)
		assert.Nil(t, err)
	}
}

func TestDecodeEventDynamicInvalidTopic(t *testing.T) {
	// Test RewardsClaimedEvent
	rewardsClaimedTopic := "0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7aa" // Example topic for RewardsClaimed event
	rewardsClaimedData := "0x000000000000000000000000000000000000000000000000095e271c526c181b"    // Example data for RewardsClaimed event
	rewardsClaimedLog := models.EventLog{
		Address:          "0x00000000000000000000000000000000000000fe",
		LogIndex:         0,
		Removed:          false,
		Topics:           []string{rewardsClaimedTopic, "0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec", "0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec"},
		Data:             rewardsClaimedData,
		TransactionHash:  "0xecd3243842f3fe4a0e8419a0ac85e4a6098bde7635de53bcad18e57f80cc6463",
		TransactionIndex: 1,
	}

	{
		name, decoded, err := DecodeEventDynamic(rewardsClaimedLog)
		assert.Equal(t, name, "")
		assert.Nil(t, decoded)
		assert.EqualError(t, err, "no event with id: 0x307831306363666362386465373233663537386139653432383265613966353231663035616534306463303866333036386466616435323861363565653363376161")
	}
}

func TestDecodeEventDynamicInvalidData(t *testing.T) {
	// Test RewardsClaimedEvent
	rewardsClaimedTopic := "0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7"  // Example topic for RewardsClaimed event
	rewardsClaimedData := "0x000000000000000000000000000000000000000000000000095e271c526c181bzz" // Example data for RewardsClaimed event
	rewardsClaimedLog := models.EventLog{
		Address:          "0x00000000000000000000000000000000000000fe",
		LogIndex:         0,
		Removed:          false,
		Topics:           []string{rewardsClaimedTopic, "0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec", "0x000000000000000000000000fc43217e71ec0a1cc480f3d210cd07cbde7374ec"},
		Data:             rewardsClaimedData,
		TransactionHash:  "0xecd3243842f3fe4a0e8419a0ac85e4a6098bde7635de53bcad18e57f80cc6463",
		TransactionIndex: 1,
	}

	{
		name, decoded, err := DecodeEventDynamic(rewardsClaimedLog)
		assert.Equal(t, name, "")
		assert.Nil(t, decoded)
		assert.EqualError(t, err, "encoding/hex: invalid byte: U+007A 'z'")
	}
}
