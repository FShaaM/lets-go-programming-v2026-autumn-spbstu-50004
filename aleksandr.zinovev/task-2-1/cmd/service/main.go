package main

import (
	"fmt"
)

func main() {
	var (
		n, k             int
		minTemp, maxTemp int
		ask              string
		t                int
	)

	if _, err := fmt.Scan(&n, &k); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		minTemp = 15
		maxTemp = 30

		for j := 0; j < k; j++ {
			fmt.Scan(&ask, &t)

			if ask == ">=" && t > minTemp {
				minTemp = t
			} else if ask == "<=" && t < maxTemp {
				maxTemp = t
			}
		}
	}
}
