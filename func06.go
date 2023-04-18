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

	var conc, err = concat(lista)

	if err != nil {
		fmt.Print("ERRO: ", err)
	} else {
		fmt.Println(conc)
	}
}

func concat(strings []string) (string, error) {
	if len(strings) != 0 {
		var conc string
		for i := range strings {
			conc += strings[i]
			if i+1 != len(strings) {
				conc += ", "
			}
		}
		return conc, nil
	} else {
		return "", fmt.Errorf("O slice está vazio.")
	}
}
