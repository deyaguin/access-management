
-- +migrate Up

CREATE TABLE services
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text NOT NULL,
  created_at date,
  updated_at date,
  deleted_at date
);

-- +migrate Down

DROP TABLE services;
