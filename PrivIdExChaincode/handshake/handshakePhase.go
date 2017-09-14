package handshake

import(
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"chaincode/PrivIdEx/PrivIdExChaincode/util"
	"encoding/json"
	"github.com/twinj/uuid"
	"fmt"
)

func InitHandshake(stub shim.ChaincodeStubInterface, args []string) (string, error){
	//validate the arguments
	if len(args) != 1{
		//err:= "Invalid number of arguments."
		return "", util.InvalidArgumentError{len(args), 1}
	}
	var handshakeRecord HandshakeRecord

	if err:= json.Unmarshal([]byte(args[0]), &handshakeRecord); err!=nil{
		return "",util.JSONDecodingError{args[0]}
	}

	//generate a unique transaction id:
	uuidByte := uuid.NewV4()
	uuidString := string(uuidByte)

	//TODO: validate the signature on the message.
	//TODO: create composite key combining the participants' ids. This will be useful in validation checks of the
	//subsequent methods, rather than only using the uuid.

	if err := stub.PutState(uuidString, []byte(args[0])); err!=nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}

	return "Message was submitted to the blockchain for processing.", nil

}