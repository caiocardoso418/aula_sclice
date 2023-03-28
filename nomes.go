package main

import "fmt"

func main() {
	nomes := []string{"jacinto, pinto, pereira"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}
}
