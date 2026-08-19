package main

import (
	"fmt"
	"os"
)

// schema_migrator - DB schema migration
func schema_migrator(path string) {
	fmt.Println("========================================")
	fmt.Println("  Schema-Migrator")
	fmt.Println("  DB schema migration")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	schema_migrator(path)
}
