package main

import ("fmt"
		// "golang.org/x/tour/pic"
)

func variadic(nums ...int){
	//nums is slice or pointer
	fmt.Printf("%#v\n", nums)
}

func main(){
	variadic(1, 3, 4)
	a := [...]int{5, 6}
	b := a[:]
	//spread operator = take each item in Slice as a function argument
	variadic(b...)

	//MAP (No Ordering) -> pointer
	// var m map[string]string
	m := map[string]string{
		"a": "apple",
		"b": "batman",
	}
	m["z"] = "zero"
	delete(m, "a")
	for k, v := range m {
		fmt.Println(k, v)
	}
	

	
}