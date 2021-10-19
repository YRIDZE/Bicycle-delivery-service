CREATE TABLE users
(
    id         INT          NOT NULL AUTO_INCREMENT,
    firstname  VARCHAR(64)  NOT NULL,
    lastname   VARCHAR(64)  NOT NULL,
    email      VARCHAR(64)  NOT NULL Unique,
    password   VARCHAR(128) NOT NULL,
    deleted    TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE TABLE suppliers
(
    id         INT         NOT NULL AUTO_INCREMENT,
    name       VARCHAR(64) NOT NULL,
    type       VARCHAR(64) NOT NULL,
    image      VARCHAR(256),
    opening    TIME      DEFAULT NULL,
    closing    TIME      DEFAULT NULL,
    deleted    TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE products
(
    id          INT           NOT NULL AUTO_INCREMENT,
    supplier_id INT           NOT NULL,
    name        VARCHAR(64)   NOT NULL,
    price       DECIMAL(9, 2) NOT NULL,
    type        VARCHAR(64)   NOT NULL,
    ingredients JSON          NOT NULL,
    image       VARCHAR(256),
    deleted     TIMESTAMP DEFAULT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (supplier_id) REFERENCES suppliers (id)
);

CREATE TABLE orders
(
    id                INT                          NOT NULL AUTO_INCREMENT,
    user_id           INT,
    address           VARCHAR(128)                 NOT NULL,
    phone_number      VARCHAR(16)                  NOT NULL,
    customer_name     VARCHAR(64)                  NOT NULL,
    customer_lastname VARCHAR(64)                  NOT NULL,
    payment_method    ENUM ('Credit Card', 'Cash') NOT NULL,
    status            ENUM ('in progress', 'done') NOT NULL,
    deleted           TIMESTAMP DEFAULT NULL,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE order_products
(
    order_id   INT           NOT NULL,
    product_id INT,
    quantity   INT           NOT NULL,
    price      DECIMAL(9, 2) NOT NULL,
    deleted    TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE uid_token
(
    user_id     INT,
    access_uid  varchar(128),
    refresh_uid varchar(128),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE cart
(
    id         INT NOT NULL AUTO_INCREMENT,
    user_id    INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted    TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE cart_products
(
    cart_id    INT           NOT NULL,
    product_id INT           NOT NULL,
    quantity   INT           NOT NULL,
    price      DECIMAL(9, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (cart_id) REFERENCES cart (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);