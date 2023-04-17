package main

import "fmt"

func main() {
	var lista = []string{}
	var tam int
	fmt.Print("Tamanho do slice: ")
	fmt.Scan(&tam)

	for i := 0; i < tam; i++ {
		var str string
		fmt.Printf("string na posição %d: ", i)
		fmt.Scan(&str)
		lista = append(lista, str)
	}

	fmt.Println(concatenar(lista))
}

func concatenar(strings []string) string {
	var conc string
	for i := range strings {
		conc += strings[i]
	}
	return conc
}
