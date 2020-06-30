.PHONY: build run clean dev-requirements snapshot copy-to-lorix test
PKGS := $(shell go list ./...)

build:
	go build -o build/kiwi-server-executable.e

run:
	build/kiwi-server-executable.e

clean:
	rm -f build/kiwi-server-executable.e 
	rm -rf dist

dev-requirements:
	go mod download
	go install github.com/goreleaser/goreleaser
	go install golang.org/x/lint/golint

snapshot:
	goreleaser --snapshot

test:
	@echo "Running tests"
	@rm -f coverage.out
	@for pkg in $(PKGS) ; do \
		golint $$pkg ; \
	done
	@go vet $(PKGS)
	@go test -p 1 -v $(PKGS) -cover -coverprofile coverage.out

copy-to-lorix:
	@if [ -z "$(host)" ]; then \
		echo "error: variable 'host' must be provides as in 'make copy-to-lorix host=xyz'"; \
		exit 1; \
	fi
	@echo "transmit kiwi-server from dist/kiwi-server_linux_arm_7 to the home folder of user 'admin' on $(host)"
	scp dist/kiwi-server_linux_arm_7/kiwi-server admin@$(host):~
