-- Passo a Passo para logar no MySQL e criar o banco de dados e as tabelas
-- 1. mysql -u root -p
-- 2. password: root_password
-- 3. SHOW DATABASES;
-- 4. USE tracking;

CREATE TABLE `users` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE
);

CREATE TABLE `orders` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    tracking_code VARCHAR(100) UNIQUE,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

