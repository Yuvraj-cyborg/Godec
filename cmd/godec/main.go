package main

import (
	"fmt"
	"github.com/Yuvraj-cyborg/godec/internal/codec"
)

func main() {
	err := codec.Decode("out.gdc")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
