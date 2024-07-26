package main

import (
	"log"

	"github.com/sommerfeld-io/source2adoc/cmd"
)

// func init() {
// }

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
