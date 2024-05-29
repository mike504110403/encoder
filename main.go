package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// hexKey := "19044cd66222478467ccde81e6647aba4bd0684012b3487c6c222bc571b2e544"
	// decoding(hexKey)
	test := "dlbm2ksagoEmVk7a2dG1gsYAsHsDs5x3KAhfzraA6fA="
	test = "a2ct5dfadvGnGA8e5sH6asJDqJrUf7b6ENnfsxlP7jA="
	//test = "GQRM1mIiR4RnzN6B5mR6ukvQaEASs0h8bCIrxXGy5UQ="
	encoded := encoding(test)
	//test2 := "7656e6824b1a82812656411ad9d1b5839d40dc7d43b04d8528085ff2bdc0b1faf7"
	if decoding(encoded) != test {
		fmt.Println("Fail!!!!!!!")
	}
}

func encoding(plaintext string) string {
	fmt.Println("Original:", plaintext)
	decodedBase64, err := base64.StdEncoding.DecodeString(plaintext)
	if err != nil {
		fmt.Println("Error decoding Base64 string:", err.Error())
	}
	fmt.Println("Decoded Base64 bytes:", decodedBase64)
	fmt.Println(len(decodedBase64))
	hexStr := hex.EncodeToString(decodedBase64)
	fmt.Println("Re-encoded hex string:", hexStr)
	return hexStr
}

func decoding(encoded string) string {
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode err:", err.Error())
	}
	base64Str := base64.StdEncoding.EncodeToString(decoded)
	fmt.Println("Decoded string:", base64Str)
	return base64Str
}
