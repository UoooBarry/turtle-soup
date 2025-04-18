declare module '*.vue' {
  import { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

declare module 'vuetify/styles' {
  import { GlobalTheme } from 'vuetify'
  const styles: GlobalTheme
  export default styles
}
