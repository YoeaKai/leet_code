package main

import (
	"fmt"
	"regexp"
)

func main() {
	// s := "web browser - Google 搜尋 - Google Chrome - testuser@gmail.com"
	// re := regexp.MustCompile("^web browser - \\S+ - Google Chrome - \\S+@\\S+$")

	s := "web browser - Google 搜尋 - Google Chrome - testuser@gmail.com"
	re := regexp.MustCompile("^web browser - Google .* - Google Chrome - .*")

	ratio := re.FindString(s)
	fmt.Println(ratio)
}
