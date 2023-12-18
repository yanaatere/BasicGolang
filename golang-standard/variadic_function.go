package main

import "fmt"

/*Variable Argument itu lah kenapa di sebut variadic argument
dan hanya bisa dilakukan dibelakang*/

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

/*Memasukkan slice ke variadic function*/
func insertSlice(slice []int) int {
	return sumAll(slice...)
}

func main() {
	total := sumAll(90, 50, 10)
	fmt.Println("Total Sum ALL", total)

	slice := []int{10, 20, 30, 50}
	totalInsertSlice := insertSlice(slice)
	fmt.Println("Total Inserted Slice", totalInsertSlice)
}
