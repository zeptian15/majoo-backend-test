CREATE TABLE users (
  id bigint(20) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name varchar(45) DEFAULT NULL,
  user_name varchar(45) DEFAULT NULL,
  password varchar(225) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);