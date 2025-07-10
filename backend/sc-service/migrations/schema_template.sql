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

CREATE TABLE {{schema_name}}.transaksi_blockchain (
    id SERIAL PRIMARY KEY,
    tx_hash VARCHAR(66) NOT NULL UNIQUE,           -- 66 karakter: "0x" + 64 karakter hash
    from_address VARCHAR(42) NOT NULL,             -- 42 karakter: "0x" + 40 karakter address
    to_address VARCHAR(42),
    value NUMERIC(78, 0) DEFAULT 0,                -- nilai dalam wei (gunakan tipe besar)
    gas_limit BIGINT,
    gas_price NUMERIC(78, 0),
    nonce INTEGER,
    data TEXT,
    chain_id INTEGER,
    block_number BIGINT,
    status BOOLEAN,                                -- TRUE = sukses, FALSE = gagal
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- waktu pencatatan transaksi
);
