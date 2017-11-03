package meander

var APIKey string

type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photo           []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

type googleResponse struct {
	Result []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

// 構造体を外部に公開するためのメソッド
func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photo,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}

// Google Place APIにデータを送信する際のフォーマットを定義
type Query struct {
	Lat          float64
	Lng          float64
	Journey      []string
	Radius       int
	CostRangeStr string
}
