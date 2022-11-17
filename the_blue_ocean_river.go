package main

func BlueOcean(arrOfInt []int, removeList int) []int {
	tempArr := []int{}
	for i := 0; i < len(arrOfInt); i++ {
		if arrOfInt[i] != removeList {
			tempArr = append(tempArr, arrOfInt[i])
		}
	}
	return tempArr
}
