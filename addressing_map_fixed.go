package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	people := map[string]Person{"mike": {"Michael"}}
	fmt.Println(people["mike"])
	mike := people["mike"]
	mike.Name = "Mikey"
	people["mike"] = mike
	fmt.Println(people["mike"])
}
