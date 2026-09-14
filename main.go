package main

import (
	"HR_system/http_server"
	simpleconn "HR_system/postgres/simple_conn"
	sqlcommands "HR_system/postgres/sql_commands"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	pool, err := simpleconn.CreatePool(ctx)
	if err != nil {
		fmt.Println("Failed connection to postgreSQL:", err)
		return
	}
	defer pool.Close()

	repo := sqlcommands.NewEmployeeRepo(pool)
	if err := repo.CreateTableEmployees(ctx); err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	handlers := http_server.NewHandlers(repo)
	server := http_server.NewServer(handlers)

	if err := server.StartServer(); err != nil {
		fmt.Println("Failed connection to server:", err)
	}
}
