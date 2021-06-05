package main

import (
	"fmt"
)

// Arrays, Slices, Maps, and later on Structs
func main() {
	GoArray() // Array is rarely used in Go. Slices are used instead. Array is an underlying function for Slices.
	GoSlice()
}

func GoArray() {
	var x [5]int
	fmt.Println("No values assigned to Array: ", x)
	x[1] = 1
	x[2] = 2
	x[3] = 3
	x[4] = 4
	fmt.Println("Values assigned to Array: ", x)
}

func GoSlice() {
	// x := []type{values} --> a composite literal
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 42}
	fmt.Println(x)
	fmt.Println(x[5:7])
	fmt.Println(x[5:])
	fmt.Println(x[:5])
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println("\nFor Loop with Range, index and value")
	for i, v := range x {
		fmt.Println("Index: ", i, "\tValue: ", v)
	}

	fmt.Println("\nFor Loop with len(x), index and value")
	for i := 0; i < len(x); i++ {
		fmt.Println("Index: ", i, "\tValue: ", x[i])
	}

	y := []int{73, 56, 9, 85, 39}
	x = append(x, 34, 67, 93)
	fmt.Println("\nAppended to 'x': ", x)
	x = append(x, y...)
	fmt.Println("\nAppended 'y' to 'x': ", x)

	// Deleting from a slice using append()
	x = append(x[:4], x[8:]...)
	fmt.Println("Deleted from slice using append(): ", x)

	// Using make() to create a slice of a specific capcity (if size is known in advance)
	// z := make([]type, length, capacity of underlying Array
	z := make([]int, 10, 12)
	fmt.Println("\nSlice 'z' created using make(): ", z, "\n\tLength of 'z': ", len(z), "\n\tCapacity of 'z': ", cap(z))
	z = append(z, 11)
	fmt.Println("\nSlice 'z' created using make(): ", z, "\n\tLength of 'z': ", len(z), "\n\tCapacity of 'z': ", cap(z))
	z = append(z, 12)
	fmt.Println("\nSlice 'z' created using make(): ", z, "\n\tLength of 'z': ", len(z), "\n\tCapacity of 'z': ", cap(z))
	z = append(z, 13)
	fmt.Println("\nSlice 'z' created using make(): ", z, "\n\tLength of 'z': ", len(z), "\n\tCapacity of 'z': ", cap(z))

	// Multi-dimensional slices
	jb := []string{"James", "Bond", "Martini", "Walther PPK"}
	fmt.Println(jb)
	saint := []string{"Simon (John)", "Templar (Rossi)", "Martini", "Walther PPK"}
	fmt.Println(saint)
	mda := [][]string{jb, saint}
	fmt.Println(mda)

	// Map
	m := map[string]int{
		"James": 32,
		"Saint": 29,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["NotStoredInMapEqualsZero"])
	// 'Comma OK idiom'
	v, ok := m["NotStoredInMapEqualsZero"]
	fmt.Println("Value: ", v, "\tOK: ", ok)

	if v, ok := m["NotStoredInMapEqualsZero"]; ok {
		fmt.Print("This is OK, well, it's not OK and will not print", v)
	}

	if v, ok := m["Saint"]; ok {
		fmt.Println("This is OK, and will print", v)
	}

	// Add to map()
	m["Naoise"] = 51
	fmt.Println("Naoise's age is:", m["Naoise"])
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("Key: ", k, "\tValue: ", v)
	}

	// Add to and delete from a map()
	m["Fiona"] = 51
	fmt.Println(m["Fiona"])
	fmt.Println(m)
	delete(m, "Fiona")
	fmt.Println()
	v, ok = m["Fiona"]
	fmt.Println("Value:", v, "\tOK:", ok)
	fmt.Println(m)
}
