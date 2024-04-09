-- +goose Up
-- +goose StatementBegin
ALTER TABLE machines
  ADD COLUMN vmid INTEGER NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE machines
  DROP COLUMN vmid;
-- +goose StatementEnd
