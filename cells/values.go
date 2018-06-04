package cells

import (
	"errors"
	"strconv"
)

type CellValue interface {
	Add(CellValue) CellValue
	AddTo(FloatValue) CellValue
	Subtract(CellValue) CellValue
	SubtractFrom(FloatValue) CellValue
	Multiply(CellValue) CellValue
	MultiplyBy(FloatValue) CellValue
	Divide(CellValue) CellValue
	DivideBy(FloatValue) CellValue
	String() string
	Error() error
}

type FloatValue struct {
	val float64
}

type ErrorValue struct {
	err error
}

func NewFloatValue(val float64) FloatValue {
	return FloatValue{val}
}

func NewErrorValue(strError string) ErrorValue {
	return ErrorValue{errors.New(strError)}
}

func (v FloatValue) String() string {
	return strconv.FormatFloat(v.val, 'g', -1, 32)
}

func (v FloatValue) Error() error {
	return nil
}
func (v FloatValue) Add(other CellValue) CellValue {
	return other.AddTo(v)
}

func (v FloatValue) AddTo(other FloatValue) CellValue {
	return NewFloatValue(v.val + other.val)
}

func (v FloatValue) Subtract(other CellValue) CellValue {
	return other.SubtractFrom(v)
}

func (v FloatValue) SubtractFrom(other FloatValue) CellValue {
	return NewFloatValue(other.val - v.val)
}

func (v FloatValue) Multiply(other CellValue) CellValue {
	return other.MultiplyBy(v)
}

func (v FloatValue) MultiplyBy(other FloatValue) CellValue {
	return NewFloatValue(v.val * other.val)
}

func (v FloatValue) Divide(other CellValue) CellValue {
	return other.DivideBy(v)
}

func (v FloatValue) DivideBy(other FloatValue) CellValue {
	if v.val == 0 {
		return NewErrorValue("Division by zero")
	}
	return NewFloatValue(other.val / v.val)
}

func (v ErrorValue) String() string {
	return "#ERR"
}

func (v ErrorValue) Error() error {
	return v.err
}

func (v ErrorValue) Add(other CellValue) CellValue {
	return v
}

func (v ErrorValue) AddTo(other FloatValue) CellValue {
	return v
}

func (v ErrorValue) Subtract(other CellValue) CellValue {
	return v
}

func (v ErrorValue) SubtractFrom(other FloatValue) CellValue {
	return v
}

func (v ErrorValue) Multiply(other CellValue) CellValue {
	return v
}

func (v ErrorValue) MultiplyBy(other FloatValue) CellValue {
	return v
}

func (v ErrorValue) Divide(other CellValue) CellValue {
	return v
}

func (v ErrorValue) DivideBy(other FloatValue) CellValue {
	return v
}
