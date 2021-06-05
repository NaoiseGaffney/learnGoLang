/*
Create your own type. Have the underlying type be an int.
create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
in func main
print out the value of the variable “x”
print out the type of the variable “x”
assign 42 to the VARIABLE “x” using the “=” OPERATOR
print out the value of the variable “x”
*/

package main

import "fmt"

type whiskey string

var w whiskey = "I Love Whiskey!"

func main() {
	fmt.Println(w)
	fmt.Printf("T", w)
	w = "I love Guinness too!"
	fmt.Println("\n", w)
}
