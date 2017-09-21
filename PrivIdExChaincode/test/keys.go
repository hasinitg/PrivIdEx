package test

//import (
//	"encoding/pem"
//	"crypto/rsa"
//	"encoding/asn1"
//	"os"
//	"fmt"
//	"encoding/gob"
//	"crypto/x509"
//)
//
///*
// * Genarate rsa keys.
// */
//
//package main

/************This file was obtained from : https://gist.github.com/sdorra/1c95de8cb80da31610d2ad767cd6f251*************/
import (
"crypto/rand"
"crypto/rsa"
"crypto/x509"
"encoding/asn1"
"encoding/gob"
"encoding/pem"
"fmt"
"os"
)

func main() {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	CheckError(err)

	publicKey := key.PublicKey

	SaveGobKey("private.key", key)
	SavePEMKey("private.pem", key)

	SaveGobKey("public.key", publicKey)
	SavePublicPEMKey("public.pem", publicKey)
}

func SaveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	CheckError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	CheckError(err)
}

func SavePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	CheckError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	CheckError(err)
}

func SavePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubkey)
	CheckError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	CheckError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}