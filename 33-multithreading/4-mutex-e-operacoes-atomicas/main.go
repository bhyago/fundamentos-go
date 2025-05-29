package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock() // Lock the mutex
		// number++ 
		// m.Unlock() // Unlock the mutex
		atomic.AddUint64(&number, 1) // Atomic operation to increment number
		time.Sleep(300 * time.Millisecond) // Simulate a delay
		w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}