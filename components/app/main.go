package main

import (
	"github.com/sommerfeld-io/source2adoc/cmd"
	"github.com/sommerfeld-io/source2adoc/internal/helper"
)

// func init() {
// }

func main() {
	err := cmd.Execute()
	helper.HandleError(err)
}
