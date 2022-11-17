package main

import (
	"fmt"
	"testing"
)

func TestBlueOcean(t *testing.T) {
	arrOfInt := []int{1, 2, 3, 4, 6, 10}
	removeList := 1

	res := BlueOcean(arrOfInt, removeList)
	fmt.Println(res)
}
