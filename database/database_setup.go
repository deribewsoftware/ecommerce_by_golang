package database

import (
	"context"
	"log"
	"time"
"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://deribewtech:golang_rest@cluster0.sklnweb.mongodb.net/"))

	if err != nil {
		log.Fatal(err)
	}
	ctx,cancel:context.WithTimeout(context.Background(),10*time.Second)
 defer cancel()
 err=client.Connect(ctx)
 if err != nil {
	log.Fatal(err)
 }
 err=client.Ping(context.TODO(),nil)
 if err != nil {
	log.Println("failed to connect to MongoDB:(")
	return nil
 }
 fmt.Println("successfully connected to MongoDB")
 return client

}
var Clinet *mongo.Client=DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {

}
func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {

}
