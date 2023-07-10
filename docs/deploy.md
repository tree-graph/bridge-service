# Deployment
Deployment diffs on different chains.

## build first
`go build main.go` to get a binary

## config.yml
- `PK` : the private key to deploy the bridge, without `0x` prefix.
- `RPC`: the json-rpc url to access the chain.

## deploy token vault
### Conflux core-space
Under the repository root, 
Run `./main dev -D` to deploy the whole bridge.

It will generate `cfx-deploy.json` under the repository root.

### Conflux consortium
Since the golang-sdk for conflux consortium chain is incompatible with conflux core-space, we use a different branch 
with different sdk version to deploy.

We may have more operations on consortium chain later, so we need a different directory, even clone this repository to another folder.

- Checkout branch `0628/consortium`
- `mkdir` consortium 
- `go build -o consortium/main main.go`
- `cd consortium`
- `cp ../config.yml .` or create one
- change RPC/PK in config.yml
- `./main dev -D`

The other things are the same with conflux core-space.

### EVM compatible chain
Use instruction in the bridge-contracts repository to deploy. Those tools refer to hardhat and ethers, in typescript/javascript.

## Deploy pegged contract for a test
We will test the bridge on the same chain at first.
For golang supported chains,  

`./main dev -t -D --tid 1 --cid 1001`
- `dev` it's a dev command
- `-t` it's a test flag
- `-D` do deployment flag. It will deploy a ERC721 contract, and a pegged ERC721 contract, and register 4 routes on the same chain.
- `--tid 1` use 1 as token id. It will mint a token with that id. so it should be different under certain case.
- `--cid 1001` use 1001 as logical chain id, see [here](overview.md) for explanation.

This command will:
- creates a ERC721 contract(erc721a)
- creates a pegged ERC721 contract(erc721b)
- register arrival/departure routes for both erc721a and erc721b, on then same chain though.
- transfer ownership of erc721a from the token vault to current account, since we need to mint tokens.
- mint an erc721a token with id 1 to current account.
- transfer `1` to the token vault, with specified information: crossing `1` to contract erc721b on logical chain 1001.

It will generate `pegInfo.json`, which holds two ERC721 contract addresses we needed later.

The bridge will find the crossing request, and claim it, see [here](operate.md)
