CREATE TABLE IF NOT EXISTS account (
  id integer NOT NULL AUTO_INCREMENT,
  name varchar(255),
  PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS charge (
  id integer NOT NULL AUTO_INCREMENT,
  account_id integer,
  cents integer,
  timestamp datetime,
  PRIMARY KEY (id)
);
