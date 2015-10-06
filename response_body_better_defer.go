package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://icanhazip.com")
	if err != nil {
		fmt.Println("Error.")
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
