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

	var str2, err = substituir(str)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(str2)
	}
}

func substituir(str string) (string, error) {
	if len(str) == 0 {
		return "", fmt.Errorf("A string está vazia")
	} else {
		str = strings.ReplaceAll(str, "A", "*")
		str = strings.ReplaceAll(str, "E", "*")
		str = strings.ReplaceAll(str, "I", "*")
		str = strings.ReplaceAll(str, "O", "*")
		str = strings.ReplaceAll(str, "U", "*")
		str = strings.ReplaceAll(str, "a", "*")
		str = strings.ReplaceAll(str, "e", "*")
		str = strings.ReplaceAll(str, "i", "*")
		str = strings.ReplaceAll(str, "o", "*")
		str = strings.ReplaceAll(str, "u", "*")
		return str, nil
	}
}
