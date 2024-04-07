CREATE TYPE ROLE AS ENUM ('admin', 'user');
CREATE TYPE SERVER_STATUS AS ENUM ('on', 'off');

CREATE TABLE users {
    id SERIAL PRIMARY KEY
    created_at TIMESTAMP
    created_by INT
    update_at TIMESTAMP
    update_by INT
    detele_at TIMESTAMP
    delete_by INT
    email VARCHAR(255) UNIQUE NOT NULL
    full_name VARCHAR(1000)
    password VARCHAR(2000)
    phone VARCHAR(20)
    avatar VARCHAR(2000)
    is_supper_admin BOOLEAN
    roles ROLE
}

CREATE TABLE servers {
    id SERIAL PRIMARY KEY
    created_at TIMESTAMP
    created_by INT
    update_at TIMESTAMP
    update_by INT
    detele_at TIMESTAMP
    delete_by INT
    "name" VARCHAR(2000) NOT NULL UNIQUE
    status SERVER_STATUS
    ipv4 VARCHAR(16)
}