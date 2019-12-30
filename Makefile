SHELL = /usr/bin/env bash

PKGNAME := github.com/ryapric/gotables
GOPATH := $$(go env GOPATH)

build:
	@printf "Building binaries...\n"
	@GOPATH=$(GOPATH) go build $(GOPATH)/src/$(PKGNAME)/...

test:
	@printf "Running tests...\n"
	@GOPATH=$(GOPATH) go test $(PKGNAME)/...

install:
	@printf 'Installing to $$GOPATH/bin...\n'
	@GOPATH=$(GOPATH) go install $(GOPATH)/src/$(PKGNAME)/...

run:
	@GOPATH=$(GOPATH) go run $(GOPATH)/src/$(PKGNAME)/main.go

# Run on a new host to configure VS Code settings. This will only set Go configs
# for VS Code's sake -- your host environment is unchanged.
vscode-settings:
	@mkdir -p ./.vscode
	@cp .vscode-settings.json ./.vscode/settings.json
	@sed -i "s;GOPATH;$(GOPATH);" ./.vscode/settings.json
