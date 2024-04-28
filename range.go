package main

func main() {
	nums := []int{2, 3, 4}
	for i, num := range nums {
		if num == 3 {
			println("index:", i)
		}
	}
}
