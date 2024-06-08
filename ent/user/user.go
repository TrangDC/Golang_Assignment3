// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWorkEmail holds the string denoting the work_email field in the database.
	FieldWorkEmail = "work_email"
	// FieldOid holds the string denoting the oid field in the database.
	FieldOid = "oid"
	// EdgeNewsEdges holds the string denoting the news_edges edge name in mutations.
	EdgeNewsEdges = "news_edges"
	// Table holds the table name of the user in the database.
	Table = "users"
	// NewsEdgesTable is the table that holds the news_edges relation/edge.
	NewsEdgesTable = "news"
	// NewsEdgesInverseTable is the table name for the News entity.
	// It exists in this package in order to avoid circular dependency with the "news" package.
	NewsEdgesInverseTable = "news"
	// NewsEdgesColumn is the table column denoting the news_edges relation/edge.
	NewsEdgesColumn = "author_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldWorkEmail,
	FieldOid,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// WorkEmailValidator is a validator for the "work_email" field. It is called by the builders before save.
	WorkEmailValidator func(string) error
	// OidValidator is a validator for the "oid" field. It is called by the builders before save.
	OidValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
