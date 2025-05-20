CREATE SEQUENCE IF NOT EXISTS sekolah_tenant_id_seq;
CREATE TABLE IF NOT EXISTS sekolah_tenant (
	id BIGINT NOT NULL DEFAULT nextval('sekolah_tenant_id_seq'::regclass),
	nama_sekolah TEXT NOT NULL,
	sekolah_tenant_id BIGINT NOT NULL,
	schema_name TEXT NOT NULL,
	-- bentuk_pendidikan_id INTEGER NOT NULL DEFAULT '0',
	created_at TIMESTAMPTZ NULL DEFAULT NULL,
	updated_at TIMESTAMPTZ NULL DEFAULT NULL,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS uni_sekolah_tenant_nama_sekolah ON sekolah_tenant (nama_sekolah);
CREATE UNIQUE INDEX  IF NOT EXISTS uni_sekolah_tenant_sekolah_id ON sekolah_tenant (sekolah_tenant_id);
CREATE UNIQUE INDEX  IF NOT EXISTS uni_sekolah_tenant_schema_name ON sekolah_tenant (schema_name);
CREATE INDEX  IF NOT EXISTS idx_sekolah_tenant_deleted_at ON sekolah_tenant (deleted_at);

CREATE SEQUENCE  IF NOT EXISTS schema_logs_id_seq;
CREATE TABLE  IF NOT EXISTS schema_logs (
	id BIGINT NOT NULL DEFAULT nextval('schema_logs_id_seq'::regclass),
	schema_name TEXT NULL DEFAULT NULL,
	created_at TIMESTAMPTZ NULL DEFAULT NULL,
	updated_at TIMESTAMPTZ NULL DEFAULT NULL,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id)
);
CREATE INDEX idx_schema_logs_deleted_at ON schema_logs (deleted_at)