package example

import "fmt"

func examArr() {
	fmt.Println("hello, world!")
	
	arr1 := []int{2, 3, 4, 5}
	fmt.Println(arr1)

	arr2 := make([]int, 4)
	arr2[0] = 30
	fmt.Println(arr2)

	txt := "today is sunday"
	// slice [ start index : end index ]
	fmt.Println(txt[0:5])
	fmt.Println(txt[0:1])
	fmt.Println(len(txt))
}