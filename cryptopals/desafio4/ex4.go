package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
)

func xor(a []byte, b byte) []byte {
	var bufResult bytes.Buffer

	for i := 0; i < len(a); i++ {
		bufResult.WriteByte(a[i] ^ b)
	}

	return bufResult.Bytes()
}

func scoreText(a []byte) int {
	frequency := map[byte]int{
		' ': 13, 'e': 12, 't': 9, 'a': 8, 'o': 8,
		'i': 7, 'n': 7, 's': 6, 'h': 6, 'r': 6,
		'd': 4, 'l': 4, 'u': 2, 'c': 2, 'm': 2,
		'f': 2, 'y': 2, 'w': 2, 'g': 2, 'p': 2,
		'b': 1, 'v': 1, 'k': 1, 'x': 1, 'j': 1,
		'q': 1, 'z': 1,
	}

	score := 0

	for _, char := range a {
		if val, exists := frequency[char]; exists {
			score += val
		} else {
			score--
		}
	}

	return score
}

func main() {
	data, err := os.Open("./4.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close() // É uma boa prática fechar o arquivo após o uso

	var itens []string

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		itens = append(itens, line)
	}

	text := ""
	bestText := 0

	for _, line := range itens {
		hexBytes1, errHexBytes1 := hex.DecodeString(line)
		if errHexBytes1 != nil {
			return
		}

		for i := 0; i < 256; i++ {
			xorResult := xor(hexBytes1, byte(i))
			score := scoreText(xorResult)
			if score >= bestText {
				text = string(xorResult)
				bestText = score
			}
		}
	}

	fmt.Print(text)
}
