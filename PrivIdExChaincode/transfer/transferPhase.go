package transfer

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	"fmt"
	"chaincode/PrivIdEx/PrivIdExChaincode/util"
)

/**
	This is the method that should be invoked by an identity provider, in order to transfer an identity asset of a user,
	which was agreed to be transferred to an identity asset consumer, during the handshake phase.
 */
func TransferAsset(stub shim.ChaincodeStubInterface, args []string, log *shim.ChaincodeLogger) (string, error) {
	//validate the arguments
	if len(args) != 1 {
		//err:= "Invalid number of arguments."
		return "", util.InvalidArgumentError{len(args), 1}
	}
	var transfRecord transferRecord

	if err := json.Unmarshal([]byte(args[0]), &transfRecord); err != nil {
		return "", util.JSONDecodingError{args[0]}
	}

	//TODO: validate the ids and signatures on the message.

	transactionKeyForConfHandshake := util.CreateTransactionKey(util.CONF_HANDSHAKE, transfRecord.TransactionID,
		transfRecord.ConsumerID, transfRecord.UserID, transfRecord.ProviderID)

	//check if there is a confHandshake record with the same key, and if not, throw an error.
	result, err := stub.GetState(transactionKeyForConfHandshake)
	if err != nil {
		return "", fmt.Errorf("Error in checking for an confHanshake record for the key: %s.", transactionKeyForConfHandshake)
	}
	if result == nil {
		return "", fmt.Errorf("A confHandshake record for the transaction key: %s does not exist.", transactionKeyForConfHandshake)
	}

	transactionKeyForTransferAsset := util.CreateTransactionKey(util.TRANSFER_ASSET, transfRecord.TransactionID,
		transfRecord.ConsumerID, transfRecord.UserID, transfRecord.ProviderID)

	//TODO: Although log level is set to Debug, it is not recognized and set to INFO by default. Therefore, making this INFO.
	log.Infof("Transaction key for transferAsset record: %s", transactionKeyForTransferAsset)

	//check if there is transfer asset record with the same key, and if so, throw an error.
	result2, err2 := stub.GetState(transactionKeyForTransferAsset)
	if err2 != nil {
		return "", fmt.Errorf("Error in checking for an existing transferAsset record for the key: %s.", transactionKeyForTransferAsset)
	}
	if result2 != nil {
		return "", fmt.Errorf("A transferAsset record for the transaction key: %s is already existing.", transactionKeyForTransferAsset)
	}

	if err3 := stub.PutState(transactionKeyForTransferAsset, []byte(args[0])); err3 != nil {
		return "", fmt.Errorf("Failed to submit the transferAsset record to the blockchain for processing: %s", args[0])
	}
	resp := fmt.Sprintf("TransferAsset record was submitted to the blockchain for processing.")
	return resp, nil
}