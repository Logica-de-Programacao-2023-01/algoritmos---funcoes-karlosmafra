package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()

	var strs = separar(str)

	fmt.Println(strs)
}

func separar(str string) []string {
	var strs = []string{}
	return strs
}
