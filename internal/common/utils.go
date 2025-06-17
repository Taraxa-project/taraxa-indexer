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
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/spiretechnology/go-pool"
	"golang.org/x/exp/constraints"
)

const DposContractAddress = "0x00000000000000000000000000000000000000fe"
const Day = 24 * 60 * 60
const Days30 = 30 * Day

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
		// debug.PrintStack()
		// log.Fatal(s, "ParseUInt ", err)
		return 0
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

func ParseFloat(s string) (v float64) {
	if len(s) == 0 {
		return
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatal(s, "ParseFloat ", err)
	}
	return v
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

type Number interface {
	constraints.Integer | constraints.Float
}

func Max[T Number](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func DayStart(timestamp uint64) uint64 {
	date := time.Unix(int64(timestamp), 0)
	return uint64(time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).Unix())
}

func DayEnd(timestamp uint64) uint64 {
	date := time.Unix(int64(timestamp), 0)
	return uint64(time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, time.UTC).Unix())
}

func MonthInterval(date *uint64) (from_date, to_date uint64) {
	if date == nil {
		to_date = uint64(time.Now().Unix())
	} else {
		to_date = *date
	}
	to_date = DayEnd(to_date - Day)
	from_date = DayStart(to_date - Days30)
	return
}
