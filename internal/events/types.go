package events

import "math/big"

type LogReward struct {
	Account   string
	Validator string
	Value     *big.Int
	EventName string
}
