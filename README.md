# PrivIdEx
Code for PrivIdEx project

**PrivIdExChaincode**   folder contains the smart contract (chaincode) implementing the **PrivIdEx protocol** in Go language.

**devmode-fabric-network** folder contains the artifacts to setup a simple blockchain network backed by hyperledger fabric in dev mode. These artifacts are slightly modified version of the artifacts provided in fabric-samples/chaincode-docker-devmode in https://github.com/hyperledger/fabric-samples.git

Following are the instructions for downloading, building and deploying the chaincode into a simple blockchain network and invoking the chaincode.

# Downloading the code:

1. First install Go in your local machine following the instructions at: https://golang.org/dl/
2. Set the GOPATH environment variable to point to your Go working directory. See: https://golang.org/doc/code.html#GOPATH for more details.
3. Create a directory named 'chaincode' inside the 'src' directory of your GOPATH.
4. Clone the PrivIdEx repository inside the 'chaincode' directory, by executing: 'git clone https://github.com/hasinitg/PrivIdEx.git'

# Building the code:

1. Import an external package used by the code, by executing: 'go get github.com/twinj/uuid'
2. Change directory into the PrivIdEx/PrivIdExChaincode
3. Execute: 'go build'

# Setting up the blockchain network

1. Hyperledger fabric blockchain network is consisted of multiple entities running in docker containers. Make sure you have installed docker in your machine and downloaded the latest hyperledger fabric docker images. 
To download the latest yperledger fabric docker images, execute this curl command: curl -sSL https://goo.gl/Gci9ZX | bash
(source: http://hyperledger-fabric.readthedocs.io/en/latest/samples.html)
