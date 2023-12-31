-- name: GetUser :one
SELECT name, is_administrator, created_on FROM users
WHERE name = $1;

-- name: UpdateUserPassword :exec
UPDATE users SET password_hash = $2 WHERE name = $1;

-- name: CreateUser :exec
INSERT INTO users (name, password_hash, is_administrator)
VALUES ($1, $2, $3);

-- name: GetPost :one
SELECT posts.title, users.name FROM posts JOIN users ON users.user_id = comments.user_id
WHERE post_id = $1;

-- name: ListComments :many
SELECT comment_id, users.name, body, created_at FROM comments JOIN users ON users.user_id = comments.user_id
WHERE post_id = $1 AND created_at < $2
ORDER BY created_at DESC
LIMIT $3;

-- name: ListPosts :many
SELECT post_id, title, created_at FROM posts
WHERE created_at < $1
ORDER BY created_at DESC
LIMIT $2;

-- name: createPost :exec
WITH current_id AS (
  INSERT INTO posts (title)
  VALUES ($2) RETURNING post_id
) INSERT INTO comments (post_id, user_id, body)
  VALUES (current_id, $1, $3);

-- name: deletePost :exec
DELETE FROM posts WHERE post_id = $1;

-- name: createComment :exec
INSERT INTO comments (post_id, user_id, body)
VALUES ($1, $2, $3);

-- name: editComment :exec
UPDATE comments SET body = $2, edited_at = CURRENT_TIME WHERE comment_id = $1;

-- name: deleteComment :exec
UPDATE comments SET body = "Comment deleted by author.", edited_at = NULL WHERE comment_id = $1;
