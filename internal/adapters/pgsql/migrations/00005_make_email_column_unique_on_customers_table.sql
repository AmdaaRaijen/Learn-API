-- +goose Up
-- +goose StatementBegin
ALTER TABLE customers ALTER COLUMN email SET NOT NULL;

ALTER TABLE customers ADD CONSTRAINT customers_email_unique UNIQUE (email);

CREATE INDEX IF NOT EXISTS idx_customers_email ON customers(email); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_customers_email;

ALTER TABLE customers ALTER COLUMN email DROP NOT NULL;
ALTER TABLE customers DROP CONSTRAINT IF EXISTS customers_email_unique;
-- +goose StatementEnd
