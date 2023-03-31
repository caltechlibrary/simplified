
# simplified

simplified is a Go package for bibliographic, software and data metadata record structure and representation used by Caltech Library suitable for mapping one specific metadata form (e.g. EPrints 3.3 EPrint XML) to another (e.g. Invenio-RDM 11 records). It is also used in our feeds system to aggregate our various repository system records into a common re-useable format.  simple record is inspited in part by the DataCite metadata record structure.

Two utilities demonstrate how you might use the simplified package.

- [simpleutil](simpleutil.1.md) will pretty print a JSON record or let you take a diff of two JSON file
- [simple2markdown](simple2markdown.1.md) is a proof of concept of rendering Markdown documents from a simple record (e.g. for a landing pages describing a metadata record).

This is a proof of concept package and not intended for production use.



