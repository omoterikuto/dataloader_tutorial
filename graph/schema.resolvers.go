package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dataloader/dataloader"
	"dataloader/entity"
	"dataloader/graph/generated"
	"dataloader/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &entity.Todo{
		Text:   input.Text,
		UserID: uint(input.UserID),
	}
	if err := r.DB.Create(&todo).Error; err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:     int(todo.ID),
		UserID: int(todo.UserID),
		Done:   todo.Done,
		Text:   todo.Text,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &entity.User{
		Name: input.Name,
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &model.User{
		ID:   int(user.ID),
		Name: user.Name,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*entity.Todo
	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	var retTodos []*model.Todo
	for _, todo := range todos {
		t := model.Todo{
			ID:     int(todo.ID),
			UserID: int(todo.UserID),
			Done:   todo.Done,
			Text:   todo.Text,
		}
		retTodos = append(retTodos, &t)
	}
	return retTodos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	//var user entity.User
	//user.ID = uint(obj.UserID)
	//if err := r.DB.First(&user).Error; err != nil {
	//	return nil, err
	//}
	//return &model.User{
	//	ID:   int(user.ID),
	//	Name: user.Name,
	//}, nil
	return dataloader.LoadUser(ctx, obj.UserID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
