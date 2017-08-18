
-- +migrate Up

CREATE TABLE group_policies
(
  group_id integer NOT NULL,
  policy_id integer NOT NULL,
  CONSTRAINT group_policies_id PRIMARY KEY (group_id, policy_id),
  CONSTRAINT group_id FOREIGN KEY (group_id)
      REFERENCES groups (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT policy_id FOREIGN KEY (policy_id)
      REFERENCES policies (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE user_groups
(
  user_id integer NOT NULL,
  group_id integer NOT NULL,
  CONSTRAINT user_group_id PRIMARY KEY (user_id, group_id),
  CONSTRAINT group_id FOREIGN KEY (group_id)
      REFERENCES groups (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT user_id FOREIGN KEY (user_id)
      REFERENCES users (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE user_policies
(
  user_id integer NOT NULL,
  policy_id integer NOT NULL,
  CONSTRAINT user_policy_id PRIMARY KEY (user_id, policy_id),
  CONSTRAINT policy_id FOREIGN KEY (policy_id)
      REFERENCES policies (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT user_id FOREIGN KEY (user_id)
      REFERENCES users (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- +migrate Down

DROP TABLE user_groups;
DROP TABLE user_policies;
DROP TABLE group_policies;
