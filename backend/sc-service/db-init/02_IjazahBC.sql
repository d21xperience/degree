-- CREATE TABLE IF NOT EXISTS blockchain_platform (
-- 	"id" UUID NOT NULL,
-- 	"nm_blockchain" VARCHAR(50) NOT NULL,
-- 	"active" BOOLEAN NULL DEFAULT 'false',
-- 	"created_at" TIMESTAMP NULL DEFAULT 'now()',
-- 	"updated_at" TIMESTAMP NULL DEFAULT 'now()',
-- 	PRIMARY KEY ("id")
-- );
-- Tabel schema_logs
CREATE TABLE schema_logs (
    id BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NULL,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    schema_name TEXT DEFAULT NULL
);

CREATE INDEX idx_schema_logs_deleted_at ON schema_logs (deleted_at);

-- Tabel sekolah_tenants
CREATE TABLE sekolah_tenants (
    id BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NULL,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    nama_sekolah TEXT DEFAULT NULL,
    user_id INTEGER DEFAULT NULL,
    sekolah_id INTEGER DEFAULT NULL,
    schema_name TEXT NOT NULL,
    sekolah_tenant_id BIGINT NOT NULL,
    CONSTRAINT uni_sekolah_tenants_schema_name UNIQUE (schema_name)
);

CREATE INDEX idx_sekolah_tenants_deleted_at ON sekolah_tenants (deleted_at);

CREATE TABLE IF NOT EXISTS ijazah_bc (
	id uuid NOT NULL,
	peserta_didik_id uuid NOT NULL,
	nama varchar(255) NULL,
	nis varchar(50) NULL,
	nisn varchar(50) NULL,
	npsn varchar(50) NULL,
	nomor_ijazah varchar(100) NULL,
	tempat_lahir varchar(100) NULL,
	tanggal_lahir date NULL,
	nama_ortuwali varchar(255) NULL,
	paket_keahlian varchar(255) NULL,
	kabupatenkota varchar(100) NULL,
	provinsi varchar(100) NULL,
	program_keahlian varchar(255) NULL,
	sekolah_penyelenggara_ujian_us varchar(255) NULL,
	sekolah_penyelenggara_ujian_un varchar(255) NULL,
	asal_sekolah varchar(255) NULL,
	tempat_ijazah varchar(100) NULL,
	tanggal_ijazah date NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT ijazah_bc_pkey PRIMARY KEY (id),
	CONSTRAINT ijazah_bc_unique UNIQUE (peserta_didik_id)
);
CREATE TABLE IF NOT EXISTS degree_data (
	id serial4 NOT NULL,
	ijazah_id uuid NOT NULL,
	degree_hash varchar(255) NULL,
	tx_hash varchar(255) NULL,
	ipfs_url text NULL,
	bc_type varchar(50) NULL,
	link_bc_explorer text NULL,
	sekolah_id uuid NOT NULL,
	tahun_ajaran_id varchar NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT degree_data_pkey PRIMARY KEY (id)
);



