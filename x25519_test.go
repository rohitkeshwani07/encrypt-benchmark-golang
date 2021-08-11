package cryptobench

import (
	"fmt"
	"testing"
)

func BenchmarkEncryptOAEP(b *testing.B) {
	msg := `merchant_id123:t:1628674459`

	for n := 0; n < b.N; n++ {
		EncryptOAEP(msg)
	}
}

func BenchmarkDecryptOAEP(b *testing.B) {
	msg := `merchant_id123:t:1628674459`
	enMsg, _ := EncryptOAEP(msg)

	for n := 0; n < b.N; n++ {
		DecryptOAEP(enMsg)
	}
}

func BenchmarkEncryptPKCS1v15(b *testing.B) {
	msg := `merchant_id123:t:1628674459`

	for n := 0; n < b.N; n++ {
		EncryptPKCS1v15(msg)
	}
}

func BenchmarkDecryptPKCS1v15(b *testing.B) {
	msg := `merchant_id123:t:1628674459`
	enMsg, _ := EncryptPKCS1v15(msg)

	for n := 0; n < b.N; n++ {
		DecryptPKCS1v15(enMsg)
	}
}

func TestEncryptOAEP(b *testing.T) {
	msg := `0123456789123:t:1628674459`

	enMsg, _ :=	EncryptOAEP(msg)
	fmt.Println(len(enMsg))
}
