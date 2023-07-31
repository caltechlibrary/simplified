// Package simplified is a package targetting intermediate bibliographic, software and data metadata representation at Caltech Library
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2021, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package simplified

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestSimplifiedRecord checks if the a record rendered in a simplified form
// can be round trip to/from an EPrint type struct
func TestSimplifiedRecord(t *testing.T) {
	// Test data is taken from   https://github.com/inveniosoftware/invenio-rdm-records/blob/396b1f2ff802e8483e30fa2e42cfe03f597d1e87/tests/conftest.py#L361
	example1Text := []byte(`{
        "pids": {
            "doi": {
                "identifier": "10.5281/inveniordm.1234",
                "provider": "datacite",
                "client": "inveniordm"
            }
        },
        "metadata": {
            "resource_type": {"id": "image-photo"},
            "creators": [
                {
                    "person_or_org": {
                        "name": "Nielsen, Lars Holm",
                        "type": "personal",
                        "given_name": "Lars Holm",
                        "family_name": "Nielsen",
                        "identifiers": [
                            {"scheme": "orcid", "identifier": "0000-0001-8135-3489"}
                        ]
                    },
                    "affiliations": [{"id": "cern"}, {"name": "free-text"}]
                }
            ],
            "title": "InvenioRDM",
            "additional_titles": [
                {
                    "title": "a research data management platform",
                    "type": {"id": "subtitle"},
                    "lang": {"id": "eng"}
                }
            ],
            "publisher": "InvenioRDM",
            "publication_date": "2018/2020-09",
            "subjects": [
                {"id": "http://id.nlm.nih.gov/mesh/A-D000007"},
                {"subject": "custom"}
            ],
            "contributors": [
                {
                    "person_or_org": {
                        "name": "Nielsen, Lars Holm",
                        "type": "personal",
                        "given_name": "Lars Holm",
                        "family_name": "Nielsen",
                        "identifiers": [
                            {"scheme": "orcid", "identifier": "0000-0001-8135-3489"}
                        ]
                    },
                    "role": { 
                        "id": "other"
                    },
                    "affiliations": [{"id": "cern"}]
                }
            ],
            "dates": [
                {"date": "1939/1945", "type": {"id": "other"}, "description": "A date"}
            ],
            "languages": [{"id": "dan"}, {"id": "eng"}],
            "identifiers": [{"identifier": "1924MNRAS..84..308E", "scheme": "bibcode"}],
            "related_identifiers": [
                {
                    "identifier": "10.1234/foo.bar",
                    "scheme": "doi",
                    "relation_type": {"id": "iscitedby"},
                    "resource_type": {"id": "dataset"}
                }
            ],
            "sizes": ["11 pages"],
            "formats": ["application/pdf"],
            "version": "v1.0",
            "rights": [
                {
                    "title": {"en": "A custom license"},
                    "description": {"en": "A description"},
                    "link": "https://customlicense.org/licenses/by/4.0/"
                },
                {"id": "cc-by-4.0"}
            ],
            "description": "<h1>A description</h1> <p>with HTML tags</p>",
            "additional_descriptions": [
                {
                    "description": "Bla bla bla",
                    "type": {"id": "methods"},
                    "lang": {"id": "eng"}
                }
            ],
            "locations": {
                "features": [
                    {
                        "geometry": {
                            "type": "Point",
                            "coordinates": [-32.94682, -60.63932]
                        },
                        "place": "test location place",
                        "description": "test location description",
                        "identifiers": [
                            {"identifier": "12345abcde", "scheme": "wikidata"},
                            {"identifier": "12345abcde", "scheme": "geonames"}
                        ]
                    }
                ]
            },
            "funding": [
                {
                    "funder": {
                        "id": "00k4n6c32"
                    },
                    "award": {"id": "00k4n6c32::755021"}
                }
            ],
            "references": [
                {
                    "reference": "Nielsen et al,..",
                    "identifier": "0000 0001 1456 7559",
                    "scheme": "isni"
                }
            ]
        },
        "ext": {
            "dwc": {
                "collectionCode": "abc",
                "collectionCode2": 1.1,
                "collectionCode3": true,
                "test": [ "abc", 1, true ]
            }
        },
        "provenance": {
            "created_by": { "user": "<USER_ID_1>" },
            "on_behalf_of": { "user": "<USER_ID_2>" }
        },
        "access": {
            "record": "public",
            "files": "restricted",
            "embargo": {
                "active": true,
                "until": "2131-01-01",
                "reason": "Only for medical doctors."
            }
        },
        "files": {
            "enabled": true,
            "total_bytes": 1114324524355,
            "count": 1,
            "order": ["big-dataset.zip"],
            "entries": {
				"big-dataset.zip": {
                    "checksum": "md5:234245234213421342",
                    "mimetype": "application/zip",
                    "size": 1114324524355,
                    "key": "big-dataset.zip",
                    "file_id": "445aaacd-9de1-41ab-af52-25ab6cb93df7"
                }
			}
        },
        "notes": ["Under investigation for copyright infringement."]
    }`)
	simpleRecord1 := new(Record)
	if err := json.Unmarshal(example1Text, &simpleRecord1); err != nil {
		t.Errorf("Unmarshal failed, %s", err)
		t.FailNow()
	}
	simpleRecord2 := new(Record)
	if err := json.Unmarshal(example1Text, &simpleRecord2); err != nil {
		t.Errorf("Unmarshal failed, %s", err)
		t.FailNow()
	}

	if !reflect.DeepEqual(simpleRecord1, simpleRecord2) {
		t.Errorf("expected deepEqual(simpleRecord1, simpleRecord2)")
		t.FailNow()
	}

	// test custom fields of rdm:journal
	journal1 := map[string]interface{}{
		"rdm:journal": map[string]interface{}{
			"title":          "Robert's Journal of Abstract Robots",
			"year":           "2023",
			"volume":         "23",
			"issue":          "4",
			"pages":          "232-244",
			"article_number": 3,
		},
	}
	journal2 := map[string]interface{}{}
	for k, v := range journal1 {
		journal2[k] = v
	}
	simpleRecord2.CustomFields = journal2

	src1, _ := json.MarshalIndent(simpleRecord1, "", "    ")
	src2, _ := json.MarshalIndent(simpleRecord2, "", "    ")

	if reflect.DeepEqual(simpleRecord1, simpleRecord2) {
		t.Errorf("simpleRecord should not have a .custom_fields element ->\n%s\n============\n%s\n", src1, src2)
		t.FailNow()
	}
	simpleRecord1.CustomFields = journal1

	src1, _ = json.MarshalIndent(simpleRecord1, "", "    ")
	src2, _ = json.MarshalIndent(simpleRecord2, "", "    ")

	if !reflect.DeepEqual(simpleRecord1, simpleRecord2) {
		t.Errorf("expected simpleRecord1 to be same as simpleRecord2 ->\n%s\n=============\n%s\n", src1, src2)
		t.FailNow()
	}
}

// TestSimplifiedFilesListing tests the structure used by
// /api/records/{record_id}/files. This is distinct from
// the "files" attributed returned by /api/records/{record_id} 
// data structure.
func TestSimplifiedFileListing(t *testing.T) {
	src := []byte(`{"enabled": true, "links": {"self": "https://authors.caltechlibrary.dev/api/records/v3szs-vn773/files", "archive": "https://authors.caltechlibrary.dev/api/records/v3szs-vn773/files-archive"}, "entries": [{"updated": "2023-07-13T00:15:29.869760+00:00", "size": 7076180, "file_id": "64383771-75c2-4e00-9b84-06ffbe3774af", "status": "completed", "key": "jcb_202212007.pdf", "mimetype": "application/pdf", "links": {"self": "https://authors.caltechlibrary.dev/api/records/v3szs-vn773/files/jcb_202212007.pdf", "content": "https://authors.caltechlibrary.dev/api/records/v3szs-vn773/files/jcb_202212007.pdf/content"}, "version_id": "10f0b6c1-9bed-472c-b2eb-090f4ae207f5", "created": "2023-07-13T00:15:29.867026+00:00", "bucket_id": "ceac4045-ebe3-415f-9b56-2849002212c7", "checksum": "md5:ae98c6b97e989fb3a19d2dfd33abf778", "metadata": null, "storage_class": "L"}, {"updated": "2023-07-13T00:15:29.874834+00:00", "size": 6248996, "file_id": "eb99d17f-8bcc-4733-ba73-f4c618526037", "status": "completed", "key": "jcb_202212007_sourcedataf1.pdf", "mimetype": "application/pdf", "links": {"self": "https://authors.caltechlibrary.dev/api/records/v3szs-vn773/files/jcb_202212007_sourcedataf1.pdf", "content": "https://authors.caltechlibrary.dev/api/records/v3szs-vn773/files/jcb_202212007_sourcedataf1.pdf/content"}, "version_id": "705aff0d-7b7d-4c3e-978a-896543641b01", "created": "2023-07-13T00:15:29.872416+00:00", "bucket_id": "ceac4045-ebe3-415f-9b56-2849002212c7", "checksum": "md5:14a92692ad011f8b126b22b65c875e25", "metadata": null, "storage_class": "L"}, {"updated": "2023-07-13T00:15:29.879525+00:00", "size": 16592, "file_id": "287470c0-1a3e-4da6-8b9b-e6a3163a3b66", "status": "completed", "key": "jcb_202212007_tables1.docx", "mimetype": "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "links": {"self": "https://authors.caltechlibrary.dev/api/records/v3szs-vn773/files/jcb_202212007_tables1.docx", "content": "https://authors.caltechlibrary.dev/api/records/v3szs-vn773/files/jcb_202212007_tables1.docx/content"}, "version_id": "703bc259-083b-4afe-921d-d76e30f5051f", "created": "2023-07-13T00:15:29.877297+00:00", "bucket_id": "ceac4045-ebe3-415f-9b56-2849002212c7", "checksum": "md5:76a7e4da98915de9be75a5c60807b5ec", "metadata": null, "storage_class": "L"}], "default_preview": null, "order": []}`)
	obj := new(Entry)
	if err := json.Unmarshal(src, &obj); err != nil {
		t.Error(err)
	}
}

// TestSimplifiedFilesStruct is the structure used when submitting
// files data via /api/records/{record_id}. This is distinct from
// the /api/records/{record_id}/files data structure.
func TestSimplifiedFilesStruct(t *testing.T) {
	src := []byte(`{
  "enabled": true,
  "entries": {
    "Goodtimebanjo-banjores-resun-rescov-guit-same-force-same-percent.mp3": {
      "key": "Goodtimebanjo-banjores-resun-rescov-guit-same-force-same-percent.mp3",
      "mimetype": "audio/mpeg",
      "size": 29635
    },
    "Vegabanjo-banjores-resun-rescov-guitar-undamped.mp3": {
      "key": "Vegabanjo-banjores-resun-rescov-guitar-undamped.mp3",
      "mimetype": "audio/mpeg",
      "size": 363922
    },
    "all-5-takes-adjusted.mp3": {
      "key": "all-5-takes-adjusted.mp3",
      "mimetype": "audio/mpeg",
      "size": 2998517
    },
    "goodtime-banjores-uncovered-covered-guitar-1st-only-12th-fret-all-damped-but-1.mp3": {
      "key": "goodtime-banjores-uncovered-covered-guitar-1st-only-12th-fret-all-damped-but-1.mp3",
      "mimetype": "audio/mpeg",
      "size": 358875
    },
    "resonator-guitar.pdf": {
      "key": "resonator-guitar.pdf",
      "mimetype": "application/pdf",
      "size": 12166506
    }
  },
  "total_bytes": 15917455,
  "count": 5
}`)
	obj := new(Entry)
	if err := json.Unmarshal(src, &obj); err != nil {
		t.Error(err)
	}
}

