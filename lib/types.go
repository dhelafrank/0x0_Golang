package lib

import (
	"fmt"
	"strings"
)

// STRINGS
var name = "John"

func StringTypes() {
	// fmt.Println(len(name))     //Use the built in len function to get the length of a string
	// fmt.Println(name[0:2])     //access individual characters using sqaured brackets and the character index you want to access
	// fmt.Println(name + " Doe") //concatenate strings using the + operator

	// var nameCopy = name[:] //Create a copy of the string
	// fmt.Println(nameCopy)

	// var second = nameCopy //Assign a string to a new variable
	// fmt.Println(second)

	// nameCopy = "Immutable" //Strings are immutable
	// fmt.Println(nameCopy)
	// fmt.Println(second) //Still give John

	//The Strings Package

	fmt.Println(strings.ToUpper(name))                //returns a new string, uppercase
	// fmt.Println(strings.ToLower(name))                //returns a new string, lowercase
	// fmt.Println(strings.HasSuffix(name, "hn"))        //checks if a string ends with a substring
	// fmt.Println(strings.HasPrefix(name, "joh"))       //checks if a string starts with a substring
	// fmt.Println(strings.Contains(name, "oh"))         //checks if a string contains a substring
	// fmt.Println(strings.Count(name, "oh"))            //counts how many times a substring appears in a string
	// fmt.Println(strings.Split(name, ","))             //used to create an array of strings from a string, dividing the original one on a specific character, like a comma or a space
	// fmt.Println(strings.ReplaceAll(name, "oh", "ho")) //used to replace a portion in a string and replace it with a new one

	fmt.Println("fmt usage check")
}
