/*
A simple demo program that serializes structure into the JSON


type ShirtSize byte

const (
    NA ShirtSize = iota
    XS
    S
    M
    L
    XL
)

type Person struct {
    Name string
    Born time.Time
    Size ShirtSize
}
Let’s try to serialize those structures above into the JSON below.

{
    "name": "Gopher",
    "birthdate": "2009/11/10",
    "shirt-size": "XS"
}
*/

package main

import (
	"fmt"
	"time"
)

type ShirtSize byte

const (
    NA ShirtSize = iota
    XS
    S
    M
    L
    XL
)

type Person struct {
    Name string
    Born time.Time
    Size ShirtSize
}

func main() {

// Add some test data
	TestDataPerson := Person{Name: "Ivan", 
							Born: time.Date(1989, 11, 10, 0, 0, 0, 0, time.UTC), 
							Size: XL}

var t = &TestDataPerson

// Switch for the shirt size
var jsonShirt string;							
	switch (t.Size) {
	case XS:
		jsonShirt = "XS"
	case S:
		jsonShirt = "S"
	case M:
		jsonShirt = "M"
    case L:
    	jsonShirt = "L"
    case XL:
    	jsonShirt = "XL"
	default:
		jsonShirt = "N/A"    	
	}

// Cast birth date to jsonBirth

	jsonBirth := fmt.Sprintf("%d/%2d/%2d", t.Born.Year(), t.Born.Month(), t.Born.Day())

// Display the list of people in the required JSON format
	
	fmt.Println(`{
    "name": "`+t.Name+`",
    "birthdate": "`+jsonBirth+`",
    "shirt-size": "`+jsonShirt+`"
}`)	

}