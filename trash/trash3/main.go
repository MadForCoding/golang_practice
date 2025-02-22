package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	sig := `Xa8PuWIgSBjgmcEhzjxKJ7ZRDV2wAM8IYxkT5saOWBtfYFs3i9k6AIgaKKTRME2u`
	body := `{"meta":{},"data":{"uid":102983960}}`
	signData := Sign([]byte(body), []byte(sig))
	fmt.Println(string(signData))
	return
}

// verify the signature of `payload` with `key`
func VerifySignature(payload, key, signature []byte) bool {
	return bytes.Equal(signature, Sign(payload, key))
}

// Sign the `payload` with `key`
func Sign(payload, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	if _, err := mac.Write(payload); err != nil {
		return nil
	}
	fmt.Println(string(mac.Sum(nil)))
	fmt.Println("-------")
	return []byte(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
}
