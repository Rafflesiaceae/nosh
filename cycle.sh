#!/usr/bin/env bash
set -eo pipefail

# cd to parent dir of current script
cd "$(dirname "${BASH_SOURCE[0]}")"

# build nosh
go build

# run quicktests
./nosh ./quicktest.nosh
