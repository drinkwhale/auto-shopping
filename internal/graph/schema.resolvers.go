package graph

import (
	"context"
	"fmt"
	"time" // 'time' 패키지를 import 합니다.

	"github.com/drinkwhale/auto-shopping/internal/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		// Go 1.25.1과 호환되는 time 패키지를 사용합니다.
		ID:   fmt.Sprintf("T%d", time.Now().UnixNano()),
		Text: input.Text,
		Done: false,
	}
	r.Resolver.todos = append(r.Resolver.todos, todo)
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.Resolver.todos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
