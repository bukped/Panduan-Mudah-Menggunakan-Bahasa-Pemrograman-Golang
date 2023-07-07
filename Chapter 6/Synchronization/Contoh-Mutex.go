import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	totalGoroutines := 5
	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	wg.Done()
}
