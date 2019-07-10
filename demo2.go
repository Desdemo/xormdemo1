package main

import "fmt"

type Un struct {
	Name string
	Age int
}

var Uns []Un

func main() {
	var u3 ,u4 []Un
	u1 := Un{Name:"sss", Age:11}
	u2 := Un{Name:"ddd", Age:22}
	Uns = append(Uns, u1)
	Uns = append(Uns, u2)
	u3 = Uns
	u4 = Uns
	u3 = append(u3, u4)
	fmt.Println(u2, Uns)
}
