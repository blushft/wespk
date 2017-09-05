export default (name, item = false) => () => import(`views/${name}${index ? '/index' : ''}.vue`)
