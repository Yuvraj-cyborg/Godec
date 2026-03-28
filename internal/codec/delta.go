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

func DeltaDecode(delta []int16) []uint8 {
	if len(delta) == 0 {
		return []uint8{}
	}

	output := make([]uint8, len(delta))
	output[0] = uint8(delta[0])

	for i := 1; i < len(delta); i++ {
		val := int(output[i-1]) + int(delta[i])

		// clamp just in case
		if val < 0 {
			val = 0
		}
		if val > 255 {
			val = 255
		}

		output[i] = uint8(val)
	}

	return output
}
