package main

import ("encoding/hex"
		"encoding/base64"
		"fmt");

func main(){
	var hexValue string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d";
	bytesHex, errorHex := hex.DecodeString(hexValue);

	if (errorHex != nil) {
		fmt.Println("Error: ", errorHex);
		return;
	}
	
	//fmt.Println("Bytes decodificados: ", bytesHex);
	var base64Value string = base64.StdEncoding.EncodeToString(bytesHex);

	fmt.Println(base64Value);
	
}