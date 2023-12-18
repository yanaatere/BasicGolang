package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // sama dengan result = result * i
	}
	return result
}

func factorialRecrusive(value int) int {
	if value == 1 {
		return 1 //Recursive harus ada kondisi berhenti, agar tidak terjadi stack overflow
	} else {
		return value * factorialRecrusive(value-1)
	}
}

func main() {
	input := 6
	loop := factorialLoop(input)
	fmt.Println("Factorial Using Loop :", loop)

	recrusive := factorialRecrusive(input)
	fmt.Println("Factorial Using Recrusive :", recrusive)
}
