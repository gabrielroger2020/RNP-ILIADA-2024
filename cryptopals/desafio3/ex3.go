package main

import ("bytes"
		"encoding/hex"
		"fmt");


func xor(a []byte, b byte) []byte{

	var bufResult bytes.Buffer;
	
	for i:= 0; i < len(a); i++{
		bufResult.WriteByte(a[i] ^ b);
	}

	//fmt.Println(string(bufResult.Bytes()));

	return bufResult.Bytes();
}


func scoreText(a []byte) int{

	frequency := map[byte]int{
		' ': 13, 'e': 12, 't': 9,  'a': 8,  'o': 8,
		'i': 7,  'n': 7,  's': 6,  'h': 6,  'r': 6,
		'd': 4,  'l': 4,  'u': 2,  'c': 2,  'm': 2,
		'f': 2,  'y': 2,  'w': 2,  'g': 2,  'p': 2,
		'b': 1,  'v': 1,  'k': 1,  'x': 1,  'j': 1,
		'q': 1,  'z': 1,
	}

	score := 0;

	for _, char := range a {
		if val, exists := frequency[char]; exists {
			score += val;
		} else {
			score--;
		}
	}

	return score;

}

func main(){
	hex1 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736";

	//Convertendo os hex para bytes e armazenando no buffer.

	hexBytes1, errHexBytes1 := hex.DecodeString(hex1);

	text := "";
	bestText := 0;

	if(errHexBytes1 != nil){
		return
	}

	for i := 0; i < 256; i++{
		//fmt.Println(byte(i));
		xorResult := xor(hexBytes1, byte(i));
		score := scoreText(xorResult);
		if(score >= bestText){
			text = string(xorResult);
			bestText = score;
		}
	} 
	
	fmt.Println(text);
}