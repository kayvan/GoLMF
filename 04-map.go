package main

import (
	"fmt"
	"maps"
)
// Map is a built-in data structure for storing key-value pairs. -- something like dictionary in python
func ma() {	
	fmt.Println("map!")
	var m0 = map[string]int{"aa": 1, "bb": 2, "cc": 3} // declare map with key and value m:= map[key type]value type{}
	var m01 map[int]string // declare a nil map
	fmt.Println(m0, m01)
	m1 := make(map[string]int) // declare map using make function -- other way to declare: m1 := map[key type]value type{}
	
	// crud on map: create read update delete
	m1["a1"] = 10 // adding item
	m1["a2"] = 20
	m1["a3"] = 30
	m1["a4"] = 40
	m1["a5"] = 50
	fmt.Println(m1["a3"]) // read an item
	xx := m1["a2"] // assign value to a variable 
	yy := m1["a8"] // If the key doesn’t exist, the zero value of the value type is returned.
	m1["a2"] = 22 // update an item
	m1["a5"] = xx + m1["a5"] // update an item by adding values in variable and key 
	delete(m1, "a4") // delete an item
	fmt.Println(m1,"xx:", xx, "yy:", yy, "len-m1:",  len(m1)) // returns map[a1:10 a2:22 a3:30 a5:70] xx: 20 yy: 0 len-m1: 4
	// len functions shows us the number of items in the map
	
	// clear map: to clear a map we can use clear function or assign it to nil
	m1 = nil // clearing by nil
	clear(m0) // clearing by clear function
	fmt.Println("m0: ", m0, "m1: ", m1, "len m1:",  len(m1)) // returns m0:  map[] m1:  map[] len m1: 0

	// check availability
	m2 := map[int]string{1: "a", 2: "b", 3: "c"}
	v1, prs1 := m2[1] // read an item and check availability
	_, prs2 := m2[4] // we use underscore to ignore the value and just check availability
	fmt.Println("m2[1]:", v1, "prs1:", prs1, "prs2:", prs2) // returns value and true for available key anf false for unavailable key

	// if and for loop
	m3 := map[int]string{1: "a", 2: "b", 3: "c"}
	if maps.Equal(m2, m3) { // using maps module to compare 2 maps
		fmt.Println("m2 and m3 are equal")
	}
	
	for k, v := range m3 {
		fmt.Printf("%d -> %s\n",k, v) // it returns key and value without (sometimes with) order
	}

	// assign map to map
	m4 := m3 // assign m3 to m4: in this form we are pointing m4 to memory address of m3 and nay change will be reflected in both m3 and m4
	m4[4] = "d" // adding an item to m4 and it will be in m3
	m3[1] = "x" // updating an item to m3 and it will be in m4
	fmt.Println("m3: ", m3, "m4: ", m4)
	// to avoid this edit problem on both maps we can use for loop (we should declare new map before loop) to navigate inside that map and put its values ​​in the new map.
	m5 := make(map[int]string)
	for a, b := range m3 {
		m5[a] = b
	}
	m5[1] = "y" // updating an item and it will be just in m5
	m3[5] = "z" // adding an item and it will be just in m3
	fmt.Println("m3: ", m3, "m5: ", m5) // returns m3:  map[1:x 2:b 3:c 4:d 5:z] m5:  map[1:y 2:b 3:c 4:d]
	// we can easily use maps module to clone a map instead of using for loop
	m6 := maps.Clone(m3) // using maps module to clone a map
	m3[2] = "h" // updating an item and it will be just in m3
	m6[6] = "g" // adding an item and it will be just in m6
	fmt.Println("m3: ", m3, "m6: ", m6) // returns m3:  map[1:x 2:h 3:c 4:d 5:z] m6:  map[1:x 2:b 3:c 4:d 5:z 6:g]

}



