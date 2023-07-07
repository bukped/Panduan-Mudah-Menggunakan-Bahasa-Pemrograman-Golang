import (
	"fmt"
)

func main() {
	message := make(chan string)

	go sendData(message)
	go receiveData(message)

	// Menunggu pengiriman dan penerimaan data selesai
	fmt.Scanln()
}

func sendData(ch chan<- string) {
	ch <- "Hello, world!"
}

func receiveData(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
