package handshake

import (
	"testing"
)

func Test(t *testing.T){
	//test sending an invalid message signature
	InitHandshake(nil, []string{})
}