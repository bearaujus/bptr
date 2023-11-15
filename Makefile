.PHONY: test
test_cmd := go test $(shell go list ./... | grep -v "go\\.work")
test:
	go mod tidy
	go clean -testcache
	$(test_cmd) --cover
	go clean -testcache
	$(test_cmd) --cover --race
