#!/bin/bash

# Checks if all project files are formatted using go fmt
#
# Usage example:
#> bin/check_go_fmt

source scripts/helpers/output.sh

unformatted=`find . -type d -name "vendor" -prune -o -type f -iregex '.*\.go' -exec goimports -l '{}' \;`
if [ ! -z "$unformatted" ]; then
    print_error "Following files are not formatted using goimports:"
    echo "$unformatted"
    exit 1
fi
print_success "All files are formatted using goimports."
