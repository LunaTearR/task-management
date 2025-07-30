package pkg

import (
	"fmt"
	"task-management-user-service/configs"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	GetDB() *gorm.DB
	Close() error
	Ping() error
	Transaction(fn func(tx *gorm.DB) error) error
	Migrate(models ...any) error
}

type database struct {
	db *gorm.DB
}

func NewPostgresDB(cfg *configs.Config) (Database, error) {
	dsn := cfg.Database.Connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("เปิดการเชื่อมต่อ GORM ไม่ได้: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("เข้าถึง SQL DB ไม่ได้: %w", err)
	}

	if cfg.Database.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	} else {
		sqlDB.SetMaxOpenConns(25) // ค่า default ตัวอย่าง
	}

	if cfg.Database.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	} else {
		sqlDB.SetMaxIdleConns(5)
	}

	if cfg.Database.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Minute)
	} else {
		sqlDB.SetConnMaxLifetime(time.Hour) // ค่า default ตัวอย่าง: connection มีอายุสูงสุด 1 ชั่วโมง
	}

	if cfg.Database.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(cfg.Database.ConnMaxIdleTime) * time.Minute)
	} else {
		sqlDB.SetConnMaxIdleTime(1 * time.Minute) // ค่า default ตัวอย่าง: connection ที่ idle นานเกิน 10 นาทีจะถูกปิด
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping ฐานข้อมูลไม่สำเร็จ: %w", err)
	}

	return &database{db: db}, nil
}

func (d *database) GetDB() *gorm.DB {
	return d.db
}

func (d *database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *database) Ping() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (d *database) Transaction(fn func(tx *gorm.DB) error) error {
	return d.db.Transaction(fn)
}

func (d *database) Migrate(models ...interface{}) error {
	return d.db.AutoMigrate(models...)
}
