package simplified

import (
	"fmt"
	"strings"
)

// An individual person record
type Person struct {
	Name         string         `json:"name,omitempty" yaml:"name,omitempty"`
	Sort         string         `json:"sort_name,omitempty" yaml:"sort_name,omitempty"`
	Family       string         `json:"family_name,omitempty" yaml:"family_name,omitempty"`
	Given        string         `json:"given_name,omitempty" yaml:"given_name,omitempty"`
	Identifiers  []*Identifier  `json:"identifiers,omitempty" yaml:"identifiers,omitempty"`
	// NOTE: Affiliations are captured at the PersonOrOrg level in RDM, this
	// is in Person which is used in our vocabulary generation.
	Affiliations []*Affiliation `json:"affiliations,omitempty" yaml:"affiliations,omitempty"`
}

func (p *Person) GetIdentifier(scheme string) string {
	for _, id := range p.Identifiers {
		if strings.Compare(id.Scheme, scheme) == 0 {
			return id.Identifier
		}
	}
	return ""
}

// Resolve takes the person object and resolves missing attributes
// based on how InvenioRDM handles things.
func (p *Person) Resolve() error {
	if p.Name != "" && p.Family == "" && p.Given == "" && strings.Contains(p.Name, ", ") {
		parts := strings.Split(p.Name, ",")
		if len(parts) == 2 {
			p.Family = strings.TrimSpace(parts[0])
			p.Given = strings.TrimSpace(parts[1])
		}
	}
	if p.Name == "" && len(p.Family) > 0 && len(p.Given) > 0 {
		p.Name = fmt.Sprintf("%s, %s", p.Family, p.Given)
	}
	if p.Sort == "" && len(p.Family) > 0 && len(p.Given) > 0 {
		p.Sort = fmt.Sprintf("%s, %s", p.Family, p.Given)
	}
	return nil
}

// HasAffiliation checks a PersonOrOrg record for a specific affiliation
func (p *Person) HasAffiliation(target *Affiliation) bool {
	for _, affiliation := range p.Affiliations {
		if (target.ID == affiliation.ID) && (target.Name == affiliation.Name) {
			return true
		}
	}
	
	return false
}
