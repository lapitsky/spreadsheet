package expreval

import (
	"reflect"
	"testing"

	"github.com/lapitsky/spreadsheet/cells"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name  string
		input [][]string
		want  [][]string
	}{
		{
			"Original test data",
			[][]string{
				[]string{" b1 b2 +  ", "2 b2 3 * -", " 3", "+"},
				[]string{"a1", "5 ", "", "7     2  /"},
				[]string{"c2 3 * ", "1 2       ", " ", "5 1 2 + 4 * + 3 -"},
			},
			[][]string{
				[]string{"-8", "-13", "3", "#ERR"},
				[]string{"-8", "5", "0", "3.5"},
				[]string{"0", "#ERR", "0", "14"},
			},
		},
		{
			"Spreadsheet with cycles",
			[][]string{
				[]string{"b1 b2 +", "4     ", "a2 c2 *"},
				[]string{"5      ", "a3 1 +", "3"},
				[]string{"c3 a1 *", "", "a2 c1 +"},
			},
			[][]string{
				[]string{"#ERR", "4", "15"},
				[]string{"5", "#ERR", "3"},
				[]string{"#ERR", "0", "20"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := cells.NewCells(tt.input)
			if got := evaluateCells(cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func evaluateCells(cs *cells.Cells) [][]string {
	result := make([][]string, len(cs.Cells))
	for i, row := range cs.Cells {
		result[i] = make([]string, len(row))
		for j, c := range row {
			result[i][j] = Evaluate(cs, &c).String()
		}
	}

	return result
}
