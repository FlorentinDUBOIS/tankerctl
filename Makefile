# project data
NAME	 := tanker
HASH   := $(shell git rev-parse --verify HEAD)
LFILES := tankerctl.go cmd/... core/...
FILES  := tankerctl.go cmd core

# compiler and compiler options
CC 		 := go build
DEP  	 := dep
CFLAGS := -race
DFLAGS := -ldflags "-X 'github.com/FlorentinDUBOIS/tankerctl/cmd.githash=$(HASH)'"
RM		 := rm -rf
FORMAT := gofmt -s -w
LINT   := gometalinter --deadline=180s --disable-all
GET 	 := go get -u

.PHONY: release
release:
	$(CC) $(DFLAGS) -o build/release/$(NAME) tankerctl.go

.PHONY: dev
dev:
	$(CC) $(CFLAGS) $(DFLAGS) -o build/debug/$(NAME) tankerctl.go

.PHONY: install
install:
	$(GET) github.com/alecthomas/gometalinter
	$(GET) github.com/golang/dep/cmd/dep

	gometalinter --install --update

.PHONY: dep
dep:
	$(DEP) ensure

.PHONY: clean
clean:
	$(RM) vendor
	$(RM) build

.PHONY: lint
lint:
	$(LINT) --enable=deadcode $(LFILES)
	$(LINT) --enable=dupl $(LFILES)
	$(LINT) --enable=errcheck $(LFILES)
	$(LINT) --enable=gas $(LFILES)
	$(LINT) --enable=goconst $(LFILES)
	$(LINT) --enable=gocyclo $(LFILES)
	$(LINT) --enable=goimports $(LFILES)
	$(LINT) --enable=golint $(LFILES)
	$(LINT) --enable=gosimple $(LFILES)
	$(LINT) --enable=gotype $(LFILES)
	$(LINT) --enable=ineffassign $(LFILES)
	$(LINT) --enable=interfacer $(LFILES)
	$(LINT) --enable=megacheck $(LFILES)
	$(LINT) --enable=misspell $(LFILES)
	$(LINT) --enable=safesql $(LFILES)
	$(LINT) --enable=staticcheck $(LFILES)
	$(LINT) --enable=structcheck $(LFILES)
	$(LINT) --enable=unconvert $(LFILES)
	$(LINT) --enable=unparam $(LFILES)
	$(LINT) --enable=unused $(LFILES)
	$(LINT) --enable=varcheck $(LFILES)

.PHONY: format
format:
	$(FORMAT) $(FILES)