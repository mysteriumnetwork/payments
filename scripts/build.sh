#!/bin/bash
set -e
scripts/gen_all.sh
go build ./...