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
)

func TestEncodingDecodingHandshakeRecord(t *testing.T){

	//create a dummy handshake message
	initHandshakeMessage := createInitHandshakeRecord()

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

//This is supposed to test the methods of the chaincode in handshake phase. However, it is not completed yet.
func TestInitHandshake_with_EncodingDecoding(t *testing.T) {
	//create a dummy handshake message
	initHandshakeMessage := createInitHandshakeRecord()

	encodedMsg, err := json.Marshal(initHandshakeMessage)
	if err!=nil{
		fmt.Println("Handshake record encoding failed..")
		t.FailNow()

	} else {
		//create a mock stub by passing identity asset
		idAsset := new(IdentityAsset)

		stub := shim.NewMockStub("testInitHSEncDec", idAsset)

		//initialize the chaincode (TODO: move this and mock stub creation to a outer bigger test case which sequences all the test cases.)
		checkInit(t, stub, [][]byte{[]byte("init")})

		//since any write to the ledger needs to be in a transactional context, the test must start the transaction before
		// invoking and end after invoking.
		stub.MockTransactionStart("t1")
		responseInitHandshake := checkInvoke(t, stub, [][]byte{[]byte("initHandshake"), []byte(encodedMsg)})
		fmt.Println("Response from initHanshake invoke : ", responseInitHandshake)
		stub.MockTransactionEnd("t1")

		//check if the message is in the ledger.
		checkQuery(t, stub, handshake.CreateTransactionKey(initHandshakeMessage.TransactionID,
			initHandshakeMessage.ConsumerID, initHandshakeMessage.UserID, initHandshakeMessage.ProviderID), string(encodedMsg))

	}
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) (string){
	response := stub.MockInit("1", args)
	if response.Status != shim.OK {
		fmt.Println("Chaincode Init failed.", string(response.Message))
		t.FailNow()
	}
	return response.Message
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) (string){
	response := stub.MockInvoke("1", args)
	if response.Status != shim.OK {
		fmt.Println("Invoke with ", string(args[0]), " failed.", string(response.Message))
	}
	return string(response.Payload)
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

func createInitHandshakeRecord() (handshake.HandshakeRecord){
	//create a transaction id (the reason why we create the transaction id by the caller, rather than generating it in
	// the chaincode is that: chaincode is run in multiple nodes and difff transaction ids could be generated)
	transactionID := ksuid.New().String()
	//fmt.Println(transactionID)

	//create a dummy handshake message for now. TODO: add actual crypto information.
	initHandshakeMessage := handshake.HandshakeRecord{transactionID, "c1", "c_PK", "u1",
		"u_PK", "p1", "p_PK", "kyc_compliance", "s1", "s2"}
	return initHandshakeMessage
}