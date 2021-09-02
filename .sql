CREATE TABLE users
(
    id         INT          NOT NULL AUTO_INCREMENT,
    firstname  VARCHAR(64)  NOT NULL,
    lastname   VARCHAR(64)  NOT NULL,
    email      VARCHAR(64)  NOT NULL,
    password   VARCHAR(128) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE suppliers
(
    id         INT         NOT NULL AUTO_INCREMENT,
    name       VARCHAR(64) NOT NULL,
--     description VARCHAR(512) NOT NULL,
--     address     VARCHAR(128) NOT NULL,
--     type        VARCHAR(64)  NOT NULL,
--     open_time   TIME         NOT NULL,
--     close_time  TIME         NOT NULL,
    logo       VARCHAR(256),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE products
(
    id          INT           NOT NULL AUTO_INCREMENT,
    supplier_id INT           NOT NULL,
    name        VARCHAR(64)   NOT NULL,
--     description VARCHAR(512),
    price       DECIMAL(9, 2) NOT NULL,
    ingredients JSON          NOT NULL,
    logo        VARCHAR(256),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (supplier_id) REFERENCES suppliers (id) on delete cascade
);

CREATE TABLE product_types
(
    id         INT         NOT NULL AUTO_INCREMENT,
    product_id INT         NOT NULL,
    name       VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES products (id) on delete cascade
);

CREATE TABLE orders
(
    id         INT          NOT NULL AUTO_INCREMENT,
    user_id    INT,
    address    VARCHAR(128) NOT NULL,
    status     ENUM('in progress', 'done') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) on delete set null
);

CREATE TABLE order_products
(
    order_id   INT NOT NULL,
    product_id INT,
    quantity   INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id) on delete cascade,
    FOREIGN KEY (product_id) REFERENCES products (id) on delete set null
);

CREATE TABLE uid_token
(
    user_id     INT,
    access_uid  varchar(128),
    refresh_uid varchar(128),
    FOREIGN KEY (user_id) REFERENCES users (id) on delete cascade
);