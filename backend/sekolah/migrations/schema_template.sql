CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE SCHEMA IF NOT EXISTS {{schema_name}};

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_sekolah (
	sekolah_id UUID NOT NULL DEFAULT gen_random_uuid(),
	nama VARCHAR(100) NOT NULL,
	npsn VARCHAR(8) NULL DEFAULT NULL,
	nss VARCHAR(60) NULL DEFAULT NULL,
	alamat TEXT NULL DEFAULT NULL,
	kd_pos VARCHAR(60) NULL DEFAULT NULL,
	telepon VARCHAR(60) NULL DEFAULT NULL,
	fax VARCHAR(60) NULL DEFAULT NULL,
	kelurahan VARCHAR(60) NULL DEFAULT NULL,
	kecamatan VARCHAR(60) NULL DEFAULT NULL,
	kab_kota VARCHAR(60) NULL DEFAULT NULL,
	propinsi VARCHAR(60) NULL DEFAULT NULL,
	website VARCHAR(100) NULL DEFAULT NULL,
	email VARCHAR(50) NULL DEFAULT NULL,
	nm_kepsek VARCHAR(100) NULL DEFAULT NULL,
	nip_kepsek VARCHAR(25) NULL DEFAULT NULL,
	niy_kepsek VARCHAR(30) NULL DEFAULT NULL,
	status_kepemilikan_id NUMERIC(1,0) NULL DEFAULT NULL,
	kode_aktivasi VARCHAR(30) NULL DEFAULT NULL,
	bentuk_pendidikan_id SMALLINT NULL DEFAULT NULL,
	jenjang_pendidikan_id NUMERIC(2,0) NULL DEFAULT NULL,
	is_dapodik BOOLEAN NULL DEFAULT FALSE,
	sekolah_id_dapo UUID NULL DEFAULT NULL,
	lama_pendidikan SMALLINT NULL DEFAULT NULL,
	PRIMARY KEY (sekolah_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_siswa (
	peserta_didik_id UUID NOT NULL DEFAULT gen_random_uuid(),
	nis VARCHAR(20) NULL DEFAULT NULL,
	nisn VARCHAR(13) NULL DEFAULT NULL,
	nm_siswa VARCHAR(100) NOT NULL,
	tempat_lahir VARCHAR(50) NULL DEFAULT NULL,
	tanggal_lahir DATE NULL DEFAULT NULL,
	jenis_kelamin VARCHAR(1) NULL DEFAULT NULL,
	agama VARCHAR(25) NULL DEFAULT NULL,
	alamat_siswa TEXT NULL DEFAULT NULL,
	telepon_siswa VARCHAR(20) NULL DEFAULT NULL,
	diterima_tanggal DATE NULL DEFAULT NULL,
	nm_ayah VARCHAR(100) NULL DEFAULT NULL,
	nm_ibu VARCHAR(100) NULL DEFAULT NULL,
	pekerjaan_ayah VARCHAR(30) NULL DEFAULT NULL,
	pekerjaan_ibu VARCHAR(30) NULL DEFAULT NULL,
	nm_wali VARCHAR(100) NULL DEFAULT NULL,
	pekerjaan_wali VARCHAR(30) NULL DEFAULT NULL,
	nik VARCHAR(30) NULL DEFAULT NULL,
	is_dapodik BOOLEAN NULL DEFAULT FALSE,
	peserta_didik_id_dapo UUID NULL DEFAULT NULL,
	PRIMARY KEY (peserta_didik_id)
);


CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_siswa_pelengkap (
	pelengkap_siswa_id UUID NOT NULL DEFAULT gen_random_uuid(),
	peserta_didik_id UUID NOT NULL,
	status_dalam_kel VARCHAR(30) NULL DEFAULT NULL,
	anak_ke NUMERIC(3,0) NULL DEFAULT NULL,
	sekolah_asal VARCHAR(100) NULL DEFAULT NULL,
	diterima_kelas VARCHAR(20) NULL DEFAULT NULL,
	alamat_ortu TEXT NULL DEFAULT NULL,
	telepon_ortu VARCHAR(20) NULL DEFAULT NULL,
	alamat_wali TEXT NULL DEFAULT NULL,
	telepon_wali VARCHAR(20) NULL DEFAULT NULL,
	foto_siswa VARCHAR(100) NULL DEFAULT NULL,
	PRIMARY KEY (pelengkap_siswa_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk (
	ptk_id UUID NOT NULL,	
	jenis_ptk_id NUMERIC(2,0) NOT NULL DEFAULT 4,
	nama VARCHAR(100) NOT NULL,
	jenis_kelamin VARCHAR(1) NULL,
	agama VARCHAR(25) NULL,
	tempat_lahir VARCHAR(60) NULL,
	tanggal_lahir DATE NULL,
	is_dapodik BOOLEAN NULL DEFAULT FALSE,
	ptk_id_dapodik UUID NULL DEFAULT NULL,
	status_keaktifan_id NUMERIC(2,0) NOT NULL DEFAULT 1,  
    soft_delete NUMERIC(1,0) NOT NULL DEFAULT 0, 
	PRIMARY KEY (ptk_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk_terdaftar (
	ptk_terdaftar_id UUID NOT NULL,
	ptk_id UUID NOT NULL,
	tahun_ajaran_id VARCHAR(4) NOT NULL,
	jenis_keluar_id CHAR(1) NULL DEFAULT NULL,
	is_dapodik BOOLEAN NULL DEFAULT FALSE,
	ptk_terdaftar_id_dapodik UUID NULL DEFAULT NULL,
	soft_delete NUMERIC(1,0) NOT NULL DEFAULT 0,
	PRIMARY KEY (ptk_terdaftar_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk_pelengkap (
	ptk_pelengkap_id UUID NOT NULL DEFAULT gen_random_uuid(),
	ptk_id UUID NOT NULL,
	gelar_depan VARCHAR(20) NULL DEFAULT NULL,
	gelar_belakang VARCHAR(20) NULL DEFAULT NULL,
	nik CHAR(16) NULL DEFAULT NULL,
	no_kk CHAR(16) NULL DEFAULT NULL,
	nuptk VARCHAR(16) NULL DEFAULT NULL,
	niy VARCHAR(18) NULL DEFAULT NULL,
	nip VARCHAR(18) NULL DEFAULT NULL,
	alamat_jalan VARCHAR(200) NULL DEFAULT NULL,
	rt VARCHAR(3) NULL DEFAULT NULL,
	rw VARCHAR(3) NULL DEFAULT NULL,
	desa_kelurahan VARCHAR(60) NULL DEFAULT NULL,
	kec VARCHAR(60) NULL DEFAULT NULL,
	kab_kota VARCHAR(60) NULL DEFAULT NULL,
	propinsi VARCHAR(60) NULL DEFAULT NULL,
	kode_pos CHAR(5) NULL DEFAULT NULL,
	no_telepon_rumah VARCHAR(20) NULL DEFAULT NULL,
	no_hp VARCHAR(20) NULL DEFAULT NULL,
	email VARCHAR(60) NULL DEFAULT NULL,
	PRIMARY KEY (ptk_pelengkap_id)
);


CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_kelas (
	rombongan_belajar_id UUID NOT NULL,
	sekolah_id UUID NULL DEFAULT NULL,
	semester_id CHAR(5) NOT NULL,
	jurusan_id VARCHAR(25) NULL DEFAULT NULL,
	ptk_id UUID NULL DEFAULT NULL,
	nm_kelas VARCHAR(30) NOT NULL,
	tingkat_pendidikan_id NUMERIC(2,0) NULL DEFAULT NULL,
	jenis_rombel NUMERIC(2,0) NULL DEFAULT NULL,
	nama_jurusan_sp VARCHAR(100) NULL DEFAULT NULL,
	jurusan_sp_id UUID NULL DEFAULT NULL,
	kurikulum_id SMALLINT NULL DEFAULT NULL,
	PRIMARY KEY (rombongan_belajar_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_anggotakelas (
	anggota_rombel_id UUID NOT NULL DEFAULT gen_random_uuid(),
	peserta_didik_id UUID NOT NULL,
	rombongan_belajar_id UUID NULL DEFAULT NULL,
	semester_id CHAR(5) NOT NULL,
	-- status_keaktifan INTEGER NOT NULL DEFAULT 0,
	PRIMARY KEY (anggota_rombel_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_pembelajaran (
	pembelajaran_id UUID NOT NULL DEFAULT gen_random_uuid(),
	rombongan_belajar_id UUID NOT NULL,
	mata_pelajaran_id INTEGER NOT NULL,
	semester_id CHAR(5) NOT NULL,
	ptk_terdaftar_id UUID NULL DEFAULT NULL,
	status_di_kurikulum NUMERIC(2,0) NULL DEFAULT NULL,
	nama_mata_pelajaran VARCHAR(50) NULL DEFAULT NULL,
	induk_pembelajaran UUID NULL DEFAULT NULL,
	is_dapo NUMERIC(1,0) NULL DEFAULT '1',
	ptk_terdaftar_id_dapo UUID NULL DEFAULT NULL,
	PRIMARY KEY (pembelajaran_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_nilaiakhir (
	id_nilai_akhir UUID NOT NULL DEFAULT gen_random_uuid(),
	anggota_rombel_id UUID NULL DEFAULT NULL,
	mata_pelajaran_id INTEGER NULL DEFAULT NULL,
	semester_id CHAR(5) NULL DEFAULT NULL,
	nilai_peng NUMERIC(5,0) NULL DEFAULT NULL,
	predikat_peng VARCHAR(1) NULL DEFAULT NULL,
	nilai_ket NUMERIC(5,0) NULL DEFAULT NULL,
	predikat_ket VARCHAR(1) NULL DEFAULT NULL,
	nilai_sik NUMERIC(2,0) NULL DEFAULT NULL,
	predikat_sik VARCHAR(15) NULL DEFAULT NULL,
	nilai_siksos NUMERIC(2,0) NULL DEFAULT NULL,
	predikat_siksos VARCHAR(15) NULL DEFAULT NULL,
	peserta_didik_id UUID NULL DEFAULT NULL,
	id_minat VARCHAR(2) NULL DEFAULT NULL,
	semester NUMERIC(1,0) NULL DEFAULT NULL,
	PRIMARY KEY (id_nilai_akhir)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_kenaikan (
	kd_kenaikan UUID NOT NULL DEFAULT gen_random_uuid(),
	semester_id CHAR(5) NOT NULL,
	anggota_rombel_id UUID NOT NULL,
	peserta_didik_id UUID NULL DEFAULT NULL,
	kenaikan NUMERIC(3,0) NULL DEFAULT NULL,
	tingkat NUMERIC(3,0) NULL DEFAULT NULL,
	PRIMARY KEY (kd_kenaikan)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.ijazah (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
	PRIMARY KEY (id),
	peserta_didik_id UUID NULL DEFAULT NULL,
    program_keahlian VARCHAR(100) NOT NULL,
    paket_keahlian VARCHAR(100) NOT NULL,
    sekolah_id UUID NOT NULL,
    npsn VARCHAR(15) NOT NULL,
    kabupaten_kota VARCHAR(100) NOT NULL,
    provinsi VARCHAR(100) NOT NULL,
    nama VARCHAR(200) NOT NULL,
    tempat_lahir VARCHAR(100) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    nis VARCHAR(20) UNIQUE NOT NULL,
    nisn VARCHAR(20) UNIQUE NOT NULL,
    nama_ortu_wali VARCHAR(200) NOT NULL,
    sekolah_penyelenggara_ujian_us VARCHAR(200) NOT NULL,
    sekolah_penyelenggara_ujian_un VARCHAR(200) NOT NULL,
    asal_sekolah VARCHAR(200) NOT NULL,
    nomor_ijazah VARCHAR(50) UNIQUE NOT NULL,
    tempat_ijazah VARCHAR(100) NOT NULL,
    tanggal_ijazah DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);



CREATE TABLE IF NOT EXISTS {{schema_name}}.data_nominasi_sementara (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
	PRIMARY KEY (id),
	peserta_didik_id UUID NULL DEFAULT NULL,
	rombongan_belajar_id UUID NULL DEFAULT NULL,
	tahun_ajaran_id VARCHAR(4) NULL DEFAULT NULL,
    program_keahlian VARCHAR(100) NOT NULL,
    paket_keahlian VARCHAR(100) NOT NULL,
    sekolah_id UUID NOT NULL,
    npsn VARCHAR(15) NOT NULL,
    kabupaten_kota VARCHAR(100) NOT NULL,
    provinsi VARCHAR(100) NOT NULL,
    nama VARCHAR(200) NOT NULL,
    tempat_lahir VARCHAR(100) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    nis VARCHAR(20) UNIQUE NOT NULL,
    nisn VARCHAR(20) UNIQUE NOT NULL,
    nama_ortu_wali VARCHAR(200) NOT NULL,
    sekolah_penyelenggara_ujian_us VARCHAR(200) NOT NULL,
    sekolah_penyelenggara_ujian_un VARCHAR(200) NOT NULL,
    asal_sekolah VARCHAR(200) NOT NULL,
    nomor_ijazah VARCHAR(50) UNIQUE NOT NULL,
    tempat_ijazah VARCHAR(100) NOT NULL,
    tanggal_ijazah DATE NOT NULL,
	is_complete BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
	CONSTRAINT data_nominasi_sementara_unique UNIQUE (peserta_didik_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_informasi_ijazah (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	tahun_ajaran_id NUMERIC NULL DEFAULT NULL,
	tempat_dikeluarkan_ijazah VARCHAR NULL DEFAULT NULL,
	tgl_dikeluarkan_ijazah DATE NULL DEFAULT NULL,
	ptk_id UUID NULL DEFAULT NULL,
	kop_sekolah_url VARCHAR NULL DEFAULT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE  IF NOT EXISTS {{schema_name}}.tabel_kategori_sekolah (
	kategori_sekolah_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	kurikulum_id SMALLINT NULL DEFAULT NULL,
	jurusan_id VARCHAR(25) NULL DEFAULT NULL,
	nama_kurikulum VARCHAR NULL DEFAULT NULL,
	nama_bidang_keahlian VARCHAR NULL DEFAULT NULL,
	nama_program_keahlian VARCHAR NULL DEFAULT NULL,
	nama_jurusan VARCHAR NULL DEFAULT NULL,
	jenjang_pendidikan_id NUMERIC(2,0) NULL DEFAULT NULL,
	tingkat_id INTEGER NULL DEFAULT NULL,
	jumlah INTEGER NULL DEFAULT NULL,
	tahun_ajaran_id NUMERIC(4,0) NOT NULL
);

CREATE TABLE {{schema_name}}.tabel_kategori_mapel (
	id int4 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1 NO CYCLE) NOT NULL,
	kategori_sekolah_id int4 NOT NULL,
	mata_pelajaran_id int4 NOT NULL,
	kurikulum_id int2 NULL,
	jurusan_id varchar(25) DEFAULT NULL::character varying NULL,
	nm_mapel varchar(100) DEFAULT NULL::character varying NULL,
	tingkat_pendidikan varchar(25) DEFAULT NULL::character varying NULL,
	tahun_ajaran_id varchar(4) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT tabel_kategori_mapel_pkey PRIMARY KEY (id),
	CONSTRAINT unique_mata_pelajaran_id UNIQUE (mata_pelajaran_id, kurikulum_id, jurusan_id, tingkat_pendidikan)
);

CREATE OR REPLACE FUNCTION {{schema_name}}.sync_mata_pelajaran_to_tabel_kategori_mapel()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO {{schema_name}}.tabel_kategori_mapel (
        kategori_sekolah_id,
        mata_pelajaran_id,
        nm_mapel,
        tingkat_pendidikan,
        jurusan_id,
        kurikulum_id,
        tahun_ajaran_id
    )
    SELECT 
        NEW.kategori_sekolah_id,
        mpk.mata_pelajaran_id,
        mp.nama,
        tp.tingkat_pendidikan_id,
        NEW.jurusan_id,
        mpk.kurikulum_id,
        NEW.tahun_ajaran_id
    FROM ref.mata_pelajaran_kurikulum mpk
    JOIN ref.mata_pelajaran mp
        ON mp.mata_pelajaran_id = mpk.mata_pelajaran_id
    JOIN ref.tingkat_pendidikan tp
        ON tp.jenjang_pendidikan_id = NEW.jenjang_pendidikan_id
    WHERE mpk.kurikulum_id = NEW.kurikulum_id 
      AND mpk.status_di_kurikulum IN (3,8,9)
      AND NOT EXISTS (
          SELECT 1 FROM {{schema_name}}.tabel_kategori_mapel tkm
          WHERE tkm.kategori_sekolah_id = NEW.kategori_sekolah_id
            AND tkm.mata_pelajaran_id = mpk.mata_pelajaran_id
            AND tkm.kurikulum_id = mpk.kurikulum_id
            AND tkm.tahun_ajaran_id = NEW.tahun_ajaran_id::VARCHAR
      )
    ON CONFLICT (mata_pelajaran_id, kurikulum_id, jurusan_id, tingkat_pendidikan)
    DO NOTHING;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE TRIGGER trigger_sync_mata_pelajaran
AFTER INSERT OR UPDATE OF kurikulum_id ON {{schema_name}}.tabel_kategori_sekolah
FOR EACH ROW
WHEN (NEW.kurikulum_id IS NOT NULL AND NEW.jenjang_pendidikan_id IS NOT NULL AND NEW.tahun_ajaran_id IS NOT NULL AND NEW.kategori_sekolah_id IS NOT NULL)
EXECUTE FUNCTION {{schema_name}}.sync_mata_pelajaran_to_tabel_kategori_mapel();