package grains

import (
	"fmt"
	"math"
)

var mymap map[int]uint64
var total uint64

// runs first
func init() {

	mymap = map[int]uint64{
		1: 1,
		2: 2,
	}

	sum := uint64(0)
	for i := 0; i < 64; i++ {

		cell := uint64(math.Pow(2, float64(i)))
		//fmt.Println("cell", i, cell)
		mymap[i+1] = uint64(cell)
		//fmt.Println("mymap", i+1, mymap[i+1])
		sum += cell
	}
	total = sum
	//fmt.Println("mymap", mymap)
	//fmt.Println("total", total)
}

func Square(input int) (uint64, error) {

	// 1 -> 1
	// 2 -> 2
	// 64 -> 18446744073709551616

	if input < 1 || input > 64 {
		return 0, fmt.Errorf("outside 1-64")
	}

	value, ok := mymap[input]
	fmt.Println("Square", input, value, ok)

	return value, nil
}

func Total() uint64 {
	return total
}
