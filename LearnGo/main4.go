package main

import ("fmt"
		// "golang.org/x/tour/pic"
)

//global declaration
type Int int

func (i Int) Set1(n int) {
	//This doesn't work because both i and n are local variable
	//i in this scope (function) is change but i in the main function doesn't affect
	i = Int(n)
}

func (i *Int) Set2(n int) {
	*i = Int(n)
}

//struct
type Rectangle struct {
	width float64
	length float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func xfunc(n int){
	//defer -> store in stacks
	//remember to execute func()
	defer func() { fmt.Println(n) }()
	defer fmt.Println(n)
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	n *= n
	fmt.Println(n)
}

func catchMe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	s := []int{}
	fmt.Println(s[1])
}

func main(){
	
	// var p = new(int)
	var p *int
	i := 42
	p = &i
	fmt.Println(p, *p, i)
	*p = 43
	fmt.Println(p, *p, i)

	//type newType T
	//new Type
	type text string

	//text is alias name of string (still the same type, can compare)
	// type text = string 

	var j Int
	fmt.Printf("%d\n", j)

	j.Set1(10)
	fmt.Println("After 1st try")
	fmt.Printf("%d\n", j)

	j.Set2(10)
	fmt.Println("After 2nd try")
	fmt.Printf("%d\n", j)

	rec := Rectangle{width: 3, length:5}
	fmt.Println(rec.Area())

	//defer is doing it before returning in function
	xfunc(20)

	catchMe()


	
}