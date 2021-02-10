package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Price holds the schema definition for the Price entity.
type Price struct {
	ent.Schema
}

// Fields of the Price.
func (Price) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Int32("min"),
		field.Int32("max"),
		field.String("price").Unique(),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the Price.
func (Price) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("price_id", RatePrice.Type).
			StorageKey(edge.Column("price_id")),
	}
}
