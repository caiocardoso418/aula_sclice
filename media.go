package main

import "fmt"

func main() {

	arr := [6]float64{3.4, 5.6, 7.2, 1.0, 4.5, 2.1}

	var total float64
	for _, value := range arr {
		total += value
	}
	mean := total / float64(len(arr))

	fmt.Printf("A média é: %.2f", mean)
}
