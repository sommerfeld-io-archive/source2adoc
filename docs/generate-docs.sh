#!/bin/bash
## Generate docs from source code using source2adoc.
## Antora-related files are written / updated through bash scripting, not through source2adoc CLI.
##
## This script is called from GitHub Actions to generate the documentation.
##
## ! Temporary file, will be replaced by a GitHub Actions workflow and
## ! https://github.com/sommerfeld-io/source2adoc/issues/145.

readonly ANTORA_MODULE_NAME="source2adoc"
readonly ANTORA_MODULE="docs/modules/$ANTORA_MODULE_NAME"
readonly ANTORA_YML="docs/antora.yml"
readonly ANTORA_YML_LINE="  - modules/source2adoc/nav.adoc"
readonly OUTPUT_DIR="$ANTORA_MODULE/pages"

(
    cd .. || exit

    # Cleanup output directory
    # optinal, but recommended because there might be files left from previous runs
    rm -rf "$OUTPUT_DIR"

    # Generate documentation
    docker run --volume "$(pwd):$(pwd)" --workdir "$(pwd)" \
        sommerfeldio/source2adoc:latest \
            --source-dir . --output-dir "$OUTPUT_DIR" \
            --exclude testdata \
            --exclude target \
            --exclude .hadolint.yml \
            --exclude .ls-lint.yml \
            --exclude .yamllint.yml \
            --exclude docs/antora.yml

    # Write module to Antora component config
    if ! grep -q "$ANTORA_YML_LINE" "$ANTORA_YML"; then
        echo "$ANTORA_YML_LINE" >> "$ANTORA_YML"
    fi

    # Generate Antora navigation file
    echo "* xref:index.adoc[]" > "$ANTORA_MODULE/nav.adoc"

    # Generate Antora module startpage with navigation
    {
        echo "= Source Code Docs"
        echo
    } > "$ANTORA_MODULE/pages/index.adoc"

    find "$OUTPUT_DIR" -type f | while read -r file; do
        path="${file#"$OUTPUT_DIR"/}"

        dir="$(dirname -- "$path")"
        filename="$(basename -- "$path")"

        if [ "$filename" != "index.adoc" ]; then
            modifiedFilename="${filename//./-}"
            modifiedFilename="${modifiedFilename//-adoc/.adoc}"

            path="$dir/$modifiedFilename"
            if [ "$dir" == "." ]; then
                path="$modifiedFilename"
            fi

            echo "* xref:$ANTORA_MODULE_NAME:${path}[]" >> "$ANTORA_MODULE/pages/index.adoc"
        fi
    done
)
