BINFILE=bank/bin.go
NAMEFILE=bank/name.go
REPOPATH=github.com/hexindai/bcbc

build:
	@echo "Generating $(BINFILE) file..."
	@awk -f scripts/make-bin-go.awk data/bin.csv > $(BINFILE)
	@echo "Success! $(BINFILE) generated"

	@echo "Generating $(NAMEFILE) file..."
	@awk -f scripts/make-name-go.awk data/name.csv > $(NAMEFILE)
	@echo "Success! $(NAMEFILE) generated"

	go fmt $(REPOPATH)/bank

	go build $(REPOPATH)

.PHONY: build

install:
	go install $(REPOPATH)