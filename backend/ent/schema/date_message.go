package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Date_Message holds the schema definition for the Date_Message entity.
type Date_Message struct {
	ent.Schema
}

// Fields of the Date_Message.
func (Date_Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at"),
	}
}

// Edges of the Date_Message.
func (Date_Message) Edges() []ent.Edge {
	// Date_Message
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
