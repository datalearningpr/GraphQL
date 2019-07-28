package main

import (
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"genre": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var authorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"books": &graphql.Field{
			Type: graphql.NewList(bookType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				client, ctx, cancel := getMongo()
				defer cancel()
				bookCollection := client.Database("gql").Collection("books")
				cur, err := bookCollection.Find(ctx, bson.M{})
				if err != nil {
					log.Fatal(err)
				}

				p, _ := params.Source.(Author)
				authorId := p.ID
				var result []Book

				for cur.Next(ctx) {
					var i Book
					err := cur.Decode(&i)
					if err != nil {
						log.Fatal(err)
					}
					if i.AuthorId == authorId.Hex() {
						result = append(result, i)
					}
				}
				return result, nil
			},
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"book": &graphql.Field{
			Type:        bookType,
			Description: "Get a book by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				client, ctx, cancel := getMongo()
				defer cancel()

				id := params.Args["id"].(string)

				bookCollection := client.Database("gql").Collection("books")
				cur, err := bookCollection.Find(ctx, bson.M{})
				if err != nil {
					log.Fatal(err)
				}
				var result Book
				for cur.Next(ctx) {
					var i Book
					err := cur.Decode(&i)
					if err != nil {
						log.Fatal(err)
					}
					if i.ID.Hex() == id {
						result = i
						break
					}
				}
				return result, nil
			},
		},
		"books": &graphql.Field{
			Type:        graphql.NewList(bookType),
			Description: "Get books",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				client, ctx, cancel := getMongo()
				defer cancel()
				bookCollection := client.Database("gql").Collection("books")
				cur, err := bookCollection.Find(ctx, bson.M{})
				if err != nil {
					log.Fatal(err)
				}

				var result []bson.M
				for cur.Next(ctx) {
					var i Book
					err := cur.Decode(&i)
					if err != nil {
						log.Fatal(err)
					}
					result = append(result, bson.M{"id": i.ID.Hex(), "name": i.Name, "genre": i.Genre, "authorId": i.AuthorId})
				}
				return result, nil
			},
		},
		"author": &graphql.Field{
			Type:        authorType,
			Description: "Get a author by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				client, ctx, cancel := getMongo()
				defer cancel()

				id := params.Args["id"].(string)

				authorCollection := client.Database("gql").Collection("authors")
				cur, err := authorCollection.Find(ctx, bson.M{})
				if err != nil {
					log.Fatal(err)
				}
				var result Author
				for cur.Next(ctx) {
					var i Author
					err := cur.Decode(&i)
					if err != nil {
						log.Fatal(err)
					}
					if i.ID.Hex() == id {
						result = i
						break
					}
				}
				return result, nil
			},
		},
		"authors": &graphql.Field{
			Type:        graphql.NewList(authorType),
			Description: "Get authors",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				client, ctx, cancel := getMongo()
				defer cancel()
				authorCollection := client.Database("gql").Collection("authors")
				cur, err := authorCollection.Find(ctx, bson.M{})
				if err != nil {
					log.Fatal(err)
				}

				var result []bson.M

				for cur.Next(ctx) {
					var i Author
					err := cur.Decode(&i)
					if err != nil {
						log.Fatal(err)
					}
					result = append(result, bson.M{"id": i.ID.Hex(), "name": i.Name, "age": i.Age})
				}
				return result, nil
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addAuthor": &graphql.Field{
			Type: authorType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"age": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				name := params.Args["name"].(string)
				age := params.Args["age"].(int)

				newAuthor := bson.M{"name": name, "age": age}

				client, ctx, cancel := getMongo()
				defer cancel()
				authorCollection := client.Database("gql").Collection("authors")
				insertResult, err := authorCollection.InsertOne(ctx, newAuthor)
				if err != nil {
					log.Fatal(err)
				}
				cur, err := authorCollection.Find(ctx, bson.M{})
				if err != nil {
					log.Fatal(err)
				}
				var result Author
				for cur.Next(ctx) {
					var i Author
					err := cur.Decode(&i)
					if err != nil {
						log.Fatal(err)
					}
					if i.ID.Hex() == insertResult.InsertedID.(primitive.ObjectID).Hex() {
						result = i
						break
					}
				}
				return result, nil
			},
		},
		"addBook": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"genre": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

				"authorId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				name := params.Args["name"].(string)
				genre := params.Args["genre"].(string)
				authorId := params.Args["authorId"].(string)

				newBook := bson.M{"name": name, "genre": genre, "authorId": authorId}

				client, ctx, cancel := getMongo()
				defer cancel()
				bookCollection := client.Database("gql").Collection("books")
				insertResult, err := bookCollection.InsertOne(ctx, newBook)
				if err != nil {
					log.Fatal(err)
				}
				cur, err := bookCollection.Find(ctx, bson.M{})
				if err != nil {
					log.Fatal(err)
				}
				var result Book
				for cur.Next(ctx) {
					var i Book
					err := cur.Decode(&i)
					if err != nil {
						log.Fatal(err)
					}
					if i.ID.Hex() == insertResult.InsertedID.(primitive.ObjectID).Hex() {
						result = i
						break
					}
				}
				return result, nil
			},
		},
	},
})
