BINFILE=bank/bin.go
NAMEFILE=bank/name.go
REPOPATH=github.com/hexindai/bcbc
DATABINFILE=data/bin.csv
DATANAMEFILE=data/name.csv

.PHONY: all
all: build test

.PHONY: test
test:
	@echo "TEST..."
	@go test $(REPOPATH)/bank

.PHONY: build
build:
	@echo "BUILD..."

	@awk -f scripts/sort-bin.awk $(DATABINFILE)
	@echo "...$(DATABINFILE) SORTED"

	@awk -f scripts/make-bin-go.awk $(DATABINFILE) > $(BINFILE)
	@echo "...$(BINFILE) GENERATED"

	@awk -f scripts/make-name-go.awk $(DATANAMEFILE) > $(NAMEFILE)
	@echo "...$(NAMEFILE) GENERATED"

	@go fmt $(REPOPATH)/bank
	@echo "...FORMATTED"

	@go build $(REPOPATH)
	@echo "...OK"

.PHONY: install
install:
	go install $(REPOPATH)
