#!/bin/sh
#
set -e

rm -rf _build/completions
mkdir -p _build/completions

for sh in bash zsh fish; do
	go run cmd/cli/main.go completion "$sh" > "_build/completions/n2xctl.$sh"
done
