package main

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Thing struct {
	ID    primitive.ObjectID `bson:"_id"`
	Hello string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}
	jake := client.Database("jake")
	c := jake.Collection("collection")

	_, err = c.InsertOne(nil, Thing{primitive.NewObjectID(), "Hi"})
	if err != nil {
		log.Fatal(err)
	}

	cursor, err := c.Find(nil, &Thing{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(nil) {
		var thing Thing
		if err = cursor.Decode(&thing); err != nil {
			log.Fatal(err)
		}
		log.Println(thing)
	}

	if cursor.Err() != nil {
		log.Fatal(err)
	}
}
