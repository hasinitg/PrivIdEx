package util

import "bytes"

/**
	This is the util method that creates a unique transaction key for each message of the protocol published in the
	block chain.
 */
func CreateTransactionKey(msgType, transactionID, consumerID, userID, providerID string) (string) {
	var buffer bytes.Buffer
	buffer.WriteString(msgType)
	buffer.WriteString(":")
	buffer.WriteString(transactionID)
	buffer.WriteString(":")
	buffer.WriteString(consumerID)
	buffer.WriteString(":")
	buffer.WriteString(userID)
	buffer.WriteString(":")
	buffer.WriteString(providerID)
	return buffer.String()
}
