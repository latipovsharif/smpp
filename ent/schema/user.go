package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Int16("balance").Default(0),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//User Month Message
		edge.To("user_messages", UserMonthMessage.Type).
			StorageKey(edge.Column("user_id")),
		// messages
		edge.To("messages", Messages.Type).
			StorageKey(edge.Column("user_id")),
		//rate in user
		edge.From("rate_id", Rate.Type).
			Ref("user").
			Unique(),
	}
}
