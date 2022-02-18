### Client

Wrappers for the client used for making RPC calls with useful funcionalities like nonce tracking, multichain clients, ... 

With it we can wrap the contract calls that we might use often using bindings from the bindings folder and objects from the crypto folder and do things like getting the hermes id from the registry smart contract, create transactions that execute smart contract calls, etc.

It is worth mentioning that the client usually creates the transactions but doesn't send them to the blockchain.
The sending is often managed by other packages (like `/transaction`) that also manage nonce and gas price.
It is still able to send the transaction himself but it needs extra logic to handle gas prices and nonce.
