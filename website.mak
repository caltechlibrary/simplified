#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = simplified

MD_PAGES = $(shell ls -1 *.md)

HTML_PAGES = $(shell ls -1 *.md | sed -E 's/.md/.html/g')

build: $(HTML_PAGES) $(MD_PAGES) pagefind

$(HTML_PAGES): $(MD_PAGES) index.html .FORCE 
	pandoc -s --to html5 $(basename $@).md -o $(basename $@).html \
		--metadata title="$(PROJECT) - $@" \
	    --lua-filter=links-to-html.lua \
	    --template=page.tmpl
	git add $(basename $@).html
	if [ "$@" = "INSTALL.html" ]; then cp INSTALL.html install.html; git add install.html; fi
	if [ "$@" = "README.html" ]; then cp README.html index.html; git add index.html; fi

pagefind: .FORCE
	pagefind --verbose --exclude-selectors="nav,header,footer" --site .
	git add pagefind

clean:
	@if [ -f index.html ]; then rm *.html; fi
	@if [ -f README.html ]; then rm *.html; fi

.FORCE:
