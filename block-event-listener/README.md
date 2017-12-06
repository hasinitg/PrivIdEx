# Registering a block event listener for the chaincode:

Here we list down the steps for registering a block event listener for our chaincode, s.t whenever a block is posted on the ledger by our chaincode, this listener will get notified and it will print the block on the command line.
The code in **PrivIdEx/block-event-listener/block-listener.go** is taken from the hyperledger fabric example at: https://github.com/hyperledger/fabric/tree/master/examples/events/block-listener

# Steps:
1. If you are in *nix system, change the etc/hosts file in order to map: 127.0.0.1 to the peer address you are listening to, - in our case, it is: peer0.org1.example.com.
2. Then start the standard fabric network, by following the steps 1-2, found in: PrivIdEx/fabric-network/README.md file.
3. Then register the block-event-listener by changing the directory to: PrivIdEx/block-event-listener and:

      i. executing the command: *'go build'*
      
      ii. executing the command: *'./block-listener -events-address=peer0.org1.example.com:7053 -events-from-chaincode=prividex  -events-mspdir=$GOPATH/src/chaincode/PrivIdEx/fabric-network/basic-network/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp  -events-mspid=Org1MSP'*
4. Then follow the remaining steps: 3-4 in PrivIdEx/fabric-network/README.md file, in order to invoke the chaincode.
5. Observe the terminal where you started the block-event-listener in step 3 above, you will be able to see the block corresponding to the chaincode invocation is printed in the terminal.
