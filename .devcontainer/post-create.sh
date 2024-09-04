#!/bin/bash
## Install tools and dependencies for the project.
## This script is run after the devcontainer is created.
##
## @see .devcontainer/devcontainer.json

echo "[INFO] Initialize pre-commit"
pre-commit install

echo "[INFO] Install gocyclo"
readonly components=(
    /workspaces/source2adoc/components/app
    /workspaces/source2adoc/components/test-acceptance
)
for dir in "${components[@]}"; do
    (
        cd "$dir" || exit
        go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
    )
done
