package main

import (
	"fmt"
	"testing"
)

func TestNeedleInHaystack(t *testing.T) {
	arrOfString := []string{"red", "blue", "yellow", "black", "grey"}
	strToSearch := "blue"
	res, err := FindNeedle(arrOfString, strToSearch)
	fmt.Println(res, err)
}
