package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func ExampleNewGCMEncrypter() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	//key := []byte("AES256Key-32Characters1234567890")
	key2 := []byte("32Characters1234")
	plaintext := []byte("1000000")

	block, err := aes.NewCipher(key2)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext)
}

func ExampleNewGCMDecrypter() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("AES256Key-32Characters1234567890")
	ciphertext, _ := hex.DecodeString("f90fbef747e7212ad7410d0eee2d965de7e890471695cddd2a5bc0ef5da1d04ad8147b62141ad6e4914aee8c512f64fba9037603d41de0d50b718bd665f019cdcd")

	nonce, _ := hex.DecodeString("bb8ef84243d2ee95a41c6c57")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", string(plaintext))
}
