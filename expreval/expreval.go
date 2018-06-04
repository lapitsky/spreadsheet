package expreval

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/lapitsky/spreadsheet/cells"
	"github.com/lapitsky/spreadsheet/stack"
)

var operations = map[byte]func(value1, value2 cells.CellValue) cells.CellValue{
	'+': add,
	'-': subtract,
	'*': multiply,
	'/': divide,
}

var validCellAddress = regexp.MustCompile(`^([a-zA-Z])([0-9]+)$`)

func Evaluate(cs *cells.Cells, c *cells.Cell) cells.CellValue {
	if c.HasValue() {
		return c.Value
	}

	if c.CycleState == cells.Pending {
		c.CycleState = cells.Cycle
		c.Value = cells.NewErrorValue("Cycle detected")
		return c.Value
	}

	c.CycleState = cells.Pending

	tokens := strings.Fields(c.Expr)

	defer func() {
		if c.CycleState != cells.Cycle {
			c.CycleState = cells.NoCycle
		}
	}()

	if len(tokens) == 0 {
		c.Value = cells.NewFloatValue(0)
		return c.Value
	}

	stack := stack.NewStack()

	for _, token := range tokens {
		if value, ok := parseValue(token); ok {
			stack.Push(cells.NewFloatValue(value))
		} else if op, ok := parseOperation(token); ok {
			value2, ok2 := stack.Pop()
			value1, ok1 := stack.Pop()
			if !ok1 || !ok2 {
				c.Value = cells.NewErrorValue("Not enough operands")
				return c.Value
			}
			stack.Push(operations[op](value1, value2))
		} else if scol, row, ok := parseCellRef(token); ok {
			stack.Push(Evaluate(cs, cs.At(scol, row)))
		} else {
			c.Value = cells.NewErrorValue("Invalid Expression")
			return c.Value
		}
	}

	if stack.Len() == 1 {
		c.Value, _ = stack.Pop()
	} else {
		c.Value = cells.NewErrorValue("Invalid Expression")
	}
	return c.Value
}

func add(value1, value2 cells.CellValue) cells.CellValue {
	return value1.Add(value2)
}

func subtract(value1, value2 cells.CellValue) cells.CellValue {
	return value1.Subtract(value2)
}

func multiply(value1, value2 cells.CellValue) cells.CellValue {
	return value1.Multiply(value2)
}

func divide(value1, value2 cells.CellValue) cells.CellValue {
	return value1.Divide(value2)
}

func parseValue(str string) (val float64, ok bool) {
	if f, err := strconv.ParseFloat(str, 32); err == nil {
		return f, true
	}
	return 0, false
}

func parseOperation(str string) (op byte, ok bool) {
	if len(str) != 1 {
		return '-', false
	}
	if _, ok := operations[str[0]]; ok {
		return str[0], true
	}
	return '-', false
}

func parseCellRef(str string) (scol byte, row int, ok bool) {
	if matches := validCellAddress.FindAllStringSubmatch(str, -1); matches != nil {
		row, _ = strconv.Atoi(matches[0][2])
		return matches[0][1][0], row - 1, true
	}
	return '-', -1, false
}
