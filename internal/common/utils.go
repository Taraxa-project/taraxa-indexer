package common

import (
	"fmt"
	"log"
	"math/big"
	"reflect"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spiretechnology/go-pool"
)

const DposContractAddress = "0x00000000000000000000000000000000000000fe"

// isn't creating threads, but limiting goroutines count. Mostly used for RPC and db related tasks
func MakeThreadPool() pool.Pool {
	return pool.New(uint(runtime.NumCPU()))
}

func ParseUInt(s string) (v uint64) {
	if len(s) == 0 {
		return
	}
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatal(s, "ParseUInt ", err)
	}
	return v
}

func ParseInt(s string) (v int64) {
	if len(s) == 0 {
		return
	}
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatal(s, "ParseUInt ", err)
	}
	return v
}

func ParseBool(s string) (v bool) {
	if len(s) == 0 {
		return
	}
	i, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatal("parseBool ", v)
	}
	return i > 0
}

func ParseStringToBigInt(v string) *big.Int {
	a := big.NewInt(0)
	a.SetString(v, 0)
	return a
}

func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 4, 64)
}

func GetYieldIntervalEnd(pbft_count uint64, block_num *uint64, interval uint64) uint64 {
	block := uint64(0)
	if block_num == nil {
		if pbft_count < interval {
			return interval
		}
		block = pbft_count - interval
	} else {
		block = *block_num
	}

	if block%interval == 0 {
		return block
	}
	block = block - block%interval + interval
	return block
}

func ParseToString(item any) (result any, err error) {
	switch val := item.(type) {
	case ethcommon.Address:
		result = strings.ToLower(val.Hex())
	case string:
		result = val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		result = fmt.Sprintf("%d", val)
	case float32:
		result = FormatFloat(float64(val))
	case float64:
		result = FormatFloat(val)
	case *big.Int:
		result = val.String()
	case []byte:
		result = fmt.Sprintf("0x%x", val)
	case bool:
		result = fmt.Sprintf("%t", val)
	default:
		// log.Error("ParseToString default ", reflect.TypeOf(item).Kind(), item)
		if reflect.TypeOf(item).Kind() == reflect.Array || reflect.TypeOf(item).Kind() == reflect.Slice {
			sliceValue := reflect.ValueOf(item)
			sliceLen := sliceValue.Len()
			sliceResult := make([]any, sliceLen)
			for j := 0; j < sliceLen; j++ {
				toDecode := sliceValue.Index(j).Interface()
				res, _ := ParseToString(toDecode)
				sliceResult[j] = res
			}
			result = sliceResult
		} else {
			return nil, fmt.Errorf("unsupported type %v", reflect.TypeOf(item))
		}
	}
	return
}

func DecodePaddedHex(hexStr string) (uint64, error) {
	// Decode the hex string to bytes.
	bytes, err := hexutil.Decode(hexStr)
	if err != nil {
		return 0, err
	}

	// Convert bytes to big.Int.
	bigInt := new(big.Int).SetBytes(bytes)
	// convert to uint64
	uint64Value := bigInt.Uint64()
	return uint64Value, nil
}

func DecodePaddedAddress(hexStr string) (common.Address, error) {
	// Decode the hex string to bytes.
	bytes, err := hexutil.Decode(hexStr)
	if err != nil {
		return common.Address{}, err
	}

	// Convert bytes to big.Int.
	bigInt := new(big.Int).SetBytes(bytes)
	// convert to uint64
	address := common.BigToAddress(bigInt)
	return address, nil
}
