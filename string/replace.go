package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("gfg gfg gfg", "g", "G", 3))
	fmt.Println(strings.Replace("gfg gfg gfg", "fg", "FG", -1))
}
