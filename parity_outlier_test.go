package main

import (
	"fmt"
	"testing"
)

func TestParityOutlier(t *testing.T) {
	numberOfArr := []int{11, 13, 15, 19, 9, 13, -21}
	res, err := Find(numberOfArr)
	fmt.Println(res, err)
}
