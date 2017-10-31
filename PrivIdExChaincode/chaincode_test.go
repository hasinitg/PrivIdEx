package main

import (
	"testing"
	"encoding/json"
	//"github.com/hyperledger/fabric/core/chaincode/shim"
	//"crypto/rsa"
	//"crypto/rand"
	"chaincode/PrivIdEx/PrivIdExChaincode/handshake"
	"github.com/segmentio/ksuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"chaincode/PrivIdEx/PrivIdExChaincode/util"
)

func TestEncodingDecodingHandshakeRecord(t *testing.T){
	transactionID := ksuid.New().String()
	//create a dummy handshake message
	initHandshakeMessage := createInitHandshakeRecord(transactionID)

	encodedMsg, err := json.Marshal(initHandshakeMessage)

	if err!=nil{
		fmt.Println("Handshake record encoding failed..")
		t.FailNow()

	} else {
		encodedMsgStr := string(encodedMsg)
		fmt.Println("Encoded message: ", string(encodedMsg))

		var handshakeRecord handshake.HandshakeRecord

		json.Unmarshal([]byte(encodedMsgStr), &handshakeRecord)
		//fmt.Println(handshakeRecord.ConsumerID)

		consumerID := handshakeRecord.ConsumerID
		if "c1" != consumerID {
			fmt.Println("ConsumerID : %v", consumerID, "was not as the expected value which is: %v.", initHandshakeMessage.ConsumerID )
		}
	}
}

/*This is the test case that simulates the different steps of the protocol implemented in prividex chaincode.*/
func TestPrivIdExChaincodeForEncodingDecoding(t *testing.T){
	//create a transaction id (the reason why we create the transaction id by the caller, rather than generating it in
	// the chaincode is that: chaincode is run in multiple nodes and difff transaction ids could be generated)
	actualTransactionID := ksuid.New().String()

	//create a mock stub by passing identity asset
	idAsset := new(IdentityAsset)

	stub := shim.NewMockStub("testInitHSEncDec", idAsset)

	//initialize the chaincode
	checkInit(t, stub, [][]byte{[]byte("init")})

	//test legitimate inithandshake step:
	testInitHandshake(t, stub, actualTransactionID)

	//test illegitimate inithandshake step
	responseStatus := testInitHandshake(t, stub, actualTransactionID)
	if responseStatus == shim.OK {
		fmt.Println("Invalid initHandshake transaction was allowed.")
		t.FailNow()
	}

}

func testInitHandshake(t *testing.T, stub *shim.MockStub, transactionID string) (int32){
	//create a dummy handshake message
	initHandshakeMessage := createInitHandshakeRecord(transactionID)

	encodedMsg, err := json.Marshal(initHandshakeMessage)
	if err!=nil{
		fmt.Println("Handshake record encoding failed..")
		t.FailNow()

	} else {
		//since any write to the ledger needs to be in a transactional context, the test must start the transaction before
		// invoking and end after invoking.
		stub.MockTransactionStart("t1")
		responseStatus, responseInitHandshake := checkInvoke(t, stub, [][]byte{[]byte("initHandshake"), []byte(encodedMsg)})
		fmt.Println("Response from initHanshake invoke : ", responseInitHandshake)
		stub.MockTransactionEnd("t1")

		if responseStatus != shim.OK {
			return responseStatus
		}

		//check if the message is in the ledger.
		checkQuery(t, stub, util.CreateTransactionKey(util.INIT_HANDSHAKE, initHandshakeMessage.TransactionID,
			initHandshakeMessage.ConsumerID, initHandshakeMessage.UserID, initHandshakeMessage.ProviderID), string(encodedMsg))
	}
	return shim.OK
}


func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) (string){
	response := stub.MockInit("1", args)
	if response.Status != shim.OK {
		fmt.Println("Chaincode Init failed.", string(response.Message))
		t.FailNow()
	}
	return response.Message
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) (int32, string){
	response := stub.MockInvoke("1", args)
	if response.Status != shim.OK {
		fmt.Println("Invoke with ", string(args[0]), " failed.", string(response.Message))
		return response.Status, ""
	}
	return response.Status, string(response.Payload)
}

func checkQuery(t *testing.T, stub *shim.MockStub, key string, expectedValue string) {
	response , err := stub.GetState(key)

	if err!=nil{
		fmt.Println("Query failed for the key %v", key)
		t.FailNow()
	}

	if response == nil{
		fmt.Println("Queried value in null.")
		t.FailNow()
	}

	actualValue := string(response)

	if expectedValue != actualValue {
		fmt.Println("Queried value: ", actualValue, "does not match the expected value: ", expectedValue)
		t.FailNow()
	}
}

func createInitHandshakeRecord(transactionID string) (handshake.HandshakeRecord){

	//create a dummy handshake message for now. TODO: add actual crypto information.
	initHandshakeMessage := handshake.HandshakeRecord{ util.INIT_HANDSHAKE, transactionID,
		"c1", "c_PK", "u1","u_PK", "p1", "p_PK",
		"kyc_compliance", "s1", "s2"}
	return initHandshakeMessage
}

func createRespHandshakeRecord(transactionID string) (handshake.HandshakeRecord){

	//create a dummy handshake message for now. TODO: add actual crypto information.
	respHandshakeMessage := handshake.HandshakeRecord{ util.RESP_HANDSHAKE, transactionID,
		"c1", "c_PK", "u1","u_PK", "p1", "p_PK",
		"kyc_compliance", "s1", "s2"}
	return respHandshakeMessage
}