package pebble

import (
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/nleeper/goment"
	"github.com/stretchr/testify/assert"
)

func TestGetter(t *testing.T) {
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	addr := storage.MakeEmptyAddressStats("test")
	if err := st.addToDBTest(addr, addr.Address, 0); err != nil {
		t.Error(err)
	}
	addr1 := storage.MakeEmptyAddressStats("test1")
	if err := st.addToDBTest(addr1, addr1.Address, 0); err != nil {
		t.Error(err)
	}
	ret := st.GetAddressStats("test")
	if !ret.IsEqual(addr) {
		t.Error("Broken DB")
	}
}

func TestGetObjects(t *testing.T) {
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	sender := "user"
	count := uint64(100)

	stats := storage.MakeEmptyAddressStats(sender)
	stats.DagsCount = count
	if err := st.addToDBTest(stats, stats.Address, 0); err != nil {
		t.Error(err)
	}

	for i := uint64(1); i <= count; i++ {
		block := models.Dag{Timestamp: i, Hash: "test" + strconv.FormatUint(i, 10), Level: 0, TransactionCount: 0}
		if err := st.addToDBTest(&block, sender, block.Timestamp); err != nil {
			t.Error(err)
		}
	}
	ret, pagination := storage.GetObjectsPage[models.Dag](st, sender, 0, uint64(count))
	assert.Equal(t, count, uint64(len(ret)))
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination = storage.GetObjectsPage[models.Dag](st, sender, 50, 100)
	assert.Equal(t, len(ret), 50)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(50))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination = storage.GetObjectsPage[models.Dag](st, sender, 0, 25)
	assert.Equal(t, len(ret), 25)
	assert.True(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(25))
}

func TestGetPaginatedWeekStats(t *testing.T) {
	stor := NewStorage("")
	defer func() { _ = stor.Close() }()

	tn, _ := goment.New()
	weekStats := stor.GetWeekStats(int32(tn.ISOWeekYear()), int32(tn.ISOWeek()))

	count := uint64(100)
	for i := uint64(1); i <= count; i++ {
		block := models.Pbft{Author: "user" + strconv.FormatUint(i, 10), Hash: "test" + strconv.FormatUint(i, 10), Number: i, Timestamp: uint64(tn.ToUnix()), TransactionCount: 0}
		weekStats.AddPbftBlock(&block)
	}

	ret, pagination := weekStats.GetPaginated(0, count)
	assert.Equal(t, uint64(len(ret)), count)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination = weekStats.GetPaginated(50, 50)
	assert.Equal(t, len(ret), 50)
	assert.False(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(50))
	assert.Equal(t, pagination.End, uint64(100))

	ret, pagination = weekStats.GetPaginated(0, 25)
	assert.Equal(t, len(ret), 25)
	assert.True(t, pagination.HasNext)
	assert.Equal(t, pagination.Start, uint64(0))
	assert.Equal(t, pagination.End, uint64(25))
}

func TestStorage(t *testing.T) {
	addr := storage.MakeEmptyAddressStats("test")
	{
		storage := NewStorage("/tmp/test")
		if err := storage.addToDBTest(addr, addr.Address, 0); err != nil {
			t.Error(err)
		}
		_ = storage.Close()
	}
	{
		storage := NewStorage("/tmp/test")
		defer func() { _ = storage.Close() }()
		ret := storage.GetAddressStats("test")
		if !ret.IsEqual(addr) {
			t.Error("Broken DB")
		}
	}
	_ = os.Remove("/tmp/test")
}

func TestCleanStorage(t *testing.T) {
	stats := storage.MakeEmptyAddressStats("test")
	stats.PbftCount = 1
	st := NewStorage("/tmp/test")
	defer func() { _ = st.Close() }()

	if err := st.addToDBTest(stats, stats.Address, 0); err != nil {
		t.Error(err)
	}

	if err := st.Clean(); err != nil {
		t.Error(err)
	}

	err := st.GetFromDB(stats, getKey(GetPrefix(stats), "test", 0))

	if err == nil {
		t.Error("Clean DB does not work")
		_ = os.Remove("/tmp/test")
	}
}

func TestBatch(t *testing.T) {
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	addr := storage.MakeEmptyAddressStats("test")
	batch := st.NewBatch()

	batch.Add(addr, addr.Address, 0)
	addr1 := storage.MakeEmptyAddressStats("test1")
	batch.Add(addr1, addr1.Address, 0)
	batch.CommitBatch()

	ret := st.GetAddressStats("test")
	if !ret.IsEqual(addr) {
		t.Error("Broken DB")
	}

	ret = st.GetAddressStats("test1")
	if !ret.IsEqual(addr1) {
		t.Error("Broken DB")
	}
}

func TestAccountsBatch(t *testing.T) {
	st := NewStorage("")
	defer func() { _ = st.Close() }()
	accounts := storage.MakeAccountBalancesMap()
	accounts.GetAccounts()["0x1111111111111111111111111111111111111111"] = big.NewInt(100)
	accounts.GetAccounts()["0x0DC0d841F962759DA25547c686fa440cF6C28C61"] = big.NewInt(50)

	batch := st.NewBatch()
	batch.SaveHoldersLeaderboard(accounts.Sorted())
	batch.CommitBatch()

	ret := st.GetAccounts()

	if len(ret.Accounts) != accounts.GetLength() {
		t.Error("Broken DB")
	}

	sorted := accounts.Sorted()
	for i, acc := range sorted.Accounts {
		assert.Equal(t, acc.Address, ret.Accounts[i].Address)
		assert.Equal(t, acc.Balance, ret.Accounts[i].Balance)
	}
}

func TestTxByHash(t *testing.T) {
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	tx := models.Transaction{
		Hash:  "0x111111",
		From:  "0x222222",
		Value: "100",
		To:    "0x00000000000000000000000000000000000000fe",
		Input: "0x5c19a95c000000000000000000000000ed4d5f4f3641cbc056e466d15dbe2403e38056f8",
	}

	batch := st.NewBatch()
	batch.AddSingleKey(tx, tx.Hash)
	batch.CommitBatch()

	ret := st.GetTransactionByHash(tx.Hash)
	err := common.ProcessTransaction(&ret)

	assert.NoError(t, err)
	assert.Equal(t, tx.Hash, ret.Hash)
	assert.Equal(t, tx.Input, ret.Input)
	assert.Equal(t, []any{"0xed4d5f4f3641cbc056e466d15dbe2403e38056f8"}, ret.Calldata.Params)
}

func fillTransactions(st *Storage, address string, count uint64) {
	for i := uint64(0); i < count; i++ {
		tx := models.Transaction{
			Hash:  "0x111111",
			From:  address,
			Value: "100",
			To:    "0x00000000000000000000000000000000000000fe",
			Input: strconv.FormatUint(i, 10),
		}
		batch := st.NewBatch()
		batch.Add(tx, tx.From, i)
		batch.CommitBatch()
	}
}

func TestForEach(t *testing.T) {
	st := NewStorage("")
	defer func() { _ = st.Close() }()

	fillTransactions(st, "test", 100)

	var tx models.Transaction
	i := uint64(0)
	st.ForEach(&tx, "test", nil, storage.Forward, func(key, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &tx)
		assert.NoError(t, err)
		assert.Equal(t, tx.Input, strconv.FormatUint(i, 10))
		i++
		return false
	})
	assert.Equal(t, uint64(100), i)
}

func TestForEachBackwards(t *testing.T) {
	st := NewStorage("")
	defer func() { _ = st.Close() }()
	count := uint64(100)
	fillTransactions(st, "test", count)

	var tx models.Transaction
	st.ForEach(&tx, "test", nil, storage.Backward, func(key, res []byte) (stop bool) {
		var tx models.Transaction
		err := rlp.DecodeBytes(res, &tx)
		assert.NoError(t, err)

		count--
		assert.Equal(t, tx.Input, strconv.FormatUint(count, 10))
		return false
	})
	assert.Equal(t, uint64(0), count)
}

func TestWasAccountActive(t *testing.T) {
	address := "0x123"
	db := NewStorage(t.TempDir())
	defer func() { _ = db.Close() }()

	batch := db.NewBatch()
	batch.Add(&models.Transaction{
		From:      address,
		Timestamp: 200,
	}, address, 1)
	batch.CommitBatch()

	assert.Equal(t, storage.WasAccountActive(db, address, 100, 200), true)
	assert.Equal(t, storage.WasAccountActive(db, address, 100, 1000), true)
	assert.Equal(t, storage.WasAccountActive(db, address, 1000, 2000), false)
	assert.Equal(t, storage.WasAccountActive(db, address, 1000, 10000), false)
}

func TestReceivedTransactionsCount(t *testing.T) {
	address := "0x123"
	db := NewStorage(t.TempDir())
	defer func() { _ = db.Close() }()

	batch := db.NewBatch()
	i := uint64(1)
	batch.Add(models.Transaction{
		To:        address,
		Timestamp: 200,
	}, address, i)
	i++
	batch.Add(models.Transaction{
		To:        address,
		Timestamp: 100,
	}, address, i)
	i++

	// shouldn't be counted
	batch.Add(models.Transaction{
		From:      address,
		Timestamp: 200,
	}, address, i)
	i++
	batch.Add(models.Transaction{
		To:        address,
		Timestamp: 201,
	}, address, i)
	i++

	batch.Add(models.Transaction{
		To:        address,
		Timestamp: 99,
	}, address, i)

	batch.CommitBatch()

	assert.Equal(t, storage.ReceivedTransactionsCount(db, address, 100, 200), uint64(2))
	assert.Equal(t, storage.ReceivedTransactionsCount(db, address, 90, 100), uint64(2))
	assert.Equal(t, storage.ReceivedTransactionsCount(db, address, 200, 202), uint64(1))
}

func TestParsePrefixUint(t *testing.T) {
	timestamp := uint64(1750667269)
	str := "0000000000" + strconv.FormatUint(timestamp, 10)
	to_date, err := strconv.ParseUint(str, 10, 64)
	assert.NoError(t, err)
	assert.Equal(t, timestamp, to_date)
}

func TestKeyUpperBound(t *testing.T) {
	// Test the keyUpperBound function logic
	keyUpperBound := func(b []byte) []byte {
		end := make([]byte, len(b))
		copy(end, b)
		for i := len(end) - 1; i >= 0; i-- {
			end[i] = end[i] + 1
			if end[i] != 0 {
				return end[:i+1]
			}
		}
		return nil // no upper-bound
	}

	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "simple prefix",
			input:    []byte("user_"),
			expected: []byte("user`"), // ASCII 95 + 1 = 96, truncated to 5 bytes
		},
		{
			name:     "numeric prefix",
			input:    []byte("123"),
			expected: []byte("124"),
		},
		{
			name:     "single byte",
			input:    []byte("a"),
			expected: []byte("b"),
		},
		{
			name:     "empty byte",
			input:    []byte{},
			expected: nil,
		},
		{
			name:     "all 0xFF bytes",
			input:    []byte{0xFF, 0xFF, 0xFF},
			expected: nil, // overflow case
		},
		{
			name:     "mixed bytes with overflow",
			input:    []byte{0x01, 0xFF, 0x02},
			expected: []byte{0x01, 0xFF, 0x03}, // Only last byte incremented
		},
		{
			name:     "hex prefix",
			input:    []byte("0x1234"),
			expected: []byte("0x1235"),
		},
		{
			name:     "address prefix",
			input:    []byte("0x00000000000000000000000000000000000000fe"),
			expected: []byte("0x00000000000000000000000000000000000000ff"),
		},
		{
			name:     "timestamp prefix",
			input:    []byte("timestamp_1234567890"),
			expected: []byte("timestamp_1234567891"),
		},
		{
			name:     "single 0xFF",
			input:    []byte{0xFF},
			expected: nil, // overflow
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := keyUpperBound(tt.input)
			assert.Equal(t, tt.expected, result,
				"keyUpperBound(%v) = %v, expected %v",
				tt.input, result, tt.expected)
		})
	}
}

func TestKeyUpperBoundRangeBehavior(t *testing.T) {
	// Test that the upper bound creates the correct range
	keyUpperBound := func(b []byte) []byte {
		end := make([]byte, len(b))
		copy(end, b)
		for i := len(end) - 1; i >= 0; i-- {
			end[i] = end[i] + 1
			if end[i] != 0 {
				return end[:i+1]
			}
		}
		return nil
	}

	tests := []struct {
		name        string
		prefix      []byte
		keys        [][]byte
		shouldMatch []bool
	}{
		{
			name:   "user prefix range",
			prefix: []byte("user_"),
			keys: [][]byte{
				[]byte("user_123"),
				[]byte("user_abc"),
				[]byte("user_"),
				[]byte("user_999"),
				[]byte("user`"),     // upper bound (truncated)
				[]byte("user_a"),    // should be included
				[]byte("other_123"), // should not match
				[]byte("users_123"), // should not match
			},
			shouldMatch: []bool{true, true, true, true, false, true, false, false},
		},
		{
			name:   "numeric prefix range",
			prefix: []byte("123"),
			keys: [][]byte{
				[]byte("123"),
				[]byte("1234"),
				[]byte("123abc"),
				[]byte("124"), // upper bound
				[]byte("122"),
				[]byte("12"),
			},
			shouldMatch: []bool{true, true, true, false, false, false},
		},
		{
			name:   "empty prefix",
			prefix: []byte{},
			keys: [][]byte{
				[]byte("any"),
				[]byte("key"),
			},
			shouldMatch: []bool{true, true}, // empty prefix matches everything
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upperBound := keyUpperBound(tt.prefix)

			for i, key := range tt.keys {
				shouldMatch := tt.shouldMatch[i]

				// Check if key is within range [prefix, upperBound)
				withinRange := true

				// Check lower bound
				if len(key) < len(tt.prefix) {
					withinRange = false
				} else {
					for j := 0; j < len(tt.prefix); j++ {
						if key[j] != tt.prefix[j] {
							withinRange = false
							break
						}
					}
				}

				// Check upper bound
				if withinRange && upperBound != nil {
					if len(key) >= len(upperBound) {
						// Compare up to upperBound length
						for j := 0; j < len(upperBound); j++ {
							if key[j] > upperBound[j] {
								withinRange = false
								break
							} else if key[j] < upperBound[j] {
								break
							}
						}
					}
				}

				assert.Equal(t, shouldMatch, withinRange,
					"Key %v should match prefix %v (upperBound: %v): expected %v, got %v",
					key, tt.prefix, upperBound, shouldMatch, withinRange)
			}
		})
	}
}

func TestKeyUpperBoundEdgeCases(t *testing.T) {
	keyUpperBound := func(b []byte) []byte {
		end := make([]byte, len(b))
		copy(end, b)
		for i := len(end) - 1; i >= 0; i-- {
			end[i] = end[i] + 1
			if end[i] != 0 {
				return end[:i+1]
			}
		}
		return nil
	}

	t.Run("all 0xFF bytes overflow", func(t *testing.T) {
		input := []byte{0xFF, 0xFF, 0xFF}
		result := keyUpperBound(input)
		assert.Nil(t, result, "All 0xFF bytes should result in nil (no upper bound)")
	})

	t.Run("single 0xFF byte overflow", func(t *testing.T) {
		input := []byte{0xFF}
		result := keyUpperBound(input)
		assert.Nil(t, result, "Single 0xFF byte should result in nil")
	})

	t.Run("zero bytes", func(t *testing.T) {
		input := []byte{0x00, 0x00, 0x00}
		result := keyUpperBound(input)
		expected := []byte{0x00, 0x00, 0x01} // Only last byte incremented
		assert.Equal(t, expected, result)
	})

	t.Run("mixed overflow pattern", func(t *testing.T) {
		input := []byte{0x01, 0xFF, 0xFF, 0x02}
		result := keyUpperBound(input)
		expected := []byte{0x01, 0xFF, 0xFF, 0x03} // Only last byte incremented
		assert.Equal(t, expected, result)
	})
}
