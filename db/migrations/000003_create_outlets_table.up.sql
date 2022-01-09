CREATE TABLE outlets (
  id bigint(20) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id int(40) NOT NULL,
  merchant_id int(40) NOT NULL,
  outlet_name varchar(45) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);