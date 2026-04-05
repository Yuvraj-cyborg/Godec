package types

type Run struct {
	Value int16
	// Count must not use int16: long runs (e.g. repeated delta 0) exceed 32767 and overflow.
	Count int
}
