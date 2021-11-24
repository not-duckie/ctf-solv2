package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage secureDelete <file name>")
		os.Exit(1)
	}

	secure_delete := Config{Iter: 1, Zero: true}
	secure_delete.DeleteFile(os.Args[1])
}
