-- +goose Up
-- +goose StatementBegin
ALTER TABLE machines
  ADD COLUMN created_at DATETIME;
ALTER TABLE machines
  ADD COLUMN updated_at DATETIME;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE machines
  DROP COLUMN created_at;
ALTER TABLE machines
  DROP COLUMN updated_at;
-- +goose StatementEnd
