
-- +migrate Up

CREATE TABLE groups
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text NOT NULL,
  creation_time date,
  created_at date,
  updated_at date,
  deleted_at date
);
-- +migrate Down

DROP TABLE groups;
