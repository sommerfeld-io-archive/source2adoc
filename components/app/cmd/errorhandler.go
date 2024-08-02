package cmd

import "log"

// handleError handles all Errors of this application. See "Error Handling in our
// Go Code" in the Development Guide (`docs/modules/ROOT/pages/development-guide.adoc`).
//
// TODO:	somehow ensure this as a unit test in the test package (e.g. the
// TODO:	signatures of all other methods). If possible ...
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
