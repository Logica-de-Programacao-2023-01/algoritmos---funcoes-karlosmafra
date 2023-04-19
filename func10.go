package main

import (
	"fmt"
	"sort"
)

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

	var cres, err = crescente(nums)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(cres)
	}
}

func crescente(nums []int) ([]int, error) {
	if len(nums) != 0 {
		sort.Ints(nums)
		return nums, nil
	} else {
		return nums, fmt.Errorf("O slice está vazio")
	}
}
