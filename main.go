package main

import (
	"log"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Thing struct {
	ID    primitive.ObjectID `bson:"_id"`
	Hello string
}

func main() {
	client, err := mongo.NewClient("mongodb://localhost")
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

	cursor, err := c.Find(nil, nil)
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
