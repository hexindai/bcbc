BINFILE := bank/bin.go
NAMEFILE := bank/name.go
DATABINFILE := data/bin.csv
DATANAMEFILE := data/name.csv
TMPBINFILE := data/bin.tmp

REPOPATH := github.com/hexindai/bcbc

ifeq ($(DEBUG), 1)
debug := 1
else
debug := 0
endif


.PHONY: all
all: build test

.PHONY: test
test:
	@echo "TEST..." \
	&&go test $(REPOPATH)/bank

.PHONY: build
build:
	@echo ">> Generating data go file" \
	&& $(call sort_bin,$(DATABINFILE)) \
	&& awk -f scripts/make-bin-go.awk $(DATABINFILE) > $(BINFILE)	\
	&& awk -f scripts/make-name-go.awk $(DATANAMEFILE) > $(NAMEFILE)

	@echo ">> Formating and building" \
	&& go fmt $(REPOPATH)/bank >/dev/null \
	&& go build $(REPOPATH)

.PHONY: install
install:
	go install $(REPOPATH)

.PHONY: add
add:
	@# check bin via api
	@# if success, append it to $(DATABINFILE)
	@# use DEBUG=1 to enable debug
	
	@echo "CHECK bin: $(bin) len: $(len)" \
	&&awk -f scripts/check-bin.awk -v bin=$(bin) -v len=$(len) -v binfile=$(DATABINFILE) -v debug=$(debug)

	@$(call sort_bin,$(DATABINFILE))

define sort_bin
	awk -f scripts/sort-bin.awk -v to=$(TMPBINFILE) $(1) \
	&& mv $(TMPBINFILE) $(1)
endef
