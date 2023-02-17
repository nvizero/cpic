-- name: CreatePost :one
INSERT INTO posts (
  title,
  link,
  img,
  state,
  content,
  created_at
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts;

-- name: UpdatePost :one
UPDATE posts
SET
  title = COALESCE(sqlc.narg(title), title),
  link = COALESCE(sqlc.narg(link), link),
  img = COALESCE(sqlc.narg(img), img),
  state = COALESCE(sqlc.narg(state), state),
  content = COALESCE(sqlc.narg(content), content)
WHERE
  link = sqlc.arg(link)
RETURNING *;
