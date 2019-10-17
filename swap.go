package piscine

func Swap(a *int, b *int) {
	a, b := a *int, b *int 
    b, a = a, b
    return []int{a, b}
}
