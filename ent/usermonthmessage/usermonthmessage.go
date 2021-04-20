// Code generated by entc, DO NOT EDIT.

package usermonthmessage

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the usermonthmessage type in the database.
	Label = "user_month_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMonth holds the string denoting the month field in the database.
	FieldMonth = "month"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// EdgeProviderID holds the string denoting the provider_id edge name in mutations.
	EdgeProviderID = "provider_id"
	// EdgeUserID holds the string denoting the user_id edge name in mutations.
	EdgeUserID = "user_id"
	// Table holds the table name of the usermonthmessage in the database.
	Table = "user_month_messages"
	// ProviderIDTable is the table the holds the provider_id relation/edge.
	ProviderIDTable = "user_month_messages"
	// ProviderIDInverseTable is the table name for the Provide entity.
	// It exists in this package in order to avoid circular dependency with the "provide" package.
	ProviderIDInverseTable = "provides"
	// ProviderIDColumn is the table column denoting the provider_id relation/edge.
	ProviderIDColumn = "provider_id"
	// UserIDTable is the table the holds the user_id relation/edge.
	UserIDTable = "user_month_messages"
	// UserIDInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserIDInverseTable = "users"
	// UserIDColumn is the table column denoting the user_id relation/edge.
	UserIDColumn = "user_id"
)

// Columns holds all SQL columns for usermonthmessage fields.
var Columns = []string{
	FieldID,
	FieldMonth,
	FieldCreateAt,
	FieldUpdateAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_month_messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"provider_id",
	"user_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultMonth holds the default value on creation for the "month" field.
	DefaultMonth func() time.Time
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
