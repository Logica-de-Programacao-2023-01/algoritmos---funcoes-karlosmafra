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
	var pares, err = verpar(nums)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(pares)
	}
}

func verpar(nums []int) ([]int, error) {
	if len(nums) != 0 {
		var pares = []int{}
		for i := range nums {
			if nums[i]%2 == 0 {
				pares = append(pares, nums[i])
			}
		}
		return pares, nil
	} else {
		return nums, fmt.Errorf("O slice está vazio")
	}
}
