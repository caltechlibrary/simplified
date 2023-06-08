package simplified

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	// 3rd Party Packages
	"gopkg.in/yaml.v3"
)

func TestPerson(t *testing.T) {
/*
- name: Adhikari, Rana X.
  family_name: Adhikari
  given_name: Rana X.
  identifiers:
    - identifier: 0000-0002-5731-5076
      scheme: orcid
    - identifier: Adhikari-R-X
      scheme: clpid
  affiliations:
    - name: Caltech
      id: 05dxps055
*/
	p := new(Person)
	p.Name = `Adhikari, Rana X.`
	p.Family = `Adhikari`
	p.Given = `Rana X.`
	p.Identifiers = []*Identifier{
		&Identifier{
			Identifier: `0000-0002-5731-5076`,
			Scheme: `orcid`,
		},
		&Identifier{
			Identifier: `Adhikari-R-X`,
			Scheme: `clpid`,
		},
	}
	p.Affiliations = []*Affiliation{
		&Affiliation{
			ID: `05dxps055`,
			Name: `Caltech`,
		},
	}

	src, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		t.Error(err)
	}
	fmt.Fprintf(os.Stderr, "DEBUG src (json) ->\n%s\n\n", src)
	src, err = yaml.Marshal(p)
	if err != nil {
		t.Error(err)
	}
	fmt.Fprintf(os.Stderr, "DEBUG src (yaml) ->\n%s\n\n", src)

/* Identifier should format in YAML like

- identifier: Adhikari-R-X
      scheme: clpid
*/
}
