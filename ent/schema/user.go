package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("user_id").
			Default(uuid.NewString()).
			Annotations(
				entgql.OrderField("ID"),
			),

		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),

		field.String("email").
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("EMAIL"),
			),

		field.String("password").
			NotEmpty(),

		field.String("status").
			NotEmpty().
			Annotations(
				entgql.OrderField("STATUS"),
			),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
