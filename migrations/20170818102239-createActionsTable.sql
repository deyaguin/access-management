
-- +migrate Up

CREATE TABLE actions
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text NOT NULL,
  service_id integer,
  created_at date,
  updated_at date,
  deleted_at date,
  CONSTRAINT service_id FOREIGN KEY (service_id)
      REFERENCES services (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- +migrate Down

DROP TABLE actions;
