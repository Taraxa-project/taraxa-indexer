# Taraxa Indexer Specs

A service that scrapes transactions, PBFT and DAG blocks from the Taraxa chains and saves them as aggregates that we need for other ecosystem applications.

## Properties

### Able to keep up with the network

The indexer needs to watch for new data from the node and save it at least as fast as the node produces more data. For this we can use language-specific threads or parallel processing.


### Able to serve data fast

The API component needs to serve data in a fast and consistent way regardless of if the current queried address has 1 or 1M transactions. We’re not optimizing for disk space here so we can save a transaction twice, for example, once for the sender and once for the receiver. We need to also have an agregate for each address that contains the number of DAG and PBFT blocks that it created and the number of total transactions that the address was mentioned in. These will be late used by the ecosystem apps  to construct the pagination.


### Be consistent at any point in time

This means that we can’t see transactions from a block that wasn’t yet saved. This can be achieved multiple ways. For example using batch writes to the db to save all transactions, DAG blocks in period and PBFT block in a transactional way. Another safeguard we can add here is to delete the last saved PBFT block on startup in case the application crashes unexpectedly during its scraping. 

### Local database, local node

From past experiences, it is more complicated to deploy this in a central place. You have to take into account situations where you connect to multiple nodes behind a load balancer and the data is no yet synced between them.

To solve this we can deploy one indexer instance with each node (in K8S same pod, different containers) and use a database that save to a local file.

## The Taraxa Indexer Service has two main components:

1. Restful API

2. Scraper

### Restful API

The Taraxa Indexer service exposes a Restful API, over the aggregates stored in the service, that can be used by various apps the Taraxa Explorer.
The API routes, requests and responses will be automatically generated from the Open API yaml specs. That way we can keep them in sync at all times.

The API has two main groups of endpoints:

1. Address  - Where we have endpoints that return all the transactions, DAG and PBFT blocks, along with their total counts, where the address was mentioned.
2. Validators - This is an aggregate over the PBFT-producing validator nodes filtered by year-week.

### Scraper

The scraper always keeps the “best block” in its internal state and tries to keep up with the node it is connected to by downloading all of the blocks until the current best block.

The connection to the Taraxa chain should use the full-duplex websocket connection that the node provides. This makes thing faster by allowing you to subscribe to node events for the “best block” part and also send on-demand RPC calls for getting data from the node.
We can have a fallback here that uses the http RPC endpoint and pools for the “best blocks” but this should be avoided if possible.

The scraper uses the ETH client from go-ethereum to connect to the Taraxa node.

The scraper needs a separate set of structs that correspond to the specs of the node's RPC responses and the ability to transform the data from the node format to the internal, storable, format.


Example flow:

1. New best block is #1234
2. Get the block from the node
3. Get all transactions in block
4. Save all transactions
5. Get all DAG blocks at that level
6. Save all DAG blocks
7. Save PBFT block

## Storage

The storage part of the scraper needs to have some specific properties in order to both save the data and serve it as fast as possible.

A good option here is to have the storage as decoupled from the rest as the service as possible in order to try multiple options.

LSM-Tree databases should be considered. The Level DB golang implementation doesn't support reverse seek but there are other options like Pebble from cockroachdb or RocksDB.

To not have three sets of structs (api, storage and chain), we can try to use the API models for database persistance by extending them with other functionality.

`oapi-codegen` allows us to generate the types separate from the go server part and we can use this to put the models in the root dir.

### Models

We should only store what we need. No need to save the full transaction data, for example.

1. Account - Used for pagination. Has 3 integers that need to be incremented when an account creates a PBFT block, sends a DAG block or send or receives a transaction.
2. Transaction
3. DAG
4. PBFT 
5. Validator - Represents a validator (address) and how many PBFT blocks they produces in a specific week. 

### A proposition for implementing the storage

In a key-value, sorted database we can constructs paths inside the key that are sortable (byte-sortable) and in ascending or descending order (harder) and iterating over them will be easier (if the database supports skipping, the pagination will also be easier).

Example:

Transaction `0xa3df3b5025517c2b27d1864cc3ddf09e1c430e6d3719da70acc33dc0fbed5aff` was sent from `0x00000000000000000000000000000000000000fe` to `0x00000000000000000000000000000000000000ff`. It was included in block #1234 at index 50

We store the transaction twice in each address namespace:

`tx:0x00000000000000000000000000000000000000fe:1234:50`

`tx:0x00000000000000000000000000000000000000ff:1234:50`

And we also increment the number of total transactions for both the first and the second address.

This way we can interate over the transactions easier when we want to display them:

```
tx:0x00000000000000000000000000000000000000fe:1234:50
tx:0x00000000000000000000000000000000000000fe:1235:1
tx:0x00000000000000000000000000000000000000fe:1235:80
tx:0x00000000000000000000000000000000000000fe:1236:22
tx:0x00000000000000000000000000000000000000fe:2000:5
```

## Other things to consider / Problems we've had in the past

1. The byte values of `0xf001937650bb4f62b57521824b2c20f5b91bea05` and `0xF001937650bb4f62b57521824B2c20f5b91bEa05` are not equal. We need to checksum the addresses both before inserting them into the database and when receiving them from a user.
Same for hashes but in that case we can just make them lowercase.

2. The indexer somehow disconnects from the websocket endpoint or it just doesn't receive data anymore. Hopefully, by using the go-ethereum client, we won't have this problem anymore but we should find a way to reconnect if this happens.

3. We also need to index the genesis block. For this, we can call the `taraxa_getConfig` RPC method, get the initial balances and add fake transactions for them.

4. Handling reorgs is hard. If we don't need to, we shouldn't do it.