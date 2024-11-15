package main

import (
	"fmt"
)


func main() {
	a := [4]int{0,1,2}
	a[1] = 1
	
	s := a[:3]
	fmt.Println(s)

	a[0] = 12
	s[0] = 21
	fmt.Println("PreAppend: ", s, a)
	
	s = append( s, 1 )
	
	fmt.Println(s, a)

	a[0] = 12
	s[0] = 21
	fmt.Println(s, a)
}
