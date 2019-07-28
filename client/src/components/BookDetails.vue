<template>
  <div id="BookDetails">
    <div v-if="$apollo.loading"></div>
    <div class="content" v-else>
    <h2 class="is-medium">{{book.name}}</h2>
    <p class="is-medium">{{book.genre}}</p>
    <p class="is-medium">{{book.author.name}}</p>
    </div>
  </div>
</template>

<script>

import { getBookQuery } from '../queries/queries';

export default {
  name: 'BookDetails',
  props: ["bookId"],
  apollo: {
    book: {
      query: getBookQuery,
      variables () {
        return {id: this.bookId}
      }
    }
  },
  watch: {
      bookId: function(val) {
          this.$apollo.queries.book.refetch();
      }
  }
}
</script>

<style>
#BookDetails {
  height: 99vh;
}

.content {
  margin-left: auto;
  margin-right: auto;
}
</style>
