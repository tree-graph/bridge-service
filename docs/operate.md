# Operations
Bridge operators/developers can interact with their bridges for deployment, maintenance or testing.

## config.yml
Config.yml example:
```
log:
  level: DEBUG
  forceColor: True
  file: ./logs/bridge.log

api:
  endpoint: 0.0.0.0:8001
store:
  mysql:
    host: 127.0.0.1:3306
    username: *
    password: *
    database: *
```

## config database
Each chain must have a Token vault contract, get it from [here](deploy.md).

Each chain has a unique logical id, explained [here](overview.md).

### config chains
The table `chains`:
- `id` a logical id, specified by the operator, should not be changed.
- `chain_id` the blockchain scope chain id. Use 1029 for conflux consortium if it's 0 in rpc respond.
- `name` a more human readable name for the chain.
- `vault_addr` the token vault contract address, keep the raw format, for example: cfx:xxx for conflux consortium.
- `rpc` the rpc url of the chain.
- `chain_type` the type of the chain: evm or cfx.
- `delay_block` how many blocks/epochs we consider as `safe` before doing work.
- `created_at` use `now()`
- `updated_at` use `now()`
- `enabled` the status of the chain, use `1` if the golang service could take care of the chain,
use `0` if golang service could not but other application such as java-plugin could. 
Event it's `0`, do not remove it, plugin application shall fetch it.
The operator must config involved chains here.

### config cursor for the event fetching
The table `chain_cursors` tracks the progress of fetching event on each chain. 
Since the chain may have run for a long time, the operator should use a recent cursor to startup.
The testing step [here](deploy.md) will show the epoch number, use its value-1 should be convenient.
- `id` the logical id of the chain.
- `block` the last processed block/epoch of the chain.
- `latest_block` the bridge take care of this field, use 0 for manual operation.
- `updated_at` use `now()`

### config secrets/private key
The table `secrets` keeps private keys used when sending claiming transactions.
For security reasons we separate it from table `chains`.
- `id` the logical id of the chain.
- `chain_id` the blockchain scope chain id.
- `private` the private key, with '0x' prefix.
- `address` the address of the private key, use '' if you don't got it, the bridge will compute and fill it.
- `comment`
- `created_at` use `now()`
- `updated_at` use `now()`

## Run the server
Run the web server, which expose api for bridge-plugin and frontend(later).  
`./main server`

## Run event fetcher
Golang event fetching worker will track crossing request for chains marked with `enabled` `1`. 
`./main worker`

## Run claiming worker
Golang claiming worker will handle requests for chains  marked with `enabled` `1`.  
`./main worker -D`

## Run plugin application
Event fetcher and claiming worker for conflux consortium is [here](https://github.com/tree-graph/bridge-plugin-treegraph/tree/fetch-event).

## About the first test crossing
We have sent A test crossing request [here](deploy.md).
If everything goes well, the bridge should send a claiming transaction, and the pegged 
contract erc721b should mint a NFT for current account.

For conflux consortium chain, a java application will do works above. 
Please refer to its docs.

## More than one chain
After archiving success on one chain, move on to multiple chains.
For example, we want to cross NFT on the conflux core-space testnet to a conflux consortium chain,
we should follow steps below(use 1001 and 1029 for logical chain ids):  
1. run `./main dev -t -D --tid 1 --cid 1001` on core-space testnet
2. run `./main dev -t -D --tid 1 --cid 1029` on consortium
3. run ```
./main dev --register --cid 1029 \
                     --local cfxtest:xxx(erc721a on 1001) \
                     --remote cfx:xxx(erc721b on 1029)
                     ``` on core-space testnet, to register routes.
4. run ```
./main dev --register --cid 1001 \
                     --remote cfxtest:xxx(erc721a on 1001) \
                     --local cfx:xxx(erc721b on 1029)
                     ``` on consortium chain, to register routes.
5. run `./main dev -t  --tid 91 --cid 1029` on core-space testnet, 
to mint a nft with id 91, and cross it to chain 1029.
6. wait java application for consortium chain claiming.
7. run `./main dev -t  --tid 91 --cid 1001 -r` on consortium chain, 
to cross that nft back to chain 1001.
8. wait golang service claiming.

