import (
	"log"
	"os"
)

func ProcessData() error {
	// Simulasi kesalahan
	err := someFunction()
	if err != nil {
		log.Printf("Error occurred during data processing: %s", err)
		return err
	}

	return nil
}

func main() {
	// Konfigurasi logger
	logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// Contoh pemanggilan fungsi
	err = ProcessData()
	if err != nil {
		log.Println("Data processing failed:", err)
	}
}
