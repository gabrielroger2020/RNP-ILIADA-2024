package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	file1, err := os.Open("random-image.png")
	if err != nil {
		fmt.Println("Erro ao abrir a imagem random:", err)
		return
	}

	file2, err := os.Open("encrypted1.png")
	if err != nil {
		fmt.Println("Erro ao abrir a imagem criptografada:", err)
		return
	}
	
	img1, _, err := image.Decode(file1)
	if err != nil {
		fmt.Println("Erro ao decodificar random-image:", err)
		return
	}

	img2, _, err := image.Decode(file2)
	if err != nil {
		fmt.Println("Erro ao decodificar encrypted1:", err)
		return
	}

	// Verifiquei se as dimensões das imagens são iguais
	if img1.Bounds() != img2.Bounds() {
		fmt.Println("As imagens não possuem as mesmas dimensões, logo não podem ser comparadas.")
		return
	}

	// Criei uma nova imagem para armazenar o resultado
	bounds := img1.Bounds()
	result := image.NewRGBA(bounds)

	// Apliquei XOR pixel a pixel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()

			r := uint8(r1>>8) ^ uint8(r2>>8)
			g := uint8(g1>>8) ^ uint8(g2>>8)
			b := uint8(b1>>8) ^ uint8(b2>>8)

			result.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// Salvei a imagem decifrada
	outputFile, err := os.Create("./decrypted_image.png")
	if err != nil {
		fmt.Println("Erro ao criar a imagem de saída:", err)
		return
	}

	err = png.Encode(outputFile, result)
	if err != nil {
		fmt.Println("Erro ao salvar a imagem decifrada:", err)
		return
	}

	fmt.Println("A imagem decifrada foi salva como decrypted_image.png")
}
