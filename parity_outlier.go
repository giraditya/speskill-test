package main

func Find(numberOfArr []int) (int, bool) {
	odds := []int{}
	evens := []int{}
	for _, v := range numberOfArr {
		if v%2 == 0 {
			odds = append(odds, v)
		} else {
			evens = append(evens, v)
		}
	}

	if len(evens) == 1 {
		return evens[0], true
	} else if len(odds) == 1 {
		return odds[0], true
	} else {
		return 0, false
	}
}
