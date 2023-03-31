---
title: "simpleutil (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-03-30
---

# NAME

simpleutil

# SYNOPSIS

simpleutil [OPTIONS] SIMPLIFIED_JSON_FILE [OUTPUT_FILENAME]

simpleutil -diff SIMPLIFIED_JSON_FILE SIMPLIFIED_JSON_FILE [OUTPUT_FILENAME]

# DESCRIPTION

simpleutil reads a simplified JSON record, validates and pretty prints
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
simpleutil my-record.json
~~~

Compare the differences between JSON records.

~~~
simpleutil -diff record-old.json record-new.json
~~~



