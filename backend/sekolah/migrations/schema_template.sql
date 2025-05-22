CREATE EXTENSION IF NOT EXISTS "pgcrypto";
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
	is_dapodik NUMERIC(1,0) NOT NULL DEFAULT 0,
	sekolah_id_dapo UUID NULL DEFAULT NULL,
	PRIMARY KEY (sekolah_id),
	CONSTRAINT FK_tabel_sekolah_ref_bentuk_pendidikan FOREIGN KEY (bentuk_pendidikan_id) REFERENCES ref.bentuk_pendidikan (bentuk_pendidikan_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_sekolah_ref_jenjang_pendidikan FOREIGN KEY (jenjang_pendidikan_id) REFERENCES ref.jenjang_pendidikan (jenjang_pendidikan_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_sekolah_ref_status_kepemilikan FOREIGN KEY (status_kepemilikan_id) REFERENCES ref.status_kepemilikan (status_kepemilikan_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_siswa (
	peserta_didik_id UUID NOT NULL DEFAULT gen_random_uuid(),
	nis VARCHAR(20) NOT NULL,
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
	is_dapodik NUMERIC(1,0) NOT NULL DEFAULT 0,
	peserta_didik_id_dapo UUID NULL DEFAULT NULL,
	PRIMARY KEY (peserta_didik_id)
);


CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_siswa_pelengkap (
	pelengkap_siswa_id UUID NOT NULL DEFAULT gen_random_uuid(),
	peserta_didik_id UUID NULL DEFAULT NULL,
	status_dalam_kel VARCHAR(30) NULL DEFAULT NULL,
	anak_ke NUMERIC(3,0) NULL DEFAULT NULL,
	sekolah_asal VARCHAR(100) NULL DEFAULT NULL,
	diterima_kelas VARCHAR(20) NULL DEFAULT NULL,
	alamat_ortu TEXT NULL DEFAULT NULL,
	telepon_ortu VARCHAR(20) NULL DEFAULT NULL,
	alamat_wali TEXT NULL DEFAULT NULL,
	telepon_wali VARCHAR(20) NULL DEFAULT NULL,
	foto_siswa VARCHAR(100) NULL DEFAULT NULL,
	PRIMARY KEY (pelengkap_siswa_id),
	CONSTRAINT FK_tabel_siswa_pelengkap_tabel_siswa FOREIGN KEY (peserta_didik_id)
		REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id)
		ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk (
	ptk_id UUID NOT NULL,
	jenis_ptk_id NUMERIC(2,0) NOT NULL DEFAULT 4,
	nama VARCHAR(100) NOT NULL,
	jenis_kelamin VARCHAR(1) NULL,
	agama VARCHAR(25) NULL,
	tempat_lahir VARCHAR(32) NULL,
	tanggal_lahir DATE NULL,
	is_dapodik NUMERIC(1,0) NOT NULL DEFAULT 0,
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
	is_dapodik NUMERIC(1,0) NOT NULL DEFAULT 0,
	ptk_terdaftar_id_dapo UUID NULL DEFAULT NULL,
	soft_delete NUMERIC(1,0) NOT NULL DEFAULT 0,
	PRIMARY KEY (ptk_terdaftar_id),
	CONSTRAINT FK_tabel_ptk_terdaftar_tabel_ptk FOREIGN KEY (ptk_id) REFERENCES {{schema_name}}.tabel_ptk (ptk_id) ON UPDATE CASCADE ON DELETE CASCADE
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
	rt NUMERIC(2,0) NULL DEFAULT NULL,
	rw NUMERIC(2,0) NULL DEFAULT NULL,
	desa_kelurahan VARCHAR(60) NULL DEFAULT NULL,
	kec VARCHAR(60) NULL DEFAULT NULL,
	kab_kota VARCHAR(60) NULL DEFAULT NULL,
	propinsi VARCHAR(60) NULL DEFAULT NULL,
	kode_pos CHAR(5) NULL DEFAULT NULL,
	no_telepon_rumah VARCHAR(20) NULL DEFAULT NULL,
	no_hp VARCHAR(20) NULL DEFAULT NULL,
	email VARCHAR(60) NULL DEFAULT NULL,
	PRIMARY KEY (ptk_pelengkap_id),
	CONSTRAINT FK_tabel_ptk_pelengkap_tabel_ptk FOREIGN KEY (ptk_id) REFERENCES {{schema_name}}.tabel_ptk (ptk_id) ON UPDATE CASCADE ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_kelas (
	rombongan_belajar_id UUID NOT NULL,
	sekolah_id UUID NOT NULL,
	semester_id CHAR(5) NOT NULL,
	jurusan_id VARCHAR(25) NULL DEFAULT NULL,
	ptk_id UUID NULL DEFAULT NULL,
	nm_kelas VARCHAR(30) NULL DEFAULT NULL,
	tingkat_pendidikan_id NUMERIC(2,0) NULL DEFAULT NULL,
	jenis_rombel NUMERIC(2,0) NULL DEFAULT NULL,
	nama_jurusan_sp VARCHAR(100) NULL DEFAULT NULL,
	jurusan_sp_id UUID NULL DEFAULT NULL,
	kurikulum_id SMALLINT NOT NULL,
	PRIMARY KEY (rombongan_belajar_id),
	CONSTRAINT FK_tabel_kelas_ref_jurusan FOREIGN KEY (jurusan_id) REFERENCES ref.jurusan (jurusan_id) ON UPDATE CASCADE ON DELETE SET NULL,
	CONSTRAINT FK_tabel_kelas_ref_kurikulum FOREIGN KEY (kurikulum_id) REFERENCES ref.kurikulum (kurikulum_id) ON UPDATE CASCADE ON DELETE SET NULL,
	CONSTRAINT FK_tabel_kelas_ref_semester FOREIGN KEY (semester_id) REFERENCES ref.semester (semester_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_kelas_ref_tingkat_pendidikan FOREIGN KEY (tingkat_pendidikan_id) REFERENCES ref.tingkat_pendidikan (tingkat_pendidikan_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_kelas_tabel_ptk FOREIGN KEY (ptk_id) REFERENCES {{schema_name}}.tabel_ptk (ptk_id) ON UPDATE CASCADE ON DELETE SET NULL

);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_anggotakelas (
	anggota_rombel_id UUID NOT NULL DEFAULT gen_random_uuid(),
	peserta_didik_id UUID NOT NULL,
	rombongan_belajar_id UUID NOT NULL,
	semester_id CHAR(5) NOT NULL,
	status_keaktifan INTEGER NOT NULL DEFAULT 0,
	PRIMARY KEY (anggota_rombel_id),
	CONSTRAINT FK_tabel_anggotakelas_tabel_anggotakelas FOREIGN KEY (rombongan_belajar_id) REFERENCES {{schema_name}}.tabel_kelas (rombongan_belajar_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_anggotakelas_tabel_siswa FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_anggotakelas_ref_semester FOREIGN KEY (semester_id) REFERENCES ref.semester (semester_id) ON UPDATE CASCADE ON DELETE SET NULL
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
	PRIMARY KEY (pembelajaran_id),
	CONSTRAINT FK_tabel_pembelajaran_ref_mata_pelajaran FOREIGN KEY (mata_pelajaran_id) REFERENCES ref.mata_pelajaran (mata_pelajaran_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_pembelajaran_ref_semester FOREIGN KEY (semester_id) REFERENCES ref.semester (semester_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_pembelajaran_tabel_kelas FOREIGN KEY (rombongan_belajar_id) REFERENCES {{schema_name}}.tabel_kelas (rombongan_belajar_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_pembelajaran_tabel_ptk_terdaftar FOREIGN KEY (ptk_terdaftar_id) REFERENCES {{schema_name}}.tabel_ptk_terdaftar (ptk_terdaftar_id) ON UPDATE CASCADE ON DELETE SET NULL
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
	PRIMARY KEY (id_nilai_akhir),
	CONSTRAINT FK_tabel_nilaiakhir_tabel_anggotakelas FOREIGN KEY (anggota_rombel_id) REFERENCES {{schema_name}}.tabel_anggotakelas (anggota_rombel_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_nilaiakhir_tabel_siswa FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_kenaikan (
	kd_kenaikan UUID NOT NULL DEFAULT gen_random_uuid(),
	semester_id CHAR(5) NOT NULL,
	anggota_rombel_id UUID NOT NULL,
	peserta_didik_id UUID NULL DEFAULT NULL,
	kenaikan NUMERIC(3,0) NULL DEFAULT NULL,
	tingkat NUMERIC(3,0) NULL DEFAULT NULL,
	PRIMARY KEY (kd_kenaikan),
	CONSTRAINT FK_tabel_kenaikan_ref_semester FOREIGN KEY (semester_id) REFERENCES ref.semester (semester_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_kenaikan_tabel_anggotakelas FOREIGN KEY (anggota_rombel_id) REFERENCES {{schema_name}}.tabel_anggotakelas (anggota_rombel_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_kenaikan_tabel_siswa FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE

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
    deleted_at TIMESTAMP NULL,
   	CONSTRAINT FK_tabel_ijazah FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE
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
   	CONSTRAINT FK_tabel_ijazah FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE,
   	CONSTRAINT FK_data_nominasi_sementara_tabel_kelas FOREIGN KEY (rombongan_belajar_id) REFERENCES {{schema_name}}.tabel_kelas (rombongan_belajar_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_informasi_ijazah (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	tahun_ajaran_id NUMERIC NULL DEFAULT NULL,
	tempat_dikeluarkan_ijazah VARCHAR NULL DEFAULT NULL,
	tgl_dikeluarkan_ijazah DATE NULL DEFAULT NULL,
	ptk_id UUID NULL DEFAULT NULL,
	kop_sekolah_url VARCHAR NULL DEFAULT NULL,
	PRIMARY KEY (id),
	CONSTRAINT FK_tabel_informasi_ijazah_tabel_ptk FOREIGN KEY (ptk_id) REFERENCES {{schema_name}}.tabel_ptk (ptk_id) ON UPDATE CASCADE ON DELETE SET NULL
);