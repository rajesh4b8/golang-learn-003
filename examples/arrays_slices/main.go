package main

import "fmt"

func main() {
	// create an array
	// var <array_name> [<size>]<data_type>
	// initialize all the elements with zero values
	var ai [10]int

	// var as [10]string
	// initilize with zero values
	as := [10]string{}

	fmt.Println(ai)
	fmt.Println(as)

	arrWithValues := [5]int{0, 1, 2, 3, 4}
	arrWithValues1 := [5]int{0: 10, 4: 30}
	fmt.Println(arrWithValues, arrWithValues1)

	arr := [10]int{1, 2, 3, 4, 5, 6, 76, 78, 8, 9}
	var slice1 []int = arr[0:3]

	fmt.Println("arr", arr, "slice1", slice1)
	arr[0] = -1
	fmt.Println("arr", arr, "slice1", slice1)

	// simple slice from array from index 0 to index 2
	slice1 = arr[1:3]
	fmt.Println("Lenth:", len(slice1), "Cotenents:", slice1, "Capacity:", cap(slice1))

	//  slice from array from index 0 to index 2
	slice1 = arr[:3]
	fmt.Println("Lenth:", len(slice1), "Cotenents:", slice1, "Capacity:", cap(slice1))

	//  slice from array from index 3 to cap of array
	slice1 = arr[3:]
	fmt.Println("Lenth:", len(slice1), "Cotenents:", slice1, "Capacity:", cap(slice1))

	//  slice from array from index 0 to cap of array
	slice1 = arr[:]
	fmt.Println("Lenth:", len(slice1), "Cotenents:", slice1, "Capacity:", cap(slice1))

	// defining any value outside of cap is going to cause error
	// slice1[10] = 20
	slice1 = append(slice1, 20)
	fmt.Println("Lenth:", len(slice1), "Cotenents:", slice1, "Capacity:", cap(slice1))
	slice1 = append(slice1, 22)
	fmt.Println("Lenth:", len(slice1), "Cotenents:", slice1, "Capacity:", cap(slice1))

	// practical usages
	s1 := make([]int, 10)
	fmt.Println("Lenth:", len(s1), "Cotenents:", s1, "Capacity:", cap(s1))

	s1 = []int{}
	fmt.Println("Lenth:", len(s1), "Cotenents:", s1, "Capacity:", cap(s1))

	s1 = []int{0, 1, 1, 2, 3, 5, 8, 11}
	fmt.Println("Lenth:", len(s1), "Cotenents:", s1, "Capacity:", cap(s1))
}
