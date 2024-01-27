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
  version INTEGER,
  group_name TEXT
);

CREATE TABLE machine_pool_assignments(
  id TEXT PRIMARY KEY,
  name TEXT,
  machine_count INTEGER,
  role TEXT,
  cluster_id TEXT,
  machine_pool_id TEXT,
  version TEXT,
  FOREIGN KEY(cluster_id) REFERENCES clusters(id),
  FOREIGN KEY(machine_pool_id) REFERENCES machine_pools(id)
);

CREATE TABLE machines(
  id TEXT PRIMARY KEY,
  name TEXT,
  memory INTEGER,
  cpus INTEGER,
  machine_pool_assignment_id TEXT,
  machine_pool_assignment_version INTEGER,
  group_name TEXT,
  node TEXT,
  state TEXT,
  version TEXT,
  machine_pool_version TEXT,
  FOREIGN KEY(machine_pool_assignment_id) REFERENCES machine_pool_assignments(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clusters;
DROP TABLE machine_pools;
DROP TABLE machines;
DROP TABLE machine_pool_assignments;
-- +goose StatementEnd
