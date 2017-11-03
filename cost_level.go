package meander

import "strings"

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

// コストの下限と上限を定義する
type CostRange struct {
	From Cost
	To   Cost
}

// 文字列からCostRange型に変換
// 入力値の例）"$...$$$"
func ParseCostRange(s string) *CostRange {
	segs := strings.Split(s, "...")
	return &CostRange{
		From: costStrings[segs[0]],
		To:   costStrings[segs[1]],
	}
}

// CostRange型から文字列に変換する
func (c CostRange) String() string {
	return c.From.String() + "..." + c.To.String()
}
