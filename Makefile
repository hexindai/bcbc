BINFILE ?= bank/bin.go
NAMEFILE ?= bank/name.go
DATABINFILE ?= data/bin.csv
DATANAMEFILE ?= data/name.csv
TMPBINFILE ?= data/bin.tmp

REPOPATH = github.com/hexindai/bcbc

.PHONY: all
all: build test

.PHONY: test
test:
	@echo "TEST..."
	@go test $(REPOPATH)/bank

.PHONY: build
build:
	@echo ">> BUILDING"

	@awk -f scripts/sort-bin.awk -v to=$(TMPBINFILE) $(DATABINFILE)
	@mv $(TMPBINFILE) $(DATABINFILE)
	@echo "...$(DATABINFILE) SORTED"

	@awk -f scripts/make-bin-go.awk $(DATABINFILE) > $(BINFILE)
	@echo "...$(BINFILE) GENERATED"

	@awk -f scripts/make-name-go.awk $(DATANAMEFILE) > $(NAMEFILE)
	@echo "...$(NAMEFILE) GENERATED"

	@echo ">> FORMATTING"
	go fmt $(REPOPATH)/bank
	@echo "FORMATTED"

	go build $(REPOPATH)
	@echo "SUCCESS"

.PHONY: install
install:
	go install $(REPOPATH)

.PHONY: add
add:
	@# check bin via api
	@# if success, append it to $(DATABINFILE)
	
	@echo "CHECK bin: $(bin) len: $(len)"
	@awk -f scripts/check-bin.awk -v bin=$(bin) -v len=$(len) -v binfile=$(DATABINFILE) #-v debug=true

	@awk -f scripts/sort-bin.awk -v to=$(TMPBINFILE) $(DATABINFILE)
	@mv $(TMPBINFILE) $(DATABINFILE)
	@echo "...$(DATABINFILE) SORTED"
