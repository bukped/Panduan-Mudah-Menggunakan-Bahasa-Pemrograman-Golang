type Animal struct {
	Name string
}

type Dog struct {
	Animal
	Breed string
}

func (a Animal) Sound() string {
	return "Animal sound"
}

func (d Dog) Sound() string {
	return "Woof"
}

func main() {
	animal := Animal{Name: "Animal"}
	dog := Dog{Animal: animal, Breed: "Labrador"}

	fmt.Println(animal.Sound()) // Output: Animal sound
	fmt.Println(dog.Sound())    // Output: Woof
}
