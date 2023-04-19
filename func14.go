package main

import "fmt"

func main() {
	var nums1 = []int{}
	var tam int
	fmt.Print("Tamanho do primeiro slice: ")
	fmt.Scan(&tam)

	for i := 0; i < tam; i++ {
		var valor int
		fmt.Printf("valor na posição %d: ", i)
		fmt.Scan(&valor)
		nums1 = append(nums1, valor)
	}

	var nums2 = []int{}
	var tam2 int
	fmt.Print("Tamanho do segundo slice: ")
	fmt.Scan(&tam2)

	for i := 0; i < tam2; i++ {
		var valor int
		fmt.Printf("valor na posição %d: ", i)
		fmt.Scan(&valor)
		nums2 = append(nums2, valor)
	}
	var com, err = comparar(nums1, nums2)

	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println(com)
	}
}

func comparar(nums1, nums2 []int) ([]int, error) {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}, fmt.Errorf("Slice vazio.")
	} else {
		var com = []int{}
		for i := range nums1 {
			x := nums1[i]
			for p := range nums2 {
				if nums2[p] == x {
					com = append(com, x)
					break
				}
			}
		}
		return com, nil
	}
}
