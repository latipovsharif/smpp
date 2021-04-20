// Code generated by entc, DO NOT EDIT.

package price

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the price type in the database.
	Label = "price"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMin holds the string denoting the min field in the database.
	FieldMin = "min"
	// FieldMax holds the string denoting the max field in the database.
	FieldMax = "max"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// EdgePriceID holds the string denoting the price_id edge name in mutations.
	EdgePriceID = "price_id"
	// Table holds the table name of the price in the database.
	Table = "prices"
	// PriceIDTable is the table the holds the price_id relation/edge.
	PriceIDTable = "rate_prices"
	// PriceIDInverseTable is the table name for the RatePrice entity.
	// It exists in this package in order to avoid circular dependency with the "rateprice" package.
	PriceIDInverseTable = "rate_prices"
	// PriceIDColumn is the table column denoting the price_id relation/edge.
	PriceIDColumn = "price_id"
)

// Columns holds all SQL columns for price fields.
var Columns = []string{
	FieldID,
	FieldMin,
	FieldMax,
	FieldPrice,
	FieldCreateAt,
	FieldUpdateAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
