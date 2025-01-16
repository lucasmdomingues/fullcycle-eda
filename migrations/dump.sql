CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at DATE);
CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance int, created_at DATE, updated_at DATE);
CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);
CREATE TABLE balances (id INTEGER PRIMARY KEY AUTOINCREMENT, account_id VARCHAR(255), amount INT, created_at DATE DEFAULT CURRENT_DATE);
