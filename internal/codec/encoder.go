package codec

import (
	"os"
	"image"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
)

func Encode(path string) error {
	file,err := os.Open(path)
	
	var pixels []uint8

	if err != nil {
		return err
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return err
	}

	bounds := img.Bounds()
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
        	r,g,b,_ := img.At(x,y).RGBA()
			R8 := uint8(r >> 8)
			G8 := uint8(g >> 8)
			B8 := uint8(b >> 8)

			brightness := (int(R8) + int(G8) + int(B8)) / 3
			pixels = append(pixels, uint8(brightness))
		}
    }

	//fmt.Println(pixels)

	delta := deltaEncode(pixels)
	rle_encoded := rleEncode(delta)
	fmt.Println(rle_encoded)

	return nil
}

func deltaEncode(input []uint8) []int16 {
	if len(input) == 0 {
		return []int16{}
	}

	output := make([]int16, len(input))
	output[0] = int16(input[0])
	for i := 1; i < len(input); i++ {
		output[i] = int16(input[i]) - int16(input[i-1])
	}
	return output
}


