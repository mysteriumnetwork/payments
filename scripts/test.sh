#!/bin/bash
set -e

source scripts/helpers/output.sh

scripts/gen_all.sh

go test $@ ./...
if [[ $? != 0 ]]; then
    print_error "Tests failed"
    exit 1
fi
print_success "All tests passed"
exit 0