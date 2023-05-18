#Cross-chain bridge service

## Add hardhat local chain

```
go run main.go addChain 31337 hardhat http://127.0.0.1:8545
```

## Start web server
```
go run main.go server
```

## Start parsing request worker
```
go run main.go worker
```