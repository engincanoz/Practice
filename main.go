package main

import (
	"fmt"

	maths "github.com/engincanoz/Practice/maths/planeAndSpace"
)

func main() {
	xt, xp := maths.GetVectors()
	fmt.Println(xt)
	fmt.Println(xp)
	s1, s2, s3 := maths.ParametricGenerator(xt, xp)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
