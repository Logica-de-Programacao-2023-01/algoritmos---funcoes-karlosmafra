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

	var res, err = aplicar(nums)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(res)
	}
}

func aplicar(nums []int) ([]int, error) {
	if len(nums) != 0 {
		dobrar := func(x int) int {
			return x * 2
		}
		for i := range nums {
			nums[i] = dobrar(nums[i])
		}
		return nums, nil
	} else {
		return nums, fmt.Errorf("O slice está vazio")
	}
}
