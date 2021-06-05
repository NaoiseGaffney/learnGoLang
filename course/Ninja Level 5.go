package main

import "fmt"

func main() {
	Person()
	Vehicle()
	AnonStruct()
}

func Person() {
	type person struct {
		first     string
		last      string
		age       int
		icFlavour []string
	}
	fmt.Println("\n--- Ninja Level 5.1 --- Struct ---")
	p1 := person{
		first: "Naoise",
		last:  "Gaffney",
		age:   51,
		icFlavour: []string{
			"Pistachio", "Liquorice", "Rum and Raisin",
		},
	}
	p2 := person{
		first: "Fiona",
		last:  "Gaffney",
		age:   51,
		icFlavour: []string{
			"Raspberry",
		},
	}

	fmt.Println("Person:", p1)
	fmt.Printf("Type: %T\n", p1)
	fmt.Println("Person:", p1.first, p1.last)
	fmt.Println("Person:", p2)
	fmt.Println("Person:", p2.first, p2.last)

	for k, v := range p1.icFlavour {
		fmt.Println("RANGE - Key:", k, "\tValue:", v)
	}

	for i := 0; i < len(p1.icFlavour); i++ {
		fmt.Println("FOR LOOP - Index:", i, "\tValue:", p1.icFlavour[i])
	}

	fmt.Println("\n--- Ninja Level 5.2 --- Map ---")
	// Using first as an index as both last names are the same and only the last value is used.
	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	fmt.Println("Printing Map m:", m)

	for k2, v2 := range m {
		fmt.Println("\nRANGE of MAP")
		fmt.Println("MAP - Key:", k2, "\tValue:", v2)
		fmt.Println(v2.first, v2.last)
		for icf, val := range v2.icFlavour {
			fmt.Println("Index:", icf, "\tValue:", val)
		}
	}
}

func Vehicle() {
	type vehicle struct {
		door   int
		colour string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type car struct {
		vehicle
		luxury bool
	}

	fmt.Println("\n--- Ninja Level 5.3 --- Nested Structs ---")
	t := truck{
		vehicle: vehicle{
			door:   2,
			colour: "Red",
		},
		fourWheel: true,
	}
	fmt.Println(t)
	fmt.Println("Door:", t.door, "\tColour:", t.colour, "\t4 x 4:", t.fourWheel)

	c := car{
		vehicle: vehicle{
			door:   4,
			colour: "Pink",
		},
		luxury: true,
	}
	fmt.Println(c)
	fmt.Println("Door:", c.door, "\tColour:", c.colour, "\tLuxury:", c.luxury)
}

func AnonStruct() {
	fmt.Println("\n--- Ninja Level 5.4 --- Anonymous Struct ---")

	s := struct {
		first     string
		last      string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Naoise",
		last:  "Gaffney",
		friends: map[string]int{
			"Mark": 123,
			"Jim":  321,
			"Ed":   456,
		},
		favDrinks: []string{
			"Gin & Tonic", "Whiskey", "Ale",
		},
	}
	fmt.Println(s)

	fmt.Println(s.first, s.last)
	for k2, v2 := range s.friends {
		fmt.Println("Index:", k2, "\tValue:", v2)
	}
	for k3, v3 := range s.favDrinks {
		fmt.Println("index:", k3, "\tValue", v3)
	}
	fmt.Println(s.first, s.last, s.friends, s.favDrinks)
}
