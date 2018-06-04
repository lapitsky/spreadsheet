package cells

import "unicode"

// Cells represents addressable spreadsheet cell
type Cells struct {
	Cells [][]Cell
}

// Cell struct represents an individual cell
type Cell struct {
	Expr       string
	Value      CellValue
	CycleState CycleDetectionState
}

type CycleDetectionState int

const (
	Initial CycleDetectionState = 0
	Pending CycleDetectionState = 1
	NoCycle CycleDetectionState = 2
	Cycle   CycleDetectionState = 3
)

var defaultCell = Cell{Expr: "", Value: NewFloatValue(0)}

// NewCells function initialized a new Cells object from a matrix of strings
func NewCells(records [][]string) *Cells {
	cells := Cells{}
	cells.init(records)
	return &cells
}

func (c *Cells) init(records [][]string) {
	c.Cells = make([][]Cell, len(records))
	for i := 0; i < len(c.Cells); i++ {
		c.Cells[i] = make([]Cell, len(records[i]))

		for j := 0; j < len(records[i]); j++ {
			c.Cells[i][j].Expr = records[i][j]
			c.Cells[i][j].CycleState = Initial
		}
	}
}

func (c *Cells) At(scol byte, row int) *Cell {
	col := int(unicode.ToUpper(rune(scol)) - 'A')
	if row < len(c.Cells) && col < len(c.Cells[row]) {
		return &c.Cells[row][col]
	}
	return &defaultCell
}

func (c *Cell) HasValue() bool {
	return c.Value != nil
}
