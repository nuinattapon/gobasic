package main

import "fmt"

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
	// fmt.Printf("%T %#v\n", employees, employees)

	var name string
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Printf("%T %#v\n", err, err)
	} else {
		fmt.Printf("%T %#v\n", name, name)
		fmt.Printf("%T %#v\n", &name, &name)
		if name == "Nattapon" {
			fmt.Printf("Hello %s\n", name)
		}
	}
}
