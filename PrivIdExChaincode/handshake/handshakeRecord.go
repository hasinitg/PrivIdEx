package handshake

// Represents the structure of messages exchanged in the three-way handshake: init, response, confirm
//init is signed by: user and consumer, response is signed by producer, confirm is signed by consumer
type HandshakeRecord struct {
	HandshakeRecordType,
	TransactionID string
	ConsumerID,
	ConsumerPublicKey,
	UserID,
	UserPublicKey,
	ProviderID,
	ProviderPublicKey []byte
	IdentityAssetName string
	//a handshake record is signed by at most two parties
	Signature1,
	Signature2 []byte
	//add a map as meta data which can be used for SLA agreements.
}
