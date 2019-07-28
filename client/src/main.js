import Vue from 'vue'
import App from './App.vue'
import ApolloClient from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory'
import VueApollo from 'vue-apollo'

Vue.config.productionTip = false

const apolloClientTask = new ApolloClient({
  link: new HttpLink({
    // this is the dev localhost address, using port 4000
    uri: 'http://localhost:4000/graphql'
  }),
  cache: new InMemoryCache(),
  connectToDevTools: true,
})

const apolloProvider = new VueApollo({
  clients: {
    task: apolloClientTask
  },
  defaultClient: apolloClientTask
})

Vue.use(VueApollo)

new Vue({
  apolloProvider,
  render: h => h(App),
}).$mount('#app')
