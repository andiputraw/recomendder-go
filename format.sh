#!/bin/bash

# Set the path to your Go workspace
WORKSPACE_PATH="."

# Use the find command to locate all Go files in the workspace
GO_FILES=$(find "$WORKSPACE_PATH" -name "*.go")

# Iterate over each Go file and run go fmt
for file in $GO_FILES; do
  go fmt "$file"
done

echo "Go fmt applied to all Go files in the workspace."
