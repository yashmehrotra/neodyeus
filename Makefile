GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

run:
	@go run main.go
test:
	@go test -v ./...

fmt:
	@gofmt -s -w ${GOFILES_NOVENDOR}

lint:
	@golint ${GOFILES_NOVENDOR}

build_neodyeus:
	@env GOOS=linux GOARCH=amd64 go build -o Neodyeus main.go

check_build_status:
	@echo "Building Neodyeus"
	@$(MAKE) build_neodyeus
	@echo "Build Success"
	@rm Neodyeus

dependency_graph:
	@echo "Making dependency graph"
	@dep status -dot | dot -T png > graph.png
