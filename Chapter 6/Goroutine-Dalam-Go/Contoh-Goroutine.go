func main() {
	go count("goroutine 1")
	go count("goroutine 2")

	// Menunggu goroutine selesai sebelum program berakhir
	time.Sleep(time.Second)
}

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Millisecond * 500)
	}
}
