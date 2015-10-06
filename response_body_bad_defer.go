package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://icanhazip.com")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error.")
		return
	}
	fmt.Println(resp.StatusCode)
}
