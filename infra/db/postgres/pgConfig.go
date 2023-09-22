package PGConfig

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func ConnectPG() (*sql.DB, error){
	if err := godotenv.Load();err!= nil{
		log.Fatal("failed to connect .env")
		return nil, err
	}
	db_dsn := os.Getenv("PG_DB_DSN")
	
	db, _:= sql.Open("postgres", db_dsn)

	if err := db.Ping(); err != nil{
		db.Close()
		log.Fatal("Failed to connect DB!")
		return nil, err
	}
	return db, nil
}