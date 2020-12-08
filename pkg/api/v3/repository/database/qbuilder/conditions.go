package qbuilder

import (
	"fmt"
	"strconv"

	"order-manager/pkg/api/v3/errs"
	"order-manager/pkg/cmn"
)

type ConditionCommon uint8

const (
	ConditionEquals ConditionCommon = iota
	ConditionNotEquals
	ConditionBiggerThan
	ConditionLessThan
	ConditionIs
	ConditionIsNot
)

var conditions = map[ConditionCommon]func(value interface{}) string{
	ConditionEquals:     condEquals,
	ConditionNotEquals:  condNotEquals,
	ConditionBiggerThan: condBiggerThan,
	ConditionLessThan:   condLessThan,
	ConditionIs:         condIs,
	ConditionIsNot:      condIsNot,
}

var conditionNames = []string{
	"Equals",
	"NotEquals",
	"BiggerThan",
	"LessThan",
	"Is",
	"IsNot",
}

func (c ConditionCommon) String() string {
	if int(c) < len(conditionNames) {
		return conditionNames[c]
	}
	return "unknown condition:" + strconv.Itoa(int(c))
}

func (c ConditionCommon) Fn(value interface{}) string {
	return conditions[c](value)
}

func condEquals(value interface{}) string {
	if value == nil {
		return ""
	}

	if target, ok := value.(Target); ok {
		return fmt.Sprintf("= %s.%s", target.Table.Name(), target.Table.ColumnName(target.FieldName))
	}

	vtype := cmn.DeriveType(value)
	if !cmn.IsPrimitiveKind(vtype) {
		return ""
	}

	if vtype.Kind().String() == cmn.KindString {
		return fmt.Sprintf("= '%s'", value)
	}
	return fmt.Sprintf("= %v", value)
}

func condNotEquals(value interface{}) string {
	if value == nil {
		return ""
	}

	if target, ok := value.(Target); ok {
		return fmt.Sprintf("<> %s.%s", target.Table.Name(), target.Table.ColumnName(target.FieldName))
	}

	vtype := cmn.DeriveType(value)
	if !cmn.IsPrimitiveKind(vtype) {
		return ""
	}

	if vtype.Kind().String() == cmn.KindString {
		return fmt.Sprintf("<> '%s'", value)
	}
	return fmt.Sprintf("<> %v", value)
}

func condBiggerThan(value interface{}) string {
	if value == nil {
		return ""
	}

	if target, ok := value.(Target); ok {
		return fmt.Sprintf("> %s.%s", target.Table.Name(), target.Table.ColumnName(target.FieldName))
	}

	vtype := cmn.DeriveType(value)
	if !cmn.IsPrimitiveKind(vtype) {
		return ""
	}

	if vtype.Kind().String() == cmn.KindString || vtype.Kind().String() == cmn.KindBool {
		return ""
	}
	return fmt.Sprintf("> %d", value)
}

func condLessThan(value interface{}) string {
	if value == nil {
		return ""
	}

	if target, ok := value.(Target); ok {
		return fmt.Sprintf("< %s.%s", target.Table.Name(), target.Table.ColumnName(target.FieldName))
	}

	vtype := cmn.DeriveType(value)
	if !cmn.IsPrimitiveKind(vtype) {
		return ""
	}

	if vtype.Kind().String() == cmn.KindString || vtype.Kind().String() == cmn.KindBool {
		return ""
	}
	return fmt.Sprintf("< %d", value)
}

func condIs(value interface{}) string {
	if value == nil {
		return fmt.Sprint("IS NULL")
	}

	vtype := cmn.DeriveType(value)
	if !cmn.IsPrimitiveKind(vtype) {
		return ""
	}

	if vtype.Kind().String() != cmn.KindBool {
		return fmt.Sprintf("IS %v", value)
	}

	return ""
}

func condIsNot(value interface{}) string {
	if value == nil {
		return fmt.Sprint("IS NOT NULL")
	}

	vtype := cmn.DeriveType(value)
	if !cmn.IsPrimitiveKind(vtype) {
		return ""
	}

	if vtype.Kind().String() != cmn.KindBool {
		return fmt.Sprintf("IS NOT %v", value)
	}

	return ""
}

type ConditionOrdersCoverage string

const (
	ConditionOrdersCoverageNot       ConditionOrdersCoverage = "not"
	ConditionOrdersCoverageFully     ConditionOrdersCoverage = "fully"
	ConditionOrdersCoveragePartially ConditionOrdersCoverage = "partially"
)

var conditionsOrders = map[ConditionOrdersCoverage]struct{}{
	ConditionOrdersCoverageNot:       {},
	ConditionOrdersCoverageFully:     {},
	ConditionOrdersCoveragePartially: {},
}

func (c ConditionOrdersCoverage) Validate() error {
	if _, ok := conditionsOrders[c]; !ok {
		return fmt.Errorf("%w: orders condition: coverage: %s", errs.ErrNotValid, c)
	}
	return nil
}
