import graphene
from flask import Flask
from flask_cors import CORS
from flask_graphql import GraphQLView
from mongoengine import connect
from schemas import Query
from schemas import Mutation

# set the URI to be the mongodb server URI
connect(
    db='gql',
    host='???'
)

app = Flask(__name__)
# enable CORS
CORS(app)
# this sets the graphql server's query and mutation definitions
schema = graphene.Schema(query=Query, mutation=Mutation)

app.add_url_rule('/graphql', view_func=GraphQLView.as_view('graphql',
                                                           schema=schema, graphiql=True))

if __name__ == "__main__":
    app.run(port=4000, debug=True)
