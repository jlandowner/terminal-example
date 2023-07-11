package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Millisecond * 100)
	for i := 1; i <= 100; i++ {
		<-ticker
		// write at previous cursor
		fmt.Printf("\x0cOn %d", i)
	}
	fmt.Println("\nAll done")
}
