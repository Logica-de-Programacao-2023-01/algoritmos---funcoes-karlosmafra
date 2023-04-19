package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&x)

	var soma, err = somar(x)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println("A soma dos dígitos é igual a", soma)
	}
}

func somar(x int) (int, error) {
	if x < 0 {
		return 0, fmt.Errorf("O número é negativo.")
	} else {
		var soma int
		numS := strconv.Itoa(x)
		for i := range numS {
			car := string(numS[i])
			n, err := strconv.Atoi(car)
			if err == nil {
				soma += n
			}
		}
		return soma, nil
	}
}
