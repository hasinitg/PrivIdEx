package confirm

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	"fmt"
	"chaincode/PrivIdEx/PrivIdExChaincode/util"
)

/**
	This is the method that should be invoked by an identity asset consumer, in order to confirm the receipt of an identity
	asset of the user, transferred by an identity asset provider, as per the request made by the identity asset consumer
	and agreed upon to be transferred by the identity provider, during the handshake phase.
 */
func ConfirmReceiptOfAsset(stub shim.ChaincodeStubInterface, args []string, log *shim.ChaincodeLogger) (string, error) {
	//validate the arguments
	if len(args) != 1 {
		//err:= "Invalid number of arguments."
		return "", util.InvalidArgumentError{len(args), 1}
	}
	var confRecord ConfirmRecord

	if err := json.Unmarshal([]byte(args[0]), &confRecord); err != nil {
		return "", util.JSONDecodingError{args[0]}
	}

	//set the record type
	confRecord.RecordType = util.CONFIRM_RECEIPT_OF_ASSET

	//TODO: validate the ids and signatures on the message.

	transactionKeyForTransferAsset := util.CreateTransactionKey(util.TRANSFER_ASSET, confRecord.TransactionID,
		confRecord.ConsumerID, confRecord.UserID, confRecord.ProviderID)

	//check if there is a confHandshake record with the same key, and if not, throw an error.
	result, err := stub.GetState(transactionKeyForTransferAsset)
	if err != nil {
		return "", fmt.Errorf("Error in checking for an transferAsset record for the key: %s.", transactionKeyForTransferAsset)
	}
	if result == nil {
		return "", fmt.Errorf("A transferAsset record for the transaction key: %s does not exist.", transactionKeyForTransferAsset)
	}

	transactionKeyForConfirmation := util.CreateTransactionKey(confRecord.RecordType, confRecord.TransactionID,
		confRecord.ConsumerID, confRecord.UserID, confRecord.ProviderID)

	//TODO: Although log level is set to Debug, it is not recognized and set to INFO by default. Therefore, making this INFO.
	//log.Infof("Transaction key for confirmReceiptOfAsset: %s", transactionKeyForConfirmation)

	//check if there is transfer asset record with the same key, and if so, throw an error.
	result2, err2 := stub.GetState(transactionKeyForConfirmation)
	if err2 != nil {
		return "", fmt.Errorf("Error in checking for an existing confirmReceiptOfAsset record for the key: %s.", transactionKeyForConfirmation)
	}
	if result2 != nil {
		return "", fmt.Errorf("A confirmReceiptOfAsset record for the transaction key: %s is already existing.", transactionKeyForConfirmation)
	}

	if err3 := stub.PutState(transactionKeyForConfirmation, []byte(args[0])); err3 != nil {
		return "", fmt.Errorf("Failed to submit the confirmReceiptOfAsset record to the blockchain for processing: %s", args[0])
	}
	resp := fmt.Sprintf("ConfirmReceiptOfAsset record was submitted to the blockchain for processing.")
	return resp, nil
}
