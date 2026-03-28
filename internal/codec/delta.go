package codec

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
