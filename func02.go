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

	fmt.Printf("A frase tem %d vogais", contar(strings.ToUpper(str)))
}

func contar(str string) int {
	var vog int
	for i := range str {
		car := string(str[i])
		if car == "A" || car == "E" || car == "I" || car == "O" || car == "U" {
			vog++
		}
	}
	return vog
}
