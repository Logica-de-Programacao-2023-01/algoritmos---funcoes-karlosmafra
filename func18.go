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

	var res, err = aplic(nums)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println("A soma dos valores triplicados é igual a", res)
	}
}

func aplic(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("O slice está vazio.")
	} else {
		var tri = func(x int) int {
			return x * 3
		}

		var soma int
		for i := range nums {
			nums[i] = tri(nums[i])
			soma += nums[i]
		}

		return soma, nil
	}
}
