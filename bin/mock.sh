#!/usr/bin/env bash

set -euo pipefail

for file in `find . -name '*.go' | grep -v /vendor`; do
	if `grep -q 'interface {' ${file}`; then
	  dest=${file//interna\//}
		mockgen -source=${file} -destination=test/mock/${dest}
	fi
done