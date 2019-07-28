<template>
  <div id="app" class="tile is-ancestor">
    <div class="tile is-6 is-vertical is-parent">
      <div id="main" class="tile is-child box">
      <h1 class="title is-2">reading list</h1>
      <BookList v-bind:books="books" @clickBook="ClickBook"></BookList>
      </div>
      <AddBook @refreshBookList="RefreshBookList"></AddBook>
    </div>
    
    <BookDetails :bookId="bookId" class="tile box"></BookDetails>
    
  </div>
</template>

<script>
import BookList from './components/BookList'
import AddBook from './components/AddBook'
import BookDetails from './components/BookDetails'
import {getBooksQuery} from './queries/queries'

export default {
  name: 'app',
  components: {
    BookList,
    AddBook,
    BookDetails
  },
  data: function() {
    return {
      // preset a value, should get from the books data array for the bookId of first book
      bookId: '5d32c47fa2b63c51249a8887'
    }
  },
  apollo: {
    books: {
      query: getBooksQuery
    }
  },
  methods: {
    RefreshBookList: function() {
    this.$apollo.queries.books.refetch();
    },
    ClickBook: function(e) {
      this.bookId = e;
    }
  }
}
</script>

<style>
#app {
  margin: 0;
  height: 100vh;
  text-align: center;
}

#main h1 {
  margin-bottom: 50px;
}
</style>
