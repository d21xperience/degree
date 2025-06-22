
-- CREATE OR REPLACE FUNCTION tabel_d4da6b98fcfd71c58f5a.move_to_ijazah_when_complete()
-- RETURNS TRIGGER AS $$
-- BEGIN
--   IF NEW.is_complete = TRUE AND OLD.is_complete = FALSE THEN
--     -- Insert ke tabel ijazah
--     INSERT INTO ijazah (
--       peserta_didik_id,
--       program_keahlian,
--       paket_keahlian,
--       sekolah_id,
--       npsn,
--       kabupaten_kota,
--       provinsi,
--       nama,
--       tempat_lahir,
--       tanggal_lahir,
--       nis,
--       nisn,
--       nama_ortu_wali,
--       sekolah_penyelenggara_ujian_us,
--       sekolah_penyelenggara_ujian_un,
--       asal_sekolah,
--       nomor_ijazah,
--       tempat_ijazah,
--       tanggal_ijazah,
--       created_at,
--       updated_at
--     )
--     VALUES (
--       NEW.peserta_didik_id,
--       NEW.program_keahlian,
--       NEW.paket_keahlian,
--       NEW.sekolah_id,
--       NEW.npsn,
--       NEW.kabupaten_kota,
--       NEW.provinsi,
--       NEW.nama,
--       NEW.tempat_lahir,
--       NEW.tanggal_lahir,
--       NEW.nis,
--       NEW.nisn,
--       NEW.nama_ortu_wali,
--       NEW.sekolah_penyelenggara_ujian_us,
--       NEW.sekolah_penyelenggara_ujian_un,
--       NEW.asal_sekolah,
--       NEW.nomor_ijazah,
--       NEW.tempat_ijazah,
--       NEW.tanggal_ijazah,
--       NOW(),
--       NOW()
--     );

--     -- Hapus data dari tabel data_nominasi_sementara
--     DELETE FROM tabel_d4da6b98fcfd71c58f5a.data_nominasi_sementara WHERE id = NEW.id;
--   END IF;

--   RETURN NULL; -- Tidak perlu mengembalikan NEW karena data dihapus
-- END;
-- $$ LANGUAGE plpgsql;


-- CREATE TRIGGER trg_move_ijazah_on_complete
-- AFTER UPDATE ON tabel_d4da6b98fcfd71c58f5a.data_nominasi_sementara
-- FOR EACH ROW
-- EXECUTE FUNCTION tabel_d4da6b98fcfd71c58f5a.move_to_ijazah_when_complete();


CREATE OR REPLACE FUNCTION tabel_d4da6b98fcfd71c58f5a.insert_nilaiakhir_on_anggotakelas()
RETURNS TRIGGER AS $$
DECLARE
  v_kurikulum_id SMALLINT;
  v_jurusan_id VARCHAR(25);
  v_tingkat_pendidikan VARCHAR(25);
  v_mapel RECORD;
BEGIN
  -- Ambil data kelas dari tabel_kelas berdasarkan rombongan_belajar_id
  SELECT 
    kurikulum_id, 
    jurusan_id, 
    tingkat_pendidikan_id::VARCHAR
  INTO 
    v_kurikulum_id, 
    v_jurusan_id, 
    v_tingkat_pendidikan
  FROM tabel_kelas
  WHERE rombongan_belajar_id = NEW.rombongan_belajar_id;

  -- Loop semua mata_pelajaran_id dari tabel_kategori_mapel yang cocok
  FOR v_mapel IN
    SELECT mata_pelajaran_id
    FROM tabel_kategori_mapel
    WHERE kurikulum_id = v_kurikulum_id
      AND jurusan_id = v_jurusan_id
      AND tingkat_pendidikan = v_tingkat_pendidikan
      AND deleted_at IS NULL
  LOOP
    INSERT INTO tabel_nilaiakhir (
      peserta_didik_id,
      anggota_rombel_id,
      semester_id,
      semester,
      mata_pelajaran_id
    ) VALUES (
      NEW.peserta_didik_id,
      NEW.anggota_rombel_id,
      NEW.semester_id,
      SUBSTRING(NEW.semester_id FROM 5 FOR 1)::NUMERIC, -- semester 1 atau 2 dari semester_id
      v_mapel.mata_pelajaran_id
    );
  END LOOP;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_insert_nilaiakhir_from_anggotakelas
AFTER INSERT ON tabel_d4da6b98fcfd71c58f5a.tabel_anggotakelas
FOR EACH ROW
EXECUTE FUNCTION tabel_d4da6b98fcfd71c58f5a.insert_nilaiakhir_on_anggotakelas();

