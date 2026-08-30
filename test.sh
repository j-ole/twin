#!/bin/bash

set -e -o pipefail

# Compile test first
echo "Building..."
go build ./...

# Linting
echo 'Linting, repro any errors locally using "golangci-lint run"...'
echo '  Linting without tests...'
golangci-lint run --tests=false
echo '  Linting with tests...'
golangci-lint run --tests=true

# Unit tests
echo "Running unit tests..."
RACE=-race
if [ "$GOARCH" == "386" ]; then
  # -race is not supported on i386
  RACE=""
fi
# Report races into files in the repo root, see RACES.md. Absolute path,
# otherwise each report lands in its own package's directory rather than
# here.
#
# Not in CI though: log_path silences the report on stdout, and the file dies
# with the runner, so there we'd lose the report entirely.
RACE_REPORTS=""
if [ -z "${CI}" ]; then
  RACE_REPORTS="log_path=${PWD}/twin-race-report"
fi
GORACE="${RACE_REPORTS}" go test $RACE -timeout 60s ./...

# Ensure we can cross compile
echo "Testing cross compilation..."
echo "  Linux i386..."
GOOS=linux GOARCH=386 go build ./...

echo "  Linux amd64..."
GOOS=linux GOARCH=amd64 go build ./...

echo "  Linux arm32..."
GOOS=linux GOARCH=arm go build ./...

echo "  macOS amd64..."
GOOS=darwin GOARCH=amd64 go build ./...

echo "  Windows amd64..."
GOOS=windows GOARCH=amd64 go build ./...

if ls twin-race-report.* 1> /dev/null 2>&1; then
  echo
  echo "ERROR: Please check RACES.md and look into the following files:"
  ls twin-race-report.*
  exit 1
fi

echo
echo "All tests passed!"
