# Running the chaincode in a standard hyperledger fabric network

In our main Readme file at https://github.com/hasinitg/PrivIdEx/blob/master/README.md, we show how to install, instantiate
and invoke our chaincode in a dev-mode hyper ledger fabric network.

Here we mention the steps to run the chaincode in a standard hyperledger fabric network involving MSP and other key nodes as well.

The contents in the  ***fabric-network*** folder is a tweaked copy of the blockchain artifacts used for the fabcar sample
described in http://hyperledger-fabric.readthedocs.io/en/latest/write_first_app.html.

## Steps:

1. Change directory to fabric-network/prividex.
2. Execute command: './startFabric.sh' which will start the fabric network defined in fabric-network/basic-network/docker-compose.yaml file,
with the additional commands specified in startFabric.sh file to install and instantiate the chaincode with the name: **prividex**,
attached to the channel named: **mychannel**.
3. Open a new terminal and enter the cli node by running the command: **docker exec -it cli bash**. 
4. Execute: 'peer chaincode invoke -n prividex -c '{"Args":["initHandshake", "{\"TransactionID\":\"0ttl5HdQCG53TR4T6ANBQHVMvcq\",\"ConsumerID\":\"c1\",\"ConsumerPublicKey\":\"c_PK\",\"UserID\":\"u1\",\"UserPublicKey\":\"u_PK\",\"ProviderID\":\"p1\",\"ProviderPublicKey\":\"p_PK\",\"IdentityAssetName\":\"kyc_compliance\",\"Signature1\":\"s1\",\"Signature2\":\"s2\"}"]}' -C mychannel' in order to invoke the 'initHandshake' method of the chaincode, with the given json input. You will receive a message that the transaction was submitted to blockchain network successfully and also can see in Terminal 1 that the handshake message is added to the ledger. You can see the log messages printed by our chaincode in Terminal 2.
5. Execute: 'peer chaincode invoke -n prividex -c '{"Args":["query","0ttl5HdQCG53TR4T6ANBQHVMvcq:c1:u1:p1"]}' -C mychannel' in order to query the posted initHandshake message by the transaction key. You will see the posted message as the query result printed on the terminal.
6. You can shutdown the network by changing the to the directory:**fabric-network/basic-network** and executing the command: **./stop.sh**.
