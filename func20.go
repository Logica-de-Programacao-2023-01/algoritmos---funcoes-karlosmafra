package main

import (
	"fmt"
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

	var m5, err = maior5(lista)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(m5)
	}
}

func maior5(lista []string) ([]string, error) {
	if len(lista) == 0 {
		return []string{}, fmt.Errorf("O slice está vazio")
	} else {
		var m5 = []string{}
		for i := range lista {
			if len(lista[i]) > 5 {
				m5 = append(m5, lista[i])
			}
		}
		return m5, nil
	}
}
