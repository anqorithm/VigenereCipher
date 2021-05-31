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
		i = flag.String("i", "input.txt", "a TXT file contains the original message")
		p = flag.String("p", "e", "an opreation for the provided file 'e' for encrypt, 'd' for decrypt")
		k = flag.String("k", "key", "a encryption key")
		o = flag.String("o", "output.txt", "a TXT file contains the result for the provided process")
	)
	flag.Parse()
	originalData := redFromFile(*i)
	switch *p {
	case "e":
		res := encrypt(originalData, *k)
		if o != nil {
			writeToFile(*o, res)
		}
	case "d":
		res := decrypt(originalData, *k)
		if o != nil {
			writeToFile(*o, res)
		}
	}
}
