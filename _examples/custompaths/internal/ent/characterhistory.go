// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/flume/enthistory"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/characterhistory"
)

// CharacterHistory is the model entity for the CharacterHistory schema.
type CharacterHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref int `json:"ref,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CharacterHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case characterhistory.FieldID, characterhistory.FieldRef, characterhistory.FieldAge:
			values[i] = new(sql.NullInt64)
		case characterhistory.FieldOperation, characterhistory.FieldName:
			values[i] = new(sql.NullString)
		case characterhistory.FieldHistoryTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CharacterHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CharacterHistory fields.
func (ch *CharacterHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case characterhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ch.ID = int(value.Int64)
		case characterhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				ch.HistoryTime = value.Time
			}
		case characterhistory.FieldRef:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				ch.Ref = int(value.Int64)
			}
		case characterhistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				ch.Operation = enthistory.OpType(value.String)
			}
		case characterhistory.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				ch.Age = int(value.Int64)
			}
		case characterhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ch.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CharacterHistory.
// Note that you need to call CharacterHistory.Unwrap() before calling this method if this CharacterHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ch *CharacterHistory) Update() *CharacterHistoryUpdateOne {
	return NewCharacterHistoryClient(ch.config).UpdateOne(ch)
}

// Unwrap unwraps the CharacterHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ch *CharacterHistory) Unwrap() *CharacterHistory {
	_tx, ok := ch.config.driver.(*txDriver)
	if !ok {
		panic("ent: CharacterHistory is not a transactional entity")
	}
	ch.config.driver = _tx.drv
	return ch
}

// String implements the fmt.Stringer.
func (ch *CharacterHistory) String() string {
	var builder strings.Builder
	builder.WriteString("CharacterHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ch.ID))
	builder.WriteString("history_time=")
	builder.WriteString(ch.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", ch.Ref))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", ch.Operation))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", ch.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ch.Name)
	builder.WriteByte(')')
	return builder.String()
}

// CharacterHistories is a parsable slice of CharacterHistory.
type CharacterHistories []*CharacterHistory
