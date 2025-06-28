package handler

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/aarondl/sqlboiler/v4/boil"
)

type DBConfig struct {
	Dbname string `toml:"dbname"`
	Host   string `toml:"host"`
	Port   int64  `toml:"port"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

func tomlRead() (*DBConfig, error) {
	path := os.Getenv("DATABASE_TOML_PATH")
	if path == "" {
		path = "command/infra/sqlboiler/config/database.toml"
	}

	m := map[string]DBConfig{}
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}

	config := m["mysql"]
	return &config, nil
}

func DBConnect() error {
	config, err := tomlRead()
	if err != nil {
		return err
	}

	rdbms := "mysql"
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Pass, config.Host, config.Port, config.Dbname)
	conn, err := sql.Open(rdbms, connectStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	MAX_IDLE_CONNS := 10
	MAX_OPEN_CONNS := 100
	CONN_MAX_LIFETIME := 300 * time.Second

	conn.SetMaxIdleConns(MAX_IDLE_CONNS)
	conn.SetMaxOpenConns(MAX_OPEN_CONNS)
	conn.SetConnMaxLifetime(CONN_MAX_LIFETIME)

	boil.SetDB(conn)
	boil.DebugMode = true

	return nil
}
