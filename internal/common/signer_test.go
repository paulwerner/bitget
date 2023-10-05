package common

import (
	"github.com/paulwerner/bitget-golang-sdk/internal"
	"fmt"
	"testing"
)

func TestSigner_Sign(t *testing.T) {
	signer := new(Signer)
	result := signer.Sign("GET", "www.bitget.com", "aaaa", internal.TimesStamp())
	fmt.Print(result)
}
