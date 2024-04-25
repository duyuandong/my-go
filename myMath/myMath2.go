package mathClass

import (
	"fmt"
	"strconv"
)

func Sub(x, y int) int {
	return x - y
}

func Sum(x string, y int) (int, int) {
	var k int
	var err error
	k, err = strconv.Atoi(x)
	fmt.Println(err)
	return k + y, y
}
