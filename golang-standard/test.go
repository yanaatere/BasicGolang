package main

import (
	"fmt"
	"time"
)

func convertTimeFormat(inputTimeStr string) string {
	// Parsing the input time string
	inputTime, err := time.Parse("03:04:05PM", inputTimeStr)
	if err != nil {
		// If there's an error, you can choose to handle it in some way.
		// For simplicity, let's print the error and return an empty string.
		fmt.Println("Error parsing time:", err)
		return ""
	}

	// Formatting the time in a different format
	outputTimeStr := inputTime.Format("15:04:05")

	return outputTimeStr
}
func main() {
	inputTimeStr := "12:01:00PM"

	// Call the function to convert the time format
	result := convertTimeFormat(inputTimeStr)

	// Print the converted time
	fmt.Println("Converted time:", result)
}
