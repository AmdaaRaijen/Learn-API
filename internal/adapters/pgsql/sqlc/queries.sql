-- name: ListProducts :many
SELECT * FROM products
WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
ORDER BY name;

-- name: FindProductByID :one
SELECT * FROM products WHERE id = $1;

-- name: UpdateProduct :one
UPDATE products SET name = $2, price = $3, quantity = $4 WHERE id = $1 RETURNING *;

-- name: CreateOrder :one
INSERT INTO orders (customer_id) VALUES ($1) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (product_id, order_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetCustomerById :one
SELECT * FROM customers WHERE id = $1;

-- name: GetCustomerByEmail :one
SELECT * FROM customers WHERE email = $1;

-- name: CreateUser :one
INSERT INTO customers (name, email, phone_number, password) VALUES ($1, $2, $3, $4) RETURNING *;