#
# Simple Makefile for conviently testing, building and deploying experiment.
#
PROJECT = simplified

VERSION = $(shell grep '"version":' codemeta.json | cut -d\"  -f 4)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)


MAN_PAGES = simplified2markdown.1 

PROGRAMS = $(shell ls -1 cmd)

PACKAGE = $(shell ls -1 *.go)

PANDOC = $(shell which pandoc)

OS = $(shell uname)

#PREFIX = /usr/local/bin
PREFIX = $(HOME)

ifneq ($(prefix),)
	PREFIX = $(prefix)
endif

EXT = 
ifeq ($(OS), Windows)
        EXT = .exe
endif

QUICK =
ifeq ($(quick), true)
	QUICK = quick=true
endif


build: version.go $(PROGRAMS)

version.go: .FORCE
	@echo 'package $(PROJECT)' >version.go
	@echo '' >>version.go
	@echo 'const (' >>version.go
	@echo '    Version = "$(VERSION)"' >>version.go
	@echo '' >>version.go
	@echo '    LicenseText = `' >>version.go
	@cat LICENSE >>version.go
	@echo '`' >>version.go
	@echo ')' >>version.go
	@echo '' >>version.go
	-git add version.go
	@if [ -f bin/codemeta ]; then ./bin/codemeta; fi

$(PROGRAMS): $(PACKAGE)
	@mkdir -p bin
	go build -o bin/$@$(EXT) cmd/$@/$@.go

man: $(MAN_PAGES)

$(MAN_PAGES): .FORCE
	mkdir -p man/man1
	$(PANDOC) $@.md --from markdown --to man -s >man/man1/$@

CITATION.cff: .FORCE
	cat codemeta.json | sed -E   's/"@context"/"at__context"/g;s/"@type"/"at__type"/g;s/"@id"/"at__id"/g' >_codemeta.json
	if [ -f $(PANDOC) ]; then echo "" | $(PANDOC) --metadata title="Cite $(PROJECT)" --metadata-file=_codemeta.json --template=codemeta-cff.tmpl >CITATION.cff; fi
	if [ -f _codemeta.json ]; then rm _codemeta.json; fi


# NOTE: macOS requires a "mv" command for placing binaries instead of "cp" due to signing process of compile
install: build man
	@echo "Installing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f ./bin/$$FNAME ]; then mv ./bin/$$FNAME $(PREFIX)/bin/$$FNAME; fi; done
	@mkdir -p $(PREFIX)/man/man1
	@for FNAME in $(MAN_PAGES); do if [ -f ./man/man1/$$FNAME ]; then cp -v ./man/man1/$$FNAME $(PREFIX)/man/man1/$$FNAME; fi; done
	@echo ""
	@echo "Make sure $(PREFIX)/bin is in your PATH"


uninstall: .FORCE
	@echo "Removing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f $(PREFIX)/bin/$$FNAME ]; then rm -v $(PREFIX)/bin/$$FNAME; fi; done
	@for FNAME in $(MAN_PAGES); do if [ -f $(PREFIX)/man/man1/$$FNAME ]; then rm -v $(PREFIX)/man/man1/$$FNAME; fi; done

index.md: .FORCE
	cp README.md index.md
	git add index.md

about.md: .FORCE
	cat codemeta.json | sed -E 's/"@context"/"at__context"/g;s/"@type"/"at__type"/g;s/"@id"/"at__id"/g' >_codemeta.json
	if [ -f $(PANDOC) ]; then echo "" | pandoc --metadata title="About $(PROJECT)" --metadata-file=_codemeta.json --template codemeta-md.tmpl >about.md; fi
	if [ -f _codemeta.json ]; then rm _codemeta.json; fi

website: index.md about.md page.tmpl *.md LICENSE css/site.css
	make -f website.mak


test: version.go eputil epfmt doi2eprintxml ep3apid
	- cd cleaner && go test -test.v
	- cd clsrules && go test -test.v
	- go test -timeout 1h -test.v
	./test_cmds.bash


clean:
	@if [ -f version.go ]; then rm version.go; fi
	@if [ -d bin ]; then rm -fR bin; fi
	@if [ -d dist ]; then rm -fR dist; fi

dist/linux-amd64:
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env  GOOS=linux GOARCH=amd64 go build -o dist/bin/$$FNAME cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-linux-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/man1/* 
	@rm -fR dist/bin

dist/macos-amd64:
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=amd64 go build -o dist/bin/$$FNAME cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macos-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/man1/*
	@rm -fR dist/bin

dist/macos-arm64:
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=arm64 go build -o dist/bin/$$FNAME cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macos-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/man1/*
	@rm -fR dist/bin

dist/windows-amd64:
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=amd64 go build -o dist/bin/$$FNAME.exe cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-windows-amd64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/man1/*
	@rm -fR dist/bin

dist/windows-arm64:
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=arm64 go build -o dist/bin/$$FNAME.exe cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-windows-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/man1/*
	@rm -fR dist/bin


dist/raspbian-arm7:
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/$$FNAME cmd/$$FNAME/$$FNAME.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-raspbian-os-arm7.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/man1/*
	@rm -fR dist/bin
  
distribute_docs:
	mkdir -p dist/man/man1
	cp -v codemeta.json dist/
	cp -v CITATION.cff dist/
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -vR man dist/

release: build man CITATION.cff distribute_docs dist/linux-amd64 dist/windows-amd64 dist/windows-arm64 dist/macos-amd64 dist/macos-arm64 dist/raspbian-arm7

status:
	git status

save:
	@if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

publish: website
	./publish.bash

.FORCE:
