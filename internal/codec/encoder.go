package codec

import (
	"os"
	"image"
	"fmt"
	"log"
)

func encode() {
	file,err := os.Open("../../assets/eren.jpg")
	
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	img, imageType, err := image.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Image format: %s\n", imageType)
	fmt.Printf("Image bounds: %v\n", img.Bounds())
}
