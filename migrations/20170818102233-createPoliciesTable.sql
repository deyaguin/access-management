
-- +migrate Up

CREATE TABLE policies
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text,
  description text
);

-- +migrate Down

DROP TABLE policies;
