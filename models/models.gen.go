// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package models

// Defines values for TransactionType.
const (
	ContractCall             TransactionType = 1
	ContractCreation         TransactionType = 2
	InternalContractCall     TransactionType = 4
	InternalContractCreation TransactionType = 5
	InternalTransfer         TransactionType = 3
	Transfer                 TransactionType = 0
)

// Account defines model for Account.
type Account struct {
	Address Address `json:"address"`
	Balance string  `json:"balance"`
}

// Address defines model for Address.
type Address = string

// AddressFilter defines model for AddressFilter.
type AddressFilter = Address

// BigInt defines model for BigInt.
type BigInt = string

// CallData defines model for CallData.
type CallData struct {
	Name   string `json:"name"`
	Params any    `json:"params"`
}

// CountResponse defines model for CountResponse.
type CountResponse struct {
	Total Uint64 `json:"total"`
}

// Dag defines model for Dag.
type Dag struct {
	Hash             Hash   `json:"hash"`
	Level            Uint64 `json:"level"`
	Timestamp        Uint64 `json:"timestamp"`
	TransactionCount Uint64 `json:"transactionCount"`
}

// DagsPaginatedResponse defines model for DagsPaginatedResponse.
type DagsPaginatedResponse = PaginatedResponse

// EventLog defines model for EventLog.
type EventLog struct {
	Address          Address  `json:"address"`
	Data             string   `json:"data"`
	LogIndex         Uint64   `json:"logIndex"`
	Name             string   `json:"name"`
	Params           any      `json:"params"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  Hash     `json:"transactionHash"`
	TransactionIndex Uint64   `json:"transactionIndex"`
}

// Hash defines model for Hash.
type Hash = string

// HoldersPaginatedResponse defines model for HoldersPaginatedResponse.
type HoldersPaginatedResponse = PaginatedResponse

// InternalTransactionsResponse defines model for InternalTransactionsResponse.
type InternalTransactionsResponse struct {
	Data []Transaction `json:"data"`
}

// OptionalUint64 defines model for OptionalUint64.
type OptionalUint64 = uint64

// PaginatedResponse defines model for PaginatedResponse.
type PaginatedResponse struct {
	End     Uint64 `json:"end"`
	HasNext bool   `json:"hasNext"`
	Start   Uint64 `json:"start"`
	Total   Uint64 `json:"total"`
}

// PaginationFilter defines model for PaginationFilter.
type PaginationFilter struct {
	Limit uint64  `json:"limit"`
	Start *uint64 `json:"start"`
}

// Pbft defines model for Pbft.
type Pbft struct {
	Author           Address `json:"author"`
	Hash             Hash    `json:"hash"`
	Number           Uint64  `json:"number"`
	Timestamp        Uint64  `json:"timestamp"`
	TransactionCount Uint64  `json:"transactionCount"`
}

// PbftsPaginatedResponse defines model for PbftsPaginatedResponse.
type PbftsPaginatedResponse = PaginatedResponse

// Period defines model for Period.
type Period struct {
	EndDate   Uint64 `json:"endDate"`
	HasNext   bool   `json:"hasNext"`
	StartDate Uint64 `json:"startDate"`
}

// StatsResponse defines model for StatsResponse.
type StatsResponse struct {
	DagsCount                Uint64          `json:"dagsCount"`
	LastDagTimestamp         *OptionalUint64 `json:"lastDagTimestamp" rlp:"nil"`
	LastPbftTimestamp        *OptionalUint64 `json:"lastPbftTimestamp" rlp:"nil"`
	LastTransactionTimestamp *OptionalUint64 `json:"lastTransactionTimestamp" rlp:"nil"`
	PbftCount                Uint64          `json:"pbftCount"`
	TransactionsCount        Uint64          `json:"transactionsCount"`
	ValidatorRegisteredBlock *OptionalUint64 `json:"validatorRegisteredBlock" rlp:"nil"`
}

// Transaction defines model for Transaction.
type Transaction struct {
	BlockNumber Uint64          `json:"blockNumber"`
	Calldata    *CallData       `json:"calldata,omitempty" rlp:"nil"`
	From        Address         `json:"from"`
	GasCost     BigInt          `json:"gasCost"`
	Hash        Hash            `json:"hash"`
	Input       string          `json:"input"`
	Status      bool            `json:"status"`
	Timestamp   Uint64          `json:"timestamp"`
	To          Address         `json:"to"`
	Type        TransactionType `json:"type"`
	Value       BigInt          `json:"value"`
}

// TransactionType defines model for Transaction.Type.
type TransactionType uint8

// TransactionLogsResponse defines model for TransactionLogsResponse.
type TransactionLogsResponse struct {
	Data []EventLog `json:"data"`
}

// TransactionsPaginatedResponse defines model for TransactionsPaginatedResponse.
type TransactionsPaginatedResponse = PaginatedResponse

// Uint64 defines model for Uint64.
type Uint64 = uint64

// Validator defines model for Validator.
type Validator struct {
	Address           Address         `json:"address"`
	PbftCount         Uint64          `json:"pbftCount"`
	Rank              Uint64          `json:"rank"`
	RegistrationBlock *OptionalUint64 `json:"registrationBlock" rlp:"nil"`
	Yield             string          `json:"yield,omitempty" rlp:"-"`
}

// ValidatorsPaginatedResponse defines model for ValidatorsPaginatedResponse.
type ValidatorsPaginatedResponse = PaginatedResponse

// Week defines model for Week.
type Week struct {
	Week *int32 `json:"week"`
	Year *int32 `json:"year"`
}

// WeekResponse defines model for WeekResponse.
type WeekResponse struct {
	EndDate   Uint64 `json:"endDate"`
	HasNext   bool   `json:"hasNext"`
	StartDate Uint64 `json:"startDate"`
	Week      *int32 `json:"week"`
	Year      *int32 `json:"year"`
}

// YieldResponse defines model for YieldResponse.
type YieldResponse struct {
	FromBlock Uint64 `json:"fromBlock"`
	ToBlock   Uint64 `json:"toBlock"`
	Yield     string `json:"yield"`
}

// AddressParam defines model for addressParam.
type AddressParam = AddressFilter

// BlockNumParam defines model for blockNumParam.
type BlockNumParam = Uint64

// HashParam defines model for hashParam.
type HashParam = Hash

// PaginationParam defines model for paginationParam.
type PaginationParam = PaginationFilter

// WeekParam defines model for weekParam.
type WeekParam = Week

// GetAddressDagsParams defines parameters for GetAddressDags.
type GetAddressDagsParams struct {
	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetAddressPbftsParams defines parameters for GetAddressPbfts.
type GetAddressPbftsParams struct {
	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetAddressTransactionsParams defines parameters for GetAddressTransactions.
type GetAddressTransactionsParams struct {
	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetAddressYieldParams defines parameters for GetAddressYield.
type GetAddressYieldParams struct {
	// BlockNumber Block Number
	BlockNumber *BlockNumParam `form:"blockNumber,omitempty" json:"blockNumber,omitempty"`
}

// GetAddressYieldForIntervalParams defines parameters for GetAddressYieldForInterval.
type GetAddressYieldForIntervalParams struct {
	// FromBlock From block number
	FromBlock *Uint64 `form:"fromBlock,omitempty" json:"fromBlock,omitempty"`

	// ToBlock To block number
	ToBlock Uint64 `form:"toBlock" json:"toBlock"`
}

// GetHoldersParams defines parameters for GetHolders.
type GetHoldersParams struct {
	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetTotalYieldParams defines parameters for GetTotalYield.
type GetTotalYieldParams struct {
	// BlockNumber Block Number
	BlockNumber *BlockNumParam `form:"blockNumber,omitempty" json:"blockNumber,omitempty"`
}

// GetValidatorsParams defines parameters for GetValidators.
type GetValidatorsParams struct {
	// Week Week to filter by
	Week *WeekParam `form:"week,omitempty" json:"week,omitempty"`

	// Pagination Pagination
	Pagination PaginationParam `form:"pagination" json:"pagination"`
}

// GetValidatorsTotalParams defines parameters for GetValidatorsTotal.
type GetValidatorsTotalParams struct {
	// Week Week to filter by
	Week *WeekParam `form:"week,omitempty" json:"week,omitempty"`
}

// GetValidatorParams defines parameters for GetValidator.
type GetValidatorParams struct {
	// Week Week to filter by
	Week *WeekParam `form:"week,omitempty" json:"week,omitempty"`
}
