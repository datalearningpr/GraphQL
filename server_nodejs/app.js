const express = require('express')
const graphqlHTTP = require('express-graphql')
const schema = require('./schema/schema')
const mongoose = require('mongoose')
var cors = require('cors')

const app = express()
// enable CORS
app.use(cors())

// set the URI to be the mongodb server URI
mongoose.connect('???', { 
    useNewUrlParser: true,
    dbName: 'gql'
})
mongoose.connection.once('open', () => {
    console.log('connected to database')
})

// this sets the graphql server's query and mutation definitions
app.use('/graphql', graphqlHTTP({
    schema: schema,
    graphiql: true
}))
app.listen(4000, () => {
    console.log('now listening on port 4000')
})