package qbuilder

import (
	"reflect"

	"order-manager/pkg/api/v3/models"
)

// Table consists:
// - name is actual name of the represented table in the postgres database
// - columns:
//   - keys are go-struct field names, for example `BaseCurrencyAmount` in models.Order
//   - values are sql column names, for example `base_currency_amount` in `orders` table
type Table struct {
	name    string
	columns map[string]string
}

func (t Table) Name() string {
	return t.name
}

func (t Table) ColumnName(fieldName string) string {
	return t.columns[fieldName]
}

func (t Table) As(alias string) Table {
	return Table{
		name:    alias,
		columns: t.columns,
	}
}

// These structs get filled its `columns` field at once in runtime
// by parsing the go structs that represent the respecting db tables.
var (
	TableOrders = Table{
		name:    "orders",
		columns: map[string]string{},
	}

	TableEmission = Table{
		name:    "emission",
		columns: map[string]string{},
	}

	TableCoverage = Table{
		name:    "coverage",
		columns: map[string]string{},
	}
)

func init() {
	var t reflect.Type

	t = reflect.TypeOf(models.Order{})
	for i := 0; i < t.NumField(); i++ {
		TableOrders.columns[t.Field(i).Name] = t.Field(i).Tag.Get("db")
	}

	t = reflect.TypeOf(models.Emission{})
	for i := 0; i < t.NumField(); i++ {
		TableEmission.columns[t.Field(i).Name] = t.Field(i).Tag.Get("db")
	}

	t = reflect.TypeOf(models.Coverage{})
	for i := 0; i < t.NumField(); i++ {
		TableCoverage.columns[t.Field(i).Name] = t.Field(i).Tag.Get("db")
	}
}
