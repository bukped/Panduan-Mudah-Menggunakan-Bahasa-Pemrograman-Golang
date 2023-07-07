func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("Cannot divide by zero")
    }
    return a / b, nil
}

result, err := Divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}

