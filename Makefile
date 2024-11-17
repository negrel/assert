.PHONY: codegen
codegen:
	./codegen.sh

.PHONY: clear
clear:
	rm -f ./*assertion*.go ./prod_*

.PHONY: lint
lint:
	shellcheck ./codegen.sh

.PHONY: test
test:
	go test -tags assert -v ./...

.PHONY: bench
bench: bench/assert bench/noassert

.PHONY: bench/noassert
bench/noassert:
	@echo "ASSERTIONS DISABLED"
	go test -bench=. -run=^$$ -v ./...

.PHONY: bench/assert
bench/assert:
	@echo "ASSERTIONS ENABLED"
	go test -bench=. -run=^$$ -tags assert -v ./...
