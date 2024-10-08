package main

import (
	"fmt"
	"reflect"
)
// data type: int uint float64 bool string rune and others
// more info about data types: https://cloud-book.gofarsi.ir/chapter-1/go-basic-types/


// more info: https://cloud-book.gofarsi.ir/chapter-1/go-variables-and-consts/
func varConst() { // we can use var or const (for immutable variables) to declare variables
	var num int = 25                        // simple declaration num := 25
	var nf float64 = 3.14                   // simple declaration nf := 3.14
	var bo bool = true                      // simple declaration bo := true
	var str string = "Hello string!"        // simple declaration str := "Hello string!"
	var x rune 								// rune is same as int32 or word in unicode
	x = 'a'                       			// simple declaration x := 'a'
	println(num, nf, bo, str, x) // here returns for rune the unicode of the character 97 which is 'a'
	
	var a, b, c, d = "hello", 6, 2.34, true // multiple declaration
	e, f, g, h := "world", 9, 5.67, false   // multiple declaration
	fmt.Println(a, b, c, d, e, f, g, h)
	
	const i int = 100          // constant declaration
	const j, k = "blabla", 0.5 // multiple constant declaration
	println(i, j, k)
	
	fmt.Println(&num) // returns pointer or memory address of num

	ref := reflect.ValueOf(str) // reflect package is used to get some infos of variable like type and kind 
	fmt.Println(ref.Type()) 	// returns string
	fmt.Println(ref.Kind()) 	// returns string
	fmt.Println(ref.Type().Size()) // returns 16 -- length of string in bytes

}

// for using some functions like fmt or reflect we need to import them -- normally vscode imports them for us when we use keywords like fmt and reflect

