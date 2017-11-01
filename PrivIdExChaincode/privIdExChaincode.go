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
	"chaincode/PrivIdEx/PrivIdExChaincode/transfer"
	"chaincode/PrivIdEx/PrivIdExChaincode/confirm"
)

/**This is the entry point to the chaincode that implements the privacy preserving identity asset exchange.**/

//Identity Asset is the type on which the methods in this chaincode are implemented.
type IdentityAsset struct {
}

var logName string = "PrivIdEx_CC_Log"
var log = shim.NewLogger(logName)

//Init is called during instantiation and upgrade of the chaincode.
func (idAsset *IdentityAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	//For the moment do nothing during instantiation.
	log.SetLevel(shim.LogDebug)
	logMessage := "PrivIdEx chaincode instantiated successfully."
	log.Info(logMessage)
	//addLogMessage(logName, logMessage, shim.LogDebug)
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
	case util.INIT_HANDSHAKE:
		log.Info("InitHandshake function was invoked.")
		result, err = handshake.InitHandshake(stub, args, log)
	case util.RESP_HANDSHAKE:
		log.Info("RespHandshake function was invoked.")
		result, err = handshake.RespHandshake(stub, args, log)
	case util.CONF_HANDSHAKE:
		log.Info("ConfHandshake function was invoked.")
		result, err = handshake.ConfHandshake(stub, args, log)
	case util.TRANSFER_ASSET:
		log.Info("TransferAsset function was invoked.")
		result, err = transfer.TransferAsset(stub, args, log)
	case util.CONFIRM_RECEIPT_OF_ASSET:
		log.Info("ConfirmReceiptOfAsset function was invoked.")
		result, err = confirm.ConfirmReceiptOfAsset(stub, args, log)
	case util.QUERY:
		log.Info("Query function was invoked.")
		result, err = query(stub, args)
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

func query(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		//err:= "Invalid number of arguments."
		return "", util.InvalidArgumentError{len(args), 1}
	}
	resultByte, err := stub.GetState(args[0])
	if err != nil {
		return "", err
	}
	respString := string(resultByte)
	//TODO: Although log level is set to Debug, it is not recognized and set to INFO by default. Therefore, making this INFO.
	log.Infof("Queried result: %s", respString)
	return respString, nil
}

// main function starts up the chaincode in the container during instantiate
func main() {

	log.SetLevel(shim.LogDebug)

	if err := shim.Start(new(IdentityAsset)); err != nil {
		fmt.Printf("Error starting identity asset chaincode: %s", err)
	}
}

//func addLogMessage(logName string, logMessage string, logType shim.LoggingLevel) {
//	//var log = shim.NewLogger(logName)
//	log.Info(logMessage)
//}
