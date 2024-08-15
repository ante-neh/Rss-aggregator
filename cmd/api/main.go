package main

import (
	"flag"
	"log"
	"os"
	"github.com/ante-neh/Rss-aggregator/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	//Load the env files
	err := godotenv.Load() 
	if err != nil{
		log.Fatal("Unable to Load environmental variables", err)
	}

	//Get the port number from .env
	port := os.Getenv("PORT")

	//get the port address from cmd
	address := flag.String("address", port, "Server address")
	flag.Parse()

	//create custom loggers to handle errors and informations gracefully 
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ltime | log.Ldate)
	errorLogger := log.New(os.Stdout, "Error: ", log.Ltime | log.Ldate | log.Lshortfile)

	//create a new server type 
	app := server.NewServer(infoLogger, errorLogger, *address)

	//start the server on port number *address
	server := app.Start() 

	app.InfoLogger.Println("Server is running on port: ", *address)
	
	err = server.ListenAndServe() 
	if err != nil{
		app.ErrorLogger.Println("Unable to Start the Server", err)
	}
}