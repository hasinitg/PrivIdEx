package handshake

type HandshakeRecord struct{
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
}
