package main

import (
	"fmt"
	"regexp"
)

func isCardValid(str string) string {
	if len(str) == 16 || len(str) == 19 {
		re := regexp.MustCompile(`(?m)^[456][0-9]{3}-?[0-9]{4}-?[0-9]{4}-?[0-9]{4}$`)
		if re.MatchString(str) {
			return "valid"
		} else {
			return "invalid"
		}
	}
	return "invalid"

}

func main() {
	cardNums := []string{
		"4123456789123456",
		"5123-4567-8912-3456",
		"61234-567-8912-3456",
		"4123356789123456",
		"5133-3367-8912-3456",
		"5123 - 3567 - 8912 - 3456",
		"6501-123456789081",
	}
	for _, str := range cardNums {
		fmt.Printf("Card: %s is %s\n", str, isCardValid(str))
	}

}
