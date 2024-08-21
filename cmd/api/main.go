package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"time"
	"github.com/ante-neh/Rss-aggregator/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//Load the env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load environmental variables", err)
	}

	//Get the port number from .env
	port := os.Getenv("PORT")
	connectionString := os.Getenv("CONN")
	//get the port address from cmd
	address := flag.String("address", port, "Server address")
	dns := flag.String("dns", connectionString, "connection string")
	flag.Parse()

	//create custom loggers to handle errors and informations gracefully
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate)
	errorLogger := log.New(os.Stdout, "Error: ", log.Ltime|log.Ldate|log.Lshortfile)

	//create a database connection
	db, err := openDb(*dns)

	if err != nil {
		errorLogger.Fatal("Unable to Connect to the database", err)
	}

	defer db.Close()

	//create a new server type
	app := server.NewServer(infoLogger, errorLogger, *address, db)

	go app.StartScrapping(10, time.Minute)


	//start the server on port number *address
	app.InfoLogger.Println("Server is running on port: ", *address)
	server := app.Start()
	

	err = server.ListenAndServe()
	if err != nil {
		app.ErrorLogger.Println("Unable to Start the Server", err)
	}
}

func openDb(dns string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
