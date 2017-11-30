package test

import (
	"testing"
	"chaincode/PrivIdEx/PrivIdExChaincode/crypt"
	"fmt"
	"encoding/json"
	"bytes"
)

type Pack struct {
	Message,
	Signature,
	PublicKey string
}

type PackB struct {
	Message string
	PublicKey, Signature []byte
}

type PackBString struct{
	Message, PublicKey, Signature string
}

func TestSignature(t *testing.T){
	message := "This is a test message"
	//read the public key from file.
	publicKeyData, err := crypt.ReadPublicKeyDataFromFile("public.consumerSigningKey.pem")

	if err!=nil{
		fmt.Println("Error in reading the public key from file.", err)
		t.FailNow()
	}

	//convert publickey data to string
	pubKeyString := string(publicKeyData)
	fmt.Println("Public Key: ", pubKeyString)

	//try to parse the public key from here
	verifierLocal, err0 := crypt.ParsePublicKey([]byte(pubKeyString))
	if err0 != nil {
		fmt.Println("Error in parsig public key string.", err0)
		t.FailNow()
	}


	//read private key from file
	privateKey, err2 := crypt.ReadPrivateKeyFromFile("private.consumerSigningKey.pem")

	if err2 != nil {
		fmt.Println("Error in reading the private key from file.", err2)
		t.FailNow()
	}

	//sign the message using the private key
	signature, err3 := privateKey.Sign([]byte(message))

	if err3 != nil {
		fmt.Println("Error in signing the data.", err3)
		t.FailNow()
	}

	//try to verify the signature using local verifier:
	err01 := verifierLocal.Verify([]byte(message), []byte(signature))
	if err01!=nil{
		fmt.Println("Error in verifying signature locally.", err01)
		t.FailNow()
	}

	sigString := string(signature)
	//fmt.Println("Signature: ", sigString)

	//try to verify signature using the local verifier, but string converted signature.
	err02 := verifierLocal.Verify([]byte(message), []byte(sigString))
	if err02!=nil{
		fmt.Println("Error in verifying signature locally from string signature. ", err02)
		t.FailNow()
	}

	//msg := Pack{message, sigString, pubKeyString}
	msgB := PackB{message, publicKeyData, signature}
	//encode msg with JSON
	jsonMsg, err4 := json.Marshal(msgB)
	if err4 != nil {
		fmt.Println("Error in encoding the message.", err4)
		t.FailNow()
	}

	//print
	fmt.Println("Encoded message: ", string(jsonMsg))

	//decode the message
	//var p Pack
	var p PackB
	err5 := json.Unmarshal(jsonMsg, &p)
	if err5 != nil {
		fmt.Println("Error in decoding.", err5)
		t.FailNow()
	}

	decodedMsg := p.Message
	decodedSig := p.Signature
	decodedPubKey := p.PublicKey

	//parse the public key
	verifier, err6 := crypt.ParsePublicKey([]byte(decodedPubKey))
	if err6!=nil {
		fmt.Println("Error in parsing public key.", err6)
		t.FailNow()
	}
	err7 := verifier.Verify([]byte(decodedMsg), []byte(decodedSig))
	if err7 != nil {
		fmt.Println("Error in verifying the signature from the json message.", err7)
		t.FailNow()
	}

	//check encoding-decoding of string values of public key, signature
	/***** NOTE: from the test below, it is confirmed that we can not use string format of public key and signature in the json message.
	Need to use byte format. This was confirmed by observing the printed encoded message.************/


	/*msgBStr := PackBString{message, pubKeyString, sigString}
	//encode msg with JSON
	jsonMsg2, err8 := json.Marshal(msgBStr)
	if err8 != nil {
		fmt.Println("Error in encoding the message.", err8)
		t.FailNow()
	}

	//print
	fmt.Println("Encoded message: ", string(jsonMsg2))

	//decode the message
	var pStr PackBString
	err9 := json.Unmarshal(jsonMsg2, &pStr)
	if err9 != nil {
		fmt.Println("Error in decoding.", err9)
		t.FailNow()
	}

	decodedMsg2 := pStr.Message
	decodedSig2 := pStr.Signature
	decodedPubKey2 := pStr.PublicKey

	//parse the public key
	verifier2, err10 := crypt.ParsePublicKey([]byte(decodedPubKey2))
	if err10!=nil {
		fmt.Println("Error in parsing public key.", err10)
		t.FailNow()
	}
	err11 := verifier2.Verify([]byte(decodedMsg2), []byte(decodedSig2))
	if err11 != nil {
		fmt.Println("Error in verifying the signature from the json message.", err11)
		t.FailNow()
	}*/
}

func TestHash(t *testing.T){
	message := "this is a test message"
	d, err := crypt.Hash([]byte(message))
	if err!=nil{
		fmt.Println("Error in computing the hash. ", err)
		t.FailNow()
	}

	b, err1 := crypt.VerifyHash(d, message)
	if err1 !=nil{
		fmt.Println("Error in verifying the hash. ", err1)
		t.FailNow()
	}
	if !b {
		fmt.Println("Given digest does not equal that of the message.")
	}
}

type EncPack struct {
	Encypted []byte
}

func TestEncryptDecrypt(t *testing.T){
	// read the public key and private key

	//read the public key from file.
	publicKey, err1 := crypt.ReadPublicKeyFromFile("public.consumerEncryptionKey.pem")

	if err1!=nil{
		fmt.Println("Error in reading the public key from file.", err1)
		t.FailNow()
	}

	//encrypt a given message
	message := "This is a test message"
	encrypted, err2 := publicKey.EncryptOAEP([]byte(message), nil)
	if err2!=nil{
		fmt.Println("Error in encryption.", err2)
		t.FailNow()
	}

	//send the cipher text in a json message
	epack := EncPack{encrypted}
	jsonEncPack, err3 := json.Marshal(epack)
	if err3!=nil{
		fmt.Println("Error in encoding encrypted.", err3)
		t.FailNow()
	}

	//decrypt the cipher text received in the json message and check if it equals the original plain text message
	var deEncPack EncPack
	err4 := json.Unmarshal(jsonEncPack, &deEncPack)
	if err4!=nil{
		fmt.Println("Error in decoding the encoded encrypted.", err3)
		t.FailNow()
	}

	//read private key from file
	privateKey, err5 := crypt.ReadPrivateKeyFromFile("private.consumerEncryptionKey.pem")
	if err5!=nil{
		fmt.Println("Error in reading the private key from file.", err3)
		t.FailNow()
	}

	plainText, err6 := privateKey.DecryptOAEP(deEncPack.Encypted, nil)
	if err6!=nil{
		fmt.Println("Error in decrypting the encrypted got from json message.", err3)
		t.FailNow()
	}

	if !bytes.Equal(plainText, []byte(message)){
		fmt.Println("Decrypted message does not equal the original message.", err3)
		t.FailNow()
	}

}

