package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
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
		field.Int32("sequence_number").Positive(),
		field.String("external_id").Optional(),
		field.String("dst").Optional(),
		field.String("message").Unique(),
		field.String("src").Unique(),
		field.Int32("state").Optional(),
		field.Int32("smsc_message_id").Optional(),
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
