# Golang Blockchain

Run `go run main.go` to print commands, run `go build main.go` to build an executable file.

### Commands

Usage: `go run main.go createblockchain -<address>`

 getbalance -address ADDRESS - get the balance for an address
 createblockchain -address ADDRESS creates a blockchain and sends genesis reward to address
 printchain - Prints the blocks in the chain
 send -from FROM -to TO -amount AMOUNT -mine - Send amount of coins. Then -mine flag is set, mine off of this node
 createwallet - Creates a new Wallet
 listaddresses - Lists the addresses in our wallet file
 reindexutxo - Rebuilds the UTXO set
 startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining
