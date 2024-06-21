package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(entgql.OrderField("ID")),

		field.String("name").
			NotEmpty().
			MaxLen(255).
			Annotations(entgql.OrderField("NAME")),

		field.String("work_email").
			MaxLen(255).
			Annotations(entgql.OrderField("WORK_EMAIL")),

		field.String("oid").
			Unique().
			MaxLen(255),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bookings", Booking.Type),
	}
}
