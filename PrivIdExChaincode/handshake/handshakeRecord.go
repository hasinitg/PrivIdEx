package handshake

type HandshakeRecord struct{
	TransactionID,
	ConsumerID,
	UserID,
	ProviderID,
	IdentityAssetName,
	Signature string
}
