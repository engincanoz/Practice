package maths

import (
	"fmt"
	"math"
)

func GetVectors() ([]int, []int) {
	var throughx, throughy, throughz int
	var parallelx, parallely, parallelz int

	fmt.Print("Enter the point your line pases through: ")
	_, err := fmt.Scan(&throughx, &throughy, &throughz)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter the point your line is parallel to: ")
	_, err = fmt.Scan(&parallelx, &parallely, &parallelz)
	if err != nil {
		fmt.Println(err)
	}
	xt := []int{throughx, throughy, throughz}
	xp := []int{parallelx, parallely, parallelz}
	return xt, xp

}

func ParametricGenerator(xt, xp []int) (string, string, string) {
	var strx, stry, strz string
	for i := 0; i < 3; i++ {
		if xp[i] < 0 {
			switch i {
			case 0:
				strx = fmt.Sprintln("x =", xt[0], "-", (math.Abs(float64(xp[0]))), "t")
			case 1:
				stry = fmt.Sprintln("y =", xt[1], "-", (math.Abs(float64(xp[1]))), "t")
			case 2:
				strz = fmt.Sprintln("z =", xt[2], "-", (math.Abs(float64(xp[2]))), "t")

			}
		}

		if xp[i] > 0 {

			switch i {
			case 0:
				strx = fmt.Sprintln("x =", xt[0], "+", xp[0], "t")
			case 1:
				stry = fmt.Sprintln("y =", xt[1], "+", xp[1], "t")
			case 2:
				strz = fmt.Sprintln("z =", xt[2], "+", xp[2], "t")

			}
		}
		if xp[i] == 0 {
			switch i {
			case 0:
				strx = fmt.Sprintln("x =", xt[0])
			case 1:
				stry = fmt.Sprintln("y =", xt[1])
			case 2:
				strz = fmt.Sprintln("z =", xt[2])

			}

		}
	}
	return strx, stry, strz
}
