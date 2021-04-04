package storage

import (
	"context"
	"errors"
	"fmt"

	db "github.com/go-ozzo/ozzo-dbx"
	"github.com/mysa-fika/go/infrastructure/log"
)

// psqlConfiguration for creating a database storage
type psqlConfiguration struct {
	Host     string
	Port     int
	DBName   string `mapsstructure:"dbname"`
	Username string
	Password string
	Schema   string
	MaxIdle  int `mapstructure:"max_idle"`

	MaxOpen int    `mapstructure:"max_open"`
	SSLmode string `mapstructure:"sslmode"`
	Debug   bool
}

func loadDBConfiguration(options map[string]interface{}) (*psqlConfiguration, error) {

	host, ok := options["host"].(string)
	if !ok {
		return nil, errors.New("invalid type for DB host")
	}

	port, ok := options["port"].(int)
	if !ok {
		return nil, errors.New("invalid type for DB port")
	}

	dbName, ok := options["dbname"].(string)
	if !ok {
		return nil, errors.New("invalid type for DB name")
	}

	username, ok := options["username"].(string)
	if !ok {
		return nil, errors.New("invalid type for DB username")
	}

	password, ok := options["password"].(string)
	if !ok {
		return nil, errors.New("invalid type for DB password")
	}

	maxIdle, ok := options["max_idle"].(int)
	if !ok {
		return nil, errors.New("invalid type for DB max_idle")
	}

	maxOpen, ok := options["max_open"].(int)
	if !ok {
		return nil, errors.New("invalid type for DB max_open")
	}

	debug, ok := options["debug"].(bool)
	if !ok {
		return nil, errors.New("invalid type for DB debug")
	}

	return &psqlConfiguration{
		Host:     host,
		Port:     port,
		DBName:   dbName,
		Username: username,
		Password: password,
		MaxIdle:  maxIdle,
		MaxOpen:  maxOpen,
		Debug:    debug,
	}, nil
}

// Connection pool + logger
type PSQL struct {
	DB     *db.DB
	Logger *log.Logger
}

func NewPSQL(ctx context.Context, options map[string]interface{}, logger *log.Logger) (*PSQL, error) {
	dbConfig, err := loadDBConfiguration(options)
	if err != nil {
		return nil, fmt.Errorf("could not initialize DB storage with error: [%s]", err)
	}

	settings := fmt.Sprintf(
		"%s:%s@%s:%d/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	cp, err := db.Open("psql", settings)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the database connection pool")
	}

	err = cp.DB().Ping()
	if err != nil {
		logger.Fatal().Err(err).Msg("Database unreachable")
	}

	go func() error {
		logger.Info().Msg("[PSQL] Starting cleanup database hook...")
		<-ctx.Done()
		logger.Info().Msg("[PSQL] Cleaning up the database...")
		return cp.Close()
	}()

	cp.DB().SetMaxOpenConns(dbConfig.MaxOpen)
	cp.DB().SetMaxIdleConns(dbConfig.MaxIdle)

	return &PSQL{DB: cp, Logger: logger}, nil
}
