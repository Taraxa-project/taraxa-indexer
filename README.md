# Description
The Taraxa Indexer saves all PBFT Blocks, DAG Blocks and Transactions on the Taraxa chain and exposes an API on top of that data that can be used on the Taraxa Explorer.

## Structure

```
.
├── api
│   ├── api_handler.go
│   ├── openapi.yaml
│   ├── server.cfg.yaml
│   └── server.gen.go
├── go.mod
├── go.sum
├── internal
│   ├── address
│   │   └── address.go
│   ├── dag
│   │   └── dag.go
│   ├── pbft
│   │   └── pbft.go
│   ├── storage
│   │   └── storage.go
│   └── tx
│       └── tx.go
├── main.go
├── models
│   └── models.gen.go
├── models.cfg.yaml
```

1. /api - contains all (http) API specs and routes

`openapi.yaml` - the OpenAPI definition that is used to generate the models and server boilerplate

`api_handler.go` - where we implement the endpoints

`server.cfg.yaml` - the config file for `oapi-codegen`

`server.gen.go` - the generated code for the API server

2. /internal - should contain all most of the code. each dir is a component that can include data models, actions to be called from the api, etc.

3. /models - stucts used both for the API and storage layers

`models.gen.go` - the generated code for types used in both the API and storage layers

4. /models.cfg.yaml - the config file for `oapi-codegen`

## Develop

`go generate ./...` will regenerate the server.gen.go and models.gen.go files

`make lint` before each commit