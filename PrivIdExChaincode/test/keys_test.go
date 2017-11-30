package test


/************This file was obtained from : https://gist.github.com/sdorra/1c95de8cb80da31610d2ad767cd6f251*************/
import (
	"testing"
"crypto/rand"
"crypto/rsa"
"crypto/x509"
//"encoding/asn1"
"encoding/gob"
"encoding/pem"
"fmt"
"os"
)

func TestGenKeys(t *testing.T) {
	reader := rand.Reader
	bitSize := 2048

	//generate user signing keys
	userSigningKey, err := rsa.GenerateKey(reader, bitSize)
	CheckError(err)

	userSigningPublicKey := userSigningKey.PublicKey

	//SaveGobKey("private.userKey", userKey)
	SavePEMKey("private.userSigningKey.pem", userSigningKey)

	//SaveGobKey("public.userKey", userPublicKey)
	SavePublicPEMKey("public.userSigningKey.pem", &userSigningPublicKey)

	//-----------------------------------------------------------------------------------------

	//generate IAP signing keys
	providerSigningKey, err := rsa.GenerateKey(reader, bitSize)
	CheckError(err)

	providerSigningPublicKey := providerSigningKey.PublicKey

	//SaveGobKey("private.providerKey", providerKey)
	SavePEMKey("private.providerSigningKey.pem", providerSigningKey)

	//SaveGobKey("public.providerKey", providerPublicKey)
	SavePublicPEMKey("public.providerSigningKey.pem", &providerSigningPublicKey)

	//-------------------------------------------------------------------------------------------

	//generate IAC signing keys
	consumerSigningKey, err := rsa.GenerateKey(reader, bitSize)
	CheckError(err)

	consumerSigningPublicKey := consumerSigningKey.PublicKey

	//SaveGobKey("private.consumerKey", consumerKey)
	SavePEMKey("private.consumerSigningKey.pem", consumerSigningKey)

	//SaveGobKey("public.consumerKey", consumerPublicKey)
	SavePublicPEMKey("public.consumerSigningKey.pem", &consumerSigningPublicKey)

	//-------------------------------------------------------------------------------------------

	//generate IAC encryption keys
	consumerEncryptionKey, err := rsa.GenerateKey(reader, bitSize)
	CheckError(err)

	consumerEncryptionPublicKey := consumerEncryptionKey.PublicKey

	//SaveGobKey("private.consumerKey", consumerKey)
	SavePEMKey("private.consumerEncryptionKey.pem", consumerEncryptionKey)

	//SaveGobKey("public.consumerKey", consumerPublicKey)
	SavePublicPEMKey("public.consumerEncryptionKey.pem", &consumerEncryptionPublicKey)

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

//func SavePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
//	asn1Bytes, err := asn1.Marshal(pubkey)
//	CheckError(err)
//
//	var pemkey = &pem.Block{
//		Type:  "PUBLIC KEY",
//		Bytes: asn1Bytes,
//	}
//
//	pemfile, err := os.Create(fileName)
//	CheckError(err)
//	defer pemfile.Close()
//
//	err = pem.Encode(pemfile, pemkey)
//	CheckError(err)
//}


func SavePublicPEMKey(fileName string, pubkey *rsa.PublicKey) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	CheckError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
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