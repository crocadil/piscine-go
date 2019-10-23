package piscine

func Index(s string, toFind string) int {
	q := []rune(s)
	o := []rune(toFind)
	count1 := 0
	count2 := 0

	for range q {
		count1++
	}
	for range o {
		count2++
	}
	if count2 == 0 {
		return 0
	}

	for index, value := range q {
		if value == o[0] && count1 >= count2+index-1 {
			creative := 1
			for i := 1; i < count2; i++ {
				if o[i] == q[index+i] {
					creative++
				}
			}
			if creative == count2 {
				return index
			}
		}
	}
	return -1
}
