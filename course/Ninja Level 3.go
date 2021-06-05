package main

import (
	"fmt"
	"time"
)

func main() {
	Print10K()
	PrintRune()
	YearsAlive()
	ModOf10to100()
	IfElseIf()
	SwitchChoices()
	SwitchChoicesCondition()
	TrueFalseStatements()
}

func Print10K() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func PrintRune() {
	for i := 33; i <= 122; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func YearsAlive() {
	birthday := 1969
	year, _, _ := time.Now().Date()
	yearsAlive := year - birthday

	for i := 0; i <= yearsAlive; i++ {
		fmt.Println(birthday + i)
	}
}

func ModOf10to100() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i, ": ", i/4, i%4)
	}
}

func IfElseIf() {
	x := "More"

	if x == "Stuff" {
		fmt.Println("X equals 'Stuff'")
	} else if x == "More Stuff" {
		fmt.Println("X equals 'More Stuff'")
	} else {
		fmt.Println("X equals nothing...")
	}
}

func SwitchChoices() {
	switch {
	case false:
		fmt.Println("Should NOT print!")
	case true:
		fmt.Println("Should print!")
	}
}

func SwitchChoicesCondition() {
	favSport := "Rugby"
	switch favSport {
	case "Football":
		fmt.Println("Football")
	case "Swimming":
		fmt.Println("Swimming")
	case "Paragliding":
		fmt.Println("Paragliding")
	case "Chess", "Rugby":
		fmt.Println("Rugby, with fallthrough")
		fallthrough
	default:
		fmt.Println("Nothing -> Default")
	}
}

func TrueFalseStatements() {
	fmt.Println("true || true:\t", true || true)
	fmt.Println("true || false:\t", true || false)
	fmt.Println("false || false:\t", false || false)
	fmt.Println("true && true:\t", true && true)
	fmt.Println("true && false:\t", true && false)
	fmt.Println("false && false:\t", false && false)
	fmt.Println("!true:\t\t", !true)
	fmt.Println("!false:\t\t", !false)
}
