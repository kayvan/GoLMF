package main

import (
	"fmt"
)
// slice is a dynamic array we let size empty and can be resized

func sli() {
	fmt.Println("Slice!")
	var nums []int
	fmt.Println(nums) // returns []
	fmt.Printf("memaddress: %p, len: %d, cap: %d \n", &nums, len(nums), cap(nums)) // returns memaddress: 0xc000008048, len: 0, cap: 0

	nums = append(nums, 1, 2, 3, 4, 5)
	fmt.Println(nums) // returns [1 2 3 4 5]
	fmt.Printf("memaddress: %p, len: %d, cap: %d \n", &nums, len(nums), cap(nums)) // returns memaddress: 0xc000008048, len: 5, cap: 5

	



}

