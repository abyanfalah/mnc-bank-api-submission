package utils

import (
	"fmt"
	"strings"
)

func Line() {
	fmt.Println(strings.Repeat("=", 50))
}

func LineText(t interface{}) {
	Line()
	fmt.Println(t)
	Line()
}
