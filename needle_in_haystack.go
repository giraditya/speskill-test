package main

func FindNeedle(arrOfStr []string, strToSearch string) (int, bool) {
	for i := 0; i < len(arrOfStr); i++ {
		if arrOfStr[i] == strToSearch {
			return i, true
		}
	}
	return 0, false
}
