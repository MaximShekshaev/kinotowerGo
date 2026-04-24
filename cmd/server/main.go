package main

import (
	"github.com/joho/godotenv"
	"github.com/MaximShekshaev/kinotowerGo/internal/core/logger"
	core_server "github.com/MaximShekshaev/kinotowerGo/internal/core/server"
	core_database "github.com/MaximShekshaev/kinotowerGo/internal/core/database"
)

func main() {
    if err := logger.Init("logs"); err != nil {
        panic("failed to init logger: " + err.Error())
    }
	db,err := core_database.NewDatabase()
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	defer db.DB.Close()


    logger.Log.Info("Starting server", "addr", ":8080")

    server := core_server.NewServer(db)

    if err := server.ListenAndServe(); err != nil {
        logger.Log.Error("Server stopped", "error", err)
    }
}
func init() {
	if err := godotenv.Load(); err != nil {	
		logger.Log.Warn("No .env file found, using environment variables")
	}
}