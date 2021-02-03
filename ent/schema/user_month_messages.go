package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserMonthMessage holds the schema definition for the UserMonthMessage entity.
type UserMonthMessage struct {
	ent.Schema
}

// Fields of the UserMonthMessage.
func (UserMonthMessage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Time("month").Default(time.Now),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the UserMonthMessage.
func (UserMonthMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider_id", Provide.Type).
			Ref("provider_id").Unique(),
		edge.From("user_id", User.Type).
			Ref("user_messages").
			Unique(),
	}
}
