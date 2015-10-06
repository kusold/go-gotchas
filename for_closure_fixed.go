package main

import (
	"fmt"
	"time"
)

func main() {
	input := []string{"Boulder", "golang", "meetup"}

	for _, value := range input {
		go func(val string) {
			fmt.Println(val)
		}(value)
	}
	time.Sleep(1 * time.Second)
}
