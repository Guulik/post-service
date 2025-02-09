package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"
	"posts/graph"
	"posts/internal/model"
)

// CreateComment is the resolver for the CreateComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.InputComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - CreateComment"))
}

// GetReplies is the resolver for the GetReplies field.
func (r *queryResolver) GetReplies(ctx context.Context, commentID string, depth int32) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: GetReplies - GetReplies"))
}

// CommentsSubscription is the resolver for the CommentsSubscription field.
func (r *subscriptionResolver) CommentsSubscription(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	panic(fmt.Errorf("not implemented: CommentsSubscription - CommentsSubscription"))
}

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment, depth int32) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - CreateComment"))
}
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }
type commentResolver struct{ *Resolver }
*/
