package main

import (
	"fmt"
	"strings"
)



func main()  {
	difficulty := 2
	hash := "1123456789"
	prefix := strings.Repeat("1", difficulty)
	fmt.Println(strings.HasPrefix(hash, prefix))
}
