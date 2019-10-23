package piscine

func TrimAtoi(s string) int {

	X := 0
	Y := 1

	for _, i := range s {
		if i >= '0' && i <= '9' {
			b := int(i - 48)
			X = X*10 + b 
		} else if i == '-' && X == 0 { //
			Y = -1 
		}

	}
	return X * Y
}
