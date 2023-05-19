// Package simplified is a package targetting intermediate bibliographic, software and data metadata representation at Caltech Library
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2023, Caltech
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

/**
 * This file implements an intermediate metadata representation
 * suitable for moving data between EPrints 3.3 and Invenio-RDM.
 *
 * See documentation and example on Invenio's structured data:
 *
 * - https://inveniordm.docs.cern.ch/reference/metadata/
 * - https://github.com/caltechlibrary/caltechdata_api/blob/ce16c6856eb7f6424db65c1b06de741bbcaee2c8/tests/conftest.py#L147
 *
 */

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

//
// Top Level Elements
//

// Record implements the top level Invenio 3 record structure
type Record struct {
	// Scheme indicates the schema and version of records
	Schema string `json:"$schema,omitempty"`
	// Interneral persistent identifier for a specific version.
	ID string `json:"id,omitempty"`

	//PID    map[string]interface{} `json:"pid,omitempty"` // Interneral persistent identifier for a specific version.

	// The internal persistent identifier for ALL versions.
	Parent *RecordIdentifier `json:"parent,omitempty"`
	// System-managed external persistent identifiers (DOI, Handles, OAI-PMH identifiers)
	ExternalPIDs map[string]*PersistentIdentifier `json:"pids,omitempty"`
	// Descriptive metadata for the resource
	Metadata *Metadata `json:"metadata,omitempty"`
	// Associated files information.
	Files *Files `json:"files,omitempty"`
	// Access control for record
	RecordAccess *RecordAccess `json:"access,omitempty"`
	// This is the place where RDM custom fields get mapped.
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// Tombstone (deasscession) information.
	Tombstone *Tombstone `json:"tombstone,omitempty"`
	// create time for record
	Created time.Time `json:"created"`
	// modified time for record
	Updated time.Time `json:"updated"`
}

//
// Second level Elements
//

// RecordIdentifier implements the scheme of "parent", a persistant
// identifier to the record.
type RecordIdentifier struct {
	ID     string  `json:"id"`               // The identifier of the parent record
	Access *Access `json:"access,omitempty"` // Access details for the record as a whole
}

// PersistentIdentifier holds an Identifier, e.g. ORCID, ROR, ISNI, GND
type PersistentIdentifier struct {
	Identifier string `json:"identifier,omitempty"` // The identifier value
	Provider   string `json:"provider,omitempty"`   // The provider idenitifier used internally by the system
	Client     string `json:"client,omitempty"`     // The client identifier used for connecting with an external registration service.
}

// RecordAccess implements a datastructure used by Invenio 3 to
// control record level accesss, e.g. in the REST API.
type RecordAccess struct {
	Record  string   `json:"record,omitempty"`  // "public" or "restricted. Read access to the record.
	Files   string   `json:"files,omitempty"`   // "public" or "restricted". Read access to the record's files.
	Embargo *Embargo `json:"embargo,omitempty"` // Embargo options for the record.
}

// Metadata holds the primary metadata about the record. This
// is where most of the EPrints 3.3.x data is mapped into.
type Metadata struct {
	ResourceType           map[string]interface{} `json:"resource_type,omitempty"` // Resource type id from the controlled vocabulary.
	Creators               []*Creator             `json:"creators,omitempty"`      //list of creator information (person or organization)
	Title                  string                 `json:"title"`
	PublicationDate        string                 `json:"publication_date,omitempty"`
	AdditionalTitles       []*TitleDetail         `json:"additional_titles,omitempty"`
	Description            string                 `json:"description,omitempty"`
	AdditionalDescriptions []*Description         `json:"additional_descriptions,omitempty"`
	Rights                 []*Right               `json:"rights,omitempty"`
	Contributors           []*Creator             `json:"contributors,omitempty"`
	Subjects               []*Subject             `json:"subjects,omitempty"`
	Languages              []map[string]interface{}    `json:"languages,omitempty"`
	Dates                  []*DateType            `json:"dates,omitempty"`
	Version                string                 `json:"version,omitempty"`
	Publisher              string                 `json:"publisher,omitempty"`
	Identifiers            []*Identifier          `json:"identifiers,omitempty"`
	RelatedIdentifiers     []*Identifier          `json:"related_identifiers,omitempty"`

	Funding []*Funder `json:"funding,omitempty"`
}

// Files
type Files struct {
	Enabled        bool              `json:"enabled,omitempty"`
	Entries        map[string]*Entry `json:"entries,omitempty"`
	DefaultPreview string            `json:"default_preview,omitempty"`
	Sizes          []string          `json:"sizes,omitempty"`
	Formats        []string          `json:"formats,omitempty"`
	Order          []string          `json:"order,omitempty"`
	Locations      *Location         `json:"locations,omitempty"`
}

type Entry struct {
	BucketID     string `json:"bucket_id,omitempty"`
	VersionID    string `json:"version_id,omitempty"`
	FileID       string `json:"file_id,omitempty"`
	Backend      string `json:"backend,omitempty"`
	StorageClass string `json:"storage_class,omitempty"`
	Key          string `json:"key,omitempty"`
	MimeType     string `json:"mimetype,omitempty"`
	Size         int    `json:"size,omitempty"`
	CheckSum     string `json:"checksum,omitempty"`
}

type Location struct {
	Feature []*Feature `json:"feature,omitempty"`
}

type Feature struct {
	Geometry    *Geometry     `json:"geometry,omitempty"`
	Identifiers []*Identifier `json:"identifiers,omitempty"`
	Place       string        `json:"place,omitempty"`
	Description string        `json:"description,omitempty"`
}

type Geometry struct {
	Type        string    `json:"type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"`
}

// Tombstone
type Tombstone struct {
	Reason    string    `json:"reason,omitempty"`
	Category  string    `json:"category,omitempty"`
	RemovedBy *User     `json:"removed_by,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

//
// Third/Fourth Level Elements
//

// Access is a third level element used by PersistentIdenitifier to
// describe access ownership of the record.
type Access struct {
	OwnedBy []*User `json:"owned_by,omitempty"`
}

// User is a data structured used in Access to describe record
// ownership or user actions.
type User struct {
	User        int    `json:"user,omitempty"`         // User (integer) identifier
	DisplayName string `json:"display_name,omitempty"` // This is my field to quickly associate the internal integer user id with a name for reporting and display.
	Email       string `json:"email,omitempty"`        // This is my field to quickly display a concact email associated with the integer user id.
}

// Embargo is a third level element used by RecordAccess to describe
// the embargo status of a record.
type Embargo struct {
	Active bool   `json:"active,omitempty"` // boolean, is the record under an embargo or not.
	Until  string `json:"until,omitempty"`  // Required if active true. ISO date string. When to lift the embargo. e.g. "2100-10-01"
	Reason string `json:"reason,omitempty"` // Explanation for the embargo
}

//
// Third level elements used in Metadata data structures
//

// Creator of a record's object
type Creator struct {
	PersonOrOrg *PersonOrOrg `json:"person_or_org,omitempty"` // The person or organization.
	Role        *Role        `json:"role,omitempty"`          // The role of the person or organization selected from a customizable controlled vocabularly.
	//	Affiliations []*Affiliation `json:"affiliations,omitempty"`  // Affiliations if `PersonOrOrg.Type` is personal.
}

// Role is an object describing a relationship to authorship
type Role struct {
	ID    string            `json:"id,omitempty"`
	Title map[string]string `json:"title,omitempty"`
	Props map[string]string `json:"props,omitempty"`
}

// PersonOrOrg holds either a person or corporate entity information
// for the creators associated with the record.
type PersonOrOrg struct {
	ID   string `json:"clpid,omitempty" yaml:"clpid,omitempty"` // The Caltech Library internal person or organizational identifier used to cross walk data across library systems. (this is not part of Invenion 3)
	Type string `json:"type,omitempty"`          // The type of name. Either "personal" or "organizational".

	GivenName  string `json:"given_name,omitempty" xml:"given_name,omitempty" yaml:"given_name,omitempty"`   // GivenName holds a peron's given name, e.g. Jane
	FamilyName string `json:"family_name,omitempty" xml:"family_name,omitempty" yaml:"family_name,omitempty"` // FamilyName holds a person's family name, e.g. Doe
	Name       string `json:"name,omitempty" xml:"name,omitempty" yaml:"name,omitempty"`               // Name holds a corporate name, e.g. The Unseen University

	// Identifiers holds a list of unique ID like ORCID, GND, ROR, ISNI
	Identifiers []*Identifier `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// Roles of the person or organization selected from a customizable controlled vocabularly.
	//Role *Role `json:"role,omitempty"`

	// Affiliations if `PersonOrOrg.Type` is personal.
	Affiliations []*Affiliation `json:"affiliations,omitempty" yaml:"affiliations,omitempty"`
}

// Affiliation describes how a person or organization is affialated
// for the purpose of the record.
type Affiliation struct {
	ID   string `json:"id,omitempty" yaml:"id,omitempty"`   // The organizational or institutional id from the controlled vocabularly
	Name string `json:"name,omitempty" yaml:"name,omitempty"` // The name of the organization or institution
	ROR string `json:"ror,omitempty" yaml:"ror,omitempty"`
}

// Identifier holds an Identifier, e.g. ORCID, ROR, ISNI, GND
// for a person for organization it holds GRID, ROR. etc.
type Identifier struct {
	Scheme       string      `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Name         string      `json:"name,omitempty" yaml:"name,omitempty"`
	Title        string      `json:"title,omitempty" yaml:"title,omitempty"`
	Number       string      `json:"number,omitempty" yaml:"number,omitempty"`
	Identifier   string      `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	RelationType *TypeDetail `json:"relation_type,omitempty" yaml:"relation_type,omitempty"`
	ResourceType *TypeDetail `json:"resource_type,omitempty" yaml:"resource_type,omitempty"`
}

func (identifier *Identifier) String() string {
	src, _ := json.Marshal(identifier)
	return fmt.Sprintf("%s", src)
}

// Type is an Invenio 3 e.g. ResourceType, title type or language
type Type struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Title map[string]string `json:"title,omitempty"`
}

// TitleDetail is used by AdditionalTitles in Metadata.
type TitleDetail struct {
	Title string `json:"title,omitempty"`
	Encoding string `json:"en,omitempty"`
	Type  *Type  `json:"type,omitempty"`
	Lang  *Type  `json:"lang,omitempty"`
}

// Description holds additional descriptions in Metadata
// element. e.g. language versions of Abstract, etc.
type Description struct {
	Description string `json:"description,omitempty"`
	Type        *Type  `json:"type,omitempty"`
	Lang        *Type  `json:"lang,omitempty"`
}

// Right holds a specific Rights element for the Metadata's
// list of Rights.
//
// NOTE: for REST API lookup by ID or Title (but not both) should
// be supported at the same end point. I.e. they both must be unique
// with in their set of field values.
type Right struct {
	ID          string       `json:"id,omitempty"`          // Identifier value
	Title       *TitleDetail `json:"title,omitempty"`       // Localized human readable title e.g., `{"en": "The ACME Corporation License."}`.
	Description *Description `json:"description,omitempty"` // Localized license description text e.g., `{"en":"This license ..."}`.
	Link        string       `json:"link,omitempty"`        // Link to full license.
}

// Subject element holds one of a list of subjects
// in the Metadata element.
type Subject struct {
	Subject string `json:"subject,omitempty"`
	ID      string `json:"id,omitempty"`
}

// DateType holds Invenio dates used in Metadata element.
type DateType struct {
	Date        string `json:"date,omitempty"`
	Type        *Type  `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

// Funder holds funding information for funding organizations in Metadata
type Funder struct {
	Funder    *Identifier      `json:"funder,omitempty"`
	Award     *AwardIdentifier `json:"award,omitempty"`
	Reference []*Identifier    `json:"references,omitempty"`
}

type AwardIdentifier struct {
	Scheme       string       `json:"scheme,omitempty"`
	Name         string       `json:"name,omitempty"`
	Title        *TitleDetail `json:"title,omitempty"`
	Number       string       `json:"number,omitempty"`
	Identifier   string       `json:"identifier,omitempty"`
	RelationType *TypeDetail  `json:"relation_type,omitempty"`
	ResourceType *TypeDetail  `json:"resource_type,omitempty"`
}

//
// Additional fourth level elements
//

// Type is an alternate expression of a type where title is map
// with additional info like language. It is used to describe relationships
// and resources in Identifiers. It is a variation of Type.
type TypeDetail struct {
	ID    string                 `json:"id,omitempty"`
	Name  string                 `json:"name,omitempty"`
	Title map[string]interface{} `json:"title,omitempty"`
}

func (rec *Record) ToString() []byte {
	src, _ := json.MarshalIndent(rec, "", "    ")
	return src
}

//
// Utility methods and functions
//

// Diff takes a new Metadata struct and compares it with
// and existing Metadata struct. It rturns two Metadata
// structs with only the different attributes sets.
//
// ```
//
//	src, err := os.ReadFile("old-record.json")
//	// ... handler error ...
//	oldRecord := new(simplified.Record)
//	if err = json.Unmarshal(src, &oldRecord) {
//	    // ... handler error ...
//	}
//	src, err = os.ReadFile("new-record.json")
//	newRecord := new(simplified.Record)
//	if err = json.Unmarshal(src, &newRecord) {
//	    // ... handler error ...
//	}
//	// Now create to new minimal Metadata structs with Diff.
//	o, n := oldRecord.Metadata.Diff(newRecord.Metadata)
//	// Convert to a JSON two cell array of old and new changes
//	src, err = json.MarshalIndent([]*Metadata{o, n}, "", "     ")
//	// ... handler error ...
//	// Print out the formatted JSON
//	fmt.Printf("%s\n", src)
//
// ```
func (m *Metadata) Diff(t *Metadata) (*Metadata, *Metadata) {
	if m == nil && t == nil {
		return nil, nil
	}
	if m == nil {
		return nil, t
	}
	if t == nil {
		return m, t
	}
	oM, nM := new(Metadata), new(Metadata)
	if !reflect.DeepEqual(m.ResourceType, t.ResourceType) {
		oM.ResourceType = m.ResourceType
		nM.ResourceType = t.ResourceType
	}
	if !reflect.DeepEqual(m.Creators, t.Creators) {
		oM.Creators = m.Creators
		nM.Creators = t.Creators
	}
	if strings.Compare(m.Title, t.Title) != 0 {
		oM.Title = m.Title
		nM.Title = t.Title
	}
	if strings.Compare(m.PublicationDate, t.PublicationDate) != 0 {
		oM.PublicationDate = m.PublicationDate
		nM.PublicationDate = t.PublicationDate
	}
	if !reflect.DeepEqual(m.AdditionalTitles, t.AdditionalTitles) {
		oM.AdditionalTitles = m.AdditionalTitles
		nM.AdditionalTitles = t.AdditionalTitles
	}
	if strings.Compare(strings.TrimSpace(m.Description), strings.TrimSpace(t.Description)) != 0 {
		oM.Description = m.Description
		nM.Description = t.Description
	}
	if !reflect.DeepEqual(m.AdditionalDescriptions, t.AdditionalDescriptions) {
		oM.AdditionalDescriptions = m.AdditionalDescriptions
		nM.AdditionalDescriptions = t.AdditionalDescriptions
	}
	if !reflect.DeepEqual(m.Rights, t.Rights) {
		oM.Rights = m.Rights
		nM.Rights = t.Rights
	}
	if !reflect.DeepEqual(m.Contributors, t.Contributors) {
		oM.Contributors = m.Contributors
		nM.Contributors = t.Contributors
	}
	if !reflect.DeepEqual(m.Subjects, t.Subjects) {
		oM.Subjects = m.Subjects
		nM.Subjects = t.Subjects
	}
	if !reflect.DeepEqual(m.Languages, t.Languages) {
		oM.Languages = m.Languages
		nM.Languages = t.Languages
	}
	if !reflect.DeepEqual(m.Dates, t.Dates) {
		oM.Dates = m.Dates
		nM.Dates = t.Dates
	}
	if strings.Compare(m.Version, t.Version) != 0 {
		oM.Version = m.Version
		nM.Version = t.Version
	}
	if strings.Compare(m.Publisher, t.Publisher) != 0 {
		oM.Publisher = m.Publisher
		nM.Publisher = t.Publisher
	}
	if !reflect.DeepEqual(m.Identifiers, t.Identifiers) {
		oM.Identifiers = m.Identifiers
		nM.Identifiers = t.Identifiers
	}
	if !reflect.DeepEqual(m.Funding, t.Funding) {
		oM.Funding = m.Funding
		nM.Funding = t.Funding
	}
	return oM, nM
}

// DiffAsJSON takes a diff of two metadata objects and
// returns a two cell array with the minimal object reflecting
// the changes from old to new.
// ```
//
//	src, err := os.ReadFile("old-record.json")
//	// ... handler error ...
//	oldRecord := new(simplified.Record)
//	if err = json.Unmarshal(src, &oldRecord) {
//	    // ... handler error ...
//	}
//	src, err = os.ReadFile("new-record.json")
//	newRecord := new(simplified.Record)
//	if err = json.Unmarshal(src, &newRecord) {
//	    // ... handler error ...
//	}
//	// Now create our JSON Diff
//	src, err := := oldRecord.Metadata.DiffAsJSON(newRecord.Metadata)
//	// ... handler error ...
//	// Print out the formatted JSON
//	fmt.Printf("%s\n", src)
//
// ```
func (m *Metadata) DiffAsJSON(t *Metadata) ([]byte, error) {
	o, n := m.Diff(t)
	src, err := json.MarshalIndent([]*Metadata{o, n}, "", "    ")
	if err != nil {
		return nil, err
	}
	return src, nil
}

// Diff takes a new Record struct and compares it with
// and existing Record struct. It rturns two Record
// structs with only the different attributes sets.
//
// ```
//
//	src, err := os.ReadFile("old-record.json")
//	// ... handler error ...
//	oldRecord := new(simplified.Record)
//	if err = json.Unmarshal(src, &oldRecord) {
//	    // ... handler error ...
//	}
//	src, err = os.ReadFile("new-record.json")
//	newRecord := new(simplified.Record)
//	if err = json.Unmarshal(src, &newRecord) {
//	    // ... handler error ...
//	}
//	// Now create to new minimal Record structs with Diff.
//	o, n := oldRecord.Diff(newRecord)
//	// Convert to a JSON two cell array of old and new changes
//	src, err = json.MarshalIndent([]*Record{o, n}, "", "     ")
//	// ... handler error ...
//	// Print out the formatted JSON
//	fmt.Printf("%s\n", src)
//
// ```
func (rec *Record) Diff(t *Record) (*Record, *Record) {
	if rec == nil && t == nil {
		return nil, nil
	}
	if rec == nil {
		return nil, t
	}
	if t == nil {
		return rec, nil
	}
	oR, nR := new(Record), new(Record)
	if strings.Compare(rec.Schema, t.Schema) != 0 {
		oR.Schema = rec.Schema
		nR.Schema = t.Schema
	}
	if strings.Compare(rec.ID, t.ID) != 0 {
		oR.ID = rec.ID
		nR.ID = t.ID
	}
	if !reflect.DeepEqual(rec.Parent, t.Parent) {
		oR.Parent = rec.Parent
		nR.Parent = t.Parent
	}
	if !reflect.DeepEqual(rec.ExternalPIDs, t.ExternalPIDs) {
		oR.ExternalPIDs = rec.ExternalPIDs
		nR.ExternalPIDs = t.ExternalPIDs
	}
	if !reflect.DeepEqual(rec.RecordAccess, t.RecordAccess) {
		oR.RecordAccess = rec.RecordAccess
		nR.RecordAccess = t.RecordAccess
	}
	if !reflect.DeepEqual(rec.Metadata, t.Metadata) {
		oR.Metadata, nR.Metadata = rec.Metadata.Diff(t.Metadata)
	}
	if !reflect.DeepEqual(rec.Files, t.Files) {
		oR.Files = rec.Files
		nR.Files = t.Files
	}
	// NOTE: The simplified Record contains the RDM CustomFields
	// map. This needs to be diffed with a map comparison function.
	if !reflect.DeepEqual(rec.CustomFields, t.CustomFields) {
		oR.CustomFields = rec.CustomFields
		nR.CustomFields = t.CustomFields
	}
	if !reflect.DeepEqual(rec.Tombstone, t.Tombstone) {
		oR.Tombstone = rec.Tombstone
		nR.Tombstone = t.Tombstone
	}
	if rec.Created.Compare(t.Created) != 0 {
		oR.Created = rec.Created
		nR.Created = t.Created
	}
	if rec.Updated.Compare(t.Updated) != 0 {
		oR.Updated = rec.Updated
		nR.Updated = t.Updated
	}
	return oR, nR
}

// DiffAsJSON takes a diff of two Record objects and
// returns a two cell array with the minimal Record reflecting
// the changes from old to new.
// ```
//
//	src, err := os.ReadFile("old-record.json")
//	// ... handler error ...
//	oldRecord := new(simplified.Record)
//	if err = json.Unmarshal(src, &oldRecord) {
//	    // ... handler error ...
//	}
//	src, err = os.ReadFile("new-record.json")
//	newRecord := new(simplified.Record)
//	if err = json.Unmarshal(src, &newRecord) {
//	    // ... handler error ...
//	}
//	// Now create our JSON Diff
//	src, err := := oldRecord.DiffAsJSON(newRecord)
//	// ... handler error ...
//	// Print out the formatted JSON
//	fmt.Printf("%s\n", src)
//
// ```
func (m *Record) DiffAsJSON(t *Record) ([]byte, error) {
	o, n := m.Diff(t)
	src, err := json.MarshalIndent([]*Record{o, n}, "", "    ")
	if err != nil {
		return nil, err
	}
	return src, nil
}
