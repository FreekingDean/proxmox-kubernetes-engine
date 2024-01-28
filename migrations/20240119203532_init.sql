-- +goose Up
-- +goose StatementBegin
CREATE TABLE clusters(
  id TEXT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  kubernetes_version TEXT NOT NULL
);

CREATE TABLE machine_pools(
  id TEXT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  image TEXT NOT NULL,
  memory INTEGER NOT NULL,
  cpus INTEGER NOT NULL,
  group_name TEXT NOT NULL
);

CREATE TABLE machine_pool_assignments(
  id TEXT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  machine_count INTEGER NOT NULL,
  role TEXT NOT NULL,

  cluster_id TEXT NOT NULL,
  machine_pool_id TEXT NOT NULL,
  FOREIGN KEY(cluster_id) REFERENCES clusters(id),
  FOREIGN KEY(machine_pool_id) REFERENCES machine_pools(id)
);

CREATE TABLE machines(
  id TEXT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  memory INTEGER NOT NULL,
  cpus INTEGER NOT NULL,
  state TEXT NOT NULL,
  image TEXT NOT NULL,
  group_name TEXT NOT NULL,
  role TEXT NOT NULL,
  node TEXT NOT NULL,

  machine_pool_assignment_id TEXT NOT NULL,
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
