package main

import (
	"fmt"

	format "example.com/task-hive/format"
)

func main() {
	fmt.Printf("%q\n", format.Truncate("sample", 0))
}

