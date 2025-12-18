package pkg

import (
	"fmt"
	"task-management-backend/configs"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	postgresDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	if cfg.Database.MaxOpenConns > 0 {
		postgresDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	} else {
		postgresDB.SetMaxOpenConns(25)
	}

	if cfg.Database.MaxIdleConns > 0 {
		postgresDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	} else {
		postgresDB.SetMaxIdleConns(25)
	}

	if cfg.Database.ConnMaxLifetime > 0 {
		postgresDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Minute)
	} else {
		postgresDB.SetConnMaxLifetime(time.Hour)
	}

	if cfg.Database.ConnMaxIdleTime > 0 {
		postgresDB.SetConnMaxIdleTime(time.Duration(cfg.Database.ConnMaxIdleTime) * time.Minute)
	} else {
		postgresDB.SetConnMaxIdleTime(1 * time.Minute)
	}

	if err := postgresDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &database{db: db}, nil
}

func (d *database) GetDB() *gorm.DB {
	return d.db
}

func (d *database) Close() error {
	postgresDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return postgresDB.Close()
}

func (d *database) Ping() error {
	postgresDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return postgresDB.Ping()
}

func (d *database) Transaction(fn func(tx *gorm.DB) error) error {
	return d.db.Transaction(fn)
}

func (d *database) Migrate(models ...any) error {
	return d.db.AutoMigrate(models...)
}
