CREATE TABLE orders (
  id varchar(255) PRIMARY KEY, 
  price float NOT NULL, 
  tax float NOT NULL,
  final_price float NOT NULL
)