
-- +migrate Up

CREATE TABLE permissions
(
  id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  resourse text NOT NULL,
  action_id integer,
  policy_id integer NOT NULL,
  access boolean,
  CONSTRAINT action_id FOREIGN KEY (action_id)
      REFERENCES actions (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT policy_id FOREIGN KEY (policy_id)
      REFERENCES policies (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- +migrate Down

DROP TABLE permissions;
