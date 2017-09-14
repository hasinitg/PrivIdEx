package main

import (
	"testing"
	//"PrivIdEx/PrivIdExChaincode/handshake"
	"encoding/json"
	//"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	//"PrivIdEx/PrivIdExChaincode/handshake"
)

func Test(t *testing.T) {
	//create a dummy handshake message
	initHandshakeMessage := HandshakeRecord{"tr1", "c1", "u1", "p1",
		"kyc_compliance", "s1"}

	encodedMsg, err := json.Marshal(initHandshakeMessage)

	//test sending an invalid message signature
	//handshake.InitHandshake(shim.ChaincodeStub{})

	if err==nil{
		fmt.Println(string(encodedMsg))
	}
}