package main

import "fmt"

func add(a, b int) int {
	return a+b
}

/*
No library

func checkString(inputVar string) {
	if n, err:= strconv.Atoi(inputVar); err != nil {
		log.Println(err)
	} else {
		log.Printlnf("%d\n")
	}
}
*/

func main(){

	//Type inference
	var a = 9
	//Can assign like this in function only
	b := "hello"
	fmt.Println(a, b, add(3, 4))

	//While loop
	i:= 0
	for i < 10 {
		fmt.Println(i)
		i += 5
	}
	
}