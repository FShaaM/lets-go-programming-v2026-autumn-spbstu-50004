package main

import (
	"fmt"
)

func main() {
	var (
		departmentsCount int
		employeesCount   int
		minTemp          int
		maxTemp          int
		ask              string
		targetTemp       int
	)

	_, err := fmt.Scan(&departmentsCount)
	if err != nil {
		return
	}

	for range departmentsCount {

		_, err = fmt.Scan(&employeesCount)
		if err != nil {
			fmt.Println("Invalid employees number")
			return
		}

		minTemp = 15
		maxTemp = 30

		for range employeesCount {

			if _, err := fmt.Scan(&ask, &targetTemp); err != nil {
				return
			}

			if ask == ">=" && targetTemp > minTemp {
				minTemp = targetTemp
			} else if ask == "<=" && targetTemp < maxTemp {
				maxTemp = targetTemp
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
