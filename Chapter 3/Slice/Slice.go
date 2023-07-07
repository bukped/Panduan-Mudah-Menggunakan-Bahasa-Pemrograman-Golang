package main

import "fmt"

func main() {
	// Membuat slice kosong dengan panjang awal 0 dan kapasitas 5
	slice := make([]int, 0, 5)

	// Menambahkan elemen ke slice menggunakan fungsi append()
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	// Mengakses dan mencetak elemen-elemen dalam slice
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Menyalin slice ke slice baru menggunakan fungsi copy()
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)

	// Mengubah elemen dalam slice baru
	newSlice[0] = 10

	fmt.Println("Original Slice:", slice)
	fmt.Println("New Slice:", newSlice)
}
