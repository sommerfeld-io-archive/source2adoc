package metadata

import (
	_ "embed"
	"strings"
)

// rawVersion is the current version, as read from the components/metadata/version file.
//
//go:embed VERSION
var rawVersion string

func Version() string {
	return strings.ReplaceAll(rawVersion, "\n", "")
}
