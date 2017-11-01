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
	"chaincode/PrivIdEx/PrivIdExChaincode/transfer"
	"io/ioutil"
	"chaincode/PrivIdEx/PrivIdExChaincode/confirm"
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

	//test duplicate inithandshake step
	responseStatus := testInitHandshake(t, stub, actualTransactionID)
	if responseStatus == shim.OK {
		fmt.Println("Repeated initHandshake transaction was allowed.")
		t.FailNow()
	}

	//test legitimate resphandshake step:
	testRespHandshake(t, stub, actualTransactionID)

	//test duplicate resphandshake step
	responseStatus2 := testRespHandshake(t, stub, actualTransactionID)
	if responseStatus2 == shim.OK {
		fmt.Println("Repeated respHandshake transaction was allowed.")
		t.FailNow()
	}

	falseTransactionID := ksuid.New().String()
	//test invalid resphandshake step:
	responseStatus3 := testRespHandshake(t, stub, falseTransactionID)
	if responseStatus3 == shim.OK {
		fmt.Println("Invalid respHandshake transaction was allowed.")
		t.FailNow()
	}

	//test legitimate confhandshake step:
	testConfHandshake(t, stub, actualTransactionID)

	//test duplicate resphandshake step
	responseStatus4 := testConfHandshake(t, stub, actualTransactionID)
	if responseStatus4 == shim.OK {
		fmt.Println("Repeated confHandshake transaction was allowed.")
		t.FailNow()
	}

	falseTransactionID2 := ksuid.New().String()
	//test invalid resphandshake step:
	responseStatus5 := testConfHandshake(t, stub, falseTransactionID2)
	if responseStatus5 == shim.OK {
		fmt.Println("Invalid confHandshake transaction was allowed.")
		t.FailNow()
	}

	//test legitimate transferAsset step:
	testTransferIDAsset(t, stub, actualTransactionID)

	//test duplicate transferAsset step
	responseStatus6 := testTransferIDAsset(t, stub, actualTransactionID)
	if responseStatus6 == shim.OK {
		fmt.Println("Repeated transferAsset transaction was allowed.")
		t.FailNow()
	}

	falseTransactionID3 := ksuid.New().String()
	//test invalid resphandshake step:
	responseStatus7 := testTransferIDAsset(t, stub, falseTransactionID3)
	if responseStatus7 == shim.OK {
		fmt.Println("Invalid transferAsset transaction was allowed.")
		t.FailNow()
	}

	//***********query the transfer asset record and compare it with the sample id asset read from the file.************
	transferAssetKey := util.CreateTransactionKey(util.TRANSFER_ASSET, actualTransactionID, "c1", "u1", "p1")
	_, response := checkInvoke(t, stub, [][]byte{[]byte(util.QUERY), []byte(transferAssetKey)})
	fmt.Printf("Response for querying transfer asset: %s \n", response)

	var transfRecord transfer.TransferRecord

	if err := json.Unmarshal([]byte(response), &transfRecord); err != nil {
		fmt.Printf("Error in decoding the transferAsset record.")
		t.FailNow()
	}

	sampleIDAssetTransferred := transfRecord.IdAsset
	actualIDAsset, err := ioutil.ReadFile("test/sample_id_asset")

	if err != nil {
		fmt.Printf("Error in reading the sample id asset from file.")
		t.FailNow()
	}

	if string(actualIDAsset) != string(sampleIDAssetTransferred){
		fmt.Printf("The transferred id asset does not match the expected.")
		t.FailNow()
	}
	//******************************************************************************************************************

	//test legitimate confirmReceiptOfAsset step:
	testConfirmReceiptOfIDAsset(t, stub, actualTransactionID)

	//test duplicate confirmReceiptOfAsset step
	responseStatus8 := testConfirmReceiptOfIDAsset(t, stub, actualTransactionID)
	if responseStatus8 == shim.OK {
		fmt.Println("Repeated confirmReceiptOfAsset transaction was allowed.")
		t.FailNow()
	}

	falseTransactionID4 := ksuid.New().String()
	//test invalid resphandshake step:
	responseStatus9 := testConfirmReceiptOfIDAsset(t, stub, falseTransactionID4)
	if responseStatus9 == shim.OK {
		fmt.Println("Invalid confirmReceiptOfAsset transaction was allowed.")
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
		fmt.Println("Encoded initHandshake message: ", string(encodedMsg))

		//since any write to the ledger needs to be in a transactional context, the test must start the transaction before
		// invoking and end after invoking.
		stub.MockTransactionStart("t1")
		responseStatus, responseInitHandshake := checkInvoke(t, stub, [][]byte{[]byte(util.INIT_HANDSHAKE), []byte(encodedMsg)})
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

func testRespHandshake(t *testing.T, stub *shim.MockStub, transactionID string) (int32){
	//create a dummy handshake message
	respHandshakeMessage := createRespHandshakeRecord(transactionID)

	encodedMsg, err := json.Marshal(respHandshakeMessage)
	if err!=nil{
		fmt.Println("Handshake record encoding failed..")
		t.FailNow()

	} else {
		fmt.Println("Encoded respHandshake message: ", string(encodedMsg))

		//since any write to the ledger needs to be in a transactional context, the test must start the transaction before
		// invoking and end after invoking.
		stub.MockTransactionStart("t1")
		responseStatus, responseMessage := checkInvoke(t, stub, [][]byte{[]byte(util.RESP_HANDSHAKE), []byte(encodedMsg)})
		fmt.Println("Response from respHanshake invoke : ", responseMessage)
		stub.MockTransactionEnd("t1")

		if responseStatus != shim.OK {
			return responseStatus
		}

		//check if the message is in the ledger.
		checkQuery(t, stub, util.CreateTransactionKey(util.RESP_HANDSHAKE, respHandshakeMessage.TransactionID,
			respHandshakeMessage.ConsumerID, respHandshakeMessage.UserID, respHandshakeMessage.ProviderID), string(encodedMsg))
	}
	return shim.OK
}

func testConfHandshake(t *testing.T, stub *shim.MockStub, transactionID string) (int32){
	//create a dummy handshake message
	confHandshakeMessage := createConfHandshakeRecord(transactionID)

	encodedMsg, err := json.Marshal(confHandshakeMessage)
	if err!=nil{
		fmt.Println("Handshake record encoding failed..")
		t.FailNow()

	} else {
		fmt.Println("Encoded confHandshake message: ", string(encodedMsg))

		//since any write to the ledger needs to be in a transactional context, the test must start the transaction before
		// invoking and end after invoking.
		stub.MockTransactionStart("t1")
		responseStatus, responseMessage := checkInvoke(t, stub, [][]byte{[]byte(util.CONF_HANDSHAKE), []byte(encodedMsg)})
		fmt.Println("Response from confHanshake invoke : ", responseMessage)
		stub.MockTransactionEnd("t1")

		if responseStatus != shim.OK {
			return responseStatus
		}

		//check if the message is in the ledger.
		checkQuery(t, stub, util.CreateTransactionKey(util.CONF_HANDSHAKE, confHandshakeMessage.TransactionID,
			confHandshakeMessage.ConsumerID, confHandshakeMessage.UserID, confHandshakeMessage.ProviderID), string(encodedMsg))
	}
	return shim.OK
}

func testTransferIDAsset(t *testing.T, stub *shim.MockStub, transactionID string) (int32){
	//create a dummy transfer message
	transferMessage := createTransferAssetRecord(transactionID, t)

	encodedMsg, err := json.Marshal(transferMessage)
	if err!=nil{
		fmt.Println("Transfer record encoding failed..")
		t.FailNow()

	} else {
		fmt.Println("Encoded transferAsset message: ", string(encodedMsg))

		//since any write to the ledger needs to be in a transactional context, the test must start the transaction before
		// invoking and end after invoking.
		stub.MockTransactionStart("t1")
		responseStatus, responseMessage := checkInvoke(t, stub, [][]byte{[]byte(util.TRANSFER_ASSET), []byte(encodedMsg)})
		fmt.Println("Response from transferAsset invoke : ", responseMessage)
		stub.MockTransactionEnd("t1")

		if responseStatus != shim.OK {
			return responseStatus
		}

		//check if the message is in the ledger.
		checkQuery(t, stub, util.CreateTransactionKey(util.TRANSFER_ASSET, transferMessage.TransactionID,
			transferMessage.ConsumerID, transferMessage.UserID, transferMessage.ProviderID), string(encodedMsg))
	}
	return shim.OK
}

func testConfirmReceiptOfIDAsset(t *testing.T, stub *shim.MockStub, transactionID string) (int32){
	//create a dummy transfer message
	confirmMessage := createConfirmReceiptOfAssetRecord(transactionID)

	encodedMsg, err := json.Marshal(confirmMessage)
	if err!=nil{
		fmt.Println("ConfirmReceiptOfAsset record encoding failed..")
		t.FailNow()

	} else {
		fmt.Println("Encoded confirmReceiptOfAsset message: ", string(encodedMsg))

		//since any write to the ledger needs to be in a transactional context, the test must start the transaction before
		// invoking and end after invoking.
		stub.MockTransactionStart("t1")
		responseStatus, responseMessage := checkInvoke(t, stub, [][]byte{[]byte(util.CONFIRM_RECEIPT_OF_ASSET), []byte(encodedMsg)})
		fmt.Println("Response from confirmReceiptOfAsset invoke : ", responseMessage)
		stub.MockTransactionEnd("t1")

		if responseStatus != shim.OK {
			return responseStatus
		}

		//check if the message is in the ledger.
		checkQuery(t, stub, util.CreateTransactionKey(util.CONFIRM_RECEIPT_OF_ASSET, confirmMessage.TransactionID,
			confirmMessage.ConsumerID, confirmMessage.UserID, confirmMessage.ProviderID), string(encodedMsg))
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
		fmt.Println("Queried value is null.")
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
		"kyc_compliance", "s1", ""}
	return respHandshakeMessage
}

func createConfHandshakeRecord(transactionID string) (handshake.HandshakeRecord){

	//create a dummy handshake message for now. TODO: add actual crypto information.
	confHandshakeMessage := handshake.HandshakeRecord{ util.CONF_HANDSHAKE, transactionID,
		"c1", "c_PK", "u1","u_PK", "p1", "p_PK",
		"kyc_compliance", "s1", ""}
	return confHandshakeMessage
}

func createTransferAssetRecord(transactionID string, t *testing.T) (transfer.TransferRecord) {
	dat, err := ioutil.ReadFile("test/sample_id_asset")
	//fmt.Printf("Sample id asset read from file: ", string(dat))

	if err != nil {
		fmt.Errorf("Error in reading the identity asset file.")
		t.FailNow()
	}

	transferRecord := transfer.TransferRecord{transactionID, "c1", "c_PK",
		"u1", "u_PK", "p1", "p_PK", "kyc_compliance",
		dat, "s1"}

	return transferRecord
}

func createConfirmReceiptOfAssetRecord(transactionID string) (confirm.ConfirmRecord){
	confirmRecord := confirm.ConfirmRecord{transactionID, "c1", "c_PK", "u1",
	"u_PK", "p1", "p_PK", "kyc_compliance", "s1"}

	return confirmRecord
}