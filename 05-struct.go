package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type adad int // type alias for int -- we can use adad instead of int 
type user struct { // struct is somehow like a object in js. we define the fields of the struct here and give them values later in program.
	id int
	lname, fname string
	age int
}
// if we have another struct (person) with same fields like user, they are not same cus they are different types 
type person struct { // struct is somehow like a object in js. we define the fields of the struct here and give them values later in program.
	id int
	lname, fname string
	age int
}

// post struct for using json from jsonholder as na example
type Post struct { // we use pascal case for fields to use them externally
	UserId int 		`json:"userId"` // using json key for marshalling and unmarshalling
	Id int			`json:"id"`
	Title string 	`json:"title"`
	Body string 	`json:"body"`
}
// function for struct to have the length of the body
func (p Post) BodyLen() int { 
	return len(p.Body)
}

func stru() {
	fmt.Println("structs!")

	// adding an empty struct
	u0 := user{} // empty struct: all values are 0 or null -- another way to define an empty struct: var u0 user
	var u1 = user{id: 01, fname: "mamad", lname:  "mamadi", age: 20} // define with fields name -- if we don't use fields name it will use the order of struct
	var u2 = user{02, "akbari", "akbar", 66} // follows the order of struct for fields -> id-lname-fname-age
	var u3 = user{03, "khali", "", 0} // we can leave some fields empty, but we need to use the order of struct and null values
	var u4 = user{age: 66, fname: "bifamil"} // if we use fields name we can leave some fields empty without order
	fmt.Println(u0, u1, u2, u3, u4)

	// anonymous struct
	var u5 = struct { // in anonymous struct we don't have type (noname type) and we define the values after struct
		id int
		lname, fname string
		age int
	}{id: 05, fname: "hasan", lname:  "hasani", age: 19} 
	fmt.Println(u5)

	// implicit conversion from anonymous struct to struct -- here we use u6 as person and give u5 values to u6
	// we can't use u4 cus it has user type and u6 has person type
	var u6 person = u5 // in this case we can convert anonymous struct to struct when both have the same fields 
	fmt.Println(u6)

	// json example from jsonholder 
	res, _ := http.Get("https://jsonplaceholder.typicode.com/posts") // get json from url -- ignore error with _ -- we need to import "net/http" package
	defer res.Body.Close() 
	body, _ := io.ReadAll(res.Body) // read json and assign to body -- we need to import "io" package
	var posts []Post
	json.Unmarshal(body, &posts) // convert json to struct -- we need to import "encoding/json" package
	fmt.Println(len(posts), posts[4].Id)
	
	for _, post := range posts { // using for range to go through all posts and print some infos
		fmt.Printf("Post ID: %d Title: %s Len: %d\n", post.Id, post.Title, post.BodyLen())
	}
	
}



