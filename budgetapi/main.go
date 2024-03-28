package main

import (
	"budgetapi/handlers"
	"budgetstoragelib/mssql"
	"fmt"
	"net/http"
)

func main() {
	storage := mssql.GetInstance("server=localhost;user id=username;password=Some!Passw0rd;port=1433;database=Budgets")
	if err := storage.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	defer storage.Close()	
	
	handlers.NewHandlerFactory(storage).Load()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
		return
	}
}