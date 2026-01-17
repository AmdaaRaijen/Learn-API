-- +goose Up
-- +goose StatementBegin
ALTER TABLE customers ADD COLUMN password TEXT;

UPDATE customers SET password = 'TEMP_INVALID_HASH';

ALTER TABLE customers ALTER COLUMN password SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE customers 
DROP COLUMN IF EXISTS password;
-- +goose StatementEnd
