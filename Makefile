# Validate every module listed in .github/ci-modules.json (the same manifest CI uses).
MODULES := $(shell jq -r '.modules[].dir' .github/ci-modules.json)

.PHONY: test vet build fmt doccheck tidy all

all: vet fmt doccheck test build

test:
	@for d in $(MODULES); do echo "== go test $$d =="; (cd $$d && go test ./...) || exit 1; done

vet:
	@for d in $(MODULES); do echo "== go vet $$d =="; (cd $$d && go vet ./...) || exit 1; done

build:
	@for d in $(MODULES); do echo "== go build $$d =="; (cd $$d && go build ./...) || exit 1; done

fmt:
	@unformatted=$$(gofmt -l $$(git ls-files '*.go')); \
	if [ -n "$$unformatted" ]; then echo "$$unformatted"; exit 1; fi

doccheck:
	go run ./tools/doccheck

# Regenerate go.mod/go.sum for every module (e.g. after adding a module to the manifest).
tidy:
	@for d in $(MODULES); do echo "== go mod tidy $$d =="; (cd $$d && go mod tidy) || exit 1; done
