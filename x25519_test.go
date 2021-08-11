package cryptobench

import (
	"testing"
)

func BenchmarkEncryptOAEP(b *testing.B) {
	msg := `Invalid string`

	for n := 0; n < b.N; n++ {
		EncryptOAEP(msg)
	}
}

func BenchmarkDecryptOAEP(b *testing.B) {
	msg := `Invalid string`
	enMsg, _ := EncryptOAEP(msg)

	for n := 0; n < b.N; n++ {
		DecryptOAEP(enMsg)
	}
}

func BenchmarkEncryptPKCS1v15(b *testing.B) {
	msg := `Invalid string`

	for n := 0; n < b.N; n++ {
		EncryptPKCS1v15(msg)
	}
}

func BenchmarkDecryptPKCS1v15(b *testing.B) {
	msg := `Invalid string`
	enMsg, _ := EncryptPKCS1v15(msg)

	for n := 0; n < b.N; n++ {
		DecryptPKCS1v15(enMsg)
	}
}
