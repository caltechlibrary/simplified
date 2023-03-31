---
title: "simple2markdown (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-31
---

# NAME

simple2markdown

# SYNOPSIS

simple2markdown [OPTIONS] SIMPLIFIED_JSON_FILE

# DESCRIPTION

simple2markdown reads a simplified JSON record and outputs Markdown. This
is primarily a test of using the simplified record in Markdown as
a visual reference for the data structures. In practice you would want
to use something like Pandoc with templates to render useful
Markdown or HTML content.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

# EXAMPLE

~~~
	simple2markdown my-record.json > my-record.md
~~~

