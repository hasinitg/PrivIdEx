package transfer

type TransferRecord struct{
	TransactionID,
	ConsumerID,
	ConsumerPublicKey,
	UserID,
	UserPublicKey,
	ProviderID,
	ProviderPublicKey,
	IdentityAssetName string
	IdAsset []byte
	Signature string
}
