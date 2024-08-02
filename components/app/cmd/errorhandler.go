package cmd

import "log"

// handleError handles all Errors of this application. This function is
// exuse assertions from "github.com/strclusively called from the CLI commands from this package. This is why
// this function is placed in the `cmd` package and is not exported.
//
// No other function or structure from any other package is allowed handle
// errors, meaning no other package should write error information to a log
// file or `stdout`. All functions and structures from all other packages
// should return errors to the caller as part of their method signature.
//
// The only exception to this rule is the test package, which only contains
// tests that are not directly related to a go code file.
//
// TODO:	somehow ensure this as a unit test in the test package (e.g. the
// TODO:	signatures of all other methods).
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
