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

	// const pi = math.Pi
	// fmt.Printf("Var type %T - value %#v\n", pi, pi)

	// s := sum(1, 2, 3, 4, 5, 6, 7)
	// fmt.Printf("Var type %T - value %#v\n", s, s)

	// name := customer.Name
	// fmt.Printf("Var type %T - value %#v\n", name, name)

	var x int
	var y *int
	x = 100
	y = &x
	fmt.Printf("Var type %T - value %#v\n", x, x)
	fmt.Printf("Var type %T - value %#v\n", y, y)
	fmt.Printf("Var type %T - value %#v\n\n", *y, *y)
	x = 200
	fmt.Printf("Var type %T - value %#v\n", x, x)
	fmt.Printf("Var type %T - value %#v\n", y, y)
	fmt.Printf("Var type %T - value %#v\n\n", *y, *y)

	*y = 50
	fmt.Printf("Var type %T - value %#v\n", x, x)
	fmt.Printf("Var type %T - value %#v\n", y, y)
	fmt.Printf("Var type %T - value %#v\n\n", *y, *y)

}

// func sum(x ...int) int {
// 	s := 0
// 	for _, v := range x {
// 		s += v
// 	}
// 	return s
// }
