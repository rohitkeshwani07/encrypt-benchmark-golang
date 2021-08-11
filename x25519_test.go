package cryptobench

import (
	"testing"
)

func BenchmarkEncrypt(b *testing.B) {
	msg := `Invalid string`

	for n := 0; n < b.N; n++ {
		Encrypt(msg)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	msg := `Invalid string`
	enMsg, _ := Encrypt(msg)

	for n := 0; n < b.N; n++ {
		Decrypt(enMsg)
	}
}
