package main

import "fmt"

func main() {
	var nums = []int{}
	var tam int
	fmt.Print("Tamanho do slice: ")
	fmt.Scan(&tam)

	for i := 0; i < tam; i++ {
		var valor int
		fmt.Printf("valor na posição %d: ", i)
		fmt.Scan(&valor)
		nums = append(nums, valor)
	}

	var x int
	fmt.Print("Valor para buscar: ")
	fmt.Scan(&x)

	var cont, err = contem(nums, x)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(cont)
	}
}

func contem(nums []int, x int) (bool, error) {
	if len(nums) == 0 {
		{
			return false, fmt.Errorf("O slice está vazio")
		}
	} else {
		cont := false
		for i := range nums {
			if nums[i] == x {
				cont = true
				break
			}
		}
		return cont, nil
	}
}
