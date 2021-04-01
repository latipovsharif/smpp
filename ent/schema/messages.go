package schema

import (
	"time"

	"entgo.io/ent/schema/edge"
	"entgo.io/ent"
    "entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Messages holds the schema definition for the Messages entity.
type Messages struct {
	ent.Schema
}

// Fields of the Messages.
func (Messages) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Int32("sequence_number"),
		field.String("external_id"),
		field.String("dst"),
		field.String("message"),
		field.String("src"),
		field.Int("state"),
		field.String("smsc_message_id"),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the Messages.
func (Messages) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_id", User.Type).
			Ref("messages").
			Unique(),
		edge.From("provider_id", Provide.Type).
			Ref("messages").
			Unique(),
	}
}
