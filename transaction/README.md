## Transaction

Everything that has to do with sending transactions successfully. Depot: It works in a similar fashion to a regular delivery depot in real life. It has Workers which check for incoming transactions, pack them and give them away to Couriers which try to deliver the transaction. Transaction state is tracked just like a regular package would be, trying to deliver it again if it was not accepted etc. It is composed of 2 main objects:

- Courier: which is in charge of signing and sending the transactions to the blockchain with the given parameters.
- Depot: It works in a similar fashion to a regular delivery depot in real life. It has Workers which check for incoming transactions, pack them and give them away to Couriers which try to deliver the transaction. Transaction state is tracked just like a regular package would be, trying to deliver it again if it was not accepted etc.  
It does so by having a queue of transactions which need to be sent or have not yet been mined and keeps checking their status and doing the necessary actions (sending, increasing gas, waiting, ...) to make sure they get delivered and then can be removed from the queue.

Transactions delivered into the depot should always get mined if they are valid and the address has enough gas.
