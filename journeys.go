package meander

import "strings"

type j struct {
	Name       string
	PlaceTypes []string
}

var Journeys = []interface{}{
	&j{Name: "ロマンティック", PlaceTypes: []string{"park", "bar"}},
	&j{Name: "ショッピング", PlaceTypes: []string{"department_store", "cafe"}},
}

func (j *j) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"journey": strings.Join(j.PlaceTypes, "|"),
	}
}
