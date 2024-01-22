-- +goose Up
-- +goose StatementBegin
CREATE TABLE clusters(
  id TEXT PRIMARY KEY,
  name TEXT,
  version TEXT
);

CREATE TABLE machine_pools(
  id TEXT PRIMARY KEY,
  name TEXT,
  memory INTEGER,
  cpus INTEGER,
  group_name TEXT
);

CREATE TABLE machine_pool_assignments(
  id INTEGER PRIMARY KEY,
  role TEXT,
  cluster_id TEXT,
  machine_pool_id TEXT,
  FOREIGN KEY(cluster_id) REFERENCES clusters(id),
  FOREIGN KEY(machine_pool_id) REFERENCES machine_pools(id)
);

CREATE TABLE machines(
  id TEXT PRIMARY KEY,
  name TEXT,
  memory INTEGER,
  cpus INTEGER,
  machine_pool_id TEXT,
  group_name TEXT,
  machine TEXT,
  state Text,
  FOREIGN KEY(machine_pool_id) REFERENCES machine_pools(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clusters;
DROP TABLE machine_pools;
DROP TABLE machine_pool_assignments;
DROP TABLE machines;
-- +goose StatementEnd
