package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	// DBName name of the db
	DBName = "glottery2"
	// notesCollection name of the collection
	notesCollection = "notes"

	// URI mongo connection uri
	URI = "mongodb://myUserAdmin:123456@127.0.0.1/admin"
	//URI = "mongodb://myTester:123456@127.0.0.1/test?authSource=test"
)

// Note ...
type Note struct {
	ID        primitive.ObjectID `bson:"_id" json:"id, omitempty"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at, omitempty"`
	UpdateAt  time.Time          `bson:"updated_at" json:"update_at, omitempty"`
}

func main() {
	ctx := context.Background()

	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return
	}

	db := client.Database(DBName)
	coll := db.Collection(notesCollection)

	note := &Note{}

	// insert
	note.ID = primitive.NewObjectID()
	note.Title = "first note"
	note.Body = "first body"
	note.CreatedAt = time.Now()
	note.UpdateAt = time.Now()
	result, err := coll.InsertOne(ctx, note)
	if err != nil {
		fmt.Println(err)
		return
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)

	notes := []interface{}{}
	note2 := Note{}
	note2.ID = primitive.NewObjectID()
	note2.Title = "second title"
	note2.Body = "second body"
	note2.CreatedAt = time.Now()
	note2.UpdateAt = time.Now()

	note3 := &Note{
		ID:        primitive.NewObjectID(),
		Title:     "third title",
		Body:      "third body",
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	notes = append(notes, note2, note3)
	results, err := coll.InsertMany(ctx, notes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(results.InsertedIDs)

	// update
	resultUpdate, err := coll.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{
			"$set": bson.M{"body": "first update body", "updated_at": time.Now()}},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resultUpdate.ModifiedCount)

	// delete
	resultDelete, err := coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resultDelete.DeletedCount)

	// query
	findResult := coll.FindOne(ctx, bson.M{"_id": objectID})
	if err := findResult.Err(); err != nil {
		fmt.Println(err)
		return
	}

	n := Note{}
	err = findResult.Decode(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n.Body)

	noteResult := []Note{}
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for cursor.Next(ctx) {
		_ = cursor.Decode(&n)
		noteResult = append(noteResult, n)
	}

	for _, el := range noteResult {
		fmt.Println(el.Title)
	}

}
