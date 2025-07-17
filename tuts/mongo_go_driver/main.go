package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func ifErrorlog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	var client, err = mongo.Connect(context.TODO(), clientOptions)

	ifErrorlog(err)
	err = client.Ping(context.TODO(), nil)

	ifErrorlog(err)

	fmt.Println("Connected to MongoDB!")
	// collection := client.Database("test").Collection("trainers")

	// filter := bson.D{
	// 	{Key: "name", Value: bson.D{
	// 		{Key: "$in", Value: bson.A{"Alice", "Bob"}},
	// 	}},
	// }

	// // Example usage: Find documents matching the filter
	// cursor, err := collection.Find(context.TODO(), filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(context.TODO())

	// for cursor.Next(context.TODO()) {
	// 	var trainer Trainer
	// 	if err := cursor.Decode(&trainer); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Found trainer: %+v\n", trainer)
	// }

	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulian City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	// insertResult, err := collection.InsertOne(context.TODO(), ash)

	// ifErrorlog(err)

	// fmt.Printf("Inserted a single document: %s\n", insertResult.InsertedID)

	// trainers := []any{misty, brock}

	// inserManyResult, err := collection.InsertMany(context.TODO(), trainers)

	// ifErrorlog(err)

	// fmt.Printf("Inserted many documents: %v\n", inserManyResult.InsertedIDs)

	// filter := bson.D{{Key: "name", Value: "Ash"}}

	// update := bson.D{
	// 	{Key: "$inc", Value: bson.D{
	// 		{Key: "age", Value: 1},
	// 	}},
	// }

	// updateRes, err := collection.UpdateOne(context.TODO(), filter, update)

	// ifErrorlog(err)

	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateRes.MatchedCount, updateRes.ModifiedCount)

	// var result Trainer

	// err = collection.FindOne(context.TODO(), filter).Decode(&result)

	// ifErrorlog(err)

	// fmt.Printf("Found a single document: %+v\n", result)

	// findOptions := options.Find()
	// findOptions.SetLimit(2)

	// var results []*Trainer

	// cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	// ifErrorlog(err)

	// for cur.Next(context.TODO()) {
	// 	var elem Trainer
	// 	err := cur.Decode(&elem)

	// 	ifErrorlog(err)

	// 	results = append(results, &elem)
	// }

	// err = cur.Err()
	// ifErrorlog(err)

	// cur.Close(context.TODO())

	// fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	// ifErrorlog(err)
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

}
