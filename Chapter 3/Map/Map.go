package main

import "fmt"

func main() {
	// Mendeklarasikan dan menginisialisasi map
	studentGrades := make(map[string]int)
	studentGrades["John"] = 90
	studentGrades["Jane"] = 85
	studentGrades["Alice"] = 95

	// Mengakses nilai dalam map
	fmt.Println("Nilai John:", studentGrades["John"])

	// Memeriksa apakah kunci ada dalam map
	grade, exists := studentGrades["Jane"]
	if exists {
		fmt.Println("Nilai Jane:", grade)
	} else {
		fmt.Println("Nilai Jane tidak ditemukan")
	}

	// Menambah atau memperbarui nilai dalam map
	studentGrades["John"] = 95
	studentGrades["Bob"] = 80

	// Menghapus pasangan kunci-nilai dari map
	delete(studentGrades, "Alice")

	// Iterasi melalui map
	for name, grade := range studentGrades {
		fmt.Printf("Nama: %s, Nilai: %d\n", name, grade)
	}
}
