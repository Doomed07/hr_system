package main

import (
	"HR_system/http"
	maincommands "HR_system/main_commands"
	"fmt"
)

func main() {
	table := maincommands.NewTableEmployees()
	hanlers := http.NewHandlers(table)
	server := http.NewServer(hanlers)

	if err:= server.StartServer(); err != nil{
		fmt.Println("Server failed:", err)
	}
}
