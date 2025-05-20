package repositories

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type SchemaRepository interface {
	InitializeDatabase(ctx context.Context, schemaFile, schemaName string) error
	SetSchema(schemaName string) error
}

type schemaRepositoryImpl struct {
	db *gorm.DB
}

func NewSchemaRepository(db *gorm.DB) SchemaRepository {
	return &schemaRepositoryImpl{db: db}
}

// ExecuteSQL menjalankan perintah SQL dari string
func (r *schemaRepositoryImpl) ExecuteSQL(query string) error {
	if err := r.db.Exec(query).Error; err != nil {
		return err
	}
	return nil
}

// LoadSQLFile membaca isi file SQL lalu mengganti placeholder {{schema_name}}
// File dibaca baris demi baris untuk mengurangi penggunaan memori.
func (r *schemaRepositoryImpl) LoadSQLFile(filePath, schemaName string) (string, error) {
	// Validasi parameter
	if filePath == "" {
		return "", fmt.Errorf("file path is empty")
	}
	if schemaName == "" {
		return "", fmt.Errorf("schema name is empty")
	}

	// Validasi schemaName untuk karakter yang aman
	isValidSchemaName := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	if !isValidSchemaName(schemaName) {
		return "", fmt.Errorf("invalid schema name: %s", schemaName)
	}

	// Buka file untuk dibaca
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Buffer untuk menampung konten SQL
	var builder strings.Builder
	scanner := bufio.NewScanner(file)

	// Iterasi setiap baris dalam file
	for scanner.Scan() {
		line := scanner.Text()

		// Ganti placeholder {{schema_name}} dengan schemaName
		line = strings.ReplaceAll(line, "{{schema_name}}", schemaName)

		// Tambahkan baris ke buffer
		builder.WriteString(line)
		builder.WriteString("\n") // Tambahkan newline untuk setiap baris
	}

	// Periksa error pada scanner
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return builder.String(), nil
}

func (r *schemaRepositoryImpl) InitializeDatabase(ctx context.Context, schemaFile, schemaName string) error {
	var err error
	log.Printf("Initializing database for schema: %s", schemaName)

	// 1. Mulai transaksi
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		log.Printf("❌ Failed to start transaction: %v", tx.Error)
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}
	defer func() {
		if r := recover(); r != nil {
			log.Printf("🚨 Panic recovered in transaction: %v", r)
			tx.Rollback()
			err = fmt.Errorf("panic occurred: %v", r)
		} else if err != nil {
			log.Printf("❌ Rolling back transaction due to error: %v", err)
			tx.Rollback()
		} else {
			log.Printf("✅ Committing transaction")
			tx.Commit()
		}
	}()

	// 2. Muat file SQL
	log.Printf("🔐 Transaction started for schema: %s", schemaName)
	sqlContent, err := r.LoadSQLFile(schemaFile, schemaName)
	if err != nil {
		log.Printf("❌ Error loading SQL file (%s): %v", schemaFile, err)
		err = fmt.Errorf("failed to load SQL file: %w", err)
		return err
	}

	if len(sqlContent) == 0 {
		log.Printf("❌ SQL content is empty for schema: %s", schemaName)
		err = fmt.Errorf("sql content is empty")
		return err
	}
	log.Printf("✅ SQL file loaded successfully")

	// 3. Eksekusi SQL satu per satu
	statements := strings.Split(sqlContent, ";")
	for i, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		log.Printf("🌀 Executing statement #%d for schema %s", i+1, schemaName)
		if execErr := tx.Exec(stmt).Error; execErr != nil {
			log.Printf("❌ Error executing statement #%d: %v", i+1, execErr)
			err = fmt.Errorf("error at statement #%d: %w", i+1, execErr)
			return err
		}
	}
	log.Printf("✅ All SQL statements executed successfully")

	// 4. Simpan ke schema_logs
	log.Printf("📝 Inserting schema log for: %s", schemaName)
	if insertErr := tx.Exec("INSERT INTO public.schema_logs (schema_name, created_at) VALUES (?, NOW())", schemaName).Error; insertErr != nil {
		log.Printf("❌ Error saving schema log for (%s): %v", schemaName, insertErr)
		err = fmt.Errorf("failed to save schema log: %w", insertErr)
		return err
	}
	log.Printf("✅ Schema log inserted successfully")

	log.Printf("🎉 Schema %s successfully initialized", schemaName)
	return nil
}

// Fungsi untuk mengganti schema dinamis
func (r *schemaRepositoryImpl) SetSchema(schemaName string) error {
	return r.db.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error
}
