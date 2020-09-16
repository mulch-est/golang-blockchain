# Golang Blockchain

Run `go run main.go` to print commands, run `go build main.go` to build an executable file.

### Commands

Usage: `go run main.go getbalance -address <ADDRESS>`

 `getbalance -address <ADDRESS>` ~gets the balance for an address
 
 `createblockchain -address <ADDRESS>` ~creates a blockchain and sends genesis reward to address
 
 `printchain` ~Prints the blocks in the chain
 
 `send -from <FROM> -to <TO> -amount <AMOUNT> -mine` ~Sends <amount> of coins. Then -mine flag is set, mine off of this node
 
 `createwallet` ~Creates a new Wallet *Does not work unless tmp folder has been added manually*
 
 `listaddresses` ~Lists the addresses in our wallet file
 
 `reindexutxo` ~Rebuilds the UTXO set
 
 `startnode -miner <ADDRESS>` ~Start a node with ID specified in NODE_ID env. var. -miner enables mining to <ADDRESS>

### Initializing the Blockchain

1. createwallet
2. createblockchain with created address
3. ready
