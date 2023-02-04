package main

import (
	"fmt"
	"math"
)

func main() {

	// type X struct {
	// 	name string
	// 	next *X
	// }
	// x := X{name: "one",
	// 	next: &X{name: "two",
	// 		next: &X{name: "three", next: nil}}}
	// fmt.Println(x.name)
	// fmt.Println(x.next.name)
	// fmt.Println(x.next.next.name)
	// fmt.Println(x.next.next.next)

	// type Employee struct {
	// 	name string
	// }
	// employees := []Employee{}
	// employees = append(employees, Employee{name: "Nattapon"})
	// fmt.Printf("Var type: %T - value: %#v\n", employees, employees)

	// var name string
	// fmt.Print("Enter your name: ")
	// _, err := fmt.Scanf("%s", &name)
	// if err != nil {
	// 	fmt.Printf("%T - %#v\n", err, err)
	// } else {
	// 	fmt.Printf("%T - %#v\n", name, name)
	// 	if name == "Nattapon" {
	// 		fmt.Printf("Hello %s\n", name)
	// 	}
	// }

	// i := 1234
	// s := strconv.Itoa(i)
	// fmt.Printf("Var type %T - value %#v\n", s, s)
	// s2 := "ABCD"
	// var err error
	// i, err = strconv.Atoi(s2)
	// fmt.Printf("Var type %T - value %#v\n", s2, s2)
	// fmt.Printf("Var type %T - value %#v\n", err, err)
	// fmt.Printf("Var type %T - value %#v\n", i, i)

	const pi = math.Pi
	fmt.Printf("Var type %T - value %#v\n", pi, pi)

}
