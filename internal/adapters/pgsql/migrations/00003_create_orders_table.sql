-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
  id BIGSERIAL PRIMARY KEY,
  customer_id BIGINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE IF NOT EXISTS order_items (
  id BIGSERIAL PRIMARY KEY,
  product_id BIGINT NOT NULL,
  order_id BIGINT NOT NULL,
  quantity INTEGER NOT NULL,
  price INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id),
  CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
