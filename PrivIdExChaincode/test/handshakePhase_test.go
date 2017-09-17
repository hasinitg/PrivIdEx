package test

import (
	"testing"
	"chaincode/PrivIdEx/PrivIdExChaincode/handshake"
	"encoding/json"
	//"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
)

//This is supposed to test the methods of the chaincode in handshake phase. However, it is not completed yet.
func Test(t *testing.T) {
	//create a dummy handshake message
	initHandshakeMessage := handshake.HandshakeRecord{"tr1", "c1", "u1", "p1",
		"kyc_compliance", "s1"}

	encodedMsg, err := json.Marshal(initHandshakeMessage)

	//test sending an invalid message signature
	//handshake.InitHandshake(shim.ChaincodeStub{})

	if err==nil{
		encodedMsgStr := string(encodedMsg)
		fmt.Println(string(encodedMsg))
		var handshakeRecord handshake.HandshakeRecord

		json.Unmarshal([]byte(encodedMsgStr), &handshakeRecord)
		fmt.Println(handshakeRecord.ConsumerID)
	}


}