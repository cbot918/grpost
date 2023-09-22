-- name: GetComments :many
SELECT comments.texts, comments.id, comments.posted_by, comments.user_name, comments.target_post
FROM comments
INNER JOIN posts
ON posts.id = comments.target_post;