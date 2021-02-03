package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Provide holds the schema definition for the Provide entity.
type Provide struct {
	ent.Schema
}

// Fields of the Provide.
func (Provide) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").Unique(),
		field.String("ip_adres").Unique(),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the Provide.
func (Provide) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider_id", UserMonthMessage.Type).
			StorageKey(edge.Column("provider_id")),
		edge.To("messages", Messages.Type).
			StorageKey(edge.Column("provider_id")),
	}
}
