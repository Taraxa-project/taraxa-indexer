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

## How to run

Before running the indexer, ensure your Taraxa node is running with the `--rpc.debug` flag enabled. This is required for the indexer to access all necessary RPC endpoints. It is better to connect to local node to ensure best performance.

You can run the indexer with the following CLI arguments (all have defaults):

- `--http_port` (default: 8080): Port for the API server.
- `--metrics_port` (default: 2112): Port for the Prometheus metrics server.
- `--blockchain_ws` (default: wss://ws.mainnet.taraxa.io): WebSocket URL to connect to the Taraxa blockchain node.
- `--data_dir` (default: ./data): Directory to store the indexer database and logs.
- `--log_level` (default: info): Log level (`trace`, `debug`, `info`, `warn`, `error`, `fatal`).
- `--sync_queue_limit` (default: 10): Limit of blocks in the sync queue.
- `--chain_stats_interval` (default: 100): Interval for saving chain stats.
- `--auth_username` (default: taraxa): Username for protected endpoints (required if `auth_password` is set).
- `--auth_password` (default: taraxa): Password for protected endpoints (required if `auth_username` is set).

### Example

```sh
go run main.go --http_port=8080 --blockchain_ws=wss://ws.your-node:port --data_dir=./data --log_level=debug
```

**Note:** Make sure your Taraxa node is running and accessible at the WebSocket URL you provide, and that it was started with the `--rpc.debug` flag.

