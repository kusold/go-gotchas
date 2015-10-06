package main

import (
	"fmt"
	"time"
)

func main() {
	input := []string{"Boulder", "golang", "meetup"}

	for _, value := range input {
		go func() {
			fmt.Println(value)
		}()
	}
	time.Sleep(1 * time.Second)
}
