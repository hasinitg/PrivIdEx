/**
 * Author: huralali@purdue.edu
 * License : Apache-2.0
 */
package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"

	"chaincode/PrivIdEx/PrivIdExChaincode/handshake"
	"chaincode/PrivIdEx/PrivIdExChaincode/util"
)

/**This is the entry point to the chaincode that implements the privacy preserving identity asset exchange.**/

//Identity Asset is the type on which the methods in this chaincode are implemented.
type IdentityAsset struct {
}

var logName string = "PrivIdEx_CC_Log"

//Init is called during instantiation and upgrade of the chaincode.
func (idAsset *IdentityAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	//For the moment do nothing during instantiation.

	logMessage := "PrivIdEx chaincode instantiated successfully."
	addLogMessage(logName, logMessage, shim.LogDebug)
	return shim.Success([]byte(logMessage));
}

// Invoke is called per transaction on the chaincode. Each transaction is
// an operation of either handshake, transfer or confirmation phase of the identity asset
// exchange protocol.
func (idAsset *IdentityAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and arguments from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	switch fn {
	case "initHandshake":
		result, err = handshake.InitHandshake(stub, args)
	case "respHandshake":
		result, err = handshake.InitHandshake(stub, args)
	case "confirmHandshake":
		result, err = handshake.InitHandshake(stub, args)
	case "transferAsset":
		result, err = handshake.InitHandshake(stub, args)
	case "confirmReceiptOfAsset":
		result, err = handshake.InitHandshake(stub, args)
	default:
		result, err = "", util.PrivIdExUnknownMethodError{fn}
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	//fmt.Println("Result: ", result)
	return shim.Success([]byte(result))
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(IdentityAsset)); err != nil {
		fmt.Printf("Error starting identity asset chaincode: %s", err)
	}
}

func addLogMessage(logName string, logMessage string, logType shim.LoggingLevel) {
	var log = shim.NewLogger(logName)
	log.SetLevel(logType)
	log.Info(logMessage)
}
