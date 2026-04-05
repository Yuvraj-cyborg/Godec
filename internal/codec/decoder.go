package codec

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/png"
	"image/color"
	"os"

	"github.com/Yuvraj-cyborg/godec/internal/types"
)

func Decode(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	magic := make([]byte, 4)
	if _, err := file.Read(magic); err != nil {
		return err
	}

	if string(magic) != "GDC1" {
		return fmt.Errorf("invalid file format")
	}

	var width int32
	var height int32

	if err := binary.Read(file, binary.LittleEndian, &width); err != nil {
		return err
	}
	if err := binary.Read(file, binary.LittleEndian, &height); err != nil {
		return err
	}

	ch := make([]byte, 1)
	if _, err := file.Read(ch); err != nil {
		return err
	}

	var runCount int32
	if err := binary.Read(file, binary.LittleEndian, &runCount); err != nil {
		return err
	}

	runs := make([]types.Run, 0, runCount)

	for i := 0; i < int(runCount); i++ {
		buf := make([]byte, 2)
		if _, err := file.Read(buf); err != nil {
			return err
		}

		value := int16(int8(buf[0]))
		count := int(buf[1])

		runs = append(runs, types.Run{
			Value: value,
			Count: count,
		})
	}

	delta := RLEDecode(runs)

	pixels := DeltaDecode(delta)

	want := int64(width) * int64(height)
	if int64(len(pixels)) != want {
		return fmt.Errorf("pixel count mismatch: got %d samples, expected %d×%d=%d (corrupt .gdc or format bug)",
			len(pixels), width, height, want)
	}

	img := image.NewGray(image.Rect(0, 0, int(width), int(height)))

	for i := range pixels {
		x := i % int(width)
		y := i / int(width)
		img.SetGray(x, y, color.Gray{Y: pixels[i]})
	}

	out, err := os.Create("output.png")
	if err != nil {
		return err
	}
	defer out.Close()

	if err := png.Encode(out, img); err != nil {
		return err
	}

	fmt.Println("Decoded image saved as output.png")

	return nil
}
