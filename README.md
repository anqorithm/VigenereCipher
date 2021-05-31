# Usage

## i pramaeter ("input.txt" as default)

```
$ go build . && ./VigenereCipher -i YourFile.txt
```

## p parameter ("e" as default)

### 1- encryption

```bash
$ go build . && ./VigenereCipher -p e
```

### 2- decryption

```bash
$ go build . && ./VigenereCipher -p d
```

## k parameter ("KEY" as default)

```bash
$ go build . && ./VigenereCipher -k YourEncryptionKey
```

## o pramater ("output.txt" as default)

```bash
$ go build . && ./VigenereCipher -o YourOutputFile
```

## provided alphabet

```go
const (
    // YOU CAN CHANGE THIS FOR ANYTHING YOU WANT
	alphabet = "ABCDEFGHIJKLMNOPKRSTUVWXYZabcdefghijklmnopkrstuvwxyz0123456789 !?.,"
)
```

# Example of usage

```
encryption key = "key"
Hello World! (input.txt) -> (encryption) -> r8UGDt6DaG7u (output.txt)
r8UGDt6DaG7u (input.txt) -> (decryption) -> Hello World! (output.txt)
```
