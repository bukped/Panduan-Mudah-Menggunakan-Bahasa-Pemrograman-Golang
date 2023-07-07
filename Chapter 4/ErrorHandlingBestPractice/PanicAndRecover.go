func DoSomething() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	// Potensi panic
	panic("Something went wrong!")
}

func main() {
	DoSomething()
	fmt.Println("Program continues after panic")
}
