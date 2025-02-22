package main

func main() {
	ExampleNewGCMEncrypter()

	//// 密钥，必须是8字节
	//key := []byte("12345678")
	//// 待解密数据
	//ciphertext := []byte("U2FsdGVkX1+T8hiY6zL5Qg==")
	//
	//block, err := des.NewCipher(key)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 解密
	//plaintext := make([]byte, len(ciphertext))
	//block.Decrypt(plaintext, ciphertext)
	//
	//fmt.Println(hex.EncodeToString(plaintext))
}

func desEncrypt(data, key string) (string, error) {
	return "", nil
}

func desDecrypt(data string) (string, error) {
	return "", nil
}
