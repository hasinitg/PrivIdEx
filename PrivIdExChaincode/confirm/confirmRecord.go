package confirm


type ConfirmRecord struct {
	RecordType,
	TransactionID,
	ConsumerID,
	ConsumerPublicKey,
	UserID,
	UserPublicKey,
	ProviderID,
	ProviderPublicKey,
	IdentityAssetName,
	Signature string
}