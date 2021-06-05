/*
func (r receiver) identifier(parameters) (returns(s)) { code }
know the difference between parameters and arguments
we define our func with parameters (if any)
we call our func and pass in arguments (in any)
Everything in Go is PASS BY VALUE

Keyword Identifier Type
POLYMORPHISM: a VALUE can be of more than one TYPE.
*/

package main

import "fmt"

func main() {
	EmptyFunc()
	ArgumentFunc("This is an argument, that becomes a function parameter.")

	s1, s2 := ArgumentFuncReturn("James", "Bond")
	fmt.Println("Message:", s1, "\tMessage Length:", s2)

	sum := UnlimitedInts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Total:", sum)

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	x := UnfurlSlice(xi...)
	fmt.Println("Total of 'xi' is:", x)

	DeferFunc() // Execute Foo() at the end of the DeferFunc() block

	MethodReceiver() // Interfaces and Polymorphism

	// Anonymous Function
	func(x int) {
		fmt.Println("Anonymous Function 'func()', the meaning of life is", x)
	}(42)

	// Function Expressions
	f := func() {
		fmt.Println("Function Expression")
	}
	f()

	g := func(x int) {
		fmt.Println("I am", x, "years young.")
	}
	g(51)

	// Returning a String
	string1 := string1()
	fmt.Println(string1)

	// Returning a Function
	func1 := func1()
	fmt.Println("func1 variable:", func1)
	fmt.Println("func1():", func1())
	fmt.Printf("%v\n", func1())
	fmt.Printf("Type 'func1': %T\n", func1)
	fmt.Printf("Type 'func1()': %T\n", func1())

	// CallBack
	CallBackFunc()

	// Closure
	ClosureExample()

	// Recursion (USE LOOPS instead of recursion)
	// fmt.Println(factorial(10))
	factNum := factorial(10)
	fmt.Println(factNum)

	// Using a FOR LOOP instead of recursion
	factNum2 := factorialLoop(10)
	fmt.Println(factNum2)
}

func EmptyFunc() {
	fmt.Println("This is a simple Go 'func'")
}

func ArgumentFunc(s string) {
	fmt.Println(s)
}

func ArgumentFuncReturn(s string, t string) (string, int) {
	a := fmt.Sprint("Hello from ", s, " ", t)
	b := len(s + t)
	return a, b

}

func UnlimitedInts(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, sum)
	}
	return sum
}

func UnfurlSlice(xi ...int) int {
	sum := 0
	for i, v := range xi {
		sum += v
		fmt.Println(i, v, sum)
	}
	return sum
}

func DeferFunc() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Foo")
}
func bar() {
	fmt.Println("Bar")
}

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func humanPerson(h human) {
	fmt.Println("I was passed into 'humanPerson'", h)
}

func MethodReceiver() {

	sa1 := secretAgent{
		person: person{
			"Naoise",
			"Gaffney",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Fiona",
			"Gaffney",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(sa2)
	sa2.speak()

	p1 := person{
		first: "Dr",
		last:  "No",
	}
	fmt.Println(p1)

	humanPerson(sa1)
	humanPerson(sa2)
	humanPerson(p1)
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak()")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak()")
}

func string1() string {
	s := "Hello, World!"
	return s
}

func func1() func() int {
	return func() int {
		return 451
	}
}

func CallBackFunc() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("odd numbers", s3)
}

func sum(xi ...int) int {
	// fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func ClosureExample() {
	a := increment()
	b := increment()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialLoop(n int) int {
	fact := 1
	if n < 0 {
		fmt.Println("Error! Factorial of a negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			fact *= i
			fmt.Println(fact, i, n)
		}
	}
	return fact
}
