package qbuilder

import (
	"fmt"
	"strings"

	"order-manager/pkg/api/v3/errs"
)

type Filter struct {
	Target    Target
	Condition ConditionCommon
	Value     interface{}
}

type Target struct {
	// This field is supposed to be set with the field name of the go-struct
	// that represents sql-column in the respecting table column.
	FieldName string
	Table     Table
}

// NewFilter provide new filtering criteria for sql-query.
// Signature is:
// - fieldName is described in comments for Target and Table
// - table is a postgres table representation
// - condition represents a part of the WHERE block kinda <>=
// - value is anything value that will be reflectively interpreted in an appropriate condition function
func NewFilter(fieldName string, table Table, condition ConditionCommon, value interface{}) Filter {
	return Filter{
		Target: Target{
			FieldName: fieldName,
			Table:     table,
		},
		Condition: condition,
		Value:     value,
	}
}

func (f Filter) Apply(b *strings.Builder) error {
	fname := f.Target.Table.ColumnName(f.Target.FieldName)
	if fname == "" {
		return fmt.Errorf("%w: qbuilder filter: target field name: %s: there is no associated column in the table: %s",
			errs.ErrNotValid,
			f.Target.FieldName,
			f.Target.Table.Name(),
		)
	}

	result := f.Condition.Fn(f.Value)
	if result == "" {
		return fmt.Errorf("something went wrong with applying the qbuilder filter: target field name: %s; condition: %s; value: %v",
			f.Target.FieldName,
			f.Condition.String(),
			f.Value,
		)
	}

	b.WriteByte(' ')
	b.WriteString(f.Target.Table.Name())
	b.WriteByte('.')
	b.WriteString(fname)
	b.WriteByte(' ')
	b.WriteString(result)

	return nil
}
