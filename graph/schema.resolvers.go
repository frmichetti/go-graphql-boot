package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/anjaneyulubatta505/gin-graphql-postgres/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// CreateQuestion is the resolver for the createQuestion field.
func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	panic(fmt.Errorf("not implemented: CreateQuestion - createQuestion"))
}

// CreateChoice is the resolver for the createChoice field.
func (r *mutationResolver) CreateChoice(ctx context.Context, input *model.ChoiceInput) (*model.Choice, error) {
	panic(fmt.Errorf("not implemented: CreateChoice - createChoice"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Questions is the resolver for the questions field.
func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	panic(fmt.Errorf("not implemented: Questions - questions"))
}

// Choices is the resolver for the choices field.
func (r *queryResolver) Choices(ctx context.Context) ([]*model.Choice, error) {
	panic(fmt.Errorf("not implemented: Choices - choices"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
