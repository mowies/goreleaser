#!/bin/bash

zig targets | jq -r '.libc[]' | grep -v freestanding | while read -r target; do
	if zig cc -target "$target" hello.c 2>/dev/null; then
		echo "Ok $target" >&2
		echo "\"${target}\","
		echo "\"$(echo "$target" | cut -f1,2 -d-)\","
	else
		echo "Er $target" >&2
	fi
done | sort | uniq
