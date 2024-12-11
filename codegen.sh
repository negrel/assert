#!/usr/bin/env bash

set -euo pipefail
shopt -s inherit_errexit

testify_file_skiplist=("assertion_forward.go" "forward_assertions.go" "doc.go" "extra")

# Copy testify files.
for f in ./testify/assert/*.go; do
	basename_f="$(basename "$f")"

	# Skip file.
	# shellcheck disable=SC2076
	if [[ " ${testify_file_skiplist[*]} " =~ " $basename_f " ]]; then
		echo "skipping $f"
		continue
	fi

	cp "$f" "$basename_f"
done

rm -f ./*_test.go

for f in ./*.go; do
	# Remove TestingT arguments in function signature.
	sed -i -E 's/(t )?TestingT, //g' "$f"

	# Remove TestingT arguments in function call.
	sed -i 's/(t, /(/g' "$f"
	sed -i -E 's/^\s+t,$//g' "$f"

	# Remove:
	# if h, ok := t.(tHelper); ok {
	# 	h.Helper()
	# }
	sed -i '/if h, ok := t.(tHelper); ok {/,/}/d' "$f"

	# Remove:
	# if n, ok := t.(interface {
	# 	Name() string
	# }); ok {
	# 	content = append(content, labeledContent{"Test", n.Name()})
	# }
	sed -i '/if n, ok := t.(interface {/,/}$/d' "$f"

	# Replace:
	# t.Errorf(.*) with panic(fmt.Errorf(.*))
	sed -i -E 's/\s+t\.Errorf\((.*)\)/panic(fmt.Errorf(\1))/g' "$f"

	# Delete:
	# if t, ok := t.(failNower); ok {
	# 	t.FailNow()
	# } else {
	# 	panic("test failed and t is missing `FailNow()`")
	# }
	sed -i '/if t, ok := t.(failNower); ok {/,/}$/d' "$f"

	# Delete YAML functions:
	sed -i '/func YAML.*{/,/^}$/d' "$f"
done

# Create prod_* files that will contain empty function.
for f in *.go; do
	if [ "$f" == "doc.go" ]; then
		continue
	fi

	# Prepend build tag to files if needed.
	if ! grep "//go:build assert" < "$f" > /dev/null; then
		sed -i '1i //go:build assert' "$f"
	fi

	cp "$f" "prod_$f"

	# Change build tag on prod file.
	sed -i 's|^//go:build assert|//go:build !assert|' "prod_$f"

	# Remove function body in prod file.
	sed -i 's/func\(.*\){$/func\1{}\n{/' "prod_$f"
	sed -i '/^{$/,/^}$/d' "prod_$f"

	# Replace:
	# func () returnType with func ()
	sed -i 's/func\(.*\) [^()]* {}$/func\1{}/' "prod_$f"
	sed -i 's/func\(.*\) \(([^)]*)\) {}$/func\1{}/' "prod_$f"
done

gofmt -w ./*.go
goimports -w ./*.go
go mod tidy
