.PHONY: test

COVERAGE_FILE := cover.out

test:
	go test -coverprofile $(COVERAGE_FILE) ./...
	go tool cover -html=$(COVERAGE_FILE)
