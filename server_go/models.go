package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// data models of Book and Author
// Golang does not provide good library for mongodb ORM
type Book struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Genre    string             `bson:"genre"`
	AuthorId string             `bson:"authorId"`
}

type Author struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  int                `bson:"age"`
}
