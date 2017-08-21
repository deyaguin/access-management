
-- +migrate Up

CREATE TABLE users
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text NOT NULL,
  password_age text,
  last_activity text,
  created_at date,
  updated_at date,
  deleted_at date
);

-- +migrate Down

DROP TABLE users;
