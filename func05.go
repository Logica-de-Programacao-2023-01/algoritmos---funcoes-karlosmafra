package main

import "fmt"

func main() {
	var nums = []int{1, 4, 7, 8, 12, 21, 18, 33, 27}
	var x int
	fmt.Print("Buscar número: ")
	fmt.Scan(&x)

	fmt.Println(nums)
	fmt.Printf("O número %d está na posição %d", x, buscar(nums, x))
}

func buscar(nums []int, x int) int {
	pos := -1
	for i := range nums {
		if nums[i] == x {
			pos = i
			break
		}
	}
	return pos
}
