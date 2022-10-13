package database

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type File struct {
	ID        primitive.ObjectID `bson:"_id"`
	Type      string             `bson:"type"`
	Data      string             `bson:"data"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type Files struct {
	collection *mongo.Collection
}

func GetFiles(db *mongo.Database) *Files {
	return &Files{
		collection: db.Collection("files"),
	}
}

func (f *Files) Upload(b64, mime string) *File {
	file := &File{
		ID:        primitive.NewObjectID(),
		Type:      mime,
		Data:      b64,
		CreatedAt: time.Now(),
	}

	_, err := f.collection.InsertOne(Ctx, file)
	if err != nil {
		fmt.Printf("err: %#v\n", err)
		return nil
	}

	return file
}

func (f *Files) Get(id primitive.ObjectID) *File {
	doc := &File{}
	err := f.collection.FindOne(Ctx, bson.D{
		bson.E{
			Key: "id", Value: id,
		},
	}).Decode(&doc)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			fmt.Printf("err: %#v\n", err)
		}
		return nil
	}

	return doc
}
