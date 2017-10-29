package confirm


type ConfirmRecord struct {
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