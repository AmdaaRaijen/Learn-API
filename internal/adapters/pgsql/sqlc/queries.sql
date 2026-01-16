-- name: ListProducts :many
SELECT * FROM products
WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
ORDER BY name;


-- name: FindProductByID :one
SELECT * FROM products WHERE id = $1;


-- name: CraeteOrder :one
INSERT INTO orders (customer_id) VALUES ($1) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (product_id, order_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING *;