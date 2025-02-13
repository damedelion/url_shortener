package migrations

const CreateTableQuery = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "urls" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  short_url VARCHAR(10),
  long_url VARCHAR(100),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT current_timestamp
);`
