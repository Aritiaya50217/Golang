package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()`-=@"

	// Encoder (เข้ารหัส)
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoder (ถอดรหัส)
	// Decoding may return an error
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64 format
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
