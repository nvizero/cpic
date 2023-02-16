-- name: CreatePost :one
INSERT INTO posts (
  title,
  link,
  img,
  content
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: UpdatePost :one
UPDATE posts
SET
  title = COALESCE(sqlc.narg(title), title),
  link = COALESCE(sqlc.narg(link), link),
  img = COALESCE(sqlc.narg(img), img),
  content = COALESCE(sqlc.narg(content), content)
WHERE
  id = sqlc.arg(id)
RETURNING *;
