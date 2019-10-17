package piscine
 

func PointOne(x *int) {

	*x = *x + 1

}

func main() {

	n := 0
	PointOne(&n)

	fmt.Println(n)

}
