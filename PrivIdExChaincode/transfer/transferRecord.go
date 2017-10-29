package transfer

type transferRecord struct{
	TransactionID,
	ConsumerID,
	ConsumerPublicKey,
	UserID,
	UserPublicKey,
	ProviderID,
	ProviderPublicKey,
	IdentityAssetName,
	Signature string
	IdAsset []byte
}
