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

	fmt.Println("A média é igual a", media(nums))
}

func media(nums []int) int {
	var soma int
	for i := range nums {
		soma += nums[i]
	}
	med := soma / len(nums)
	return med
}
