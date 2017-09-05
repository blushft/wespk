<template>
  <section class="section">
    <div class="container">
      <todo-list v-bind:todos="todos"></todo-list>
    </div>
  </section>
</template>

<script>
import TodoList from './TodoList'
import axios from 'axios'

export default {
  name: 'todo',
  components: {
    TodoList
  },
  data () {
    return {
      todos: [],
      errors: []
    }
  },
  created () {
    axios.get('/api/v1/todos')
    .then(response => {
      this.todos = response.data
    })
    .catch(e => {
      this.errors.push(e)
    })
  }

}
</script>
