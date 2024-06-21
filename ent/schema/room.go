package schema

import (
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Room struct {
	ent.Schema
}

func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(entgql.OrderField("ID")),

		field.String("name").
			NotEmpty().
			MaxLen(255).
			Annotations(entgql.OrderField("NAME")),

		field.String("color").
			NotEmpty().
			MinLen(7).
			MaxLen(7).
			Match(regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)).
			Annotations(entgql.OrderField("COLOR")),

		field.UUID("office_id", uuid.UUID{}),

		field.String("description").
			Optional(),

		field.String("image_url").
			Optional(),
	}
}

func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bookings", Booking.Type),

		edge.From("office", Office.Type).
			Ref("rooms").
			Unique().
			Field("office_id").
			Required(),
	}
}
