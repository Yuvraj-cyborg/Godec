package format

import (
	"encoding/binary"
	"os"

	"github.com/Yuvraj-cyborg/godec/internal/types"
)

func WriteGDC(path string, width int, height int, runs []types.Run) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte("GDC1")); err != nil {
		return err
	}

	if err := binary.Write(file, binary.LittleEndian, int32(width)); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, int32(height)); err != nil {
		return err
	}

	if _, err := file.Write([]byte{1}); err != nil {
		return err
	}

	if err := binary.Write(file, binary.LittleEndian, int32(len(runs))); err != nil {
		return err
	}

	for _, run := range runs {
		value := int8(run.Value)

		remaining := run.Count

		for remaining > 255 {
			if _, err := file.Write([]byte{byte(value), 255}); err != nil {
				return err
			}
			remaining -= 255
		}

		if _, err := file.Write([]byte{byte(value), uint8(remaining)}); err != nil {
			return err
		}
	}

	return nil
}
