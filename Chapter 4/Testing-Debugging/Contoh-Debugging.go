package main

func FindMax(slice []int) int {
	max := slice[0]
	for _, num := range slice {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	slice := []int{3, 9, 2, 7, 5}

	max := FindMax(slice)

	println("Max number:", max)
}
