package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()

	var strs, err = separar(str)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(strs)
	}
}

func separar(str string) ([]string, error) {

	if len(str) == 0 {
		return []string{}, fmt.Errorf("A string está vazia")
	}

	strs := strings.Split(str, " ")

	return strs, nil
}
