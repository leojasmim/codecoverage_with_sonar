package main

const (
	//EVEN PAR
	EVEN = iota
	//ODD - IMPAR
	ODD
)

func evenOrOdd(v int) int {
	if v%2 == 0 {
		return EVEN
	}
	return ODD
}
