package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// this sets the graphql server's query and mutation definitions
var Schema, err = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

// this function helps to connect to mongodb
// it turns out connecting to mongodb is not as easy as other languages
func getMongo() (*mongo.Client, context.Context, context.CancelFunc) {
	// set the URI to be the mongodb server URI
	client, err := mongo.NewClient(options.Client().ApplyURI("???"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, cancel
}

func init() {
	// due to recursive definition error, have to put the author field of bookType in init function
	bookType.AddFieldConfig("author", &graphql.Field{
		Type: authorType,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			client, ctx, cancel := getMongo()
			defer cancel()

			authorCollection := client.Database("gql").Collection("authors")
			cur, err := authorCollection.Find(ctx, bson.M{})
			if err != nil {
				log.Fatal(err)
			}
			p, _ := params.Source.(Book)
			authorId := p.AuthorId
			var result Author
			for cur.Next(ctx) {
				var i Author
				err := cur.Decode(&i)
				if err != nil {
					log.Fatal(err)
				}
				if i.ID.Hex() == authorId {
					result = i
					break
				}
			}
			return result, nil
		},
	})
}

func main() {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		GraphiQL: true,
		Pretty:   true,
	})

	mux := http.NewServeMux()
	mux.Handle("/graphql", h)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("now listening on port 4000"))
	})
	// enable CORS
	handlerMux := cors.Default().Handler(mux)
	http.ListenAndServe(":4000", handlerMux)
}
