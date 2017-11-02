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
4. Open a new terminal and run the block-event-listener as described in the README file at: https://github.com/hasinitg/PrivIdEx/tree/master/block-event-listener in order to get notifications on the blocks published to the blockchain.

5. **Simulating the protocol:**

i. *InitHandshake Phase:*
Execute: 'peer chaincode invoke -n prividex -c '{"Args":["initHandshake", "{\\"HandshakeRecordType\\":\\"initHandshake\\",\\"TransactionID\\":\\"0ttl5HdQCG53TR4T6ANBQHVMvcq\\",\\"ConsumerID\\":\\"c1\\",\\"ConsumerPublicKey\\":\\"c_PK\\",\\"UserID\\":\\"u1\\",\\"UserPublicKey\\":\\"u_PK\\",\\"ProviderID\\":\\"p1\\",\\"ProviderPublicKey\\":\\"p_PK\\",\\"IdentityAssetName\\":\\"kyc_compliance\\",\\"Signature1\\":\\"s1\\",\\"Signature2\\":\\"s2\\"}"]}' -C mychannel'. 

*RespHandshake Phase*

ii. Execute: 'peer chaincode invoke -n prividex -c '{"Args":["respHandshake", "{\\"HandshakeRecordType\\":\\"respHandshake\\",\\"TransactionID\\":\\"0ttl5HdQCG53TR4T6ANBQHVMvcq\\",\\"ConsumerID\\":\\"c1\\",\\"ConsumerPublicKey\\":\\"c_PK\\",\\"UserID\\":\\"u1\\",\\"UserPublicKey\\":\\"u_PK\\",\\"ProviderID\\":\\"p1\\",\\"ProviderPublicKey\\":\\"p_PK\\",\\"IdentityAssetName\\":\\"kyc_compliance\\",\\"Signature1\\":\\"s1\\",\\"Signature2\\":\\"\\"}"]}' -C mychannel'

*ConfHandshake Phase*

iii. Execute: 'peer chaincode invoke -n prividex -c '{"Args":["confHandshake", "{\\"HandshakeRecordType\\":\\"confHandshake\\",\\"TransactionID\\":\\"0ttl5HdQCG53TR4T6ANBQHVMvcq\\",\\"ConsumerID\\":\\"c1\\",\\"ConsumerPublicKey\\":\\"c_PK\\",\\"UserID\\":\\"u1\\",\\"UserPublicKey\\":\\"u_PK\\",\\"ProviderID\\":\\"p1\\",\\"ProviderPublicKey\\":\\"p_PK\\",\\"IdentityAssetName\\":\\"kyc_compliance\\",\\"Signature1\\":\\"s1\\",\\"Signature2\\":\\"\\"}"]}' -C mychannel'

*TransferAsset Phase*

iv. Execute: 'peer chaincode invoke -n prividex -c '{"Args":["transferAsset", "{\\"TransactionID\\":\\"0ttl5HdQCG53TR4T6ANBQHVMvcq\\",\\"ConsumerID\\":\\"c1\\",\\"ConsumerPublicKey\\":\\"c_PK\\",\\"UserID\\":\\"u1\\",\\"UserPublicKey\\":\\"u_PK\\",\\"ProviderID\\":\\"p1\\",\\"ProviderPublicKey\\":\\"p_PK\\",\\"IdentityAssetName\\":\\"kyc_compliance\\",\\"IdAsset\\":\\"ewogICJpZEFzc2V0IiA6IHsKICAgICJuYW1lT2ZJZEFzc2V0IiA6ICJLWUNfZm9yX0JhbmtpbmciLAogICAgIm5hbWVPZlVzZXIiICAgIDogIkNoZXJyeSBCZXJyeSIsCiAgICAicGFzc3BvcnQiIDogewogICAgICAibnVtYmVyIiA6ICJONTg3ODY1IiwKICAgICAgImV4cGlyYXRpb24tZGF0ZSIgOiAiMDYtMDctMjAyMiIsCiAgICAgICJpc3N1aW5nLWNvdW50cnkiIDogIlVESyIKICAgIH0sCiAgICAiU1NOIiA6IHsKICAgICAgIm51bWJlciIgOiAiNjc4MC05NS0zMjQ1IiwKICAgICAgImV4cGlyYXRpb24tZGF0ZSIgOiAiMjAtMDctMjAxMCIKICAgIH0KCiAgfQp9\\",\\"Signature1\\":\\"s1\\"}"]}' -C mychannel'

6. You can shutdown the network by changing the to the directory:**fabric-network/basic-network** and executing the command: **./stop.sh**. If you want to make changes to the chaincode and redeploy the chaincode in a fresh blockchain network, execute the command: **./teardown.sh**.
