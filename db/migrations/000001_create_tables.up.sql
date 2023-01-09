CREATE TABLE IF NOT EXISTS operation_types (
    id SERIAL PRIMARY KEY NOT NULL,
    description TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY NOT NULL,
    document_number TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY NOT NULL,
    account_id INTEGER NOT NULL,
    operation_type_id INTEGER NOT NULL,
    amount DECIMAL(13, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (account_id) REFERENCES accounts (id) ON DELETE RESTRICT,
    FOREIGN KEY (operation_type_id) REFERENCES operation_types (id) ON DELETE RESTRICT
);