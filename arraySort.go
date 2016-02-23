/*
A simple demo program that sorts a collection of people by their age or their name as decided by the PersonSorter struct.
Author: Peter Sabev, 23.02.2016
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

type Person struct {
    Age int
    Name string
}

type PersonSorter struct {
    People []Person // List of Person
    By string // Sort by what?
}

// Get the length as number of elements
func (p PersonSorter) Len() int {
	return len(p.People)
}

// Compare elements
func (p PersonSorter) Less(i, j int) bool {
	if (strings.ToUpper(p.By)=="AGE") { // If AGE is set, sort by age
	return p.People[i].Age < p.People[j].Age;
	} else { // otherwise sort by name
	return p.People[i].Name < p.People[j].Name;	
	}

}

// Swap elements
func (p PersonSorter) Swap(i, j int) {
	p.People[i], p.People[j] = p.People[j], p.People[i]
}

// Main
func main() {

// Add some test data
	TestDataPeople := []Person{
		{Name: "Ivan", Age: 3},
		{Name: "Dragan", Age: 11},
		{Name: "Peter", Age: 33},
		{Name: "George", Age: 36},
		{Name: "Maria", Age: 20},
		}

// Assign test data
	pp := PersonSorter {
		People: TestDataPeople,
		By: "Name", // Sort by NAME or AGE
	}

// Sort pp
	sort.Sort(pp)

// Display the list of people
	for i, c := range pp.People {
		fmt.Println(i+1,":", c.Name, c.Age)
	}
}