package events

import "math/big"

type LogReward struct {
	Account   string
	Validator string
	Value     *big.Int
	EventName string
}

type RewardsClaimedEvent struct {
	Account   string
	Validator string
	Amount    *big.Int
}

type CommissionRewardsClaimedEvent struct {
	Account   string
	Validator string
	Amount    *big.Int
}
