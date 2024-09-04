package codefiles

const (
	LanguageYml          = "yml"
	LanguageDockerfile   = "Dockerfile"
	LanguageVagrant      = "Vagrantfile"
	LanguageMake         = "Makefile"
	LanguageBash         = "sh"
	LanguageNotSupported = "not-supported"

	// DocumentationPartMetadata represents the meta information of a code file like the filename and path.
	DocumentationPartMetadata = "meta"

	// DocumentationPartHeader represents the header documentation of a code file.
	DocumentationPartHeader = "header"
)

// TestSourceDir is the path to the test data directory for use in testcases.
const TestSourceDir = "/workspaces/source2adoc/testdata/common"

// TestOutputDir is the path to the test output directory for use in testcases.
const TestOutputDir = "/workspaces/source2adoc/target/unit-test"
