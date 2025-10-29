package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	URI = "mongodb://localhost:27017"
	timeout = 3 // 3 seconds
	logFolder = "/Users/venkata.koushik/Pro/netV/logs"
	logfile = logFolder+"Mongo.log"  // This woul dbe updated later 
)

func main(){
	os.Mkdir(logFolder , 0666)
	file, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE , 0666)
	if err!=nil {
		panic( err)
	}
	log.SetOutput(file)
	log.SetPrefix("MONGOCLIENT::")
	_, cancel := context.WithTimeout(context.Background(), timeout* time.Second)
	defer cancel()
	clientOps := options.Client().ApplyURI(URI)
	_ , err = mongo.Connect(clientOps)

	if err!=nil{
		log.Fatalln("[ERROR] Cannot connect to Mongo. Error :", err)
	}
	log.Println("Connected to Mongo DB :",URI)



}