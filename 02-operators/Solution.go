package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	// Declare totalCost, tip and tax
	var total_cost int32
	var tip float64
	var tax float64

	// Read and save an mealCost, tipPercent, and taxPercent to your variables.
	fmt.Scan(&meal_cost)
	fmt.Scan(&tip_percent)
	fmt.Scan(&tax_percent)

	// Calculate tip, tax and total cost
	tip = meal_cost * (float64(tip_percent) / 100)
	tax = meal_cost * (float64(tax_percent) / 100)
	total_cost = int32(math.Round((meal_cost + tip + tax)))

	fmt.Println(total_cost)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	meal_cost, err := strconv.ParseFloat(readLine(reader), 64)
	checkError(err)

	tip_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tip_percent := int32(tip_percentTemp)

	tax_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tax_percent := int32(tax_percentTemp)

	solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}