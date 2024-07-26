package codefiles

const (
	LANGUAGE_YML        = "yml"
	LANGUAGE_DOCKERFILE = "Dockerfile"
	LANGUAGE_VAGRANT    = "Vagrantfile"
	LANGUAGE_MAKE       = "Makefile"
	LANGUAGE_BASH       = "sh"
	LANGUAGE_INVALID    = "invalid"

	// DOCUMENTATION_PART_META represents the meta information of a code file like the filename and path.
	DOCUMENTATION_PART_META = "meta"

	// DOCUMENTATION_PART_HEADER represents the header documentation of a code file.
	DOCUMENTATION_PART_HEADER = "header"
)

// TEST_SOURCE_DIR is the path to the test data directory for use in testcases.
const TEST_SOURCE_DIR = "/workspaces/source2adoc/components/app/testdata"

// TEST_OUTPUT_DIR is the path to the test output directory for use in testcases.
const TEST_OUTPUT_DIR = "/workspaces/source2adoc/components/app/target"
