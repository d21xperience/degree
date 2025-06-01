CREATE SCHEMA IF NOT EXISTS {{schema_name}};
CREATE TABLE IF NOT EXISTS {{schema_name}}.accounts ( 
    id BIGSERIAL PRIMARY KEY,
    address TEXT NOT NULL,
    username VARCHAR(100) NULL DEFAULT NULL,
    type TEXT NULL DEFAULT 'IMPORTED',
    user_id INTEGER NOT NULL,
    password VARCHAR(100) NULL DEFAULT NULL,
    keystroke_filename TEXT NULL DEFAULT NULL,
    network_id BIGINT NOT NULL,
    organization TEXT NULL DEFAULT NULL,
    is_active BOOLEAN NULL DEFAULT NULL,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    UNIQUE (address)
);

CREATE INDEX idx_accounts_network_id ON {{schema_name}}.accounts (network_id);
CREATE INDEX idx_accounts_user_id ON {{schema_name}}.accounts (user_id);

CREATE TABLE IF NOT EXISTS {{schema_name}}.contracts (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    contract_address TEXT NOT NULL,
    abi TEXT NOT NULL,
    bytecode TEXT,
    network_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    UNIQUE (contract_address)
);

CREATE INDEX idx_contracts_network_id ON {{schema_name}}.contracts (network_id);

CREATE TABLE IF NOT EXISTS {{schema_name}}.transactions (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    contract_id BIGINT,
    network_id BIGINT NOT NULL,
    method VARCHAR(100),
    input_data TEXT,
    gas_used BIGINT,
    status VARCHAR(20),
    timestamp TIMESTAMPTZ,
    UNIQUE (tx_hash)
);