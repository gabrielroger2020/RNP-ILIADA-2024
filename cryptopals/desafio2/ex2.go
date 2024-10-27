package main

import ("bytes"
		"encoding/hex"
		"fmt");


func xor(a []byte, b []byte){

	//Verifico se os buffers possuem o mesmo tamanho!
	if(len(a) != len(b)){
		fmt.Println("Os buffers não possuem o mesmo tamanho!");
		return;
	}else{
		//fmt.Println("Os buffers possuem o mesmo tamanho!");
	}

	var bufResult bytes.Buffer;

	for i:= 0; i < len(a); i++{
		bufResult.WriteByte(a[i]^b[i]);
	}

	fmt.Println(hex.EncodeToString(bufResult.Bytes()));

	return;
} 

func main(){
	hex1 := "1c0111001f010100061a024b53535009181c";
	hex2 := "686974207468652062756c6c277320657965";

	//Convertendo os hex para bytes e armazenando no buffer.

	hexBytes1, errHexBytes1 := hex.DecodeString(hex1);
	hexBytes2, errHexBytes2 := hex.DecodeString(hex2);

	if(errHexBytes1 != nil || errHexBytes2 != nil){
		return
	}

	buf1 := bytes.NewBuffer(hexBytes1);
	buf2 := bytes.NewBuffer(hexBytes2);

	xor(buf1.Bytes(),buf2.Bytes());
}