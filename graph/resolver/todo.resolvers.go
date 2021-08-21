package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/xdorro/golang-fiber-project/ent"
	"github.com/xdorro/golang-fiber-project/graph/generated"
	"github.com/xdorro/golang-fiber-project/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*ent.Todo, error) {
	return ent.FromContext(ctx).Todo.Create().
		SetText(*input.Text).
		SetStatus(string(input.Status)).
		SetPriority(*input.Priority).
		Save(ctx)
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input model.UpdateTodoInput) (*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTodos(ctx context.Context, ids []string, input model.UpdateTodoInput) ([]*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	log.Println("Case todos")
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
}

func (r *todoResolver) Status(ctx context.Context, obj *ent.Todo) (model.TodoStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) Parent(ctx context.Context, obj *ent.Todo) (*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) Children(ctx context.Context, obj *ent.Todo) ([]*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
