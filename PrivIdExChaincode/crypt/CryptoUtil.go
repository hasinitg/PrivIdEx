package crypt

//*************The code is mainly from here: https://play.golang.org/p/bzpD7Pa9mr with minor changes******************
import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"fmt"
	"crypto/sha256"
	"errors"
	"bytes"
)

//-----------------------Private Key Signature Functions------------------------------------------

// A PrivKey creates signatures that verify against a public key.
type PrivKey interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Sign(data []byte) ([]byte, error)
	DecryptOAEP(cipherText, label []byte) ([]byte, error)
}

// PrivKey interface is implemented by rsaPrivateKey type
type rsaPrivateKey struct {
	*rsa.PrivateKey
}

// Sign signs data with rsa-sha256
func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d)
}

func ReadPrivateKeyFromFile(path string) (PrivKey, error){
	//read from file
	privKeyData, err := ioutil.ReadFile(path)
	if err != nil{
		return nil, err
	}
	//parse key
	return parsePrivateKey(privKeyData)

}

// parsePrivateKey parses a PEM encoded private key.
func parsePrivateKey(pemBytes []byte) (PrivKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("No key found.")
	}

	var rawkey interface{}
	switch block.Type {
	case "PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("Unsupported key type %q", block.Type)
	}
	return newPrivKeyFromKey(rawkey)
}

func newPrivKeyFromKey(k interface{}) (PrivKey, error) {
	var sshKey PrivKey
	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, fmt.Errorf("Unsupported key type %T", k)
	}
	return sshKey, nil
}
//-------------------------Public Key Signature Functions---------------------------------------------------



// A PubKey verifies signatures against signed data.
type PubKey interface {
	//Verifies signature on data
	Verify(data[]byte, sig []byte) error
	EncryptOAEP(plainText, label []byte) ([]byte, error)
}

// PubKey interface is implemented by rsa public key
type rsaPublicKey struct {
	*rsa.PublicKey
}

// Verifies signature on data
func (r *rsaPublicKey) Verify(message []byte, sig []byte) error {
	h := sha256.New()
	h.Write(message)
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, d, sig)
}

//Reads the content of a public key file.
func ReadPublicKeyDataFromFile(path string) ([]byte, error){
	//read from file
	pubKeyData, err := ioutil.ReadFile(path)
	if err != nil{
		return nil, err
	}
	return pubKeyData, nil

}

func ReadPublicKeyFromFile(path string) (PubKey, error){
	//read from file
	pubKeyData, err := ioutil.ReadFile(path)
	if err != nil{
		return nil, err
	}
	//parse key
	return ParsePublicKey(pubKeyData)

}

func ParsePublicKey(pemBytes []byte) (PubKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("No key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "PUBLIC KEY":
		rsa, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("Unsupported key type %q", block.Type)
	}

	return newVerifierFromKey(rawkey)
}

func newVerifierFromKey(k interface{}) (PubKey, error) {
	var sshKey PubKey
	switch t := k.(type) {
	case *rsa.PublicKey:
		sshKey = &rsaPublicKey{t}
	default:
		return nil, fmt.Errorf("Unsupported key type %T", k)
	}
	return sshKey, nil
}


//--------------------------------------Hash Functions---------------------------------------------------

func Hash(message []byte) ([]byte, error){
	h := sha256.New()

	_, err := h.Write(message)

	if err!=nil{
		return nil, err
	}

	messageDigest := h.Sum(nil)

	return messageDigest, nil;
}

func VerifyHash(messageDigest []byte, message string) (bool, error){
	d, err := Hash([]byte(message))

	if err !=nil{
		return false, err
	}

	return bytes.Equal(messageDigest, d), nil
}

//---------------------------------Encrypt Functions-----------------------------------------------------

func (publicKey *rsaPublicKey) EncryptOAEP(plainText, label []byte) ([]byte, error){

	rng := rand.Reader
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rng, publicKey.PublicKey, plainText, label)
	if err!=nil{
		return nil, err
	}
	return cipherText, nil
}

func (privateKey *rsaPrivateKey) DecryptOAEP(cipherText, label []byte) ([]byte, error){
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, privateKey.PrivateKey, cipherText, label)
	if err!=nil{
		return nil, err
	}
	return plaintext, nil
}