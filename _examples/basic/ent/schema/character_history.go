// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/flume/enthistory"

	"time"
)

// CharacterHistory holds the schema definition for the CharacterHistory entity.
type CharacterHistory struct {
	ent.Schema
}

// Annotations of the CharacterHistory.
func (CharacterHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "character_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
	}
}

// Fields of the CharacterHistory.
func (CharacterHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.Int("ref").
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(enthistory.OpType("")).
			Immutable(),
		field.Int("updated_by").
			Optional().
			Immutable().
			Nillable(),
	}

	return append(historyFields, Character{}.Fields()...)
}

// Mixin of the CharacterHistory.
func (CharacterHistory) Mixin() []ent.Mixin {
	return Character{}.Mixin()
}
