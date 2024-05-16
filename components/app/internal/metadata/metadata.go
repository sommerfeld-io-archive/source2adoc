package metadata

import (
	_ "embed"
	"strings"
)

// rawVersion is the current version, as read from the components/app/internal/metadata/VERSION file.
//
//go:embed VERSION
var rawVersion string

// rawCommitSha is the commit sha, as read from the components/app/internal/metadata/COMMIT_SHA file.
//
//go:embed COMMIT_SHA
var rawCommitSha string

func Version() string {
	return strings.ReplaceAll(rawVersion, "\n", "")
}

func CommitSha() string {
	return strings.ReplaceAll(rawCommitSha, "\n", "")
}
