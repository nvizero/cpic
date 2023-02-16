package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"cpic/util"

	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {

	arg := CreatePostParams{
		Title:   sql.NullString{String: util.RandomString(21), Valid: true},
		Link:    sql.NullString{String: util.RandomString(21), Valid: true},
		Img:     util.RandomString(21),
		Content: util.RandomString(19),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, post.Title, arg.Title)
	require.Equal(t, post.Link, arg.Link)
	require.Equal(t, post.Img, arg.Img)

	require.Equal(t, post.Content, arg.Content)
}

func TestGetPost(t *testing.T) {
	post, err := testQueries.GetPost(context.Background(), 4)
	require.NoError(t, err)
	require.NotEmpty(t, post)
}

func TestUpdatePost(t *testing.T) {
	arg := UpdatePostParams{
		Title:   sql.NullString{String: util.RandomString(21), Valid: true},
		Link:    sql.NullString{String: util.RandomString(21), Valid: true},
		Img:     sql.NullString{String: util.RandomString(21), Valid: true},
		Content: sql.NullString{String: util.RandomString(21), Valid: true},
		ID:      5,
	}
	post, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	fmt.Println(post)
}
