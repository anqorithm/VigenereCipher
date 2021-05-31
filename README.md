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

## provided alphabit

```go
const (
    // YOU CAN CHANGE THIS FOR ANYTHING YOU WANT
	alphabet = "ABCDEFGHIJKLMNOPKRSTUVWXYZabcdefghijklmnopkrstuvwxyz0123456789 !?.,"
)
```

# Example of usage

```
Hello World! -> (encryption) -> Ri9vsTgsAvhU
Ri9vsTgsAvhU -> (decryption) -> Hello World!
```
