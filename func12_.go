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

	var mai, err = maiusculas(lista)

	if err != nil {
		fmt.Print("ERRO: ", err)
	} else {
		fmt.Println(mai)
	}
}

func maiusculas(strings []string) ([]string, error) {
	if len(strings) != 0 {
		var mai = []string{}
		return mai, nil
	} else {
		return []string{}, fmt.Errorf("O slice está vazio.")
	}
}
