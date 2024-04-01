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
	if simpleRecord1.ExternalPIDs == nil {
		t.Errorf("Expected unmarshaled record to have pids -> %s", example1Text)
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

// TestSimplifiedVersionStruct
func TestSimplifiedVersionsStruct(t *testing.T) {
	src1 := []byte(`{"is_published": true, "revision_id": 7, "is_draft": false, "pids": {"oai": {"identifier": "oai:authors.library.caltech.edu:rd9fg-k5282", "provider": "oai"}}, "custom_fields": {"caltech:groups": [{"id": "Division-of-Geological-and-Planetary-Sciences", "title": {"en": "Division of Geological and Planetary Sciences"}}, {"id": "Caltech-Center-for-Environmental-Microbial-Interactions-(CEMI)", "title": {"en": "Caltech Center for Environmental Microbial Interactions (CEMI)"}}, {"id": "Division-of-Biology-and-Biological-Engineering", "title": {"en": "Division of Biology and Biological Engineering"}}], "journal:journal": {"issn": "0027-8424", "issue": "51", "title": "Proceedings of the National Academy of Sciences", "volume": "120", "pages": "e2302156120"}}, "files": {"enabled": true, "count": 6, "entries": {"pnas.2302156120.sd03.xlsx": {"mimetype": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "key": "pnas.2302156120.sd03.xlsx", "metadata": null, "ext": "xlsx", "size": 19676}, "pnas.2302156120.sd01.xlsx": {"mimetype": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "key": "pnas.2302156120.sd01.xlsx", "metadata": null, "ext": "xlsx", "size": 14335}, "pnas.2302156120.sapp.pdf": {"mimetype": "application/pdf", "key": "pnas.2302156120.sapp.pdf", "metadata": null, "ext": "pdf", "size": 5195716}, "pnas.2302156120.sd04.xlsx": {"mimetype": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "key": "pnas.2302156120.sd04.xlsx", "metadata": null, "ext": "xlsx", "size": 10516}, "osorio-rodriguez-et-al-2023-microbially-induced-precipitation-of-silica-by-anaerobic-methane-oxidizing-consortia-and.pdf": {"mimetype": "application/pdf", "key": "osorio-rodriguez-et-al-2023-microbially-induced-precipitation-of-silica-by-anaerobic-methane-oxidizing-consortia-and.pdf", "metadata": null, "ext": "pdf", "size": 4353482}, "pnas.2302156120.sd02.xlsx": {"mimetype": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "key": "pnas.2302156120.sd02.xlsx", "metadata": null, "ext": "xlsx", "size": 22773}}, "order": [], "total_bytes": 9616498, "default_preview": "osorio-rodriguez-et-al-2023-microbially-induced-precipitation-of-silica-by-anaerobic-methane-oxidizing-consortia-and.pdf"}, "status": "published", "metadata": {"description": "<p>Authigenic carbonate minerals can preserve biosignatures of microbial anaerobic oxidation of methane (AOM) in the rock record. It is not currently known whether the microorganisms that mediate sulfate-coupled AOM\u2014often occurring as multicelled consortia of anaerobic methanotrophic archaea (ANME) and sulfate-reducing bacteria (SRB)\u2014are preserved as microfossils. Electron microscopy of ANME-SRB consortia in methane seep sediments has shown that these microorganisms can be associated with silicate minerals such as clays [Chen <i>et al</i>., <i>Sci. Rep.</i> 4 , 1\u20139 (2014)], but the biogenicity of these phases, their geochemical composition, and their potential preservation in the rock record is poorly constrained. Long-term laboratory AOM enrichment cultures in sediment-free artificial seawater [Yu <i>et al</i>., <i>Appl. Environ. Microbiol.</i> 88 , e02109-21 (2022)] resulted in precipitation of amorphous silicate particles (~200 nm) within clusters of exopolymer-rich AOM consortia from media undersaturated with respect to silica, suggestive of a microbially mediated process. The use of techniques like correlative fluorescence in situ hybridization (FISH), scanning electron microscopy with energy dispersive X-ray spectroscopy (SEM-EDS), and nanoscale secondary ion mass spectrometry (nanoSIMS) on AOM consortia from methane seep authigenic carbonates and sediments further revealed that they are enveloped in a silica-rich phase similar to the mineral phase on ANME-SRB consortia in enrichment cultures. Like in cyanobacteria [Moore <i>et al</i>., <i>Geology</i> 48 , 862\u2013866 (2020)], the Si-rich phases on ANME-SRB consortia identified here may enhance their preservation as microfossils. The morphology of these silica-rich precipitates, consistent with amorphous-type clay-like spheroids formed within organic assemblages, provides an additional mineralogical signature that may assist in the search for structural remnants of microbial consortia in rocks which formed in methane-rich environments from Earth and other planetary bodies.</p>", "funding": [{"funder": {"id": "01bj3aw27", "name": "United States Department of Energy"}, "award": {"number": "DE-C0020373", "title": {"en": " "}}}, {"funder": {"id": "021nxhr62", "name": "National Science Foundation"}, "award": {"number": "OCE-1634002", "title": {"en": " "}}}, {"funder": {"id": "006wxqw41", "name": "Gordon and Betty Moore Foundation"}, "award": {"number": "3780", "title": {"en": " "}}}, {"funder": {"id": "01cmst727", "name": "Simons Foundation"}}, {"funder": {"id": "05dxps055", "name": "California Institute of Technology"}, "award": {"number": "Caltech Center for Environmental Microbial Interactions", "title": {"en": " "}}}, {"funder": {"id": "021nxhr62", "name": "National Science Foundation"}, "award": {"number": "NSF Graduate Research Fellowship", "title": {"en": " "}}}, {"funder": {"id": "01sdtdd95", "name": "Canadian Institute for Advanced Research"}}], "title": "Microbially induced precipitation of silica by anaerobic methane-oxidizing consortia and implications for microbial fossil preservation", "related_identifiers": [{"scheme": "url", "relation_type": {"id": "issupplementedby", "title": {"de": "Wird erg\u00e4nzt durch", "en": "Is supplemented by"}}, "identifier": "https://www.pnas.org/doi/suppl/10.1073/pnas.2302156120/suppl_file/pnas.2302156120.sapp.pdf"}, {"scheme": "url", "relation_type": {"id": "issupplementedby", "title": {"de": "Wird erg\u00e4nzt durch", "en": "Is supplemented by"}}, "identifier": "https://www.pnas.org/doi/suppl/10.1073/pnas.2302156120/suppl_file/pnas.2302156120.sd01.xlsx"}, {"scheme": "url", "relation_type": {"id": "issupplementedby", "title": {"de": "Wird erg\u00e4nzt durch", "en": "Is supplemented by"}}, "identifier": "https://www.pnas.org/doi/suppl/10.1073/pnas.2302156120/suppl_file/pnas.2302156120.sd02.xlsx"}, {"scheme": "url", "relation_type": {"id": "issupplementedby", "title": {"de": "Wird erg\u00e4nzt durch", "en": "Is supplemented by"}}, "identifier": "https://www.pnas.org/doi/suppl/10.1073/pnas.2302156120/suppl_file/pnas.2302156120.sd03.xlsx"}, {"scheme": "url", "relation_type": {"id": "issupplementedby", "title": {"de": "Wird erg\u00e4nzt durch", "en": "Is supplemented by"}}, "identifier": "https://www.pnas.org/doi/suppl/10.1073/pnas.2302156120/suppl_file/pnas.2302156120.sd04.xlsx"}], "identifiers": [{"scheme": "issn", "identifier": "1091-6490"}, {"scheme": "doi", "identifier": "10.1073/pnas.230215612"}], "resource_type": {"id": "publication-article", "title": {"de": "Zeitschriftenartikel", "en": "Journal Article"}}, "additional_descriptions": [{"type": {"id": "copyright", "title": {"de": "Sonstige", "en": "Copyright and License"}}, "description": "<p>\u00a9 2023 the Author(s). Published by PNAS. This article is distributed under <a href=\"https://creativecommons.org/licenses/by-nc-nd/4.0/\">Creative Commons Attribution-NonCommercial-NoDerivatives License 4.0 (CC BY-NC-ND)</a>.</p>"}, {"type": {"id": "acknowledgement", "title": {"de": "Sonstige", "en": "Acknowledgement"}}, "description": "<p>We would like to acknowledge Jake Bailey, Tom Bristow, Ted Present, Paul Myrow, Kelsey Moore, and Jen Glass for contributing to discussions about this work and the thoughtful comments from K. Konhauser and anonymous reviewers who improved this work. We would like to thank Chi Ma for assistance with SEM-EDS, Yunbin Guan for assistance with NanoSIMS, and Nathan Dalleska for help with ICP-MS analysis. We are additionally grateful to Alice Michel for discussions regarding sample preparation for FISH-SEM. We further acknowledge the crews of R/V <i>Atlantis</i> and R/V <i>Western Flyer</i> for assistance in sample collection and processing. Funding for this work was provided by the US Department of Energy's Office of Science (DE-SC0020373), the NSF BIO-OCE grant (#1634002), a Gordon and Betty Moore Foundation Marine Microbiology Investigator grant (#3780), the Simons Collaboration for the Origin of Life, and a grant from the Center for Environmental Microbial Interactions (to V.J.O.). K.S.M. was supported in part by a NSF Graduate Research Fellowship and a Schlanger Ocean Drilling Fellowship. V.J.O. is a CIFAR science fellow in the Earth 4D program. Portions of this work were developed from the doctoral dissertation of Kyle Metcalfe, Symbiotic Diversity and Mineral-Associated Microbial Ecology in Marine Microbiomes; Anne Dekas, Diazotrophy in the Deep: An Analysis of the Distribution, Magnitude, Geochemical Controls, and Biological Mediators of Deep-Sea Benthic Nitrogen Fixation; and Daniela Osorio Rodriguez, Microbial Transformations of Sulfur: Environmental and (Paleo) Ecological Implications.</p>"}, {"type": {"id": "contributions", "title": {"de": "Sonstige", "en": "Contributions"}}, "description": "<p>D.O.-R., K.S.M., J.P.G., and V.J.O. designed research; D.O.-R., K.S.M., S.E.M., H.Y., A.E.D., M.E., and T.D. performed research; D.O.-R., K.S.M., S.E.M., L.A., J.P.G., and V.J.O. analyzed data; and D.O.-R. and K.S.M. wrote the paper.</p>"}, {"type": {"id": "data-availability", "title": {"de": "Sonstige", "en": "Data Availability"}}, "description": "<p>All data and custom scripts were collected and stored using Git version control. Code for raw data processing, analysis, and figure generation is available in the GitHub repository (<a href=\"https://github.com/daniosro/Si_biomineralization_ANME_SRB\">https://github.com/daniosro/Si_biomineralization_ANME_SRB</a>) (<a href=\"https://www.pnas.org/doi/10.1073/pnas.2302156120?ai=11xml&amp;ui=cdok&amp;af=T#core-r92\">92</a>). All other data are included in the manuscript and/or <a href=\"http://www.pnas.org/lookup/doi/10.1073/pnas.2302156120#supplementary-materials\">supporting information</a>.</p>"}, {"type": {"id": "coi", "title": {"de": "Sonstige", "en": "Conflict of Interest"}}, "description": "<p>he authors declare no competing interest.</p>"}], "publisher": "National Academy of Sciences", "subjects": [{"subject": "Multidisciplinary"}], "rights": [{"description": {"en": ""}, "id": "cc-by-nc-nd-4.0", "title": {"en": "Creative Commons Attribution Non Commercial No Derivatives 4.0 International"}, "props": {"scheme": "spdx", "url": "https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode"}}], "creators": [{"person_or_org": {"type": "personal", "identifiers": [{"scheme": "orcid", "identifier": "0000-0001-6676-4124"}, {"scheme": "clpid", "identifier": "Osorio-Rodriguez-Daniela"}], "given_name": "Daniela", "name": "Osorio-Rodriguez, Daniela", "family_name": "Osorio-Rodriguez"}}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "clpid", "identifier": "Metcalfe-Kyle-S"}, {"scheme": "orcid", "identifier": "0000-0002-2963-765X"}], "given_name": "Kyle S.", "name": "Metcalfe, Kyle S.", "family_name": "Metcalfe"}}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "orcid", "identifier": "0000-0002-8199-7011"}, {"scheme": "clpid", "identifier": "McGlynn-Shawn-E"}], "given_name": "Shawn E.", "name": "McGlynn, Shawn E.", "family_name": "McGlynn"}}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "clpid", "identifier": "Yu-Hang"}, {"scheme": "orcid", "identifier": "0000-0002-7600-1582"}], "given_name": "Hang", "name": "Yu, Hang", "family_name": "Yu"}}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "orcid", "identifier": "0000-0001-9548-8413"}, {"scheme": "clpid", "identifier": "Dekas-Anne-E"}], "given_name": "Anne E.", "name": "Dekas, Anne E.", "family_name": "Dekas"}}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "orcid", "identifier": "0000-0001-8893-8455"}, {"scheme": "clpid", "identifier": "Ellisman-Mark"}], "given_name": "Mark", "name": "Ellisman, Mark", "family_name": "Ellisman"}}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "clpid", "identifier": "Deerinck-Tom"}], "given_name": "Tom", "name": "Deerinck, Tom", "family_name": "Deerinck"}}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "orcid", "identifier": "0000-0002-8566-1486"}, {"scheme": "clpid", "identifier": "Aristilde-Ludmilla"}], "given_name": "Ludmilla", "name": "Aristilde, Ludmilla", "family_name": "Aristilde"}}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "clpid", "identifier": "Grotzinger-J-P"}, {"scheme": "orcid", "identifier": "0000-0001-9324-1257"}], "given_name": "John P.", "name": "Grotzinger, John P.", "family_name": "Grotzinger"}, "affiliations": [{"id": "05dxps055", "name": "California Institute of Technology"}]}, {"person_or_org": {"type": "personal", "identifiers": [{"scheme": "clpid", "identifier": "Orphan-V-J"}, {"scheme": "orcid", "identifier": "0000-0002-5374-6178"}], "given_name": "Victoria J.", "name": "Orphan, Victoria J.", "family_name": "Orphan"}, "affiliations": [{"id": "05dxps055", "name": "California Institute of Technology"}]}], "languages": [{"id": "eng", "title": {"en": "English"}}], "publication_date": "2023-12-19"}, "links": {"self": "https://authors.library.caltech.edu/api/records/rd9fg-k5282", "self_html": "https://authors.library.caltech.edu/records/rd9fg-k5282", "self_iiif_manifest": "https://authors.library.caltech.edu/api/iiif/record:rd9fg-k5282/manifest", "self_iiif_sequence": "https://authors.library.caltech.edu/api/iiif/record:rd9fg-k5282/sequence/default", "files": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/files", "archive": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/files-archive", "latest": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/versions/latest", "latest_html": "https://authors.library.caltech.edu/records/rd9fg-k5282/latest", "draft": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/draft", "versions": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/versions", "access_links": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/access/links", "reserve_doi": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/draft/pids/doi", "communities": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/communities", "communities-suggestions": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/communities-suggestions", "requests": "https://authors.library.caltech.edu/api/records/rd9fg-k5282/requests"}, "created": "2023-12-12T20:21:52.455369+00:00", "id": "rd9fg-k5282", "updated": "2024-01-09T22:24:02.634266+00:00", "stats": {"this_version": {"data_volume": 0.0, "downloads": 0, "unique_views": 0, "views": 0, "unique_downloads": 0}, "all_versions": {"data_volume": 0.0, "downloads": 0, "unique_views": 0, "views": 0, "unique_downloads": 0}}, "parent": {"id": "xkx1h-9ks92", "access": {"owned_by": [{"user": 13}], "links": []}, "communities": {"default": "aedd135f-227e-4fdf-9476-5b3fd011bac6", "ids": ["aedd135f-227e-4fdf-9476-5b3fd011bac6"]}}, "versions": {"is_latest": true, "is_latest_draft": true, "index": 2}, "access": {"embargo": {"active": false, "reason": null}, "status": "open", "files": "public", "record": "public"}}`)

	rec := new(Record)
	if err := json.Unmarshal(src1, &rec); err != nil {
		t.Error(err)
	}
	if rec.Versions == nil {
		t.Errorf(".versions should not be nil")
		t.FailNow()
	}
	if ! rec.Versions.IsLatest {
		t.Errorf(".version.is_latest should be true, got false")
	}
	if ! rec.Versions.IsLatestDraft {
		t.Errorf(".version.is_latest_draft should be true, got false")
	}
	if rec.Versions.Index != 2 {
		t.Errorf(".version.index to be 2, got %d", rec.Versions.Index)
	}

}
