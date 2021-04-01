package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Rate holds the schema definition for the Rate entity.
type Rate struct {
	ent.Schema
}

// Fields of the Rate.
func (Rate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").Unique(),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the Rate.
func (Rate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rate_id", RatePrice.Type).
			StorageKey(edge.Column("rate_id")),
	}
}
