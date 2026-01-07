package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworm("testdata/bookworkhub.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading bookworms: %v\n", err)
		os.Exit(1)

	}
	fmt.Println(bookworms)

}
