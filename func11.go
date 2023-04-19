package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número primo: ")
	fmt.Scan(&x)

	var pri, err = primo(x)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(pri)
	}
}

func primo(x int) (bool, error) {
	if x < 0 {
		return false, fmt.Errorf("O número é negativo")
	} else if x%2 != 0 && x%3 != 0 && x%5 != 0 && x%7 != 0 || x == 2 || x == 3 || x == 5 || x == 7 {
		return true, nil
	} else {
		return false, nil
	}
}
