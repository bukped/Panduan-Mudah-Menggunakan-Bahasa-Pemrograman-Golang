func panggilFungsiLain(fungsi func(int, int) int) {
    hasil := fungsi(3, 4)
    fmt.Println("Hasil:", hasil)
}
func tambah(a int, b int) int {
    return a + b
}
panggilFungsiLain(tambah)  
// Memanggil fungsi "panggilFungsiLain" dengan argumen fungsi "tambah"

