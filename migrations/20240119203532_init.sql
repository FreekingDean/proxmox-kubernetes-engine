-- +goose Up
-- +goose StatementBegin
CREATE TABLE clusters(
  id TEXT PRIMARY KEY,
  name TEXT,
  version TEXT
);

CREATE TABLE node_pools(
  id TEXT PRIMARY KEY,
  name TEXT,
  memory INTEGER,
  cpus INTEGER
);

CREATE TABLE node_pool_assignments(
  id INTEGER PRIMARY KEY,
  role TEXT,
  cluster_id TEXT,
  node_pool_id TEXT,
  FOREIGN KEY(cluster_id) REFERENCES clusters(id),
  FOREIGN KEY(node_pool_id) REFERENCES node_pools(id)
);

CREATE TABLE machines(
  id TEXT PRIMARY KEY,
  name TEXT,
  memory INTEGER,
  cpus INTEGER,
  node_pool_id TEXT,
  FOREIGN KEY(node_pool_id) REFERENCES node_pools(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clusters;
DROP TABLE node_pools;
DROP TABLE node_pool_assignments;
DROP TABLE machines;
-- +goose StatementEnd
