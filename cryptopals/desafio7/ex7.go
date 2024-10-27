package main

import (
    "crypto/aes"
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    // Lê o arquivo que contém o texto criptografado (em Base64)
    encryptedBase64, err := ioutil.ReadFile("7.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Decodifica o conteúdo Base64
    encryptedData, err := base64.StdEncoding.DecodeString(string(encryptedBase64))
    if err != nil {
        log.Fatal(err)
    }

    // Define a chave
    key := []byte("YELLOW SUBMARINE")

    // Descriptografa o texto
    decryptedData, err := decryptAES128ECB(encryptedData, key)
    if err != nil {
        log.Fatal(err)
    }

    // Imprime o texto descriptografado
    fmt.Print(string(decryptedData))
}

func decryptAES128ECB(ciphertext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    // O tamanho do bloco AES é 16 bytes
    blockSize := aes.BlockSize
    decrypted := make([]byte, len(ciphertext))

    for i := 0; i < len(ciphertext); i += blockSize {
        block.Decrypt(decrypted[i:i+blockSize], ciphertext[i:i+blockSize])
    }

    return removePKCS7Padding(decrypted), nil
}

func removePKCS7Padding(data []byte) []byte {
    paddingLen := int(data[len(data)-1])
    return data[:len(data)-paddingLen]
}
