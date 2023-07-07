package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagian dengan nol tidak diizinkan")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
	} else {
		fmt.Println("Hasil pembagian:", result)
	}
}
