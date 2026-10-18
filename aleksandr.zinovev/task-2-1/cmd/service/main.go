package main

import (
	"fmt"
)

func main() {
	var n, k int

	if _, err := fmt.Scan(&n, &k); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
		}
	}
}
