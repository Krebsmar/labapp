package main

import "time"

func generateLoad() {
	for {
		// Perform some computation
		for i := 0; i < 1e6; i++ {
			_ = i * i
		}

		// Sleep for a while to avoid consuming too much CPU
		time.Sleep(100 * time.Millisecond)
	}
}
