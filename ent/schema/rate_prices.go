package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RatePrice holds the schema definition for the RatePrice entity.
type RatePrice struct {
	ent.Schema
}

// Fields of the RatePrice.
func (RatePrice) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the RatePrice.
func (RatePrice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("id_rate", Rate.Type).
			Ref("rate_id").
			Unique(),
		edge.From("id_price", Price.Type).
			Ref("price_id").
			Unique(),
		edge.To("user", User.Type).
			StorageKey(edge.Column("rate_id")),
	}
}
