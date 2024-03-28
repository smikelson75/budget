# API: Budget API

This project is a simple RESTful API written in GO.

## Workspace

The workspace contains 3 projects:

- budgetlib: The library that contains the business logic for Budgeting.
- budgetstoragelib: The library that contains the logic for accessing the database backend. See the [budget-db](https://github.com/smikelson75/budget-db) repository for the database code for SQL Server.
- budgetapi: The API that exposes the budgetlib as a RESTful API.

## Running the API

To run the API, you need to have the database running. The database code is in the [budget-db](https://github.com/smikelson75/budget-db) repository. Once the database is running, you can run the API using the following command:

```bash
go run ./budgetapi
```

## Current configuration

