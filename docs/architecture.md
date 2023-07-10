#System Architecture
For now, it's a centralized bridge. 

## Main steps
The main steps to cross a NFT from one chain to another are listed below:

- User Bob gains a NFT from contract C0.
- Bob sends that NFT to TokenVault through `C0.safeTransferFrom`, with target chain id and target contract encoded in data parameter.
- The TokenVault contract emits a CrossRequest event.
- The bridge fetches the CrossRequest event, saves it to the database.
- The claiming worker(in the bridge, or in a stand alone bridge-plugin) finds that request, and claims it on the destination chain for Bob.

## Chain ID
The bridge supports multiple chains. Since different consortium chains may have the same chain id, we use a
logical id for each chain.

## Token Vault
On each chain we deploy only one token vault contract, to keep assets and run claiming business. It's also a route described below.

## Route
The TokenVault contract inherits Router contract, which maintains route information. One route information describes 
a behavior when a NFT leaving from or arriving to that chain. Operator must register routes on both source chain and destination chain.

## Claiming nonce
The bridge must prevent replaying claiming transactions. To archive it,
the bridge maintains nonce records on source chains, one nonce for each user on each destination chain.
The bridge increases it when user crossing a NFT.
On the destination chain, the bridge checks that nonce, reverts transactions with mismatched claiming nonce.


## Consortium plugin
To support different chains, we use plugin mechanism. A conflux consortium plugin in java is available [here](https://github.com/tree-graph/bridge-plugin-treegraph/tree/fetch-event).