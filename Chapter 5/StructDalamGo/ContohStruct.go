type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	p := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	fmt.Println(p.Name) // Output: John Doe
}
