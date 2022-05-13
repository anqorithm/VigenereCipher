package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPKRSTUVWXYZabcdefghijklmnopkrstuvwxyz0123456789 ?!:;.,-_<>*'|=+{}()[]&#@/%\\"
)

func encrypt(originalMsg, key string) string {
	res := ""
	for i := 0; i < len(originalMsg); i++ {
		msgChar := strings.Index(alphabet, string([]rune(originalMsg)[i]))
		keyChar := strings.Index(alphabet, string([]rune(key)[i%len(key)]))
		encryptedChar := (msgChar + keyChar) % len(alphabet)
		res += string(alphabet[encryptedChar])
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

func redFromFile(filePath string) string {
	result, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

func writeToFile(filePath, data string) {
	file, err0 := os.Create(filePath)
	if err0 != nil {
		log.Fatal(err0)
	}
	defer file.Close()
	_, err1 := file.WriteString(data)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("Done successfully!.")
}

func main() {
	var (
		input     = flag.String("i", "input.txt", "a TXT file contains the original message")
		operation = flag.String("p", "e", "an operation for the provided file 'e' for encrypt, 'd' for decrypt")
		key       = flag.String("k", "key", "a encryption key")
		output    = flag.String("o", "output.txt", "a TXT file contains the result for the provided process")
	)
	flag.Parse()
	originalData := redFromFile(*input)
	switch *operation {
	case "e":
		result := encrypt(originalData, *key)
		if output != nil {
			writeToFile(*output, result)
		}
	case "d":
		result := decrypt(originalData, *key)
		if output != nil {
			writeToFile(*output, result)
		}
	}
}
