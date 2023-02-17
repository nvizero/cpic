package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"cpic/util"

	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {

	arg := CreatePostParams{
		Title:     sql.NullString{String: util.RandomString(21), Valid: true},
		Link:      sql.NullString{String: util.RandomString(21), Valid: true},
		State:     sql.NullBool{Bool: false, Valid: true},
		Img:       util.RandomString(21),
		Content:   util.RandomString(25),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	post, err := testQueries.CreatePost(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, post.Title, arg.Title)
	require.Equal(t, post.Link, arg.Link)
	require.Equal(t, post.Img, arg.Img)

	require.Equal(t, post.Content, arg.Content)
}

func TestGetPosts(t *testing.T) {
	posts, err := testQueries.GetPosts(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, posts)
	require.NotEmpty(t, len(posts))
	fmt.Println(len(posts))
}

func TestGetPost(t *testing.T) {
	post, err := testQueries.GetPost(context.Background(), 116)
	require.NoError(t, err)
	require.NotEmpty(t, post)
}

func TestUpdatePost(t *testing.T) {
	arg := UpdatePostParams{
		Title:   sql.NullString{String: util.RandomString(21), Valid: true},
		Link:    sql.NullString{String: "/doc_Y2JocHhaWFBTRWxYZlJtTkRrWmFDZz09", Valid: true},
		Img:     sql.NullString{String: util.RandomString(21), Valid: true},
		Content: sql.NullString{String: util.RandomString(21), Valid: true},
		ID:      117,
	}
	post, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)
}
