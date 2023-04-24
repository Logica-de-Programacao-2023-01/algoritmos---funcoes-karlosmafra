package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x)

	var slc, err = primos(x)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(slc)
	}
}

func primos(x int) ([]int, error) {
	if x < 0 {
		return []int{}, fmt.Errorf("O número é negativo")
	} else {
		var pri = []int{}
		for i := 2; i <= x; i++ {
			if i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 || i == 2 || i == 3 || i == 5 || i == 7 {
				pri = append(pri, i)
			}
		}
		return pri, nil
	}
}
