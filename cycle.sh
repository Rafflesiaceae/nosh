#!/bin/bash
set -eo pipefail
go build
./nosh ./quicktest.star
