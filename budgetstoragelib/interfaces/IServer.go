package interfaces

import "database/sql"

type IServer interface {
	Connect() error
	Database() *sql.DB	
	Close() error
}