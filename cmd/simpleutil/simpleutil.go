package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/simplified"
)

var (
	helpText = `---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-03-30
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] SIMPLIFIED_JSON_FILE [OUTPUT_FILENAME]

{app_name} -diff SIMPLIFIED_JSON_FILE SIMPLIFIED_JSON_FILE [OUTPUT_FILENAME]

# DESCRIPTION

{app_name} reads a simplified JSON record, validates and pretty prints
the JSON output.  This is primarily a test of using the simplified record.
The "-diff" option will take to JSON records and perform a diff operation
returning a JSON array with the first cell holding the in difference from
the first file and the second one holding the attribitutes in difference
for the second file.

You can use a filename of "-" to read input from standard input.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-diff 
: will difference two simple records in JSON files.


# EXAMPLES

Pretty print a simplified JSON record.

~~~
{app_name} my-record.json
~~~

Compare the differences between JSON records.

~~~
{app_name} -diff record-old.json record-new.json
~~~


`
)

func fmtTxt(src string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(src, "{app_name}", appName), "{version}", version)
}

func main() {
	var (
		showHelp bool
		showLicense bool
		showVersion bool
		diffRecords bool 

		err error
	)
	
	appName := path.Base(os.Args[0])
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&diffRecords, "diff", false, "display difference between two JSON records")
	flag.Parse()

	args := flag.Args()

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if showHelp {
		fmt.Fprintf(out, "%s\n", fmtTxt(helpText, appName, simplified.Version))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", fmtTxt(simplified.LicenseText, appName, simplified.Version))
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s\n", appName, simplified.Version)
		os.Exit(0)
	}

	if len(args) == 0 {
		fmt.Fprintf(eout, "expected the name of a simplified record JSON document or '-' to read from standard input")
		os.Exit(1)
	}

	if diffRecords {
		in1 := os.Stdin
		in2 := os.Stdin
		if len(args) > 0 && args[0] != "-" {
			in1, err = os.Open(args[0])
			if err != nil {
				fmt.Fprintf(eout, "%s\n", err)
				os.Exit(1)
			}
			defer in1.Close()
		}
		if len(args) > 1 && args[1] != "-" {
			in2, err = os.Open(args[1])
			if err != nil {
				fmt.Fprintf(eout, "%s\n", err)
				os.Exit(1)
			}	
			defer in2.Close()
		}
		if len(args) > 2 && args[2] != "-" {
			out, err = os.Create(args[2])
			if err != nil {
				fmt.Fprintf(eout, "%s\n", err)
				os.Exit(1)
			}	
			defer out.Close()
		}
		src1, err := io.ReadAll(in1)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		src2, err := io.ReadAll(in2)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}

		rec1, rec2 := new(simplified.Record), new(simplified.Record)

		err = json.Unmarshal(src1, &rec1)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		err = json.Unmarshal(src2, &rec2)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}

		if src, err := rec1.DiffAsJSON(rec2); err != nil  {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		} else {
			fmt.Fprintf(out, "%s\n", src)
			os.Exit(0)
		}
	} else {
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
		fmt.Fprintf(out, "%s\n", record.AsMarkdown())

		os.Exit(0)
	}
}
