package piscine

func StrRev(s string) string {

	c := ""
	for range s {
		c = s + c
	}
	return c
}
