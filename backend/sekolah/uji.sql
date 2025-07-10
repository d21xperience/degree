-- CREATE TABLE tabel_d4da6b98fcfd71c58f5a.transaksi_blockchain (
--     id SERIAL PRIMARY KEY,
--     tx_hash VARCHAR(66) NOT NULL UNIQUE,           -- 66 karakter: "0x" + 64 karakter hash
--     from_address VARCHAR(42) NOT NULL,             -- 42 karakter: "0x" + 40 karakter address
--     to_address VARCHAR(42),
--     value NUMERIC(78, 0) DEFAULT 0,                -- nilai dalam wei (gunakan tipe besar)
--     gas_limit BIGINT,
--     gas_price NUMERIC(78, 0),
--     nonce INTEGER,
--     data TEXT,
--     chain_id INTEGER,
--     block_number BIGINT,
--     status BOOLEAN,                                -- TRUE = sukses, FALSE = gagal
--     timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- waktu pencatatan transaksi
-- );
CREATE OR REPLACE FUNCTION tabel_d4da6b98fcfd71c58f5a.insert_pembelajaran_on_kelas()
RETURNS TRIGGER AS $$
DECLARE
  v_kurikulum_id SMALLINT;
  v_jurusan_id VARCHAR(25);
  v_tingkat_pendidikan VARCHAR(25);
  v_rombongan_belajar_id VARCHAR;
  v_mapel RECORD;
BEGIN
  -- Ambil data kelas dari tabel_kelas berdasarkan rombongan_belajar_id
  SELECT 
    kurikulum_id, 
    jurusan_id, 
    tingkat_pendidikan_id::VARCHAR,
    rombongan_belajar_id
  INTO 
    v_kurikulum_id, 
    v_jurusan_id, 
    v_tingkat_pendidikan,
    v_rombongan_belajar_id
  FROM tabel_kelas
  WHERE rombongan_belajar_id = NEW.rombongan_belajar_id;

  -- Loop semua mata_pelajaran_id dari tabel_kategori_mapel yang cocok
  FOR v_mapel IN
    SELECT mata_pelajaran_id
    FROM tabel_kategori_mapel
    WHERE kurikulum_id = v_kurikulum_id
      AND jurusan_id = v_jurusan_id
      AND tingkat_pendidikan = v_tingkat_pendidikan
    --   AND deleted_at IS NULL
  LOOP
    INSERT INTO tabel_pembelajaran (
      rombongan_belajar_id,
      mata_pelajaran_id,
      semester_id,
      nama_mata_pelajaran,
      ptk_terdaftar_id,
    ) VALUES (
    --   NEW.peserta_didik_id,
      v_rombongan_belajar_id,
      v_mapel.mata_pelajaran_id,
      NEW.anggota_rombel_id,
      NEW.semester_id,
      SUBSTRING(NEW.semester_id FROM 5 FOR 1)::NUMERIC, -- semester 1 atau 2 dari semester_id
    );
  END LOOP;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_insert_pembelajaran_from_kelas
AFTER INSERT ON tabel_d4da6b98fcfd71c58f5a.tabel_kelas
FOR EACH ROW
EXECUTE FUNCTION tabel_d4da6b98fcfd71c58f5a.insert_pembelajaran_on_kelas();