
# simplified

simplified is a Go package for bibliographic, software and data metadata. It is used internally by Caltech Library as suitable for mapping one specific metadata form (e.g. EPrints 3.3 EPrint XML) to another (e.g. Invenio-RDM 11 records). It is also used in our feeds system which provides an aggregated view our institutional author, thesis and data repositories.  The simple record structure is inspired in part by the DataCite and Invenio RDM metadata models.

Two utilities demonstrate how you might use the simplified package in a 
Go program.

- [simpleutil](simpleutil.1.md) will pretty print a JSON record or let you take a diff of two JSON file
- [simple2markdown](simple2markdown.1.md) is a proof of concept of rendering Markdown documents from a simple record (e.g. for a landing pages describing a metadata record).


## References

- [DataCite's Metadata Schema v4](https://schema.datacite.org/meta/kernel-4.4/)
- [Invenio RDM metadata model](https://inveniordm.docs.cern.ch/reference/metadata/)
- [CrossRef's API](https://api.crossref.org/swagger-ui/index.html)
