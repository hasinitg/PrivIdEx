package transfer

type TransferRecord struct{
	RecordType,
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
