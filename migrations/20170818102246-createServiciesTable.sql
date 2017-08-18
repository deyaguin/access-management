
-- +migrate Up

CREATE TABLE services
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text NOT NULL
);

-- +migrate Down

DROP TABLE services;
