package meander

import "fmt"

type j struct {
	Name       string
	PlaceTypes []string
}

var Journeys = []interface{}{
	&j{Name: "ロマンティック", PlaceTypes: []string{"park", "bar"}},
	&j{Name: "ショッピング", PlaceTypes: []string{"department_store", "cafe"}},
}
