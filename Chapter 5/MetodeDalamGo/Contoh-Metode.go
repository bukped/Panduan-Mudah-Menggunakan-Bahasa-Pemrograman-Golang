type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func main() {
	rect := Rectangle{
		Length: 5.0,
		Width:  3.0,
	}

	fmt.Println(rect.Area()) // Output: 15.0
}
