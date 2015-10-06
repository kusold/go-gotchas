package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June"}
	var keys []int

	for key := range m {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}
