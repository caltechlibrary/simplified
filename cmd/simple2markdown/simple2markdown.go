package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	// Caltech Library Packages
	"github.com/caltechlibrary/simplified"
)

var (
	helpText = `---
title: "{app_name} (1) user manual | {version} {release_hash}"
author: "R. S. Doiel"
pubDate: {release_date}
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] SIMPLIFIED_JSON_FILE

# DESCRIPTION

{app_name} reads a simplified JSON record and outputs Markdown. This
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
	{app_name} my-record.json > my-record.md
~~~
`
)

func main() {
	var (
		showHelp bool
		showLicense bool
		showVersion bool

		newline bool

		err error
	)
	
	appName := path.Base(os.Args[0])
	// NOTE: the following are set when version.go is generated
	version := simplified.Version
	releaseDate := simplified.ReleaseDate
	releaseHash := simplified.ReleaseHash
	fmtHelp := simplified.FmtHelp

	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&newline, "newline", true, "add a tailing newline")
	flag.Parse()

	args := flag.Args()

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if showHelp {
		fmt.Fprintf(out, "%s\n", fmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", fmtHelp(simplified.LicenseText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s %s\n", appName, version, releaseHash)
		os.Exit(0)
	}

	if len(args) == 0 {
		fmt.Fprintf(eout, "expected the name of a simplified record JSON document or '-' to read from standard input")
		os.Exit(1)
	}

	if len(args) > 0 && args[0] != "-" {
		in, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer in.Close()
	}
	if len(args) > 1 && args[1] != "-" {
		out, err = os.Create(args[1])
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}	
		defer out.Close()
	}

	src, err := io.ReadAll(in)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
	record := new(simplified.Record)
	err = json.Unmarshal(src, &record)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(out, "%s", record.AsMarkdown())
	if newline {
		fmt.Fprintln(out)
	}
	os.Exit(0)
}
