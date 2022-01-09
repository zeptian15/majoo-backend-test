CREATE TABLE outlets (
  id bigint(20) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id int(40) NOT NULL,
  merchant_id int(40) NOT NULL,
  outlet_name varchar(45) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

-- Populate Database
INSERT INTO outlets ( id, user_id, merchant_id, outlet_name )
VALUES
(1, 1, 1, "Outlet 1"),
(2, 2, 2, "Outlet 2"),
(3, 1, 1, "Outlet 3");