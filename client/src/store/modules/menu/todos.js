import lazyLoading from './lazyLoading'

export default {
  name: 'Todos',
  path: '/',
  meta: {
    icon: 'fa-edit',
    expanded: false,
    link: 'todos/index.vue'
  },
  component: lazyLoading('todos', true),
  children: []
}
