# Golang Blockchain

Run `go run main.go` to print commands, run `go build main.go` to build an executable file.

### Commands

Usage: `go run main.go getbalance -address <ADDRESS>`

 `getbalance -address <ADDRESS>` ~gets the balance of ADDRESS
 
 `createblockchain -address <ADDRESS>` ~creates a blockchain and sends genesis reward to ADDRESS
 
 `printchain` ~Prints the blocks in the chain
 
 `send -from <FROM> -to <TO> -amount <AMOUNT> -mine` ~Sends AMOUNT of coins from FROM to TO. Then -mine flag is set, mine off of this node
 
 `createwallet` ~Creates a new wallet
 
 `listaddresses` ~Lists the addresses in our wallet file
 
 `reindexutxo` ~Rebuilds the UTXO set
 
 `startnode -miner <ADDRESS>` ~Start a node with ID specified in NODE_ID env. var. -miner enables mining to ADDRESS

### Initializing the Blockchain

1. Create wallet<br>
`go run main.go createwallet`

2. Create blockchain with genesis reward sent to created wallet<br>
`go run main.go createblockchain -address <CREATED_ADDRESS>`

3.  Check validity<br>
of address `go run main.go listaddresses` OR `go run main.go getbalance -address <CREATED_ADDRESS>`<br>
of chain `go run main.go printchain`

If you received the error **NODE_ID env is not set!**, it can be fixed using the command `set NODE_ID=<node-id>` on Windows using [cmder](https://github.com/cmderdev/cmder).

## Limitations

1. Not truly decentralized, one node is set as the centralized node.

2. When a block download attempt is made, the entire blockchain is downloaded.

3. When a send attempt is made, the program checks valid wallets from local memory (inside the node).
