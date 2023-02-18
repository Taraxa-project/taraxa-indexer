//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=models.cfg.yaml ../../petstore-expanded.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ../../petstore-expanded.yaml

package api

import (
	"sync"
)

type IndexerApi struct {
	Lock sync.Mutex
}

func NewIndexerApi() *IndexerApi {
	return &IndexerApi{}
}
