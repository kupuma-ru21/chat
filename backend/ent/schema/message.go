package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}).Default(uuid.New),
		field.String("content").NotEmpty(),
		field.Time("created_at"),
		field.UUID("date_id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("messages").
			Field("user_id").
			Unique().
			Required(),
		edge.From("date_message", Date_Message.Type).
			Ref("messages").
			Field("date_id").
			Unique().
			Required(),
	}
}
