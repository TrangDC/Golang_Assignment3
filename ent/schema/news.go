package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// News holds the schema definition for the News entity.
type News struct {
	ent.Schema
}

// Fields of the News. NOTE : Part of the public API ( ultimately exposed to end team
func (News) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("TITLE")),
		field.Enum("status").Values("draft", "published", "hidden").Default("draft"),
		field.String("description").MaxLen(512).Optional(),
		field.Text("content").NotEmpty(),
		field.UUID("author_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the News
func (News) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author_edge", User.Type).Ref("news_edges").Unique().Field("author_id"),
		// edge.From("author", Author.Type).Ref("news_edges").Field("author_id"),
	}
}

func (News) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		SlugMixin{},
	}
}
