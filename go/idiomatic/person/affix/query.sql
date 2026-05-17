-- name: GetAffix :one
SELECT * FROM affix
WHERE id = ? LIMIT 1;

-- name: CreateAffix :exec
INSERT INTO affix (id, 'name', 'description', category_id)
VALUES (?, ?, ?, ?);

-- name: UpdateAffix :exec
UPDATE affix SET
'name' = ?, 'description' = ?, category_id = ? 
WHERE id = ?;

-- name: GetAAbbreviations :many
SELECT abbreviation FROM affix_abbreviation
WHERE id = ? ORDER BY order;

-- name: GetCategory :one
SELECT * FROM affix_category
WHERE id = ?;
