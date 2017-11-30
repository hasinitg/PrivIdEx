# PrivIdEx

**PrivIdExChaincode**   folder contains the smart contract (chaincode) implementing the **PrivIdEx protocol** in Go language.

**devmode-fabric-network** folder contains the artifacts to setup a simple blockchain network backed by hyperledger fabric in **dev mode**. These artifacts are slightly modified version of the artifacts provided in fabric-samples/chaincode-docker-devmode in https://github.com/hyperledger/fabric-samples.git

Following are the instructions for downloading, building and deploying the chaincode into a simple blockchain network and invoking the chaincode.

# Downloading the code:

1. First install Go in your local machine following the instructions at: https://golang.org/dl/ (make sure to install go version 1.8.3 or newer)

2. Set the GOPATH environment variable to point to your Go working directory. See: https://golang.org/doc/code.html#GOPATH for more details.

3. Create a directory named 'chaincode' inside the 'src' directory of your GOPATH.

4. Clone the PrivIdEx repository inside the 'chaincode' directory, by executing: 'git clone https://github.com/hasinitg/PrivIdEx.git'

# Obtaining the required dependencies:

1. From the 'src' directory of your GOPATH, execute:

i. go get github.com/hyperledger/fabric

ii. go get github.com/segmentio/ksuid

# Building the code:

1. Change directory into the GOPATH/src/chaincode/PrivIdEx/PrivIdExChaincode

2. Execute: 'go build'

3. To run the test cases, execute: 'go test' 

# Setting up the blockchain network

1. Hyperledger fabric blockchain network is consisted of multiple entities running in docker containers. Make sure you have installed docker in your machine and downloaded the latest hyperledger fabric docker images. 

To download the latest hyperledger fabric docker images, execute this curl command: curl -sSL https://goo.gl/Gci9ZX | bash
(source: http://hyperledger-fabric.readthedocs.io/en/latest/samples.html)

2. Change directory to PrivIdEx/devmode-fabric-network. The file named: docker-compose-simple.yaml file has the definition of the blockchain network. This network is supported by some pre-created artifacts such as channel artifacts to make the development and testing of chaincode easier.

3. We will execute the commands on the network in three different terminals, when deploying and testing our chaincode, as it is also done in this tutorial: http://hyperledger-fabric.readthedocs.io/en/latest/chaincode4ade.html

4. If you have existing docker containers running in your machine, which may conflict with the docker containers used in this tutorial, please stop and remove them (see: https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes for commands).

#### Terminal 1:
1. Execute: *'docker-compose -f docker-compose-simple.yaml up'* - this will start a blockchain network with orderer, peer, chaincode container and a cli container.

#### Terminal 2:
1. Execute: *'docker exec -it chaincode bash'*. This will enter you into the chaincode container. Since the script in: docker-compose-simple.yaml file has mapped your local working directory into the chaincode container's working directory, you can see the PrivIdEx folder inside the chaincode container's working directory.
2. Change directory into PrivIdEx/PrivIdExChaincode
3. Execute: 'go build' to build the chaincode inside the container.
3. Execute: *'CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=prividexcc:0 ./PrivIdExChaincode'* in order to run the chaincode with the chaincode id: prividexcc. At this point, the chaincode is only started and is not associated with any channel.

#### Terminal 3:
1. Execute : *'docker exec -it cli bash'* to enter into the cli container.
3. Execute: *'peer chaincode install -p chaincode/PrivIdEx/PrivIdExChaincode -n prividexcc -v 0'* in order to install the chaincode.
4. Execute: *'peer chaincode instantiate -n prividexcc -v 0 -c '{"Args":[]}' -C myc'* in order to instrantiate the chaincode and associate it with the channel named: 'myc'.
5. Execute: *'peer chaincode invoke -n prividexcc -c '{"Args":["initHandshake", "{\\"TransactionID\\":\\"0ttl5HdQCG53TR4T6ANBQHVMvcq\\",\\"ConsumerID\\":\\"c1\\",\\"ConsumerPublicKey\\":\\"c_PK\\",\\"UserID\\":\\"u1\\",\\"UserPublicKey\\":\\"u_PK\\",\\"ProviderID\\":\\"p1\\",\\"ProviderPublicKey\\":\\"p_PK\\",\\"IdentityAssetName\\":\\"kyc_compliance\\",\\"Signature1\\":\\"s1\\",\\"Signature2\\":\\"s2\\"}"]}' -C myc'* in order to invoke the 'initHandshake' method of the chaincode, with the given json input. 
You will receive a message that the transaction was submitted to blockchain network successfully and also can see in Terminal 1 that the handshake message is added to the ledger.
You can see the log messages printed by our chaincode in Terminal 2.
6. Execute: *'peer chaincode invoke -n prividexcc -c '{"Args":["query","0ttl5HdQCG53TR4T6ANBQHVMvcq:c1:u1:p1"]}' -C myc'* in order to query the posted initHandshake message by the transaction key. You will see the posted message as the query result printed on the terminal. 

### Stop the fabric network:
In a separate terminal, execute : 'docker-compose -f docker-compose-simple.yaml down' gracefully shutdown the blockchain network, during the chaincode development process.

## Next Steps:

To run the chaincode in a full hyperledger-fabric network and to simulate the complete protocol via CLI interface, please follow the instructions at: https://github.com/hasinitg/PrivIdEx/tree/master/fabric-network
