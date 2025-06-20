CREATE OR REPLACE FUNCTION tabel_d4da6b98fcfd71c58f5a.sync_mata_pelajaran_to_tabel_kategori_mapel()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO tabel_d4da6b98fcfd71c58f5a.tabel_kategori_mapel (
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
          SELECT 1 FROM tabel_d4da6b98fcfd71c58f5a.tabel_kategori_mapel tkm
          WHERE tkm.kategori_sekolah_id = NEW.kategori_sekolah_id
            AND tkm.mata_pelajaran_id = mpk.mata_pelajaran_id
            AND tkm.kurikulum_id = mpk.kurikulum_id
            AND tkm.tahun_ajaran_id = NEW.tahun_ajaran_id::VARCHAR
      )
    ON CONFLICT ( mata_pelajaran_id, kurikulum_id, jurusan_id, tingkat_pendidikan )
    DO NOTHING;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE TRIGGER trigger_sync_mata_pelajaran
AFTER INSERT OR UPDATE OF kurikulum_id ON tabel_d4da6b98fcfd71c58f5a.tabel_kategori_sekolah
FOR EACH ROW
WHEN (NEW.kurikulum_id IS NOT NULL AND NEW.jenjang_pendidikan_id IS NOT NULL AND NEW.tahun_ajaran_id IS NOT NULL AND NEW.kategori_sekolah_id IS NOT NULL)
EXECUTE FUNCTION tabel_d4da6b98fcfd71c58f5a.sync_mata_pelajaran_to_tabel_kategori_mapel();