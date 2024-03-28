package mssql

import (
	"database/sql"
	"errors"
	"sync"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

var (
	server *Server
	once   sync.Once
)

type Server struct {
	connectionString string
	database         *sql.DB
	isOpen					 bool
}

func GetInstance(connectionString string) *Server {
	once.Do(func() {
		server = &Server{connectionString: connectionString, isOpen: false}
	})
	return server
}

func (s *Server) Connect() error {
	// Make sure to use sqlserver as the driver name
	// https://pkg.go.dev/github.com/microsoft/go-mssqldb#readme-query-parameter-token-replace-driver-quot-mssql-quot
	if s.isOpen {
		return nil
	}
	
	db, err := sql.Open("sqlserver", s.connectionString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)

	s.database = db
	return nil
}

func (s *Server) Close() error {
	if s.database == nil {
		return errors.New("database connection is not open")
	}

	if err := s.database.Close(); err != nil {
		return err
	}

	return nil
}

func (s Server) Database() *sql.DB {
	return s.database
}