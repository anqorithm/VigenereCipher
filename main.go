package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPKRSTUVWXYZabcdefghijklmnopkrstuvwxyz0123456789 !?.,"
)

func encrypt(originalMsg, key string) string {
	res := ""
	for i := 0; i < len(originalMsg); i++ {
		msgChar := strings.Index(alphabet, string([]rune(originalMsg)[i]))
		keyChar := strings.Index(alphabet, string([]rune(key)[i%len(key)]))
		encryptedChat := (msgChar + keyChar) % len(alphabet)
		res += string(alphabet[encryptedChat])
	}
	return res
}

func decrypt(encryptedMsg, key string) string {
	res := ""
	for i := 0; i < len(encryptedMsg); i++ {
		msgChar := strings.Index(alphabet, string([]rune(encryptedMsg)[i]))
		keyChar := strings.Index(alphabet, string([]rune(key)[i%len(key)]))
		decryptedChar := ((msgChar - keyChar) % len(alphabet))
		if decryptedChar < 0 {
			decryptedChar += len(alphabet)
		}
		res += string(alphabet[decryptedChar])
	}
	return res
}

func readTheOriginalData(filePath string) string {
	result, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

func main() {
	var (
		inputFile = flag.String("inputFile", "input.txt", "a TXT file contains the original message")
		process   = flag.String("process", "e", "an opreation for the provided file 'e' for encrypt, 'd' for decrypt")
		key       = flag.String("key", "RED", "a encryption key")
	)
	flag.Parse()
	originalData := readTheOriginalData(*inputFile)
	switch *process {
	case "e":
		res := encrypt(originalData, *key)
		fmt.Println(res)
	case "d":
		res := decrypt(originalData, *key)
		fmt.Println(res)
	}
}
