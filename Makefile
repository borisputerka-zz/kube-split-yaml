export GO111MODULE=on

.PHONY: bin
bin: fmt vet
	go build -o kube-split-yaml github.com/borisputerka/kube-split-yaml

.PHONY: fmt
fmt:
	go fmt ./... ./pkg/...

.PHONY: vet
vet:
	go vet ./... ./pkg/...

.PHONY: lint
lint:
	@golangci-lint run ./...
