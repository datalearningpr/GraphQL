<template>
  <div id="AddBook">
    <form id="add-book" v-on:submit.prevent="onSubmit()">
        <div class="field is-horizontal">
          <div class="field-label">
            <label class="label">Book Name:</label>
          </div>
          <div class="field-body">
            <input class="input" type="text" v-model="name"/>
          </div>
        </div>
        <div class="field is-horizontal">
          <div class="field-label">
            <label class="label">Genre:</label>
          </div>
          <div class="field-body">
            <input class="input" type="text" v-model="genre"/>
          </div>
        </div>
        <div class="field is-horizontal">
          <div class="field-label">
            <label class="label">Author:</label>
          </div>
          <div class="field-body">
            <select class="select is-8" v-model='authorId'>
                <option>Select author</option>
                <option v-for="item in authors" v-bind:key="item.id" v-bind:value="item.id">
                    {{ item.name }}
                </option>
            </select>
          </div>
        </div>
        <button class="button is-success is-rounded">+</button>
    </form>
  </div>
</template>

<script>

import {getAuthorsQuery, addBookMutation} from '../queries/queries'

export default {
  name: 'AddBook',
  mounted: async function() {
      const authorsQuery =  await this.$apollo.query({
          query: getAuthorsQuery
        })
      this.authors = authorsQuery.data.authors
  },
  data: function() {
      return {
        authors: null,
        name: '',
        genre: '',
        authorId: ''
      }
  },
  methods: {
      onSubmit: function() {
        this.$apollo.mutate({
            mutation: addBookMutation,
            variables: {
                name: this.name,
                genre: this.genre,
                authorId: this.authorId
            }
        }).then((data) => {
            console.log(data)
            this.$emit('refreshBookList')
        })
      }
  }
}
</script>

<style>
</style>
