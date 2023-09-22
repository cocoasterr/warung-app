CREATE TABLE users (
    id VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE,
    fullName VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    
    CONSTRAINT user_id PRIMARY KEY (id)
);

CREATE INDEX idx_user_id ON users (id);
CREATE INDEX idx_user_username ON users (username);


CREATE TABLE product (
    id VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    stock INT,
    created_at TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP,
    updated_by VARCHAR(255),
    
    CONSTRAINT product_id PRIMARY KEY (id),
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users (username),
    CONSTRAINT fk_updated_by FOREIGN KEY (updated_by) REFERENCES users (username)
);

CREATE INDEX idx_product_id ON product (id);
CREATE INDEX idx_product_name ON product (name);

