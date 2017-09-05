<template>
  <div>
    <p>Completed Tasks: {{todos.filter(todo => {return todo.done === true}).length}}</p>
    <p>Pending Tasks: {{todos.filter(todo => {return todo.done === false}).length}}</p>
    <div class="columns">
      <div class="column" v-for="todo in todos">
        <div class="box">
          <article class="media">
            <div class="media-left">
            <figure class="image is-64x64">
              <img src="../assets/logo.png">
            </figure>
            </div>
            <div class="media-content">
              <div class="content">
                <p>
                  <strong>{{ todo.Title }}</strong>
                  <small>{{ todo.Created }}</small>
                  <br>
                  {{ todo.Project }}
                </p>
              </div>
              <nav class="level is-mobile">
                <div class="level-left">
                  <a class="level-item">
                    <b-icon pack="fa" icon="pencil-square-o" size="is-small"></b-icon>
                  </a>
                </div>
              </nav>
            </div>
          </article>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['todos']
}

</script>
<style></style>
