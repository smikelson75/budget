module budgetstoragelib

go 1.22.0

replace budgetlib => ../budgetlib

require (
	budgetlib v0.0.0-00010101000000-000000000000
	github.com/microsoft/go-mssqldb v1.7.0
)

require (
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.19.0
)
