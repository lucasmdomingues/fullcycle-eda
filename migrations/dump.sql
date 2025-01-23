USE wallet;

CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at DATE);
CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance int, created_at DATE, updated_at DATE);
CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);
CREATE TABLE balances (id INTEGER PRIMARY KEY AUTO_INCREMENT, account_id VARCHAR(255), amount INT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);

INSERT INTO customers(id, name, email, created_at) VALUES ('a245d69a-60b3-46a9-957e-20b981c2b623', "John Doe", "johnd@gmail.com", CURRENT_TIMESTAMP);
INSERT INTO customers(id, name, email, created_at) VALUES ('46244fed-16e5-478a-bdc6-e3fc80890553', "Jane Doe", "janed@gmail.com", CURRENT_TIMESTAMP);

INSERT INTO accounts (id, customer_id, balance, created_at, updated_at)
    VALUES ("c21cfb5c-a830-426a-b976-f77d25b2e2bb", "a245d69a-60b3-46a9-957e-20b981c2b623", 1000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO accounts (id, customer_id, balance, created_at, updated_at)
    VALUES ("f1d40484-063c-4909-b552-edbc9bbb71a2", "46244fed-16e5-478a-bdc6-e3fc80890553", 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
