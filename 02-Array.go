package main

import (
	"fmt"
)

// in go array and struct have fixed size and can't be resized. on the other hand slice and map have dynamic size and can be resized.

func arr() {
	fmt.Println("Array!")
	// var arInt [5]int = [5]int{1, 2, 3, 4, 5}
	arInt := [5]int{1, 2, 3, 4, 5} // fixed int array with 5 elements	
	arSrt := [3]string{"a", "b", "c"}
	fmt.Println(arInt , arSrt)
	fmt.Printf("Array arSrt: %v, length arInt: %d, capacity arSrt: %d \n", arSrt, len(arInt), cap(arSrt)) // returns Array arSrt: [a b c], length arInt: 5, capacity arSrt: 3
	fmt.Printf("arInt type: %T, array memAddress: %p, second element memAddress: %p \n", arInt, &arInt ,&arInt[1]) // returns arInt type: [5]int, array memAddress: 0xc00000e3f0, second element memAddress: 0xc00000e3f8
	// in printf we can use %v for value and %d for length and capacity. %T for type and %p for memory address
	// in type we see the size included and that means [5]int is different from [6]int

	arr1 := [5]int{} // all	elements will be 0
	arr2int := [6]int{1, 2, 3} // example: if we have 6 elements and give 3 values it will give 0 for the remaining elements in int array
	arr2str := [10]string{"a", "b", "c"} // the same for string for 10 elements and 3 values it will be empty string
	fmt.Println(arr1, arr2int, arr2str) // returns [0 0 0 0 0] [1 2 3 0 0 0] [a b c       ]
	arr3 := [5]int{1: 20, 4: 40} // we can also give values to specific index in the array. in this case we will give 20 to index 1 and 40 to index 4. in this case we will get [0 20 0 40 0}
	arr4 := [...]string{"hello", "world"} // if we give ... in the array it fix the size of the array with the last value index. in this case size =2
	//arr4[2] = "goodbye" // it will return an error because we can't change the size of the array.
	arr5 := [...]int{1: 15, 9: 50} // an array with size 10 and 2 elements with values 15 and 50
	arr5[0]	= 10 // /add/change value of first element
	fmt.Println(arr3, arr4, arr5)
	
	// using for range with array
	for range arr4 {
		fmt.Println("haha") // print haha 2 times
	}
	for index, value := range arr3 { // range have 2 values index and value and it returns first index and then value
		fmt.Println("index: ", index, " value: ", value)
	}

	for  _, value := range arr3 { // in this case we don't need index and we used _ to ignore it. if we don't use _ it will print index instead of value
		fmt.Println( "value: ", value) 
	}
}

