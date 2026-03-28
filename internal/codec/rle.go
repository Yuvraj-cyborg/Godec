package codec

import (
	"github.com/Yuvraj-cyborg/godec/internal/types"
)

func rleEncode(input []int16) []types.Run {
	var output []types.Run

	var current int16 = input[0]
	var count int16 = 1

	for i := 1; i < len(input); i++ {
		if input[i] == current {
			count++
		} else {
			output = append(output, types.Run{Value: current, Count: count})
			current = input[i]
			count = 1
		}
	}

	// emit last run
	output = append(output, types.Run{Value: current, Count: count})

	return output
}

