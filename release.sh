#!/bin/bash

# Rewrite references from github.com/peay/gogen-avro to gopkg.in/peay/gogen-avro.<version>

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <version>"
  exit 1
fi
 
GITHUB_REPO="github.com/peay/gogen-avro"
VERSION="$1"
GOPKG_REPO="gopkg.in/peay/gogen-avro.$VERSION"

sed -i "s|$GITHUB_REPO|$GOPKG_REPO|" container/*.go generator/*.go types/*.go main.go example/*/*.go test.sh 
