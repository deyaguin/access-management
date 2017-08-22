
-- +migrate Up

CREATE TABLE policies
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text,
  description text,
  created_at date,
  updated_at date,
  deleted_at date
);

-- +migrate Down

DROP TABLE policies;
