package main

import "fmt"

func main() {
	fmt.Println(selection("y"))
}

func selection(in string) (choice string) {
	switch in {
	case "y", "yeah":
		choice = "yes"
		break
	default:
		choice = "no"
	}
	return
}
