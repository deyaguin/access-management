
-- +migrate Up

CREATE TABLE groups
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text NOT NULL,
  creation_time date
);
-- +migrate Down

DROP TABLE groups;
