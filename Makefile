
.PHONY: codegen
codegen:
	./codegen.sh

.PHONY: clear
clear:
	rm -f ./*.go

.PHONY: lint
lint:
	shellcheck ./codegen.sh

