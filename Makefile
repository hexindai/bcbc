BINFILE=bank/bin.go
NAMEFILE=bank/name.go
REPOPATH=github.com/hexindai/bcbc

.PHONY: all
all: test build

.PHONY: test
test:
	go test $(REPOPATH)/bank

.PHONY: build
build:
	@echo "Generating $(BINFILE) file..."
	@awk -f scripts/make-bin-go.awk data/bin.csv > $(BINFILE)
	@echo "Success! $(BINFILE) generated"

	@echo "Generating $(NAMEFILE) file..."
	@awk -f scripts/make-name-go.awk data/name.csv > $(NAMEFILE)
	@echo "Success! $(NAMEFILE) generated"

	go fmt $(REPOPATH)/bank

	go build $(REPOPATH)

.PHONY: install
install:
	go install $(REPOPATH)
