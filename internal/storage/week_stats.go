package storage

import (
	"sort"

	"github.com/Taraxa-project/taraxa-indexer/models"
)

type WeekStats struct {
	Validators []models.Validator
	Total      uint32
	Key        []byte `rlp:"-"`
}

func MakeEmptyWeekStats() *WeekStats {
	data := new(WeekStats)
	return data
}

func (w *WeekStats) Sort() {
	sort.Slice(w.Validators, func(i, j int) bool {
		return w.Validators[i].PbftCount > w.Validators[j].PbftCount
	})
}

func (w *WeekStats) AddPbftBlock(block *models.Pbft) {
	w.Total++
	for k, v := range w.Validators {
		if v.Address == block.Author {
			w.Validators[k].PbftCount++
			return
		}
	}
	w.Validators = append(w.Validators, models.Validator{Address: block.Author, PbftCount: 1})
}

func (w *WeekStats) GetPaginated(from, count uint64) ([]models.Validator, *models.PaginatedResponse) {
	pagination := new(models.PaginatedResponse)
	pagination.Total = uint64(len(w.Validators))
	if from > pagination.Total {
		from = pagination.Total
	}
	pagination.Start = from
	end := from + count
	pagination.HasNext = (end < pagination.Total)
	if end > pagination.Total {
		end = pagination.Total
	}
	pagination.End = end

	w.Sort()
	var validators []models.Validator

	for k, v := range w.Validators[from:end] {
		v.Rank = uint64(k + 1)
		validators = append(validators, v)
	}

	return validators, pagination
}
