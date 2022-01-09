CREATE TABLE merchants (
  id bigint(20) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id int(40) NOT NULL,
  merchant_name varchar(45) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

-- Populate Database
INSERT INTO merchants ( id, user_id, merchant_name )
VALUES
(1, 1, "Merchant 1"),
(2, 2, "Merchant 2");