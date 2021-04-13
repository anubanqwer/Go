package main

import ("fmt"
		// "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	
	twoArray := make([][]uint8, dy)
	
	for i := range twoArray {
		oneArray := make([]uint8, dx)
		for j := range oneArray {
			oneArray[j] = uint8 ((i+j)/2)
		}			
		twoArray[i] = oneArray
	}
	return twoArray	
}

func main(){

	//mutable
	// var a [5]string
	//auto count when compile
	var b = [...]string{"a", "b", "c"}
	fmt.Printf("%#v\n", b)

	//just use index for passing unused variable (i)
	for _, v := range b {
		// _ = i
		fmt.Println(v)
	}

	//Slice is immutable
	//pointer, default value is nil
	var s []string
	if( s==nil ) {
		fmt.Println("YES")
	}
	
	s = b[0:2]
	// for i, v := range s {
	// 	fmt.Println(i, v)
	// }

	//Length is always <= Capacity
	// slice2 := make([]int, 0)
	// slice3 := []int{1, 2, 5}

	s = append(s,  "d")
	/*
		Append "d" into slice means change the value of 
		index 3 in real array (b) to "d"
	*/
	for i, v := range b {
		fmt.Println(i, v)
	}

	
}