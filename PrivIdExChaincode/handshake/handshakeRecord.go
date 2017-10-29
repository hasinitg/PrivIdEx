package handshake

type HandshakeRecord struct{
	HandshakeRecordType,
	TransactionID,
	ConsumerID,
	ConsumerPublicKey,
	UserID,
	UserPublicKey,
	ProviderID,
	ProviderPublicKey,
	IdentityAssetName,
	Signature1,
	Signature2 string
	//add a map as meta data which can be used for SLA agreements.
}
