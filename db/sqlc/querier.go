// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	GetPost(ctx context.Context, id int64) (Post, error)
	GetPosts(ctx context.Context) ([]Post, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
}

var _ Querier = (*Queries)(nil)
