package main

import (
	"fmt"
	"sort"
)

func main() {
	var lista = []string{}
	var tam int
	fmt.Print("Tamanho do slice: ")
	fmt.Scan(&tam)

	for i := 0; i < tam; i++ {
		var valor string
		fmt.Printf("String na posição %d: ", i)
		fmt.Scan(&valor)
		lista = append(lista, valor)
	}

	var alf, err = alfabetica(lista)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(alf)
	}
}

func alfabetica(lista []string) ([]string, error) {
	if len(lista) == 0 {
		return []string{}, fmt.Errorf("O slice está vazio")
	} else {
		sort.Strings(lista)
		return lista, nil
	}
}
