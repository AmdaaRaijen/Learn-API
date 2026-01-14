-- name: ListProducts :many
SELECT * FROM products
WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
ORDER BY name;


-- name: FindProductByID :one
SELECT * FROM products WHERE id = $1;