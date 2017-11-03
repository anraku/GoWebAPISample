package meander

type Cost int8

// Cost型のenumを作成
const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

// Costと文字列のマッピングを保持する
var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

// Cost型の値から対応するmapのkeyを返す
func (l Cost) String() string {
	for s, v := range costStrings {
		if l == v {
			return s
		}
	}
	return "不正な値です"
}

// 文字列から対応するコスト（数値）を返す
func ParseCost(s string) Cost {
	return costStrings[s]
}
