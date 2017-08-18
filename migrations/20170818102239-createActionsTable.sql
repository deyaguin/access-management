
-- +migrate Up

CREATE TABLE actions
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  name text NOT NULL,
  service_id integer,
  CONSTRAINT service_id FOREIGN KEY (service_id)
      REFERENCES services (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- +migrate Down

DROP TABLE actions;
