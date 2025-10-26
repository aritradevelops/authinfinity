package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// implements Database
type Postgres struct {
	uri       string
	db        *gorm.DB
	underline *sql.DB
}

func NewPostgres(uri string) Database {
	return &Postgres{
		uri: uri,
	}
}

func (p *Postgres) Connect() error {
	logFile, err := os.OpenFile("gorm.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(p.uri), &gorm.Config{

		Logger: logger.New(log.New(logFile, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		}),
	})
	if err != nil {
		return err
	}
	underline, err := db.DB()
	if err != nil {
		return err
	}
	p.db = db
	p.underline = underline
	return nil
}
func (p *Postgres) Disconnect() error {
	if p.db == nil {
		return NotInitializedErr("Postgres")
	}
	err := p.underline.Close()
	return err
}
func (p *Postgres) Health() error {
	if p.db == nil {
		return NotInitializedErr("Postgres")
	}
	err := p.underline.Ping()
	return err
}

func (p *Postgres) Db() *gorm.DB {
	return p.db
}
