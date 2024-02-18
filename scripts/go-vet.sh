#!/bin/bash
set -e

# Get the list of Go packages in the current project
packages=$(go list ./... | grep -vE '(^|/)(venv|tmp)/')

# Run go vet on each package
for pkg in $packages; do
  echo "Running go vet on package: $pkg"
  go vet "$pkg"
done

echo "go vet completed successfully."
