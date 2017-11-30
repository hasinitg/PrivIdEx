package handshake

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"chaincode/PrivIdEx/PrivIdExChaincode/util"
	"encoding/json"
	"fmt"
	"bytes"
	"chaincode/PrivIdEx/PrivIdExChaincode/discovery"
)

/**
	This is the method that should be invoked by an identity consumer, for requesting an identity asset of a user,
	from an identity provider.
 */
func InitHandshake(stub shim.ChaincodeStubInterface, args []string, log *shim.ChaincodeLogger) (string, error) {
	//validate the arguments
	if len(args) != 1 {
		//err:= "Invalid number of arguments."
		return "", util.InvalidArgumentError{len(args), 1}
	}
	var initHandshakeRecord HandshakeRecord

	if err := json.Unmarshal([]byte(args[0]), &initHandshakeRecord); err != nil {
		return "", util.JSONDecodingError{args[0]}
	}
	//set the handshake record type
	initHandshakeRecord.HandshakeRecordType = util.INIT_HANDSHAKE

	//TODO: validate the ids and signatures on the message.

	transactionKey := util.CreateTransactionKey(initHandshakeRecord.HandshakeRecordType, initHandshakeRecord.TransactionID,
		string(initHandshakeRecord.ConsumerID), string(initHandshakeRecord.UserID), string(initHandshakeRecord.ProviderID))

	//TODO: Although log level is set to Debug, it is not recognized and set to INFO by default. Therefore, making this INFO.
	//log.Infof("Transaction key for initHandshake message: %s", transactionKey)

	//check if there is another transaction with the same key, and if so, throw an error.
	resultByte, err := stub.GetState(transactionKey)
	if err != nil {
		return "", fmt.Errorf("Error in checking for an existing record for the key: %s.", transactionKey)
	}
	if resultByte != nil {
		return "", fmt.Errorf("An initHandshake record for the transaction key: %s is already existing.", transactionKey)
	}

	if err := stub.PutState(transactionKey, []byte(args[0])); err != nil {
		return "", fmt.Errorf("Failed to submit the initHandshake message to the blockchain for processing: %s", args[0])
	}
	resp := fmt.Sprintf("InitHandshake message was submitted to the blockchain for processing.")
	return resp, nil
}

/**
	This is the method that should be invoked by an identity provider, who got a request from an identity consumer, requesting
	an identity asset of a user.
 */
func RespHandshake(stub shim.ChaincodeStubInterface, args []string, log *shim.ChaincodeLogger) (string, error){
	//validate the arguments
	if len(args) != 1 {
		return "", util.InvalidArgumentError{len(args), 1}
	}
	var respHandshakeRecord HandshakeRecord

	if err := json.Unmarshal([]byte(args[0]), &respHandshakeRecord); err != nil {
		return "", util.JSONDecodingError{args[0]}
	}

	//set the handshake record type
	respHandshakeRecord.HandshakeRecordType = util.RESP_HANDSHAKE

	//TODO: validate the ids and signatures on the message.

	transactionKeyForInitHandshake := util.CreateTransactionKey(util.INIT_HANDSHAKE, respHandshakeRecord.TransactionID,
		string(respHandshakeRecord.ConsumerID), string(respHandshakeRecord.UserID), string(respHandshakeRecord.ProviderID))

	//check if there is initHandshake record with the same key, and if not, throw an error.
	result, err := stub.GetState(transactionKeyForInitHandshake)
	if err != nil {
		return "", fmt.Errorf("Error in checking for an initHanshake record for the key: %s.", transactionKeyForInitHandshake)
	}
	if result == nil {
		return "", fmt.Errorf("An initHandshake record for the transaction key: %s does not exist.", transactionKeyForInitHandshake)
	}

	transactionKeyForRespHandshake := util.CreateTransactionKey(util.RESP_HANDSHAKE, respHandshakeRecord.TransactionID,
		string(respHandshakeRecord.ConsumerID), string(respHandshakeRecord.UserID), string(respHandshakeRecord.ProviderID))

	//TODO: Although log level is set to Debug, it is not recognized and set to INFO by default. Therefore, making this INFO.
	//log.Infof("Transaction key for respHandshake message: %s", transactionKeyForRespHandshake)

	//check if there is respHandshake record with the same key, and if so, throw an error.
	result2, err2 := stub.GetState(transactionKeyForRespHandshake)
	if err2 != nil {
		return "", fmt.Errorf("Error in checking for an existing respHandshake record for the key: %s.", transactionKeyForRespHandshake)
	}
	if result2 != nil {
		return "", fmt.Errorf("A respHandshake record for the transaction key: %s is already existing.", transactionKeyForRespHandshake)
	}

	if err3 := stub.PutState(transactionKeyForRespHandshake, []byte(args[0])); err3 != nil {
		return "", fmt.Errorf("Failed to submit the respHandshake message to the blockchain for processing: %s", args[0])
	}
	resp := fmt.Sprintf("RespHandshake message was submitted to the blockchain for processing.")
	return resp, nil
}

/**
	This is the method that should be invoked by an identity consumer, who initiated an handshake message for requesting an identity asset of a user
	from an identity consumer and who got a response handshake message from that identity provider.
 */
func ConfHandshake(stub shim.ChaincodeStubInterface, args []string, log *shim.ChaincodeLogger) (string, error){
	//validate the arguments
	if len(args) != 1 {
		return "", util.InvalidArgumentError{len(args), 1}
	}
	var confHandshakeRecord HandshakeRecord

	if err := json.Unmarshal([]byte(args[0]), &confHandshakeRecord); err != nil {
		return "", util.JSONDecodingError{args[0]}
	}

	//set the handshake record type
	confHandshakeRecord.HandshakeRecordType = util.CONF_HANDSHAKE

	//TODO: validate the ids and signatures on the message.

	transactionKeyForRespHandshake := util.CreateTransactionKey(util.RESP_HANDSHAKE, confHandshakeRecord.TransactionID,
		string(confHandshakeRecord.ConsumerID), string(confHandshakeRecord.UserID), string(confHandshakeRecord.ProviderID))

	//check if there is respHandshake record with the same key, and if not, throw an error.
	result, err := stub.GetState(transactionKeyForRespHandshake)
	if err != nil {
		return "", fmt.Errorf("Error in checking for an respHanshake record for the key: %s.", transactionKeyForRespHandshake)
	}
	if result == nil {
		return "", fmt.Errorf("A respHandshake record for the transaction key: %s does not exist.", transactionKeyForRespHandshake)
	}

	transactionKeyForConfHandshake := util.CreateTransactionKey(util.CONF_HANDSHAKE, confHandshakeRecord.TransactionID,
		string(confHandshakeRecord.ConsumerID), string(confHandshakeRecord.UserID), string(confHandshakeRecord.ProviderID))

	//TODO: Although log level is set to Debug, it is not recognized and set to INFO by default. Therefore, making this INFO.
	//log.Infof("Transaction key for respHandshake message: %s", transactionKeyForConfHandshake)

	//check if there is confHandshake record with the same key, and if so, throw an error.
	result2, err2 := stub.GetState(transactionKeyForConfHandshake)
	if err2 != nil {
		return "", fmt.Errorf("Error in checking for an existing confHandshake record for the key: %s.", transactionKeyForConfHandshake)
	}
	if result2 != nil {
		return "", fmt.Errorf("A confHandshake record for the transaction key: %s is already existing.", transactionKeyForConfHandshake)
	}

	if err3 := stub.PutState(transactionKeyForConfHandshake, []byte(args[0])); err3 != nil {
		return "", fmt.Errorf("Failed to submit the confHandshake message to the blockchain for processing: %s", args[0])
	}
	resp := fmt.Sprintf("ConfHandshake message was submitted to the blockchain for processing.")
	return resp, nil
}

func CreateUserSignedMessageInInitHandshakeMessage(dresp discovery.DiscoveryResponse)(string){
	var buffer bytes.Buffer
	buffer.WriteString(string(dresp.ConsumerAnonymousID))
	buffer.WriteString(":")
	buffer.WriteString(string(dresp.ConsumerAnonymousPublicKey))
	buffer.WriteString(":")
	buffer.WriteString(dresp.IdentityAssetName)
	buffer.WriteString(":")
	buffer.WriteString(string(dresp.UserAnonymousID))
	buffer.WriteString(":")
	buffer.WriteString(string(dresp.ProviderAnonymousID))
	buffer.WriteString(":")
	buffer.WriteString(string(dresp.ProviderAnonymousPubKey))
	return buffer.String()

}


func CreateConsumerSignedMessageInInitHandshakeMessage(initHandshakeRecord HandshakeRecord) (string){
	dresp := discovery.DiscoveryResponse{[]byte (initHandshakeRecord.ConsumerID),
									[]byte (initHandshakeRecord.ConsumerPublicKey),
										initHandshakeRecord.IdentityAssetName,
										[]byte (initHandshakeRecord.UserID),
										[]byte (initHandshakeRecord.ProviderID),
										[]byte (initHandshakeRecord.ProviderPublicKey),
										nil}
	userMessage := CreateUserSignedMessageInInitHandshakeMessage(dresp)

	var buffer bytes.Buffer
	buffer.WriteString(initHandshakeRecord.TransactionID)
	buffer.WriteString(userMessage)
	buffer.WriteString(string(initHandshakeRecord.Signature1))
	return buffer.String()
}

