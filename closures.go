package main

func initSum() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := initSum()

	println(nextInt())
	println(nextInt())
	println(nextInt())

	newInts := initSum()
	println(newInts())
}
