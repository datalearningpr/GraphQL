<template>
  <div id="BookList">
    <ul id="book-list" class="columns is-multiline">
        <li class="column box" v-for="item in books" v-bind:key="item.id" @click="$emit('clickBook', item.id)">
            {{ item.name }}
        </li>
    </ul>
  </div>
</template>

<script>

export default {
  name: 'BookList',
  props: ["books"],
}
</script>

<style>
.box {
  margin: 5px;
}
</style>
