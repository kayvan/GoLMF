package main

import (
	"fmt"
)

// slice is a dynamic array we let size empty and can be resized

func sli() {
	fmt.Println("Slice!")
	var nums []int // it creates an null/nil slice
	// nums := []int{} // this is not a null slice but it doesn't have any value
	var sl1 = make([]int, 3) // using make function to create a slice with default length 3 (values will be 0) 
	sl2 := make([]string, 2, 10) // using make function to create a slice with default length 2 and capacity 10 with value ""
	// in make capacity can't be less than length
	fmt.Println(nums, sl1, sl2) // returns []
	fmt.Printf("memaddress: %p, len: %d, cap: %d \n", &nums, len(nums), cap(nums)) // returns memaddress: 0xc000008048, len: 0, cap: 0

	//add modify elements
	nums = append(nums, 1, 2, 3, 4, 5) // add elements to the slice
	sl1 = append(sl1, 10, 20, 30, 40, 50) // add elements to the slice after all 3 zeros 
	sl1[1] = 100 // change value
	sl2 = append(sl2, "a", "bb", "ca", "d")
	sl2[1] = "x"
	fmt.Println(nums, sl1, sl2) // returns [1 2 3 4 5]
	fmt.Printf("memaddress: %p, len: %d, cap: %d \n", &nums, len(nums), cap(nums)) // returns memaddress: 0xc000008048, len: 5, cap: 5

	// slicing: create slice from array and it use the array as a backend; if we change a value in slice it will be changed in the array too
	// this changes will happen if we append to the slice till we reach the cap of array.
	// after appending more than cap of array, the slice will be copied in another address and doesn't affect the original array any more.
	var arr3 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sl3 = arr3[2:7] // it creates a slice from index 2 to 7 
	fmt.Printf("memarr3: %p, len: %d, cap: %d \n", &arr3, len(arr3), cap(arr3)) // 
	fmt.Println(sl3) // returns [3 4 5 6 7]
	fmt.Printf("memsl3: %p, len: %d, cap: %d \n", &sl3, len(sl3), cap(sl3)) // memory address is the address of the array[2]
	sl3[0] = 100 // change value in both slice[0] and array[2]
	arr3[4] = 200 // change value in both slice[2] and array[4]
	sl3 = append(sl3, 45, 55, 65)
	fmt.Println(arr3) // returns [1 2 100 4 200 6 7 45 55 65]
	fmt.Println(sl3) // returns [100 4 200 6 7 45 55 65]
	sl3 = append(sl3, 75) // after this append we are out of cap of array so it will be copied in another address
	sl3[1] = 400 // this change will be just in slice and not in array
	fmt.Println(arr3) // returns [1 2 100 4 200 6 7 45 55 65]
	fmt.Println(sl3) // returns [100 400 200 6 7 45 55 65 75]


}

