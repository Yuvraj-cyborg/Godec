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
						pixels = append(pixels, uint8(r>>8))
						pixels = append(pixels, uint8(g>>8))
						pixels = append(pixels, uint8(b>>8))
				}
    }

	fmt.Println(pixels)

	return nil

}
