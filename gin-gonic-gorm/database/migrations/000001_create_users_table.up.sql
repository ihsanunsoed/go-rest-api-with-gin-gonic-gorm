CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) ,
    address VARCHAR(255) ,
    email VARCHAR(255) NOT NULL,
    born_date TIMESTAMP
);