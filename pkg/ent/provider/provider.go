// Code generated by entc, DO NOT EDIT.

package provider

const (
	// Label holds the string label denoting the provider type in the database.
	Label = "provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeEmotes holds the string denoting the emotes edge name in mutations.
	EdgeEmotes = "emotes"
	// Table holds the table name of the provider in the database.
	Table = "providers"
	// EmotesTable is the table that holds the emotes relation/edge.
	EmotesTable = "emotes"
	// EmotesInverseTable is the table name for the Emote entity.
	// It exists in this package in order to avoid circular dependency with the "emote" package.
	EmotesInverseTable = "emotes"
	// EmotesColumn is the table column denoting the emotes relation/edge.
	EmotesColumn = "provider_emotes"
)

// Columns holds all SQL columns for provider fields.
var Columns = []string{
	FieldID,
	FieldName,
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
