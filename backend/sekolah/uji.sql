CREATE TABLE transaksi_blockchain (
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
