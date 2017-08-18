
-- +migrate Up

CREATE TABLE users
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text NOT NULL,
  password_age text,
  last_activity text
);

-- +migrate Down

DROP TABLE users;
