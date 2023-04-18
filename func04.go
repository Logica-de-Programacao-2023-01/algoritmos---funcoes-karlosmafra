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

	fmt.Println("O segundo maior valor é", calcular2(nums))
}

func calcular2(nums []int) int {
	m1 := nums[0]
	m2 := nums[0]
	for i := range nums {
		if nums[i] > m1 {
			m1 = nums[i]
		}
	}
	for i := range nums {
		if nums[i] > m2 && nums[i] != m1 {
			m2 = nums[i]
		}
	}
	return m2
}
