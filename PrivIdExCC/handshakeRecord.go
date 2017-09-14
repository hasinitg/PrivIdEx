package main

type HandshakeRecord struct{
	TransactionID,
	ConsumerID,
	UserID,
	ProviderID,
	IdentityAssetName,
	Signature string
}
