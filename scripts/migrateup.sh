#!/bin/bash
set -e

if [ -f .env ]; then
    source .env
fi

cd sql/schema
"$(go env GOPATH)/bin/goose" up
