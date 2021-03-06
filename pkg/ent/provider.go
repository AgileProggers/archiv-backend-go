// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/provider"
)

// Provider is the model entity for the Provider schema.
type Provider struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProviderQuery when eager-loading is set.
	Edges ProviderEdges `json:"edges"`
}

// ProviderEdges holds the relations/edges for other nodes in the graph.
type ProviderEdges struct {
	// Emotes holds the value of the emotes edge.
	Emotes []*Emote `json:"emotes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmotesOrErr returns the Emotes value or an error if the edge
// was not loaded in eager-loading.
func (e ProviderEdges) EmotesOrErr() ([]*Emote, error) {
	if e.loadedTypes[0] {
		return e.Emotes, nil
	}
	return nil, &NotLoadedError{edge: "emotes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Provider) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case provider.FieldID:
			values[i] = new(sql.NullInt64)
		case provider.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Provider", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Provider fields.
func (pr *Provider) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provider.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case provider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		}
	}
	return nil
}

// QueryEmotes queries the "emotes" edge of the Provider entity.
func (pr *Provider) QueryEmotes() *EmoteQuery {
	return (&ProviderClient{config: pr.config}).QueryEmotes(pr)
}

// Update returns a builder for updating this Provider.
// Note that you need to call Provider.Unwrap() before calling this method if this Provider
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Provider) Update() *ProviderUpdateOne {
	return (&ProviderClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Provider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Provider) Unwrap() *Provider {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Provider is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Provider) String() string {
	var builder strings.Builder
	builder.WriteString("Provider(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Providers is a parsable slice of Provider.
type Providers []*Provider

func (pr Providers) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
