# Variables
BINARY_NAME         := "hgmx"
MAIN_PACKAGE_PATH   := "cmd/hgmx/main.go"
VERSION_VAR_PATH    := "main.Version"
GIT_VERSION         := `git describe --tags --always || echo dev`

default:
    @just --list

build:
    @go build -ldflags="-X {{VERSION_VAR_PATH}}={{GIT_VERSION}}" -o {{BINARY_NAME}} {{MAIN_PACKAGE_PATH}}

clean:
    @rm -f {{BINARY_NAME}}

# e.g., run info -v
run *ARGS:
    @echo
    @go run {{MAIN_PACKAGE_PATH}} {{ARGS}}

describe:
    @echo "{{GIT_VERSION}}"

patch:
    #!/usr/bin/env bash
    set -euo pipefail
    LATEST_TAG=$(git describe --tags --abbrev=0)
    MAJOR_MINOR=$(echo $LATEST_TAG | awk -F. '{print $1"."$2}')
    CURRENT_PATCH=$(echo $LATEST_TAG | awk -F. '{print $3}')
    NEW_PATCH=$((CURRENT_PATCH + 1))
    NEW_TAG="${MAJOR_MINOR}.${NEW_PATCH}"
    echo "$LATEST_TAG -> $NEW_TAG"
    git tag $NEW_TAG

patch-undo:
    #!/usr/bin/env bash
    set -euo pipefail
    # Get the most recently created tag
    LATEST_TAG=$(git tag --sort=-creatordate | tail -n 1)
    if [[ -z "$LATEST_TAG" ]]; then
      echo "Error: No tags found to delete."
      exit 1
    fi
    git tag -d $LATEST_TAG

release:
    @git push
    @git push origin --tags
