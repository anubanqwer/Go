package main

import "fmt"

type Rectangle struct {
	width float64
	length float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

type Triangle struct {
	width float64
	length float64
}

func (r Triangle) Area() float64 {
	return r.width * r.length * 0.5
}


type Object interface {
	Area() float64
}

func printArea(Obj Object) {
	fmt.Println(Obj.Area())
}

func main(){

	//interface
	var i interface{}

	i = 10
	fmt.Printf("type is %T, value is %v\n", i, i)

	var n int
	//Failed because go is static type and i is still type interface
	//n = i
	n = i.(int)
	fmt.Printf("This is n! => type is %T, value is %v\n", n, n)

	i = "ten"
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = struct {
		number int
		text string
	} {
		number: 10,
		text: "ten",
	}
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = func() string {
		return "10"
	}
	fmt.Printf("type is %T, value is %v\n", i, i)

	r := Rectangle{3, 4}
	t := Triangle{3, 4}
	// fmt.Printf("Print Area Rectangle %f", printArea(r))
	// fmt.Printf("Print Area Triangle %f", printArea(t))
	//Same signature (Object, Rectangle and Triangle have method .Area() )
	printArea(r)
	printArea(t)
	
}