package handshake

import(
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"chaincode/PrivIdEx/PrivIdExChaincode/util"
	"encoding/json"
	"fmt"
	"bytes"
)

func InitHandshake(stub shim.ChaincodeStubInterface, args []string, log *shim.ChaincodeLogger) (string, error){
	//validate the arguments
	if len(args) != 1{
		//err:= "Invalid number of arguments."
		return "", util.InvalidArgumentError{len(args), 1}
	}
	var initHandshakeRecord HandshakeRecord

	if err:= json.Unmarshal([]byte(args[0]), &initHandshakeRecord); err!=nil{
		return "",util.JSONDecodingError{args[0]}
	}

	//TODO: validate the ids and signatures on the message.

	//TODO: create composite key combining the participants' ids. This will be useful in validation checks of the
	//subsequent methods, rather than only using the uuid.
	transactionKey := CreateTransactionKey(initHandshakeRecord.TransactionID, initHandshakeRecord.ConsumerID, initHandshakeRecord.UserID,
		initHandshakeRecord.ProviderID)

	//TODO: Although log level is set to Debug, it is not recognized and set to INFO by default. Therefore, making this INFO.
	log.Infof("Transaction key for initHandshake message: %s", transactionKey)

	if err := stub.PutState(transactionKey, []byte(args[0])); err!=nil {
		return "", fmt.Errorf("Failed to submit the initHandshake message to the blockchain for processing: %s", args[0])
	}
	resp := fmt.Sprintf("InitHandshake message was submitted to the blockchain for processing.")
	return resp, nil
}

func CreateTransactionKey(transactionID, consumerID, userID, providerID string) (string){
	var buffer bytes.Buffer
	buffer.WriteString(transactionID)
	buffer.WriteString(":")
	buffer.WriteString(consumerID)
	buffer.WriteString(":")
	buffer.WriteString(userID)
	buffer.WriteString(":")
	buffer.WriteString(providerID)
	return buffer.String()
}